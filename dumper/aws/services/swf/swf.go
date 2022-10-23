package swf

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/swf"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := swf.NewFromConfig(awscfg)
    
	out(describeActivityType(client), "swf", "ActivityType")
    
	out(describeDomain(client), "swf", "Domain")
    
	out(describeWorkflowExecution(client), "swf", "WorkflowExecution")
    
	out(describeWorkflowType(client), "swf", "WorkflowType")
    
}

func describeActivityType(client *swf.Client) map[string]interface{} {
	input := &swf.DescribeActivityTypeInput{}
	result, err := client.DescribeActivityType(context.TODO(), input)
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

func describeDomain(client *swf.Client) map[string]interface{} {
	input := &swf.DescribeDomainInput{}
	result, err := client.DescribeDomain(context.TODO(), input)
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

func describeWorkflowExecution(client *swf.Client) map[string]interface{} {
	input := &swf.DescribeWorkflowExecutionInput{}
	result, err := client.DescribeWorkflowExecution(context.TODO(), input)
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

func describeWorkflowType(client *swf.Client) map[string]interface{} {
	input := &swf.DescribeWorkflowTypeInput{}
	result, err := client.DescribeWorkflowType(context.TODO(), input)
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

