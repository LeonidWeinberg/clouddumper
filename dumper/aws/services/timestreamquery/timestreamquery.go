package timestreamquery

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := timestreamquery.NewFromConfig(awscfg)
    
	out(describeEndpoints(client), "timestreamquery", "Endpoints")
    
	out(describeScheduledQuery(client), "timestreamquery", "ScheduledQuery")
    
}

func describeEndpoints(client *timestreamquery.Client) map[string]interface{} {
	input := &timestreamquery.DescribeEndpointsInput{}
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

func describeScheduledQuery(client *timestreamquery.Client) map[string]interface{} {
	input := &timestreamquery.DescribeScheduledQueryInput{}
	result, err := client.DescribeScheduledQuery(context.TODO(), input)
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

