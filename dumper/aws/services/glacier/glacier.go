package glacier

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := glacier.NewFromConfig(awscfg)
    
	out(describeJob(client), "glacier", "Job")
    
	out(describeVault(client), "glacier", "Vault")
    
}

func describeJob(client *glacier.Client) map[string]interface{} {
	input := &glacier.DescribeJobInput{}
	result, err := client.DescribeJob(context.TODO(), input)
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

func describeVault(client *glacier.Client) map[string]interface{} {
	input := &glacier.DescribeVaultInput{}
	result, err := client.DescribeVault(context.TODO(), input)
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

