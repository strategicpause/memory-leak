package socket

import (
	"github.com/urfave/cli"
	"golang.org/x/sys/unix"
	"time"
)

const (
	NumSocketsName    = "num-sockets"
	SizeName          = "size"
	PauseDurationName = "pause"

	KiB = 1024
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
		cli.IntFlag{
			Name:  SizeName,
			Usage: "Number of KiB to write per per socket.",
			Value: 4,
		},
	}
}

type Params struct {
	NumSockets            int64
	NetworkAddressDomain  int
	ConnectionType        int
	CommunicationProtocol int
	DataSize              int
	PauseTimeInSeconds    time.Duration
}

func action(ctx *cli.Context) error {
	params := &Params{
		NumSockets:            ctx.Int64(NumSocketsName),
		NetworkAddressDomain:  unix.AF_INET,
		ConnectionType:        unix.SOCK_STREAM,
		CommunicationProtocol: unix.IPPROTO_TCP,
		DataSize:              ctx.Int(SizeName) * KiB,
		PauseTimeInSeconds:    ctx.Duration(PauseDurationName),
	}
	return tcpLeak(params)
}
