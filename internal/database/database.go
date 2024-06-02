package database

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	cfg "github.com/jallenmanaloto/soha-bot/config"
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
