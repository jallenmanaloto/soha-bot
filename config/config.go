package config

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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
