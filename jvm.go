package main

import (
	"fmt"
	"os"
	"syscall"
)

// AttachableJvm represents a JVM attachable through a socket
type AttachableJvm interface {
	Pid() int
	Cwd() string
	AttachSocketAddr() string
	SigQuit() error
}

// JvmProcess represents a Java Virtual Machine process
type JvmProcess struct {
	pid int
}

// Pid returns the JVM process identifier
func (jvm *JvmProcess) Pid() int {
	return jvm.pid
}

// Sigquit sends a SIGQUIT signal to the JVM
func (jvm *JvmProcess) SigQuit() error {
	proc, _ := os.FindProcess(jvm.pid)
	return proc.Signal(syscall.SIGQUIT)
}

// Cwd returns the current working directory for a given pid
func (jvm *JvmProcess) Cwd() string {
	cwd := fmt.Sprintf("/proc/%d/cwd", jvm.pid)

	if _, err := os.Stat(cwd); os.IsNotExist(err) {
		cwd = os.TempDir()
	}

	return cwd
}

// AttachSocketAddr returns the name of the attach file to create
func (jvm *JvmProcess) AttachSocketAddr() string {
	return fmt.Sprintf("%s/.java_pid%d", os.TempDir(), jvm.pid)
}

// NewJvmProcess creates a new instance of a Jvm Process given a pid
func NewJvmProcess(pid int) AttachableJvm {
	return &JvmProcess{pid: pid}
}
