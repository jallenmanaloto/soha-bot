package database

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	cfg "github.com/jallenmanaloto/soha-bot/config"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/internal/database/utils"
	"github.com/jallenmanaloto/soha-bot/models"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

type Service struct {
	db        *dynamodb.Client
	TableName string
}

var (
	dbInstance *Service
)

func New() *Service {
	if dbInstance != nil {
		return dbInstance
	}

	table := os.Getenv("TABLENAME")
	ddbClient := dynamodb.NewFromConfig(cfg.Config)
	dbInstance = &Service{
		db:        ddbClient,
		TableName: table,
	}
	return dbInstance
}

func SearchManhwas(exprName string, value string, op string) ([]models.Manhwa, error) {
	var manhwas []models.Manhwa
	var err error
	var filtEx expression.ConditionBuilder
	var response *dynamodb.ScanOutput

	input := &dynamodb.ScanInput{
		TableName: aws.String(dbInstance.TableName),
	}

	filtEx = utils.GenerateFilterExpression(exprName, value, op)
	expr, err := expression.NewBuilder().WithFilter(filtEx).Build()
	if err != nil {
		logger.Log.Errorf(constants.ErrorBuildExpression, err)
		return nil, err
	}
	input.ExpressionAttributeNames = expr.Names()
	input.ExpressionAttributeValues = expr.Values()
	input.FilterExpression = expr.Filter()

	response, err = dbInstance.db.Scan(context.Background(), input)
	if err != nil {
		logger.Log.Errorf(constants.ErrorScan, err)
		return nil, err
	}

	for _, item := range response.Items {
		var manhwa models.Manhwa
		err = attributevalue.UnmarshalMap(item, &manhwa)
		if err != nil {
			logger.Log.Errorf(constants.ErrorUnmarshalItem, err)
			return nil, err
		}
		manhwas = append(manhwas, manhwa)
	}

	return manhwas, nil
}

func SearchSubscribedToManhwa(manhwaId string) ([]models.ServerManhwa, error) {
	var serverManhwas []models.ServerManhwa
	var err error
	var response *dynamodb.ScanOutput

	input := &dynamodb.ScanInput{
		TableName: aws.String(dbInstance.TableName),
	}

	filtEx := expression.And(
		expression.Name("PK").Equal(expression.Value("SERVER")),
		expression.BeginsWith(expression.Name("SK"), "SERVER#"),
		expression.Name("TitleId").Equal(expression.Value(manhwaId)),
	)
	expr, err := expression.NewBuilder().WithFilter(filtEx).Build()
	if err != nil {
		logger.Log.Errorf(constants.ErrorBuildExpression, err)
		return nil, err
	}
	input.ExpressionAttributeValues = expr.Values()
	input.ExpressionAttributeNames = expr.Names()
	input.FilterExpression = expr.Filter()

	response, err = dbInstance.db.Scan(context.Background(), input)
	if err != nil {
		logger.Log.Errorf(constants.ErrorScan, err)
		return nil, err
	}

	for _, item := range response.Items {
		var serverManhwa models.ServerManhwa

		err = attributevalue.UnmarshalMap(item, &serverManhwa)
		if err != nil {
			logger.Log.Errorf(constants.ErrorUnmarshalItem, err)
			return nil, err
		}

		serverManhwas = append(serverManhwas, serverManhwa)
	}

	return serverManhwas, nil
}
