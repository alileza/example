package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/alileza/example/command"
)

func main() {
	app := &cli.App{
		Name:  "example",
		Usage: "Example web server",
		Commands: []*cli.Command{
			command.ServeCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stdout, "ERR: %v\n", err)
		os.Exit(1)
	}
}
