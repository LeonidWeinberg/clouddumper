package schemas

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := schemas.NewFromConfig(awscfg)
    
	out(describeCodeBinding(client), "schemas", "CodeBinding")
    
	out(describeDiscoverer(client), "schemas", "Discoverer")
    
	out(describeRegistry(client), "schemas", "Registry")
    
	out(describeSchema(client), "schemas", "Schema")
    
}

func describeCodeBinding(client *schemas.Client) map[string]interface{} {
	input := &schemas.DescribeCodeBindingInput{}
	result, err := client.DescribeCodeBinding(context.TODO(), input)
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

func describeDiscoverer(client *schemas.Client) map[string]interface{} {
	input := &schemas.DescribeDiscovererInput{}
	result, err := client.DescribeDiscoverer(context.TODO(), input)
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

func describeRegistry(client *schemas.Client) map[string]interface{} {
	input := &schemas.DescribeRegistryInput{}
	result, err := client.DescribeRegistry(context.TODO(), input)
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

func describeSchema(client *schemas.Client) map[string]interface{} {
	input := &schemas.DescribeSchemaInput{}
	result, err := client.DescribeSchema(context.TODO(), input)
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

