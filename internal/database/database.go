package database

import (
	// "context"

	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	cfg "github.com/jallenmanaloto/soha-bot/config"
	// "github.com/jallenmanaloto/soha-bot/pkg/logger"
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

	ddbClient := dynamodb.NewFromConfig(cfg.Config)
	dbInstance = &Service{
		db: ddbClient,
	}
	return dbInstance
}
