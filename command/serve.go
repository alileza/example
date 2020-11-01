package command

import (
	"github.com/alileza/example/server"
	"github.com/sirupsen/logrus"
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
		logger := logrus.New()

		srv := server.Server{
			Logger:              logger,
			ListenAddress:       ServeInputs.ListenAddress,
			SwaggerDocsFilePath: ServeInputs.SwaggerPath,
		}

		return srv.Run(c.Context)
	},
}
