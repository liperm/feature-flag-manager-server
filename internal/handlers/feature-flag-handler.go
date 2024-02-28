package handlers

import (
	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/models"
	"github.com/liperm/ff-manager-server/internal/repository"
)

func CreateBooleanFeatureFlag(request pb.CreateBooleanFeatureFlagRequest) (*pb.CreateBooleanFeatureFlagResponse, error) {
	featureFlag := newGenericFeatureFlag[bool](request.Name, true, request.OnActiveValues, request.Enviroments)

	result, err := repository.PersistFeatureFlag[bool](*featureFlag)

	return &pb.CreateBooleanFeatureFlagResponse{Id: result}, err
}

func newGenericFeatureFlag[T models.FeatureFlagType](ffName string, active bool, values []T, environmentRequest []*pb.Environment) *models.FeatureFlag[T] {
	envs := newEnvironments(environmentRequest)

	featureFlag := models.NewFeatureFlag[T](ffName, active, values, envs)

	return featureFlag
}

func newEnvironments(requestedEnviroment []*pb.Environment) []models.Environment {
	var environments []models.Environment
	for _, r := range requestedEnviroment {
		newEnviroment := models.NewEnvironment(r.Name, r.Active)
		environments = append(environments, *newEnviroment)
	}

	return environments
}
