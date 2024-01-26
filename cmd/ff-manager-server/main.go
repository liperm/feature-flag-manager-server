package main

import (
	"log"
	"net"

	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/controllers"
	"github.com/liperm/ff-manager-server/internal/db"
	"google.golang.org/grpc"
)

func main() {
	db.Init()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterFeatureFlagServer(s, &controllers.FeatureFlagServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
