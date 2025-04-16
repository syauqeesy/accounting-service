package application

import (
	"fmt"
	"log"
	"net"

	"github.com/syauqeesy/accounting-service/common"
	"github.com/syauqeesy/accounting-service/user/configuration"
	"github.com/syauqeesy/accounting-service/user/repository"
	grpc_service "github.com/syauqeesy/accounting-service/user/service/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcApplication struct {
	configuration *configuration.Configuration
	repository    *repository.Repository
	database      *databaseApplication
	tcpSignal     *common.GracefullShutdown
	server        *grpc.Server
}

func (a *grpcApplication) Init() error {
	a.server = grpc.NewServer()

	reflection.Register(a.server)

	a.database = &databaseApplication{
		configuration: a.configuration,
	}

	err := a.database.Init()
	if err != nil {
		return err
	}

	err = a.database.Run()
	if err != nil {
		return err
	}

	a.repository = repository.New(a.database.database)

	grpc_service.New(a.server, a.configuration, a.repository)

	a.tcpSignal = common.NewGracefullShutdown()

	return nil
}

func (a *grpcApplication) Run() error {
	go func() {
		a.tcpSignal.Wait()

		log.Println("grpc server exited")

		a.Close()
	}()

	fmt.Println("grpc server started")

	tcpServer, err := net.Listen("tcp", a.configuration.GRPC.Port)
	if err != nil {
		return err
	}

	err = a.server.Serve(tcpServer)
	if err != nil {
		return err
	}

	return nil
}

func (a *grpcApplication) Close() error {
	a.server.GracefulStop()

	return nil
}
