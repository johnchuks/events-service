package main

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"github.com/johnchuks/events-service/pb"
	"github.com/johnchuks/events-service/transports"
  )

func main() {
	grpcServer := transport.NewGRPCServer()
	go func() {
        baseServer := grpc.NewServer()
        pb.RegisterEventServiceServer(baseServer, grpcServer)
        log.Info("msg", "Server started successfully 🚀")
        baseServer.Serve(grpcListener)
    }()

}