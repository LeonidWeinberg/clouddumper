package kms

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := kms.NewFromConfig(awscfg)
    
	out(describeCustomKeyStores(client), "kms", "CustomKeyStores")
    
	out(describeKey(client), "kms", "Key")
    
}

func describeCustomKeyStores(client *kms.Client) map[string]interface{} {
	input := &kms.DescribeCustomKeyStoresInput{}
	result, err := client.DescribeCustomKeyStores(context.TODO(), input)
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

func describeKey(client *kms.Client) map[string]interface{} {
	input := &kms.DescribeKeyInput{}
	result, err := client.DescribeKey(context.TODO(), input)
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

