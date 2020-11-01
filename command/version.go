package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/alileza/example/version"
)

var VersionCommand *cli.Command = &cli.Command{
	Name:        "version",
	Description: "Show application version",
	Usage:       "Show application version",
	Flags:       []cli.Flag{},
	Action: func(c *cli.Context) error {
		fmt.Println(version.Print())
		return nil
	},
}
