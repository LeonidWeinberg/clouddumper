package iotanalytics

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iotanalytics.NewFromConfig(awscfg)
    
	out(describeChannel(client), "iotanalytics", "Channel")
    
	out(describeDataset(client), "iotanalytics", "Dataset")
    
	out(describeDatastore(client), "iotanalytics", "Datastore")
    
	out(describeLoggingOptions(client), "iotanalytics", "LoggingOptions")
    
	out(describePipeline(client), "iotanalytics", "Pipeline")
    
}

func describeChannel(client *iotanalytics.Client) map[string]interface{} {
	input := &iotanalytics.DescribeChannelInput{}
	result, err := client.DescribeChannel(context.TODO(), input)
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

func describeDataset(client *iotanalytics.Client) map[string]interface{} {
	input := &iotanalytics.DescribeDatasetInput{}
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

func describeDatastore(client *iotanalytics.Client) map[string]interface{} {
	input := &iotanalytics.DescribeDatastoreInput{}
	result, err := client.DescribeDatastore(context.TODO(), input)
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

func describeLoggingOptions(client *iotanalytics.Client) map[string]interface{} {
	input := &iotanalytics.DescribeLoggingOptionsInput{}
	result, err := client.DescribeLoggingOptions(context.TODO(), input)
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

func describePipeline(client *iotanalytics.Client) map[string]interface{} {
	input := &iotanalytics.DescribePipelineInput{}
	result, err := client.DescribePipeline(context.TODO(), input)
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

