package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagType interface{ bool | int64 }

type Environment struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Active bool               `bson:"active"`
}

type FeatureFlag[T FeatureFlagType] struct {
	Id             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Enviroments    []Environment      `bson:"enviroments"`
	OnActiveValues []T                `bson:"on_active_values"`
	Active         bool               `bson:"active"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}

func NewEnvironment(name string, active bool) *Environment {
	return &Environment{
		Id:     primitive.NewObjectID(),
		Name:   name,
		Active: active,
	}
}

func NewFeatureFlag[T FeatureFlagType](name string, active bool, values []T, envs []Environment) *FeatureFlag[T] {
	return &FeatureFlag[T]{
		Id:             primitive.NewObjectID(),
		Name:           name,
		Active:         active,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Enviroments:    envs,
		OnActiveValues: values,
	}
}
