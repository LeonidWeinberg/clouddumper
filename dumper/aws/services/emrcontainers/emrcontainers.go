package emrcontainers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := emrcontainers.NewFromConfig(awscfg)
    
	out(describeJobRun(client), "emrcontainers", "JobRun")
    
	out(describeManagedEndpoint(client), "emrcontainers", "ManagedEndpoint")
    
	out(describeVirtualCluster(client), "emrcontainers", "VirtualCluster")
    
}

func describeJobRun(client *emrcontainers.Client) map[string]interface{} {
	input := &emrcontainers.DescribeJobRunInput{}
	result, err := client.DescribeJobRun(context.TODO(), input)
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

func describeManagedEndpoint(client *emrcontainers.Client) map[string]interface{} {
	input := &emrcontainers.DescribeManagedEndpointInput{}
	result, err := client.DescribeManagedEndpoint(context.TODO(), input)
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

func describeVirtualCluster(client *emrcontainers.Client) map[string]interface{} {
	input := &emrcontainers.DescribeVirtualClusterInput{}
	result, err := client.DescribeVirtualCluster(context.TODO(), input)
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

