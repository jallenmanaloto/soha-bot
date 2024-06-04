package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/internal/database/utils"
	"github.com/jallenmanaloto/soha-bot/models"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

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

func SearchServerManhwasByTitle(keys constants.Keys, exprName string, value string, op string) ([]models.ServerManhwa, error) {
	var serverManhwas []models.ServerManhwa
	var err error
	var filtCon expression.ConditionBuilder
	var response *dynamodb.QueryOutput

	input := &dynamodb.QueryInput{
		TableName: aws.String(dbInstance.TableName),
	}

	keyCon := expression.KeyAnd(expression.Key("PK").Equal(expression.Value(keys.PK)), expression.Key("SK").Equal(expression.Value(keys.SK)))
	filtCon = utils.GenerateFilterExpression(exprName, value, op)

	expr, err := expression.NewBuilder().WithKeyCondition(keyCon).WithFilter(filtCon).Build()

	if err != nil {
		logger.Log.Errorf(constants.ErrorBuildExpression, err)
		return nil, err
	}
	input.ExpressionAttributeNames = expr.Names()
	input.ExpressionAttributeValues = expr.Values()
	input.FilterExpression = expr.Filter()
	input.KeyConditionExpression = expr.KeyCondition()

	response, err = dbInstance.db.Query(context.TODO(), input)
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
