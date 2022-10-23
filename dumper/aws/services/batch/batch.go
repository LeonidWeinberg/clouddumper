package batch

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := batch.NewFromConfig(awscfg)
    
	out(describeComputeEnvironments(client), "batch", "ComputeEnvironments")
    
	out(describeJobDefinitions(client), "batch", "JobDefinitions")
    
	out(describeJobQueues(client), "batch", "JobQueues")
    
	out(describeJobs(client), "batch", "Jobs")
    
	out(describeSchedulingPolicies(client), "batch", "SchedulingPolicies")
    
}

func describeComputeEnvironments(client *batch.Client) map[string]interface{} {
	input := &batch.DescribeComputeEnvironmentsInput{}
	result, err := client.DescribeComputeEnvironments(context.TODO(), input)
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

func describeJobDefinitions(client *batch.Client) map[string]interface{} {
	input := &batch.DescribeJobDefinitionsInput{}
	result, err := client.DescribeJobDefinitions(context.TODO(), input)
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

func describeJobQueues(client *batch.Client) map[string]interface{} {
	input := &batch.DescribeJobQueuesInput{}
	result, err := client.DescribeJobQueues(context.TODO(), input)
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

func describeJobs(client *batch.Client) map[string]interface{} {
	input := &batch.DescribeJobsInput{}
	result, err := client.DescribeJobs(context.TODO(), input)
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

func describeSchedulingPolicies(client *batch.Client) map[string]interface{} {
	input := &batch.DescribeSchedulingPoliciesInput{}
	result, err := client.DescribeSchedulingPolicies(context.TODO(), input)
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

