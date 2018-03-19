package tests

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	//"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/server"
)

// Tests assume the `basecoind` and `basecli` binaries
// have been built and are located in `./build`

//nolint
var (
	gopath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "cosmos", "cosmos-sdk")

	basecoind = "build/basecoind"
	basecli   = "build/basecli"

	basecoindPath = filepath.Join(gopath, basecoind)
	basecliPath   = filepath.Join(gopath, basecli)

	basecoindDir = "./tmp-basecoind-tests"
	basecliDir   = "./tmp-basecli-tests"

	ACCOUNTS = []string{"alice", "bob", "charlie", "igor"}
	alice    = ACCOUNTS[0]
	bob      = ACCOUNTS[1]
	charlie  = ACCOUNTS[2]
	igor     = ACCOUNTS[3]
)

func TestMain(m *testing.M) {
	// setup should start server
	// and make keys, then run tests

	m.Run()

	cleanUp()

}

// Init Basecoin Test
func TestInitBasecoin(t *testing.T) {
	var err error

	password := "some-random-password"

	initBasecoind := exec.Command(basecoindPath, "init", "--home", basecoindDir)
	cmdWriter, err := initBasecoind.StdinPipe()
	assert.Nil(t, err)

	buf := new(bytes.Buffer)
	initBasecoind.Stdout = buf

	err = initBasecoind.Start()
	assert.Nil(t, err)

	_, err = cmdWriter.Write([]byte(password))
	assert.Nil(t, err)
	cmdWriter.Close()

	err = initBasecoind.Wait()
	assert.Nil(t, err)
}

func _TestSendCoins(t *testing.T) {
	startServer(t)

	// send some coins
	// [zr] where dafuq do I get a FROM (oh, use --name)

	sendTo := fmt.Sprintf("--to=%s", bob)
	sendFrom := fmt.Sprintf("--from=%s", alice)

	cmdOut, err := exec.Command(basecliPath, "send", sendTo, "--amount=1000mycoin", sendFrom, "--seq=0").Output()
	assert.Nil(t, err)

	fmt.Printf("sent: %s", string(cmdOut))
}

// Init Basecoin Test
func initServerForTest(t *testing.T) {

	var err error

	password := "some-random-password"
	usePassword := exec.Command("echo", password)

	initBasecoind := exec.Command(basecoindPath, "init", "--home", basecoindDir)

	initBasecoind.Stdin, err = usePassword.StdoutPipe()
	assert.Nil(t, err)

	initBasecoind.Stdout = os.Stdout

	err = initBasecoind.Start()
	assert.Nil(t, err)
	err = usePassword.Run()
	assert.Nil(t, err)
	err = initBasecoind.Wait()
	assert.Nil(t, err)

	makeKeys(t)
}

// expects TestInitBaseCoin to have been run
func startNodeServerForTest(t *testing.T, home string) *exec.Cmd {
	cmdName := basecoindPath
	cmdArgs := []string{"start", "--home", home}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	assert.Nil(t, err)

	// FIXME: if there is a nondeterministic node start failure,
	// we should probably make this read the logs to wait for RPC
	time.Sleep(time.Second * 2)

	return cmd
}

// expects TestInitBaseCoin to have been run
func startLCDServerForTest(t *testing.T, home, chainID string) (cmd *exec.Cmd, port string) {
	cmdName := basecliPath
	port = strings.Split(server.FreeTCPAddr(t), ":")[2]
	cmdArgs := []string{
		"rest-server",
		"--home",
		home,
		"--bind",
		fmt.Sprintf("localhost:%s", port),
		"--chain-id",
		chainID,
	}
	cmd = exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	assert.Nil(t, err)
	time.Sleep(time.Second * 2) // TODO: LOL
	return cmd, port
}

func makeKeys(t *testing.T) {
	for _, acc := range ACCOUNTS {
		makeKeys := exec.Command(basecliPath, "keys", "add", acc, "--home", basecliDir)
		cmdWriter, err := makeKeys.StdinPipe()
		assert.Nil(t, err)

		makeKeys.Stdout = os.Stdout
		err = makeKeys.Start()
		assert.Nil(t, err)

		cmdWriter.Write([]byte("1234567890"))
		cmdWriter.Close()

		err = makeKeys.Wait()
		assert.Nil(t, err)
	}
}

// expects TestInitBaseCoin to have been run
func startServer(t *testing.T) {
	// straight outta https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
	cmdName := basecoindPath
	cmdArgs := []string{"start", "--home", basecoindDir}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	assert.Nil(t, err)

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("running [basecoind start] %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	assert.Nil(t, err)

	err = cmd.Wait()
	assert.Nil(t, err)

	time.Sleep(5 * time.Second)

	// TODO return cmd.Process so that we can later do something like:
	// cmd.Process.Kill()
	// see: https://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang
}

// clean the directories
func cleanUp() {
	// ignore errors b/c the dirs may not yet exist
	os.RemoveAll(basecoindDir)
	os.RemoveAll(basecliDir)
}
