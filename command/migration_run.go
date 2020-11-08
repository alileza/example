package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var MigrationRunInputs struct {
	Path       string
	Datasource string
	Action     string
}

var MigrationRunCommand *cli.Command = &cli.Command{
	Name:        "run",
	Description: "Database migration command",
	Usage:       "Database migration command",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Destination: &MigrationRunInputs.Path,
			Value:       "file://" + os.Getenv("PWD") + "/migrations",
		},
		&cli.StringFlag{
			Name:        "datasource",
			Destination: &MigrationRunInputs.Datasource,
			Value:       "postgres://example:example@localhost:5432/example?sslmode=disable",
		},
		&cli.StringFlag{
			Name:        "action",
			Destination: &MigrationRunInputs.Action,
		},
	},
	Action: func(c *cli.Context) error {
		m, err := migrate.New(MigrationRunInputs.Path, MigrationRunInputs.Datasource)
		if err != nil {
			return fmt.Errorf("Failed to initiate database migration.\n > input=%+v\n > error=%w", MigrationRunInputs, err)
		}

		switch MigrationRunInputs.Action {
		case "up":
			if err := m.Up(); err != migrate.ErrNoChange {
				return err
			}
			return nil
		case "down":
			if err := m.Steps(-1); err != nil {
				return fmt.Errorf("Failed execute step down database migration: %w", err)
			}
			return nil
		case "reset":
			if err := m.Down(); err != nil {
				return fmt.Errorf("Failed execute reset database migration: %w", err)
			}
			return nil
		}
		return errors.New("unexpected migrate action (--action=up|down|reset)")
	},
}
