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

func CreateMigrationFiles(filename string, dbConfig config.DB) error {
	InitPostgres(dbConfig, migrationFilePath)
	return Create(filename)
}

func RunDatabaseMigrations(dbConfig config.DB) error {
	InitPostgres(dbConfig, migrationFilePath)
	return Run()
}

func RollbackLatestMigration(dbConfig config.DB) error {
	InitPostgres(dbConfig, migrationFilePath)
	return RollbackLatest()
}

var (
	appMigrationFilesPath string
	appMigrate            *migrate.Migrate
)

func InitPostgres(dbConfig config.DB, migrationFilesPath string) {
	appMigrationFilesPath = migrationFilesPath

	var err error
	appMigrate, err = migrate.New("file://"+migrationFilesPath, dbConfig.GetConnectionString())
	if err != nil {
		log.Err(err).Msg("migration failed")
		panic("failed to init migration")
	}
}

func Create(filename string) error {
	if len(filename) == 0 {
		return errors.New("Migration filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", appMigrationFilesPath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", appMigrationFilesPath, timeStamp, filename)

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

func Run() error {
	err := appMigrate.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Migrations successful")
	return nil
}

func RollbackLatest() error {
	err := appMigrate.Steps(-1)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Rollback successful")
	return nil
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
