package controllers

import (
	"context"

	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/handlers"
	"github.com/liperm/ff-manager-server/pkg/logger"
)

type FeatureFlagServer struct {
	pb.UnimplementedFeatureFlagServer
	logger logger.CustomLogger
}

func NewFeatureFlagServer(logger logger.CustomLogger) *FeatureFlagServer {
	return &FeatureFlagServer{logger: logger}
}

func (s *FeatureFlagServer) CreateBooleanFeatureFlag(ctx context.Context, request *pb.CreateBooleanFeatureFlagRequest) (*pb.CreateFeatureFlagResponse, error) {
	result, err := handlers.CreateBooleanFeatureFlag(*request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *FeatureFlagServer) CreateInt64FeatureFlag(ctx context.Context, request *pb.CreateInt64FeatureFlagRequest) (*pb.CreateFeatureFlagResponse, error) {
	result, err := handlers.CreateInt64FeatureFlag(*request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
