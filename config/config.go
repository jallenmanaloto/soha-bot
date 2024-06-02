package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var Config aws.Config

func NewClientConfig() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-southeast-1"),
	)
	if err != nil {
		fmt.Printf("ERROR Unable to load AWS SDK config: %v\n", err)
	}
	Config = cfg
}

func LoadEnvironmentVariables() {
	parameters := []string{"PORT", "LOKI_LOG_URL", "LOKI_USERNAME", "LOKI_PASSWORD"}
	svc := ssm.NewFromConfig(Config)

	for _, parameter := range parameters {
		paramOut, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
			Name:           aws.String(parameter),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			fmt.Printf("Unable to get parameters: %v", err)
		}
		os.Setenv(parameter, *paramOut.Parameter.Value)
	}
}
