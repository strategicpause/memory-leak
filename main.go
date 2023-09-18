package main

import (
	"github.com/strategicpause/memory-leak/command/memory"
	"github.com/strategicpause/memory-leak/command/socket"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: RegisterCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func RegisterCommands() cli.Commands {
	return cli.Commands{
		memory.Register(),
		socket.Register(),
	}
}
