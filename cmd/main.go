package main

import (
	"os"
	"net"
	"syscall"
	"os/signal"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"github.com/johnchuks/events-service/pb"
	"github.com/johnchuks/events-service/transports"
	"github.com/johnchuks/events-service/service"
	"github.com/johnchuks/events-service/endpoints"
  )

func main() {
	eventService := service.NewService(os.Getenv("DATABASE_HOST"),"5432", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"),os.Getenv("DATABASE_NAME"))
	eventEndpoint := endpoints.MakeEndpoints(eventService)
	grpcServer := transport.NewGRPCServer(eventEndpoint)

	errs := make(chan error)
    go func() {
        c := make(chan os.Signal)
        signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
        errs <- fmt.Errorf("%s", <-c)
    }()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
        log.Error("An Error occcurred", err)
        os.Exit(1)
    }
	server := grpc.NewServer()
	pb.RegisterEventServiceServer(server, grpcServer)
	log.Info("gRPC Server started successfully 🚀")
	server.Serve(grpcListener)

	log.Error(<-errs)

}