package kinesisanalyticsv2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := kinesisanalyticsv2.NewFromConfig(awscfg)
    
	out(describeApplication(client), "kinesisanalyticsv2", "Application")
    
	out(describeApplicationSnapshot(client), "kinesisanalyticsv2", "ApplicationSnapshot")
    
	out(describeApplicationVersion(client), "kinesisanalyticsv2", "ApplicationVersion")
    
}

func describeApplication(client *kinesisanalyticsv2.Client) map[string]interface{} {
	input := &kinesisanalyticsv2.DescribeApplicationInput{}
	result, err := client.DescribeApplication(context.TODO(), input)
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

func describeApplicationSnapshot(client *kinesisanalyticsv2.Client) map[string]interface{} {
	input := &kinesisanalyticsv2.DescribeApplicationSnapshotInput{}
	result, err := client.DescribeApplicationSnapshot(context.TODO(), input)
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

func describeApplicationVersion(client *kinesisanalyticsv2.Client) map[string]interface{} {
	input := &kinesisanalyticsv2.DescribeApplicationVersionInput{}
	result, err := client.DescribeApplicationVersion(context.TODO(), input)
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

