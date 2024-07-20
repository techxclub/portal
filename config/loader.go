package config

import (
	"context"
	"os"

	"github.com/sonnes/go-envconfig"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// LoadConfig reads the YAML and env configurations to populate `cfg`.
// env vars overwrite YAML config values.
// Pass empty string to `path`, if there's no config file.
func LoadConfig(path string, cfg interface{}) error {
	// first priority is a YAML config file
	if path != "" {
		ymlData, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(ymlData, cfg)
		if err != nil {
			return err
		}
	}

	// then the environment variables override any set values
	ctx := context.Background()
	return envconfig.Load(ctx, cfg)
}

// GenerateDefaultsFile dumps application's default configuration into a YAML file
func GenerateDefaultsFile(ctx *cli.Context) error {
	path := ctx.String("config-file")

	defaultCfg := &Config{}
	defaultCfg.SetDefaults()
	return saveConfig(path, defaultCfg)
}

func saveConfig(path string, cfg interface{}) error {
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, d, 0o600)
	if err != nil {
		return err
	}

	return nil
}
