package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/syauqeesy/accounting-service/common"
	common_http "github.com/syauqeesy/accounting-service/common/http"
	"github.com/syauqeesy/accounting-service/invoice/configuration"
	"github.com/syauqeesy/accounting-service/invoice/handler"
	grpc_outbound "github.com/syauqeesy/accounting-service/invoice/outbound/grpc"
	"github.com/syauqeesy/accounting-service/invoice/repository"
	"github.com/syauqeesy/accounting-service/invoice/service"
)

type httpApplication struct {
	configuration          *configuration.Configuration
	mux                    *http.ServeMux
	server                 *http.Server
	httpSignal             *common.GracefullShutdown
	database               *databaseApplication
	repository             *repository.Repository
	service                *service.Service
	handler                *handler.Handler
	grpcOutboundService    *grpc_outbound.GRPCOutboundService
	grpcOutboundConnection *grpc_outbound.GRPCOutboundConnection
}

func (a *httpApplication) Init() error {

	a.mux = http.NewServeMux()

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

	a.grpcOutboundConnection, a.grpcOutboundService = grpc_outbound.New(a.configuration)

	a.service = service.New(a.configuration, a.repository, a.grpcOutboundService)

	a.handler = handler.New(a.mux, a.configuration, a.service)

	a.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		common_http.WriteHttpResponse(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), nil)
	})

	a.server = &http.Server{
		Addr:    a.configuration.HTTP.Port,
		Handler: a.mux,
	}

	a.httpSignal = common.NewGracefullShutdown()

	return nil
}

func (a *httpApplication) Run() error {
	go func() {
		a.httpSignal.Wait()

		a.Close()
	}()

	fmt.Println("http server started")

	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (a *httpApplication) Close() error {
	a.database.Close()

	fmt.Println("shutting down http server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	fmt.Println("http server exited")

	return nil
}
