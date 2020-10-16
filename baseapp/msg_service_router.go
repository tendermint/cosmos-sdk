package baseapp

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"

	gogogrpc "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgServiceRouter routes fully-qualified Msg service methods to their handler.
type MsgServiceRouter struct {
	interfaceRegistry codectypes.InterfaceRegistry
	routes            map[string]MsgServiceHandler
}

var _ gogogrpc.Server = &MsgServiceRouter{}

// NewMsgServiceRouter creates a new MsgServiceRouter.
func NewMsgServiceRouter() *MsgServiceRouter {
	return &MsgServiceRouter{
		routes: map[string]MsgServiceHandler{},
	}
}

// MsgServiceHandler defines a function type which handles Msg service message.
type MsgServiceHandler = func(ctx sdk.Context, req sdk.MsgRequest) (*sdk.Result, error)

// Handler returns the MsgServiceHandler for a given query route path or nil
// if not found.
func (msr *MsgServiceRouter) Handler(methodName string) MsgServiceHandler {
	return msr.routes[methodName]
}

// RegisterService implements the gRPC Server.RegisterService method. sd is a gRPC
// service description, handler is an object which implements that gRPC service.
func (msr *MsgServiceRouter) RegisterService(sd *grpc.ServiceDesc, handler interface{}) {
	// Adds a top-level query handler based on the gRPC service name.
	for _, method := range sd.Methods {
		fqMethod := fmt.Sprintf("/%s/%s", sd.ServiceName, method.MethodName)
		methodHandler := method.Handler

		// NOTE: This is how we pull the concrete request type for each handler for registering in the InterfaceRegistry.
		// This approach is maybe a bit hacky, but less hacky than reflecting on the handler object itself.
		// We use a no-op interceptor to avoid actually calling into the handler itself.
		_, _ = methodHandler(nil, context.Background(), func(i interface{}) error {
			msg, ok := i.(proto.Message)
			if !ok {
				// We panic here because there is no other alternative and the app cannot be initialized correctly
				// this should only happen if there is a problem with code generation in which case the app won't
				// work correctly anyway.
				panic(fmt.Errorf("can't register request type %T for service method %s", i, fqMethod))
			}

			msr.interfaceRegistry.RegisterCustomTypeURL((*sdk.MsgRequest)(nil), fqMethod, msg)
			return nil
		}, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			return nil, nil
		})

		msr.routes[fqMethod] = func(ctx sdk.Context, req sdk.MsgRequest) (*sdk.Result, error) {
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Call the method handler from the service description with the handler object.
			res, err := methodHandler(handler, sdk.WrapSDKContext(ctx), func(_ interface{}) error {
				// We don't do any decoding here because the decoding was already done.
				return nil
			}, func(goCtx context.Context, _ interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
				goCtx = context.WithValue(goCtx, sdk.SdkContextKey, ctx)
				return handler(goCtx, req)
			})
			if err != nil {
				return nil, err
			}

			resMsg, ok := res.(proto.Message)
			if !ok {
				return nil, fmt.Errorf("can't proto encode %T", resMsg)
			}

			return sdk.WrapServiceResult(ctx, resMsg, err)
		}
	}
}

// SetInterfaceRegistry sets the interface registry for the router.
func (msr *MsgServiceRouter) SetInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) {
	msr.interfaceRegistry = interfaceRegistry
}
