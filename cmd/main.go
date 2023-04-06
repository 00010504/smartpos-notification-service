package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/Invan2/invan_notification_service/config"
	"github.com/Invan2/invan_notification_service/genproto/notification_service"
	"github.com/Invan2/invan_notification_service/pkg/logger"
	"github.com/Invan2/invan_notification_service/services/listeners"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()

	log := logger.New(cfg.LogLevel, cfg.ServiceName)

	server := grpc.NewServer()

	notification_service.RegisterSMSServer(server, listeners.NewSMSService(log, cfg))

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.HttpPort))
	if err != nil {
		log.Error("error while create listener", logger.Error(err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		server.GracefulStop()
	}()

	if err := server.Serve(lis); err != nil {
		log.Error("error while serve", logger.Error(err))
		return
	}
}
