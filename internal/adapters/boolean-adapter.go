package adapters

import (
	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/models"
)

type BooleanAdapter struct{}

func (a *BooleanAdapter) GetFeatureFlag(r pb.CreateBooleanFeatureFlagRequest) models.FeatureFlag[bool] {
	enviroments := a.getEnviroments(r.Enviroments)
	featureFlag := models.NewFeatureFlag[bool](r.Name, r.Active, enviroments)
	return *featureFlag
}

func (a *BooleanAdapter) getEnviroments(requestedEnviroment []*pb.BooleanEnviroment) []models.Environment[bool] {
	var enviroments []models.Environment[bool]
	for _, r := range requestedEnviroment {
		newEnviroment := models.NewEnvironment[bool](r.Name, r.Active, r.OnActiveValues)
		enviroments = append(enviroments, *newEnviroment)
	}

	return enviroments
}
