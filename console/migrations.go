package console

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/config"
	"github.com/techx/portal/errors"
)

const (
	migrationFilePath = "./migrations"
)

func CreateMigrationFiles(filename string) error {
	if len(filename) == 0 {
		return errors.New("Migration filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migrationFilePath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migrationFilePath, timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}
	fmt.Printf("Created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	fmt.Printf("Created %s\n", downMigrationFilePath)

	return nil
}

func RunDatabaseMigrations(dbConfig config.DB) error {
	appMigrate := getAppMigrate(dbConfig)
	err := appMigrate.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Migrations successful")
	return nil
}

func RollbackLatestMigration(dbConfig config.DB) error {
	appMigrate := getAppMigrate(dbConfig)
	err := appMigrate.Steps(-1)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Rollback successful")
	return nil
}

func getAppMigrate(dbConfig config.DB) *migrate.Migrate {
	var err error
	appMigrate, err := migrate.New("file://"+migrationFilePath, dbConfig.GetConnectionString())
	if err != nil {
		log.Err(err).Msg("migration failed")
		panic("failed to init migration")
	}

	appMigrate.Log = migrateLogger{}
	return appMigrate
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
