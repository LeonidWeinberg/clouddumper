package iot1clickprojects

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iot1clickprojects.NewFromConfig(awscfg)
    
	out(describePlacement(client), "iot1clickprojects", "Placement")
    
	out(describeProject(client), "iot1clickprojects", "Project")
    
}

func describePlacement(client *iot1clickprojects.Client) map[string]interface{} {
	input := &iot1clickprojects.DescribePlacementInput{}
	result, err := client.DescribePlacement(context.TODO(), input)
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

func describeProject(client *iot1clickprojects.Client) map[string]interface{} {
	input := &iot1clickprojects.DescribeProjectInput{}
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

