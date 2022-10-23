package autoscalingplans

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := autoscalingplans.NewFromConfig(awscfg)
    
	out(describeScalingPlanResources(client), "autoscalingplans", "ScalingPlanResources")
    
	out(describeScalingPlans(client), "autoscalingplans", "ScalingPlans")
    
}

func describeScalingPlanResources(client *autoscalingplans.Client) map[string]interface{} {
	input := &autoscalingplans.DescribeScalingPlanResourcesInput{}
	result, err := client.DescribeScalingPlanResources(context.TODO(), input)
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

func describeScalingPlans(client *autoscalingplans.Client) map[string]interface{} {
	input := &autoscalingplans.DescribeScalingPlansInput{}
	result, err := client.DescribeScalingPlans(context.TODO(), input)
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

