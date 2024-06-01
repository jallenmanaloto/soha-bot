package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
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
		logger.Log.Errorf("ERROR Unable to load AWS SDK config: %v\n", err)
	}

	ddbClient := dynamodb.NewFromConfig(cfg)
	dbInstance = &Service{
		db: ddbClient,
	}
	return dbInstance
}
