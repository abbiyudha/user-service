package utils

import (
	"CodingTestUser/configs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(config *configs.AppConfig) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(fmt.Sprintf("mongodb://%v:%v,",
		config.Database.Address,
		config.Database.Port,
	))
	client, err := mongo.NewClient(clientOptions)
	var ctx = context.Background()
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("db_project"), nil
}
