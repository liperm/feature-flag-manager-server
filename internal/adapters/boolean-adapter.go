package adapters

import (
	"github.com/liperm/ff-manager-server/api/pb"
	"github.com/liperm/ff-manager-server/internal/models"
)

type BooleanAdapter struct {
	adapter *BooleanAdapter
}

func (a *BooleanAdapter) GetFeatureFlag(r pb.CreateBooleanFeatureFlagRequest) models.FeatureFlag[bool] {
	enviroments := a.getEnviroments(r.Enviroments)
	featureFlag := models.NewFeatureFlag[bool](r.Name, r.Active, enviroments)
	return *featureFlag
}

func (a *BooleanAdapter) getEnviroments(requestedEnviroment []*pb.BooleanEnviroment) []models.Enviroment[bool] {
	var enviroments []models.Enviroment[bool]
	for _, r := range requestedEnviroment {
		newEnviroment := models.NewEnviroment[bool](r.Name, r.Active, r.OnActiveValues)
		enviroments = append(enviroments, *newEnviroment)
	}

	return enviroments
}
