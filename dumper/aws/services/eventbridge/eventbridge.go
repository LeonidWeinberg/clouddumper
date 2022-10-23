package eventbridge

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := eventbridge.NewFromConfig(awscfg)
    
	out(describeApiDestination(client), "eventbridge", "ApiDestination")
    
	out(describeArchive(client), "eventbridge", "Archive")
    
	out(describeConnection(client), "eventbridge", "Connection")
    
	out(describeEndpoint(client), "eventbridge", "Endpoint")
    
	out(describeEventBus(client), "eventbridge", "EventBus")
    
	out(describeEventSource(client), "eventbridge", "EventSource")
    
	out(describePartnerEventSource(client), "eventbridge", "PartnerEventSource")
    
	out(describeReplay(client), "eventbridge", "Replay")
    
	out(describeRule(client), "eventbridge", "Rule")
    
}

func describeApiDestination(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeApiDestinationInput{}
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

func describeArchive(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeArchiveInput{}
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

func describeConnection(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeConnectionInput{}
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

func describeEndpoint(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeEndpointInput{}
	result, err := client.DescribeEndpoint(context.TODO(), input)
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

func describeEventBus(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeEventBusInput{}
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

func describeEventSource(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeEventSourceInput{}
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

func describePartnerEventSource(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribePartnerEventSourceInput{}
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

func describeReplay(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeReplayInput{}
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

func describeRule(client *eventbridge.Client) map[string]interface{} {
	input := &eventbridge.DescribeRuleInput{}
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

