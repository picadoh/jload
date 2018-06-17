package main

import (
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestLoadAgent(t *testing.T) {

	pid := 1234

	stub := &AttachableJvmStub{pid: pid}

	subject := NewAgentLoader(stub)

	testAgent := "testagent"

	go func() {
		subject.LoadAgent(testAgent)
	}()

	serverListener, err := net.Listen("unix", stub.AttachSocketAddr())
	if err != nil {
		panic(err)
	}
	defer serverListener.Close()
	defer os.Remove(stub.AttachSocketAddr())

	for {
		conn, _ := serverListener.Accept()
		defer conn.Close()

		buf, _ := ioutil.ReadAll(conn)

		expectedMessage := "1\000load\000instrument\000false\000" + testAgent + "\000"

		if msg := string(buf[:]); msg != expectedMessage {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, expectedMessage)
		}
		return
	}

}

type AttachableJvmStub struct {
	ServerListener net.Listener
	pid            int
}

func (stub *AttachableJvmStub) Pid() int {
	return stub.pid
}

func (stub *AttachableJvmStub) Cwd() string {
	return os.TempDir()
}

func (stub *AttachableJvmStub) AttachSocketAddr() string {
	return os.TempDir() + string(os.PathSeparator) + "jload_loader_test.sock"
}

func (stub *AttachableJvmStub) SigQuit() error {
	// open socket
	stub.ServerListener, _ = net.Listen("unix", stub.AttachSocketAddr())
	return nil
}
