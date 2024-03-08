package main

import (
	"log"
	"net"

	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/db"
	"github.com/liperm/ff-manager-server/internal/interceptors"
	controllers "github.com/liperm/ff-manager-server/internal/servers"
	"github.com/liperm/ff-manager-server/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	db.Init()
	logger.Init()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggerInterceptor))
	ffServer := controllers.NewFeatureFlagCreationServer(logger.Logger)
	pb.RegisterFeatureFlagCreationServer(s, ffServer)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
