package repository

import (
	"context"
	"errors"

	"github.com/liperm/ff-manager-server/internal/db"
	"github.com/liperm/ff-manager-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PersistFeatureFlag[T models.FeatureFlagType](featureFlag models.FeatureFlag[T]) (string, error) {
	result, err := db.Collection.InsertOne(db.Ctx, featureFlag)
	if err != nil {
		return "", err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id.Hex(), nil
	}

	return "", errors.New("insert_one_unknown_error")
}

func GetFeatureFlagByName[T models.FeatureFlagType](name string) (models.FeatureFlag[T], error) {
	filter := bson.D{{"name", bson.D{{"$eq", name}}}}
	opts := options.FindOne()

	var featureFlag models.FeatureFlag[T]
	err := db.Collection.FindOne(context.TODO(), filter, opts).Decode(&featureFlag)
	if err != nil {
		return featureFlag, err
	}

	return featureFlag, nil
}
