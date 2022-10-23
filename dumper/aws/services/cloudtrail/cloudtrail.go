package cloudtrail

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudtrail.NewFromConfig(awscfg)
    
	out(describeQuery(client), "cloudtrail", "Query")
    
	out(describeTrails(client), "cloudtrail", "Trails")
    
}

func describeQuery(client *cloudtrail.Client) map[string]interface{} {
	input := &cloudtrail.DescribeQueryInput{}
	result, err := client.DescribeQuery(context.TODO(), input)
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

func describeTrails(client *cloudtrail.Client) map[string]interface{} {
	input := &cloudtrail.DescribeTrailsInput{}
	result, err := client.DescribeTrails(context.TODO(), input)
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

