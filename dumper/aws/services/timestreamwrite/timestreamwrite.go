package timestreamwrite

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := timestreamwrite.NewFromConfig(awscfg)
    
	out(describeDatabase(client), "timestreamwrite", "Database")
    
	out(describeEndpoints(client), "timestreamwrite", "Endpoints")
    
	out(describeTable(client), "timestreamwrite", "Table")
    
}

func describeDatabase(client *timestreamwrite.Client) map[string]interface{} {
	input := &timestreamwrite.DescribeDatabaseInput{}
	result, err := client.DescribeDatabase(context.TODO(), input)
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

func describeEndpoints(client *timestreamwrite.Client) map[string]interface{} {
	input := &timestreamwrite.DescribeEndpointsInput{}
	result, err := client.DescribeEndpoints(context.TODO(), input)
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

func describeTable(client *timestreamwrite.Client) map[string]interface{} {
	input := &timestreamwrite.DescribeTableInput{}
	result, err := client.DescribeTable(context.TODO(), input)
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

