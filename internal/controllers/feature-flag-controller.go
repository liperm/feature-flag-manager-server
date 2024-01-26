package controllers

import (
	"context"

	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/handlers"
)

type FeatureFlagServer struct {
	pb.UnimplementedFeatureFlagServer
}

func (s *FeatureFlagServer) CreateBooleanFeatureFlag(ctx context.Context, request *pb.CreateBooleanFeatureFlagRequest) (*pb.CreateBooleanFeatureFlagResponse, error) {
	return handlers.CreateBooleanFeatureFlag(*request)
}
