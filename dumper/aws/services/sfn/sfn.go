package sfn

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := sfn.NewFromConfig(awscfg)
    
	out(describeActivity(client), "sfn", "Activity")
    
	out(describeExecution(client), "sfn", "Execution")
    
	out(describeStateMachine(client), "sfn", "StateMachine")
    
	out(describeStateMachineForExecution(client), "sfn", "StateMachineForExecution")
    
}

func describeActivity(client *sfn.Client) map[string]interface{} {
	input := &sfn.DescribeActivityInput{}
	result, err := client.DescribeActivity(context.TODO(), input)
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

func describeExecution(client *sfn.Client) map[string]interface{} {
	input := &sfn.DescribeExecutionInput{}
	result, err := client.DescribeExecution(context.TODO(), input)
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

func describeStateMachine(client *sfn.Client) map[string]interface{} {
	input := &sfn.DescribeStateMachineInput{}
	result, err := client.DescribeStateMachine(context.TODO(), input)
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

func describeStateMachineForExecution(client *sfn.Client) map[string]interface{} {
	input := &sfn.DescribeStateMachineForExecutionInput{}
	result, err := client.DescribeStateMachineForExecution(context.TODO(), input)
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

