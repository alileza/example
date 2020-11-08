package command

import (
	"github.com/urfave/cli/v2"
)

var MigrationCommand *cli.Command = &cli.Command{
	Name:        "migration",
	Description: "Database migration command",
	Usage:       "Database migration command",
	Subcommands: []*cli.Command{
		MigrationRunCommand,
		MigrationNewCommand,
	},
	Flags: MigrationRunCommand.Flags,
}
