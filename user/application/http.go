package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	common "github.com/syauqeesy/accounting-service/common/gracefull-http-shutdown"
	"github.com/syauqeesy/accounting-service/configuration"
	"github.com/syauqeesy/accounting-service/handler"
	"github.com/syauqeesy/accounting-service/service"
)

type httpApplication struct {
	configuration *configuration.Configuration
	mux           *http.ServeMux
	server        *http.Server
	httpSignal    *common.GracefullHTTPShutdown
	service       *service.Service
	handler       *handler.Handler
}

func (a *httpApplication) Init() error {
	a.service = service.New(a.configuration)

	a.mux = http.NewServeMux()

	a.handler = handler.New(a.mux, a.configuration, a.service)

	a.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user service")
	})

	a.server = &http.Server{
		Addr:    a.configuration.HTTP.Port,
		Handler: a.mux,
	}

	a.httpSignal = common.NewGracefullHTTPShutdown()

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
	fmt.Println("shutting down http server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	fmt.Println("http server exited")

	return nil
}
