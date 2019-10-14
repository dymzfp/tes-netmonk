package db 

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_HOST  =  "mongodb://user:pass@54.255.163.114:27018"
	DB_NAME  =  "db_assessment"
)

func Connect() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(DB_HOST)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	return client.Database(DB_NAME), nil
}