package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/techx/portal/config"
	"github.com/techx/portal/console"
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
		{
			Name:        "migrate:run",
			Description: "Running Migration",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config-file",
					Aliases: []string{"c"},
					Usage:   "YAML config file",
					Value:   "",
				},
			},
			Action: func(ctx *cli.Context) error {
				applicationContext, err := initApplicationContext(ctx)
				if err != nil {
					return err
				}
				return console.RunDatabaseMigrations(applicationContext.Config.DB)
			},
		},
		{
			Name:        "migrate:rollback",
			Description: "Rollback Migration",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config-file",
					Aliases: []string{"c"},
					Usage:   "YAML config file",
					Value:   "",
				},
			},
			Action: func(ctx *cli.Context) error {
				applicationContext, err := initApplicationContext(ctx)
				if err != nil {
					return err
				}
				return console.RollbackLatestMigration(applicationContext.Config.DB)
			},
		},
		{
			Name:        "migrate:create",
			Description: "Create up and down migration files with timestamp",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config-file",
					Aliases: []string{"c"},
					Usage:   "YAML config file",
					Value:   "",
				},
				&cli.StringFlag{
					Name:    "migration-name",
					Aliases: []string{"f"},
					Usage:   "migration file name",
					Value:   "",
				},
			},
			Action: func(ctx *cli.Context) error {
				applicationContext, err := initApplicationContext(ctx)
				if err != nil {
					return err
				}

				migrationFileName := ctx.String("migration-name")
				return console.CreateMigrationFiles(migrationFileName, applicationContext.Config.DB)
			},
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
