package console

import (
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
)

var _ migrate.Logger = migrateLogger{}

type migrateLogger struct{}

func (ml migrateLogger) Printf(format string, v ...interface{}) {
	format = strings.TrimSuffix(format, "\n")
	if strings.HasPrefix(format, "error: ") {
		log.Error().Msgf(strings.TrimPrefix(format, "error: "), v...)
		return
	}

	log.Info().Msgf(format, v...)
}

func (ml migrateLogger) Verbose() bool { return false }
