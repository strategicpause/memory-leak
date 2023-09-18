package socket

import (
	"github.com/urfave/cli"
	"golang.org/x/sys/unix"
	"time"
)

const (
	NumSocketsName            = "num-sockets"
	CommunicationProtocolName = "comm-protocol"
	PauseDurationName         = "pause"
)

func Register() cli.Command {
	return cli.Command{
		Name:   "socket",
		Usage:  "Reproduces a socket leak.",
		Action: action,
		Flags:  flags(),
	}
}

func flags() []cli.Flag {
	return []cli.Flag{
		cli.Int64Flag{
			Name:  NumSocketsName,
			Usage: "specify the number of sockets to create.",
			Value: 9223372036854775807,
		},
		cli.DurationFlag{
			Name:  PauseDurationName,
			Usage: "Time between allocations in seconds.",
			Value: time.Second,
		},
	}
}

func action(ctx *cli.Context) error {
	params := &Params{
		NumSockets:            ctx.Int64(NumSocketsName),
		NetworkAddressDomain:  unix.AF_INET,
		ConnectionType:        unix.SOCK_STREAM,
		CommunicationProtocol: unix.IPPROTO_TCP,
		PauseTimeInSeconds:    ctx.Duration(PauseDurationName),
	}
	return tcpLeak(params)
}
