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

func CreateServerManhwa(serverManhwa models.ServerManhwa) error {
	item, err := attributevalue.MarshalMap(serverManhwa)
	if err != nil {
		logger.Log.Errorf(constants.ErrorMarshalItem, err)
		return err
	}

	_, err = dbInstance.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(dbInstance.TableName),
		Item:      item,
	})
	if err != nil {
		logger.Log.Errorf(constants.ErrorPutItem, err)
		return err
	}
	return nil
}

func SearchManhwas(exprName string, value string, op string) ([]models.Manhwa, error) {
	var manhwas []models.Manhwa
	var err error
	var filtEx expression.ConditionBuilder
	var response *dynamodb.ScanOutput

	input := &dynamodb.ScanInput{
		TableName: aws.String(dbInstance.TableName),
	}

	// filtEx := expression.Name("Title").Contains(value)
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
