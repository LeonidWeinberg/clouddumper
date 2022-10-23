package cloudwatchevents

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudwatchevents.NewFromConfig(awscfg)
    
	out(describeApiDestination(client), "cloudwatchevents", "ApiDestination")
    
	out(describeArchive(client), "cloudwatchevents", "Archive")
    
	out(describeConnection(client), "cloudwatchevents", "Connection")
    
	out(describeEventBus(client), "cloudwatchevents", "EventBus")
    
	out(describeEventSource(client), "cloudwatchevents", "EventSource")
    
	out(describePartnerEventSource(client), "cloudwatchevents", "PartnerEventSource")
    
	out(describeReplay(client), "cloudwatchevents", "Replay")
    
	out(describeRule(client), "cloudwatchevents", "Rule")
    
}

func describeApiDestination(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribeApiDestinationInput{}
	result, err := client.DescribeApiDestination(context.TODO(), input)
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

func describeArchive(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribeArchiveInput{}
	result, err := client.DescribeArchive(context.TODO(), input)
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

func describeConnection(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribeConnectionInput{}
	result, err := client.DescribeConnection(context.TODO(), input)
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

func describeEventBus(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribeEventBusInput{}
	result, err := client.DescribeEventBus(context.TODO(), input)
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

func describeEventSource(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribeEventSourceInput{}
	result, err := client.DescribeEventSource(context.TODO(), input)
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

func describePartnerEventSource(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribePartnerEventSourceInput{}
	result, err := client.DescribePartnerEventSource(context.TODO(), input)
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

func describeReplay(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribeReplayInput{}
	result, err := client.DescribeReplay(context.TODO(), input)
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

func describeRule(client *cloudwatchevents.Client) map[string]interface{} {
	input := &cloudwatchevents.DescribeRuleInput{}
	result, err := client.DescribeRule(context.TODO(), input)
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

