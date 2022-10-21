package main

import (
	"context"

	"github.com/CheckmarxDev/dumper/internals/resources/ec2"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	awscfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	ec2.DescribeInstances(awscfg)
}
