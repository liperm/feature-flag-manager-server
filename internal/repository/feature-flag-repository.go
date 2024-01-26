package repository

import (
	"errors"

	"github.com/liperm/ff-manager-server/internal/db"
	"github.com/liperm/ff-manager-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFeatureFlag[T models.FeatureFlagType](featureFlag models.FeatureFlag[T]) (string, error) {
	result, err := db.Collection.InsertOne(db.Ctx, featureFlag)
	if err != nil {
		return "", err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id.Hex(), nil
	}

	return "", errors.New("insert_one_unknown_error")
}
