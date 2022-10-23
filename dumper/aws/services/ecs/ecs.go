package ecs

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ecs.NewFromConfig(awscfg)
    
	out(describeCapacityProviders(client), "ecs", "CapacityProviders")
    
	out(describeClusters(client), "ecs", "Clusters")
    
	out(describeContainerInstances(client), "ecs", "ContainerInstances")
    
	out(describeServices(client), "ecs", "Services")
    
	out(describeTaskDefinition(client), "ecs", "TaskDefinition")
    
	out(describeTaskSets(client), "ecs", "TaskSets")
    
	out(describeTasks(client), "ecs", "Tasks")
    
}

func describeCapacityProviders(client *ecs.Client) map[string]interface{} {
	input := &ecs.DescribeCapacityProvidersInput{}
	result, err := client.DescribeCapacityProviders(context.TODO(), input)
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

func describeClusters(client *ecs.Client) map[string]interface{} {
	input := &ecs.DescribeClustersInput{}
	result, err := client.DescribeClusters(context.TODO(), input)
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

func describeContainerInstances(client *ecs.Client) map[string]interface{} {
	input := &ecs.DescribeContainerInstancesInput{}
	result, err := client.DescribeContainerInstances(context.TODO(), input)
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

func describeServices(client *ecs.Client) map[string]interface{} {
	input := &ecs.DescribeServicesInput{}
	result, err := client.DescribeServices(context.TODO(), input)
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

func describeTaskDefinition(client *ecs.Client) map[string]interface{} {
	input := &ecs.DescribeTaskDefinitionInput{}
	result, err := client.DescribeTaskDefinition(context.TODO(), input)
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

func describeTaskSets(client *ecs.Client) map[string]interface{} {
	input := &ecs.DescribeTaskSetsInput{}
	result, err := client.DescribeTaskSets(context.TODO(), input)
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

func describeTasks(client *ecs.Client) map[string]interface{} {
	input := &ecs.DescribeTasksInput{}
	result, err := client.DescribeTasks(context.TODO(), input)
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

