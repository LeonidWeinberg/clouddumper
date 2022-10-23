package applicationautoscaling

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := applicationautoscaling.NewFromConfig(awscfg)
    
	out(describeScalableTargets(client), "applicationautoscaling", "ScalableTargets")
    
	out(describeScalingActivities(client), "applicationautoscaling", "ScalingActivities")
    
	out(describeScalingPolicies(client), "applicationautoscaling", "ScalingPolicies")
    
	out(describeScheduledActions(client), "applicationautoscaling", "ScheduledActions")
    
}

func describeScalableTargets(client *applicationautoscaling.Client) map[string]interface{} {
	input := &applicationautoscaling.DescribeScalableTargetsInput{}
	result, err := client.DescribeScalableTargets(context.TODO(), input)
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

func describeScalingActivities(client *applicationautoscaling.Client) map[string]interface{} {
	input := &applicationautoscaling.DescribeScalingActivitiesInput{}
	result, err := client.DescribeScalingActivities(context.TODO(), input)
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

func describeScalingPolicies(client *applicationautoscaling.Client) map[string]interface{} {
	input := &applicationautoscaling.DescribeScalingPoliciesInput{}
	result, err := client.DescribeScalingPolicies(context.TODO(), input)
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

func describeScheduledActions(client *applicationautoscaling.Client) map[string]interface{} {
	input := &applicationautoscaling.DescribeScheduledActionsInput{}
	result, err := client.DescribeScheduledActions(context.TODO(), input)
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

