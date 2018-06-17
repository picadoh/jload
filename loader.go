package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// AgentLoader represents a component capable of loading an agent
type AgentLoader interface {
	LoadAgent(agent string) error
}

// JvmAgentLoader represents a component capable of loading an agent into a JVM
type JvmAgentLoader struct {
	jvm AttachableJvm
}

func (loader *JvmAgentLoader) openAttachListener() error {
	attachFilePath := attachFilePath(
		loader.jvm.Cwd(),
		loader.jvm.Pid(),
	)

	var err error

	// create attach file
	if _, err = os.Stat(attachFilePath); os.IsNotExist(err) {
		log.Printf("Creating file: %s\n", attachFilePath)
		_, err = os.Create(attachFilePath)
	}

	if err == nil {
		// send sigquit so JVM loads that file
		loader.jvm.SigQuit()
	}

	return err
}

func attachFilePath(cwd string, pid int) string {
	return fmt.Sprintf("%s/.attach_pid%d", cwd, pid)
}

func writeLoadMessage(socket SocketClient, agent string) (string, error) {
	msgArgs := []string{
		"1",          // protocol version
		"load",       // command
		"instrument", // instrumentation
		"false",      // path may be relative
		agent,        // agent (path to jar)
	}

	msg := strings.Join(msgArgs, " ")

	var err error

	for _, msgArg := range msgArgs {
		err := socket.WriteString(msgArg)
		if err != nil {
			break
		}
	}

	return msg, err
}

func (loader *JvmAgentLoader) sendCommand(agent string) error {
	var err error

	socketAddr := loader.jvm.AttachSocketAddr()
	socket := NewUnixSocketClient(socketAddr)

	err = socket.TryDial()

	if err == nil {
		defer socket.Close()

		msg, err := writeLoadMessage(socket, agent)

		if err != nil {
			log.Fatalf("Could not send %s to socket at %s\n", msg, socketAddr)
		} else {
			log.Printf("Sent %s to socket at %s\n", msg, socketAddr)
		}
	}

	return err
}

// LoadAgent loads a givne agent into the JVM
func (loader *JvmAgentLoader) LoadAgent(agent string) error {
	// $TMPDIR/.attach_pid<PID> file
	loader.openAttachListener()

	// send load command
	return loader.sendCommand(agent)
}

// NewAgentLoader creates a new agent loader for a given JVM
func NewAgentLoader(jvm AttachableJvm) AgentLoader {
	return &JvmAgentLoader{jvm: jvm}
}
