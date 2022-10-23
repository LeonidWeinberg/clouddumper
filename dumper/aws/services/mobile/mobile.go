package mobile

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mobile"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mobile.NewFromConfig(awscfg)
    
	out(describeBundle(client), "mobile", "Bundle")
    
	out(describeProject(client), "mobile", "Project")
    
}

func describeBundle(client *mobile.Client) map[string]interface{} {
	input := &mobile.DescribeBundleInput{}
	result, err := client.DescribeBundle(context.TODO(), input)
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

func describeProject(client *mobile.Client) map[string]interface{} {
	input := &mobile.DescribeProjectInput{}
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

