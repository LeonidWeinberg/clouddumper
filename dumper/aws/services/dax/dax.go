package dax

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := dax.NewFromConfig(awscfg)
    
	out(describeClusters(client), "dax", "Clusters")
    
	out(describeDefaultParameters(client), "dax", "DefaultParameters")
    
	out(describeEvents(client), "dax", "Events")
    
	out(describeParameterGroups(client), "dax", "ParameterGroups")
    
	out(describeParameters(client), "dax", "Parameters")
    
	out(describeSubnetGroups(client), "dax", "SubnetGroups")
    
}

func describeClusters(client *dax.Client) map[string]interface{} {
	input := &dax.DescribeClustersInput{}
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

func describeDefaultParameters(client *dax.Client) map[string]interface{} {
	input := &dax.DescribeDefaultParametersInput{}
	result, err := client.DescribeDefaultParameters(context.TODO(), input)
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

func describeEvents(client *dax.Client) map[string]interface{} {
	input := &dax.DescribeEventsInput{}
	result, err := client.DescribeEvents(context.TODO(), input)
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

func describeParameterGroups(client *dax.Client) map[string]interface{} {
	input := &dax.DescribeParameterGroupsInput{}
	result, err := client.DescribeParameterGroups(context.TODO(), input)
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

func describeParameters(client *dax.Client) map[string]interface{} {
	input := &dax.DescribeParametersInput{}
	result, err := client.DescribeParameters(context.TODO(), input)
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

func describeSubnetGroups(client *dax.Client) map[string]interface{} {
	input := &dax.DescribeSubnetGroupsInput{}
	result, err := client.DescribeSubnetGroups(context.TODO(), input)
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

