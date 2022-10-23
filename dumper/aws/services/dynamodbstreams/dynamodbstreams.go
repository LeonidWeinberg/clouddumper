package dynamodbstreams

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := dynamodbstreams.NewFromConfig(awscfg)
    
	out(describeStream(client), "dynamodbstreams", "Stream")
    
}

func describeStream(client *dynamodbstreams.Client) map[string]interface{} {
	input := &dynamodbstreams.DescribeStreamInput{}
	result, err := client.DescribeStream(context.TODO(), input)
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

