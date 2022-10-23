package iotthingsgraph

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iotthingsgraph.NewFromConfig(awscfg)
    
	out(describeNamespace(client), "iotthingsgraph", "Namespace")
    
}

func describeNamespace(client *iotthingsgraph.Client) map[string]interface{} {
	input := &iotthingsgraph.DescribeNamespaceInput{}
	result, err := client.DescribeNamespace(context.TODO(), input)
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

