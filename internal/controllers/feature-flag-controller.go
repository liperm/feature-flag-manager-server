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

func (s *FeatureFlagServer) CreateBooleanFeatureFlag(ctx context.Context, request *pb.CreateBooleanFeatureFlagRequest) (*pb.CreateBooleanFeatureFlagResponse, error) {
	s.logger.Request(request)

	result, err := handlers.CreateBooleanFeatureFlag(*request)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	s.logger.Response(result)
	return result, nil
}
