package container

import "reflect"

type Option interface {
	isOption()
}

func Provide(constructors ...interface{}) Option {
	panic("TODO")
}

func ProvideWithScope(scope Scope, constructors ...interface{}) Option {
	panic("TODO")
}

func AutoGroupTypes(types ...reflect.Type) Option {
	panic("TODO")
}

func OnePerScopeTypes(types ...reflect.Type) Option {
	panic("TODO")
}

func Error(err error) Option {
	panic("TODO")
}

func Options(opts ...Option) Option {
	panic("TODO")
}
