package command

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var MigrationNewInputs struct {
	Filename  string
	OutputDir string
}

var MigrationNewCommand *cli.Command = &cli.Command{
	Name:        "new",
	Description: "Create a new database migration files",
	Usage:       "Create a new database migration files",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "filename",
			Destination: &MigrationNewInputs.Filename,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "out",
			Destination: &MigrationNewInputs.OutputDir,
			Required:    true,
		},
	},
	Action: func(c *cli.Context) error {
		filename := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), MigrationNewInputs.Filename)
		migrationPath := strings.ReplaceAll(MigrationNewInputs.OutputDir, "file://", "")

		err := ioutil.WriteFile(fmt.Sprintf("%s/%s.%s.sql", migrationPath, filename, "up"), []byte(""), 0755)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s/%s.%s.sql", migrationPath, filename, "down"), []byte(""), 0755)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("%s/%s.%s.sql", migrationPath, filename, "down"), "created!")
		return nil
	},
}
