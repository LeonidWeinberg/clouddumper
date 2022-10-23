package lakeformation

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := lakeformation.NewFromConfig(awscfg)
    
	out(describeResource(client), "lakeformation", "Resource")
    
	out(describeTransaction(client), "lakeformation", "Transaction")
    
}

func describeResource(client *lakeformation.Client) map[string]interface{} {
	input := &lakeformation.DescribeResourceInput{}
	result, err := client.DescribeResource(context.TODO(), input)
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

func describeTransaction(client *lakeformation.Client) map[string]interface{} {
	input := &lakeformation.DescribeTransactionInput{}
	result, err := client.DescribeTransaction(context.TODO(), input)
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

