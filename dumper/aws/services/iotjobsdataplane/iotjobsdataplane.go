package iotjobsdataplane

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iotjobsdataplane.NewFromConfig(awscfg)
    
	out(describeJobExecution(client), "iotjobsdataplane", "JobExecution")
    
}

func describeJobExecution(client *iotjobsdataplane.Client) map[string]interface{} {
	input := &iotjobsdataplane.DescribeJobExecutionInput{}
	result, err := client.DescribeJobExecution(context.TODO(), input)
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

