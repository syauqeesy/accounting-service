package application

import (
	"errors"
	"fmt"

	"github.com/syauqeesy/accounting-service/configuration"
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

func Run(applicationType string) error {
	application, err := New(applicationType)
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

func New(applicationType string) (Application, error) {
	configuration, err := configuration.Load("./config.json")
	if err != nil {
		return nil, err
	}

	switch applicationType {
	case ApplicationHTTP:
		return &httpApplication{
			configuration: configuration,
		}, nil
	default:
		return nil, errors.New("invalid application type")
	}
}
