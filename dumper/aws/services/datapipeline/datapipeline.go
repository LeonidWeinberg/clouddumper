package datapipeline

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := datapipeline.NewFromConfig(awscfg)
    
	out(describeObjects(client), "datapipeline", "Objects")
    
	out(describePipelines(client), "datapipeline", "Pipelines")
    
}

func describeObjects(client *datapipeline.Client) map[string]interface{} {
	input := &datapipeline.DescribeObjectsInput{}
	result, err := client.DescribeObjects(context.TODO(), input)
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

func describePipelines(client *datapipeline.Client) map[string]interface{} {
	input := &datapipeline.DescribePipelinesInput{}
	result, err := client.DescribePipelines(context.TODO(), input)
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

