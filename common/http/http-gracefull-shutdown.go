package http

import (
	"os"
	"os/signal"
	"syscall"
)

type GracefullHTTPShutdown struct {
	channel chan os.Signal
}

func NewGracefullHTTPShutdown() *GracefullHTTPShutdown {
	chanServer := make(chan os.Signal, 1)
	signal.Notify(chanServer, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	return &GracefullHTTPShutdown{
		channel: chanServer,
	}
}

func (sc *GracefullHTTPShutdown) Wait() {
	defer close(sc.channel)

	<-sc.channel
	signal.Stop(sc.channel)
}
