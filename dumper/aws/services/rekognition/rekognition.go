package rekognition

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := rekognition.NewFromConfig(awscfg)
    
	out(describeCollection(client), "rekognition", "Collection")
    
	out(describeDataset(client), "rekognition", "Dataset")
    
	out(describeProjectVersions(client), "rekognition", "ProjectVersions")
    
	out(describeProjects(client), "rekognition", "Projects")
    
	out(describeStreamProcessor(client), "rekognition", "StreamProcessor")
    
}

func describeCollection(client *rekognition.Client) map[string]interface{} {
	input := &rekognition.DescribeCollectionInput{}
	result, err := client.DescribeCollection(context.TODO(), input)
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

func describeDataset(client *rekognition.Client) map[string]interface{} {
	input := &rekognition.DescribeDatasetInput{}
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

func describeProjectVersions(client *rekognition.Client) map[string]interface{} {
	input := &rekognition.DescribeProjectVersionsInput{}
	result, err := client.DescribeProjectVersions(context.TODO(), input)
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

func describeProjects(client *rekognition.Client) map[string]interface{} {
	input := &rekognition.DescribeProjectsInput{}
	result, err := client.DescribeProjects(context.TODO(), input)
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

func describeStreamProcessor(client *rekognition.Client) map[string]interface{} {
	input := &rekognition.DescribeStreamProcessorInput{}
	result, err := client.DescribeStreamProcessor(context.TODO(), input)
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

