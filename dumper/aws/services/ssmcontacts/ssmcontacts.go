package ssmcontacts

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ssmcontacts.NewFromConfig(awscfg)
    
	out(describeEngagement(client), "ssmcontacts", "Engagement")
    
	out(describePage(client), "ssmcontacts", "Page")
    
}

func describeEngagement(client *ssmcontacts.Client) map[string]interface{} {
	input := &ssmcontacts.DescribeEngagementInput{}
	result, err := client.DescribeEngagement(context.TODO(), input)
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

func describePage(client *ssmcontacts.Client) map[string]interface{} {
	input := &ssmcontacts.DescribePageInput{}
	result, err := client.DescribePage(context.TODO(), input)
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

