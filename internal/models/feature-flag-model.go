package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagType interface{ bool | int64 }

type Environment[T FeatureFlagType] struct {
	Id             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Active         bool               `bson:"active"`
	OnActiveValues []T                `bson:"on_active_values"`
}

type FeatureFlag[T FeatureFlagType] struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Enviroments []Environment[T]   `bson:"enviroments"`
	Active      bool               `bson:"active"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func NewEnvironment[T FeatureFlagType](name string, active bool, values []T) *Environment[T] {
	return &Environment[T]{
		Id:             primitive.NewObjectID(),
		Name:           name,
		Active:         active,
		OnActiveValues: values,
	}
}

func NewFeatureFlag[T FeatureFlagType](name string, active bool, envs []Environment[T]) *FeatureFlag[T] {
	return &FeatureFlag[T]{
		Id:          primitive.NewObjectID(),
		Name:        name,
		Active:      active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Enviroments: envs,
	}
}
