package emr

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := emr.NewFromConfig(awscfg)
    
	out(describeCluster(client), "emr", "Cluster")
    
	out(describeJobFlows(client), "emr", "JobFlows")
    
	out(describeNotebookExecution(client), "emr", "NotebookExecution")
    
	out(describeReleaseLabel(client), "emr", "ReleaseLabel")
    
	out(describeSecurityConfiguration(client), "emr", "SecurityConfiguration")
    
	out(describeStep(client), "emr", "Step")
    
	out(describeStudio(client), "emr", "Studio")
    
}

func describeCluster(client *emr.Client) map[string]interface{} {
	input := &emr.DescribeClusterInput{}
	result, err := client.DescribeCluster(context.TODO(), input)
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

func describeJobFlows(client *emr.Client) map[string]interface{} {
	input := &emr.DescribeJobFlowsInput{}
	result, err := client.DescribeJobFlows(context.TODO(), input)
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

func describeNotebookExecution(client *emr.Client) map[string]interface{} {
	input := &emr.DescribeNotebookExecutionInput{}
	result, err := client.DescribeNotebookExecution(context.TODO(), input)
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

func describeReleaseLabel(client *emr.Client) map[string]interface{} {
	input := &emr.DescribeReleaseLabelInput{}
	result, err := client.DescribeReleaseLabel(context.TODO(), input)
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

func describeSecurityConfiguration(client *emr.Client) map[string]interface{} {
	input := &emr.DescribeSecurityConfigurationInput{}
	result, err := client.DescribeSecurityConfiguration(context.TODO(), input)
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

func describeStep(client *emr.Client) map[string]interface{} {
	input := &emr.DescribeStepInput{}
	result, err := client.DescribeStep(context.TODO(), input)
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

func describeStudio(client *emr.Client) map[string]interface{} {
	input := &emr.DescribeStudioInput{}
	result, err := client.DescribeStudio(context.TODO(), input)
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

