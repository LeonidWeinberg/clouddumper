package cloudwatchlogs

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudwatchlogs.NewFromConfig(awscfg)
    
	out(describeDestinations(client), "cloudwatchlogs", "Destinations")
    
	out(describeExportTasks(client), "cloudwatchlogs", "ExportTasks")
    
	out(describeLogGroups(client), "cloudwatchlogs", "LogGroups")
    
	out(describeLogStreams(client), "cloudwatchlogs", "LogStreams")
    
	out(describeMetricFilters(client), "cloudwatchlogs", "MetricFilters")
    
	out(describeQueries(client), "cloudwatchlogs", "Queries")
    
	out(describeQueryDefinitions(client), "cloudwatchlogs", "QueryDefinitions")
    
	out(describeResourcePolicies(client), "cloudwatchlogs", "ResourcePolicies")
    
	out(describeSubscriptionFilters(client), "cloudwatchlogs", "SubscriptionFilters")
    
}

func describeDestinations(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeDestinationsInput{}
	result, err := client.DescribeDestinations(context.TODO(), input)
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

func describeExportTasks(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeExportTasksInput{}
	result, err := client.DescribeExportTasks(context.TODO(), input)
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

func describeLogGroups(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeLogGroupsInput{}
	result, err := client.DescribeLogGroups(context.TODO(), input)
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

func describeLogStreams(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeLogStreamsInput{}
	result, err := client.DescribeLogStreams(context.TODO(), input)
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

func describeMetricFilters(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeMetricFiltersInput{}
	result, err := client.DescribeMetricFilters(context.TODO(), input)
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

func describeQueries(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeQueriesInput{}
	result, err := client.DescribeQueries(context.TODO(), input)
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

func describeQueryDefinitions(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeQueryDefinitionsInput{}
	result, err := client.DescribeQueryDefinitions(context.TODO(), input)
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

func describeResourcePolicies(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeResourcePoliciesInput{}
	result, err := client.DescribeResourcePolicies(context.TODO(), input)
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

func describeSubscriptionFilters(client *cloudwatchlogs.Client) map[string]interface{} {
	input := &cloudwatchlogs.DescribeSubscriptionFiltersInput{}
	result, err := client.DescribeSubscriptionFilters(context.TODO(), input)
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

