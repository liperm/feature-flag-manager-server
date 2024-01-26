package handlers

import (
	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/adapters"
	"github.com/liperm/ff-manager-server/internal/repository"
)

func CreateBooleanFeatureFlag(request pb.CreateBooleanFeatureFlagRequest) (*pb.CreateBooleanFeatureFlagResponse, error) {
	adapter := adapters.BooleanAdapter{}

	featureFlag := adapter.GetFeatureFlag(request)
	result, err := repository.CreateFeatureFlag[bool](featureFlag)

	return &pb.CreateBooleanFeatureFlagResponse{Id: result}, err
}
