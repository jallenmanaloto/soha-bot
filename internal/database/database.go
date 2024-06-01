package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Service struct {
	db *dynamodb.Client
}

var (
	dbInstance *Service
)

func New() *Service {
	if dbInstance != nil {
		return dbInstance
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-southeast-1"),
	)
	if err != nil {
		log.Fatalf("Unable to load AWS sdk config: %v", err)
	}

	ddbClient := dynamodb.NewFromConfig(cfg)
	dbInstance = &Service{
		db: ddbClient,
	}
	return dbInstance
}
