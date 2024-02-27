package adapters

import (
	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/models"
)

type BooleanAdapter struct{}

func (a *BooleanAdapter) GetFeatureFlag(r pb.CreateBooleanFeatureFlagRequest) models.FeatureFlag[bool] {
	environments := a.getEnviroments(r.Enviroments)
	featureFlag := models.NewFeatureFlag[bool](r.Name, r.Active, r.OnActiveValues, environments)
	return *featureFlag
}

func (a *BooleanAdapter) getEnviroments(requestedEnviroment []*pb.Environment) []models.Environment {
	var enviroments []models.Environment
	for _, r := range requestedEnviroment {
		newEnviroment := models.NewEnvironment(r.Name, r.Active)
		enviroments = append(enviroments, *newEnviroment)
	}

	return enviroments
}
