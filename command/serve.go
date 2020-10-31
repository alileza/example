package command

import (
	"log"

	"github.com/alileza/example/server"
	"github.com/urfave/cli/v2"
)

var ServeInputs struct {
	ListenAddress string
	SwaggerPath   string
}

var ServeCommand *cli.Command = &cli.Command{
	Name:        "serve",
	Description: "Start serving services",
	Usage:       "Start serving services",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "addr",
			Destination: &ServeInputs.ListenAddress,
			Value:       "0.0.0.0:9000",
		},
		&cli.StringFlag{
			Name:        "swagger-path",
			Destination: &ServeInputs.SwaggerPath,
			Value:       "./autogen/docs/example.swagger.json",
		},
	},
	Action: func(c *cli.Context) error {
		srv := server.Server{
			ListenAddress:       ServeInputs.ListenAddress,
			SwaggerDocsFilePath: ServeInputs.SwaggerPath,
		}

		log.Printf("Start serving on %s", ServeInputs.ListenAddress)
		return srv.Run(c.Context)
	},
}
