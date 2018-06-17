package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// SocketClient represents a general dialable socket
type SocketClient interface {
	TryDial() error
	WriteString(str string) error
	Close()
}

// UnixSocketClient represents a Unix socket
type UnixSocketClient struct {
	address string
	conn    net.Conn
}

// MaxAttemptsToDialSocket is the Number of maximum attempts to dial the java_pid socket
const MaxAttemptsToDialSocket = 10

// WaitToRetryDialingSocketNs is the amount of time to wait before retry dialing java_pid socket
const WaitToRetryDialingSocketNs = 1e9

// TryDial tries dialing an Unix Socket given an socket path and retries if it fails
func (socket *UnixSocketClient) TryDial() error {
	var err error

	for i := 0; i < MaxAttemptsToDialSocket; i++ {
		time.Sleep(WaitToRetryDialingSocketNs)
		log.Printf("Waiting for socket at '%s' to become available", socket.address)
		socket.conn, err = net.Dial("unix", socket.address)

		if err != nil {
			err = fmt.Errorf("Could not connect to socket after %d attemps. Reason: %s", MaxAttemptsToDialSocket, err)
			break
		} else if socket.conn != nil {
			log.Println("Connected")
			break
		}
	}

	return err
}

// Close closes the underlying socket connection
func (socket *UnixSocketClient) Close() {
	if socket.conn != nil {
		socket.conn.Close()
	}
}

// WriteString writes a null-terminated string into the socket
func (socket *UnixSocketClient) WriteString(str string) error {
	_, err := socket.conn.Write([]byte(str + "\000"))
	return err
}

// NewUnixSocketClient creates a new client for unix socket
func NewUnixSocketClient(address string) SocketClient {
	return &UnixSocketClient{address: address, conn: nil}
}
