package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	common_http "github.com/syauqeesy/accounting-service/common/http"
	"github.com/syauqeesy/accounting-service/user/configuration"
	"github.com/syauqeesy/accounting-service/user/handler"
	"github.com/syauqeesy/accounting-service/user/repository"
	"github.com/syauqeesy/accounting-service/user/service"
)

type httpApplication struct {
	configuration *configuration.Configuration
	mux           *http.ServeMux
	server        *http.Server
	httpSignal    *common_http.GracefullHTTPShutdown
	database      *databaseApplication
	repository    *repository.Repository
	service       *service.Service
	handler       *handler.Handler
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

	a.service = service.New(a.configuration, a.repository)

	a.handler = handler.New(a.mux, a.configuration, a.service)

	a.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		common_http.WriteHttpResponse(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), nil)
	})

	a.server = &http.Server{
		Addr:    a.configuration.HTTP.Port,
		Handler: a.mux,
	}

	a.httpSignal = common_http.NewGracefullHTTPShutdown()

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
