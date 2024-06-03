package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/config"
	_ "github.com/techx/portal/docs"
	"github.com/techx/portal/version"
	"github.com/urfave/cli/v2"
)

func main() {
	commands := []*cli.Command{
		{
			Name:  "start",
			Usage: "starts the API server",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config-file",
					Aliases: []string{"c"},
					Usage:   "YAML config file",
					Value:   "",
				},
			},
			Action: startAPIServer,
		},
		{
			Name:  "generate-config",
			Usage: "dumps application's default configuration into a YAML file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config-file",
					Aliases: []string{"c"},
					Usage:   "YAML config file",
					Value:   "application.yml",
				},
			},
			Action: config.GenerateDefaultsFile,
		},
	}

	app := &cli.App{
		Name:     "portal",
		Version:  version.Version,
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Err(err).Msg("")
		os.Exit(1)
	}
}
