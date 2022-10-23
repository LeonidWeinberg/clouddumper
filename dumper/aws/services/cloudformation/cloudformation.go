package cloudformation

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudformation.NewFromConfig(awscfg)
    
	out(describeAccountLimits(client), "cloudformation", "AccountLimits")
    
	out(describeChangeSet(client), "cloudformation", "ChangeSet")
    
	out(describeChangeSetHooks(client), "cloudformation", "ChangeSetHooks")
    
	out(describePublisher(client), "cloudformation", "Publisher")
    
	out(describeStackDriftDetectionStatus(client), "cloudformation", "StackDriftDetectionStatus")
    
	out(describeStackEvents(client), "cloudformation", "StackEvents")
    
	out(describeStackInstance(client), "cloudformation", "StackInstance")
    
	out(describeStackResource(client), "cloudformation", "StackResource")
    
	out(describeStackResourceDrifts(client), "cloudformation", "StackResourceDrifts")
    
	out(describeStackResources(client), "cloudformation", "StackResources")
    
	out(describeStackSet(client), "cloudformation", "StackSet")
    
	out(describeStackSetOperation(client), "cloudformation", "StackSetOperation")
    
	out(describeStacks(client), "cloudformation", "Stacks")
    
	out(describeType(client), "cloudformation", "Type")
    
	out(describeTypeRegistration(client), "cloudformation", "TypeRegistration")
    
}

func describeAccountLimits(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeAccountLimitsInput{}
	result, err := client.DescribeAccountLimits(context.TODO(), input)
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

func describeChangeSet(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeChangeSetInput{}
	result, err := client.DescribeChangeSet(context.TODO(), input)
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

func describeChangeSetHooks(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeChangeSetHooksInput{}
	result, err := client.DescribeChangeSetHooks(context.TODO(), input)
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

func describePublisher(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribePublisherInput{}
	result, err := client.DescribePublisher(context.TODO(), input)
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

func describeStackDriftDetectionStatus(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackDriftDetectionStatusInput{}
	result, err := client.DescribeStackDriftDetectionStatus(context.TODO(), input)
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

func describeStackEvents(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackEventsInput{}
	result, err := client.DescribeStackEvents(context.TODO(), input)
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

func describeStackInstance(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackInstanceInput{}
	result, err := client.DescribeStackInstance(context.TODO(), input)
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

func describeStackResource(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackResourceInput{}
	result, err := client.DescribeStackResource(context.TODO(), input)
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

func describeStackResourceDrifts(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackResourceDriftsInput{}
	result, err := client.DescribeStackResourceDrifts(context.TODO(), input)
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

func describeStackResources(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackResourcesInput{}
	result, err := client.DescribeStackResources(context.TODO(), input)
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

func describeStackSet(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackSetInput{}
	result, err := client.DescribeStackSet(context.TODO(), input)
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

func describeStackSetOperation(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStackSetOperationInput{}
	result, err := client.DescribeStackSetOperation(context.TODO(), input)
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

func describeStacks(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeStacksInput{}
	result, err := client.DescribeStacks(context.TODO(), input)
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

func describeType(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeTypeInput{}
	result, err := client.DescribeType(context.TODO(), input)
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

func describeTypeRegistration(client *cloudformation.Client) map[string]interface{} {
	input := &cloudformation.DescribeTypeRegistrationInput{}
	result, err := client.DescribeTypeRegistration(context.TODO(), input)
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

