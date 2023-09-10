package tcp

import (
	"github.com/urfave/cli"
)

const Port = ":8080"

func Register() cli.Command {
	return cli.Command{
		Name:   "tcp",
		Usage:  "Reproduces a TCP socket leak.",
		Action: action,
		Flags:  flags(),
	}
}

func flags() []cli.Flag {
	return []cli.Flag{}
}

func action(_ *cli.Context) error {
	return tcpLeak()
}
