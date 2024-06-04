package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func DeleteServerManhwa(keys constants.Keys) error {
	smKeys, err := attributevalue.MarshalMap(keys)
	if err != nil {
		logger.Log.Errorf(constants.ErrorMarshalItem, err)
	}

	_, err = dbInstance.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(dbInstance.TableName),
		Key:       smKeys,
	})

	return err
}

func SearchServerManhwas(keys constants.Keys) ([]models.ServerManhwa, error) {
	var serverManhwas []models.ServerManhwa
	var err error
	var response *dynamodb.QueryOutput

	input := &dynamodb.QueryInput{
		TableName: aws.String(dbInstance.TableName),
	}

	keyCon := expression.KeyAnd(expression.Key("PK").Equal(expression.Value(keys.PK)), expression.Key("SK").BeginsWith(keys.SK))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCon).Build()

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

func UpdateServerManhwaCh(keys constants.Keys, titleCh string) (map[string]interface{}, error) {
	var err error
	var response *dynamodb.UpdateItemOutput
	var attribMap map[string]interface{}

	avKeys, err := attributevalue.MarshalMap(keys)
	if err != nil {
		logger.Log.Errorf(constants.ErrorMarshalItem, err)
	}

	update := expression.Set(expression.Name("TitleCh"), expression.Value(titleCh))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		logger.Log.Errorf(constants.ErrorBuildExpression, err)
	} else {
		response, err = dbInstance.db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName:                 aws.String(dbInstance.TableName),
			Key:                       avKeys,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueAllNew,
		})

		if err != nil {
			logger.Log.Errorf(constants.ErrorUpdateItem, err)
		} else {
			err := attributevalue.UnmarshalMap(response.Attributes, &attribMap)
			if err != nil {
				logger.Log.Errorf(constants.ErrorUnmarshalItem, err)
			}
		}
	}

	return attribMap, err
}
