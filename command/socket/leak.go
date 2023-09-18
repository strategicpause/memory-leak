package socket

import (
	"fmt"
	"github.com/strategicpause/memory-leak/metrics"
	"golang.org/x/sys/unix"
	"time"
)

const (
	// MaxConnections defines the maximum number of connections that can be accepted by a given socket. This value is
	// defined by  /proc/sys/net/core/somaxconn.
	MaxConnections = 4096
	StartPort      = 9090
	// By default each established socket connection will write 1 KiB to the buffer.
	KiB = 1024
)

var (
	LocalAddr = [4]byte{127, 0, 0, 1}
)

type Params struct {
	NumSockets            int64
	NetworkAddressDomain  int
	ConnectionType        int
	CommunicationProtocol int
	PauseTimeInSeconds    time.Duration
}

func tcpLeak(params *Params) error {
	PrintParams(params)

	var stopChan chan bool
	var err error

	nextPort := StartPort
	currentPort := nextPort

	for i := int64(0); i < params.NumSockets; i++ {
		// Setup a new service every 4096 connections.
		if i%MaxConnections == 0 {
			if stopChan != nil {
				stopChan <- true
			}
			currentPort = nextPort
			stopChan, err = resetServer(params, currentPort)
			if err != nil {
				return err
			}
			nextPort = nextPort + 1
		}

		clientFd, err := unix.Socket(params.NetworkAddressDomain, params.ConnectionType, params.CommunicationProtocol)
		if err != nil {
			return err
		}

		if err = unix.Connect(clientFd, &unix.SockaddrInet4{Port: currentPort, Addr: LocalAddr}); err != nil {
			return err
		}

		_, err = unix.Write(clientFd, make([]byte, KiB))

		if i%100 == 0 {
			metrics.PrintSocketStats()
			time.Sleep(params.PauseTimeInSeconds)
		}
	}

	return nil
}

// resetServer will create a new service socket after 4096 connections (which is the maximum backlog value for listen).
func resetServer(params *Params, port int) (chan bool, error) {
	stopChan := make(chan bool, 1)
	serverFd := Must(unix.Socket(params.NetworkAddressDomain, params.ConnectionType, params.CommunicationProtocol))

	serviceAddr := &unix.SockaddrInet4{
		Port: port,
		Addr: LocalAddr,
	}

	if err := unix.Bind(serverFd, serviceAddr); err != nil {
		return stopChan, err
	}

	if err := unix.Listen(serverFd, MaxConnections); err != nil {
		return stopChan, err
	}
	fmt.Printf("Service bound to 127.0.0.1:%d.\n", port)

	go func() {
		for {
			select {
			case <-stopChan:
				unix.Close(serverFd)
				return
			default:
				unix.Accept(serverFd)
			}
		}
	}()

	return stopChan, nil
}

func PrintParams(params *Params) {
	fmt.Printf("NumSockets: %v\n", params.NumSockets)
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}
