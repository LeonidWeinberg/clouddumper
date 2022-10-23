package lookoutequipment

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := lookoutequipment.NewFromConfig(awscfg)
    
	out(describeDataIngestionJob(client), "lookoutequipment", "DataIngestionJob")
    
	out(describeDataset(client), "lookoutequipment", "Dataset")
    
	out(describeInferenceScheduler(client), "lookoutequipment", "InferenceScheduler")
    
	out(describeLabel(client), "lookoutequipment", "Label")
    
	out(describeLabelGroup(client), "lookoutequipment", "LabelGroup")
    
	out(describeModel(client), "lookoutequipment", "Model")
    
}

func describeDataIngestionJob(client *lookoutequipment.Client) map[string]interface{} {
	input := &lookoutequipment.DescribeDataIngestionJobInput{}
	result, err := client.DescribeDataIngestionJob(context.TODO(), input)
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

func describeDataset(client *lookoutequipment.Client) map[string]interface{} {
	input := &lookoutequipment.DescribeDatasetInput{}
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

func describeInferenceScheduler(client *lookoutequipment.Client) map[string]interface{} {
	input := &lookoutequipment.DescribeInferenceSchedulerInput{}
	result, err := client.DescribeInferenceScheduler(context.TODO(), input)
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

func describeLabel(client *lookoutequipment.Client) map[string]interface{} {
	input := &lookoutequipment.DescribeLabelInput{}
	result, err := client.DescribeLabel(context.TODO(), input)
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

func describeLabelGroup(client *lookoutequipment.Client) map[string]interface{} {
	input := &lookoutequipment.DescribeLabelGroupInput{}
	result, err := client.DescribeLabelGroup(context.TODO(), input)
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

func describeModel(client *lookoutequipment.Client) map[string]interface{} {
	input := &lookoutequipment.DescribeModelInput{}
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

