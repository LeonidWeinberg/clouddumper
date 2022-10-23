package redshiftdata

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := redshiftdata.NewFromConfig(awscfg)
    
	out(describeStatement(client), "redshiftdata", "Statement")
    
	out(describeTable(client), "redshiftdata", "Table")
    
}

func describeStatement(client *redshiftdata.Client) map[string]interface{} {
	input := &redshiftdata.DescribeStatementInput{}
	result, err := client.DescribeStatement(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeTable(client *redshiftdata.Client) map[string]interface{} {
	input := &redshiftdata.DescribeTableInput{}
	result, err := client.DescribeTable(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

