package lookoutvision

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lookoutvision"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := lookoutvision.NewFromConfig(awscfg)
    
	out(describeDataset(client), "lookoutvision", "Dataset")
    
	out(describeModel(client), "lookoutvision", "Model")
    
	out(describeModelPackagingJob(client), "lookoutvision", "ModelPackagingJob")
    
	out(describeProject(client), "lookoutvision", "Project")
    
}

func describeDataset(client *lookoutvision.Client) map[string]interface{} {
	input := &lookoutvision.DescribeDatasetInput{}
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

func describeModel(client *lookoutvision.Client) map[string]interface{} {
	input := &lookoutvision.DescribeModelInput{}
	result, err := client.DescribeModel(context.TODO(), input)
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

func describeModelPackagingJob(client *lookoutvision.Client) map[string]interface{} {
	input := &lookoutvision.DescribeModelPackagingJobInput{}
	result, err := client.DescribeModelPackagingJob(context.TODO(), input)
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

func describeProject(client *lookoutvision.Client) map[string]interface{} {
	input := &lookoutvision.DescribeProjectInput{}
	result, err := client.DescribeProject(context.TODO(), input)
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

