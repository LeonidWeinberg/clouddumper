package cloud9

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloud9.NewFromConfig(awscfg)
    
	out(describeEnvironmentMemberships(client), "cloud9", "EnvironmentMemberships")
    
	out(describeEnvironmentStatus(client), "cloud9", "EnvironmentStatus")
    
	out(describeEnvironments(client), "cloud9", "Environments")
    
}

func describeEnvironmentMemberships(client *cloud9.Client) map[string]interface{} {
	input := &cloud9.DescribeEnvironmentMembershipsInput{}
	result, err := client.DescribeEnvironmentMemberships(context.TODO(), input)
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

func describeEnvironmentStatus(client *cloud9.Client) map[string]interface{} {
	input := &cloud9.DescribeEnvironmentStatusInput{}
	result, err := client.DescribeEnvironmentStatus(context.TODO(), input)
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

func describeEnvironments(client *cloud9.Client) map[string]interface{} {
	input := &cloud9.DescribeEnvironmentsInput{}
	result, err := client.DescribeEnvironments(context.TODO(), input)
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

