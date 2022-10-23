package machinelearning

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := machinelearning.NewFromConfig(awscfg)
    
	out(describeBatchPredictions(client), "machinelearning", "BatchPredictions")
    
	out(describeDataSources(client), "machinelearning", "DataSources")
    
	out(describeEvaluations(client), "machinelearning", "Evaluations")
    
	out(describeMLModels(client), "machinelearning", "MLModels")
    
	out(describeTags(client), "machinelearning", "Tags")
    
}

func describeBatchPredictions(client *machinelearning.Client) map[string]interface{} {
	input := &machinelearning.DescribeBatchPredictionsInput{}
	result, err := client.DescribeBatchPredictions(context.TODO(), input)
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

func describeDataSources(client *machinelearning.Client) map[string]interface{} {
	input := &machinelearning.DescribeDataSourcesInput{}
	result, err := client.DescribeDataSources(context.TODO(), input)
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

func describeEvaluations(client *machinelearning.Client) map[string]interface{} {
	input := &machinelearning.DescribeEvaluationsInput{}
	result, err := client.DescribeEvaluations(context.TODO(), input)
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

func describeMLModels(client *machinelearning.Client) map[string]interface{} {
	input := &machinelearning.DescribeMLModelsInput{}
	result, err := client.DescribeMLModels(context.TODO(), input)
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

func describeTags(client *machinelearning.Client) map[string]interface{} {
	input := &machinelearning.DescribeTagsInput{}
	result, err := client.DescribeTags(context.TODO(), input)
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

