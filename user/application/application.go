package application

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/syauqeesy/accounting-service/user/configuration"
)

type Application interface {
	Init() (err error)
	Run() (err error)
	Close() (err error)
}

const (
	ApplicationHTTP      = "http"
	ApplicationGRPC      = "grpc"
	ApplicationMigration = "migration"
	ApplicationSeeder    = "seeder"
)

func Run(applicationType string, arguments []string) error {
	application, err := New(applicationType, arguments)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer application.Close()

	err = application.Init()
	if err != nil {
		return err
	}

	err = application.Run()
	if err != nil {
		return err
	}

	return nil
}

func New(applicationType string, arguments []string) (Application, error) {
	path, err := filepath.Abs("./config.json")
	if err != nil {
		return nil, err
	}

	configuration, err := configuration.Load(path)
	if err != nil {
		return nil, err
	}

	switch applicationType {
	case ApplicationHTTP:
		return &httpApplication{
			configuration: configuration,
		}, nil
	case ApplicationMigration:
		application := &migrationApplication{
			configuration: configuration,
			commandType:   arguments[0],
		}

		if len(arguments) > 1 {
			application.commandArgument = arguments[1]
		}

		return application, nil
	default:
		return nil, errors.New("invalid application type")
	}
}
