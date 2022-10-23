package cognitosync

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cognitosync.NewFromConfig(awscfg)
    
	out(describeDataset(client), "cognitosync", "Dataset")
    
	out(describeIdentityPoolUsage(client), "cognitosync", "IdentityPoolUsage")
    
	out(describeIdentityUsage(client), "cognitosync", "IdentityUsage")
    
}

func describeDataset(client *cognitosync.Client) map[string]interface{} {
	input := &cognitosync.DescribeDatasetInput{}
	result, err := client.DescribeDataset(context.TODO(), input)
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

func describeIdentityPoolUsage(client *cognitosync.Client) map[string]interface{} {
	input := &cognitosync.DescribeIdentityPoolUsageInput{}
	result, err := client.DescribeIdentityPoolUsage(context.TODO(), input)
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

func describeIdentityUsage(client *cognitosync.Client) map[string]interface{} {
	input := &cognitosync.DescribeIdentityUsageInput{}
	result, err := client.DescribeIdentityUsage(context.TODO(), input)
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

