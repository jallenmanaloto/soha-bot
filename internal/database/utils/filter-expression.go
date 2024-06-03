package utils

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
)

func GenerateFilterExpression(exprName string, exprVal string, op string) expression.ConditionBuilder {
	var filtEx expression.ConditionBuilder

	switch op {
	case constants.CONTAINS:
		filtEx = expression.Name(exprName).Contains(exprVal)
	case constants.EQUALTO:
		filtEx = expression.Name(exprName).Equal(expression.Value(exprVal))
	}

	return filtEx
}
