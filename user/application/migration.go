package application

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/syauqeesy/accounting-service/user/configuration"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type migrationApplication struct {
	configuration         *configuration.Configuration
	migrator              *migrate.Migrate
	databaseConnectionUrl string
	commandType           string
	commandArgument       string
	migrationPath         string
}

const (
	MigrationGenerate = "generate"
	MigrationUp       = "up"
	MigrationDown     = "down"
)

func (a *migrationApplication) Init() (err error) {

	a.databaseConnectionUrl = fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s", a.configuration.Database.User, a.configuration.Database.Password, a.configuration.Database.Host, a.configuration.Database.Port, a.configuration.Database.Name)

	a.migrationPath, err = filepath.Abs("migration")
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	if a.commandType != MigrationUp && a.commandType != MigrationDown {
		return nil
	}

	fmt.Println(a.migrationPath)
	m, err := migrate.New("file://"+a.migrationPath, a.databaseConnectionUrl)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	a.migrator = m

	return nil
}

func (a *migrationApplication) Run() error {
	switch a.commandType {
	case MigrationGenerate:
		err := a.Generate(a.commandArgument)
		if err != nil {
			return err
		}

		return nil
	case MigrationUp:
		err := a.migrator.Up()
		if err != nil {
			return err
		}

		return nil
	case MigrationDown:
		err := a.migrator.Down()
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (a *migrationApplication) Close() error {
	if a.migrator == nil {
		return nil
	}

	sourceErr, databaseErr := a.migrator.Close()
	if sourceErr != nil {
		return sourceErr
	}

	if databaseErr != nil {
		return databaseErr
	}

	return nil
}

func (a *migrationApplication) Generate(name string) error {
	timestamp := time.Now().Unix()

	up := filepath.Join(a.migrationPath, fmt.Sprintf("%d_%s.up.sql", timestamp, name))
	down := filepath.Join(a.migrationPath, fmt.Sprintf("%d_%s.down.sql", timestamp, name))

	err := os.MkdirAll(a.migrationPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}

	for _, path := range []string{up, down} {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err)
		}

		file.Close()
	}

	fmt.Println(" -", up)
	fmt.Println(" -", down)

	return nil
}
