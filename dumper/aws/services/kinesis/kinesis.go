package kinesis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := kinesis.NewFromConfig(awscfg)
    
	out(describeLimits(client), "kinesis", "Limits")
    
	out(describeStream(client), "kinesis", "Stream")
    
	out(describeStreamConsumer(client), "kinesis", "StreamConsumer")
    
	out(describeStreamSummary(client), "kinesis", "StreamSummary")
    
}

func describeLimits(client *kinesis.Client) map[string]interface{} {
	input := &kinesis.DescribeLimitsInput{}
	result, err := client.DescribeLimits(context.TODO(), input)
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

func describeStream(client *kinesis.Client) map[string]interface{} {
	input := &kinesis.DescribeStreamInput{}
	result, err := client.DescribeStream(context.TODO(), input)
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

func describeStreamConsumer(client *kinesis.Client) map[string]interface{} {
	input := &kinesis.DescribeStreamConsumerInput{}
	result, err := client.DescribeStreamConsumer(context.TODO(), input)
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

func describeStreamSummary(client *kinesis.Client) map[string]interface{} {
	input := &kinesis.DescribeStreamSummaryInput{}
	result, err := client.DescribeStreamSummary(context.TODO(), input)
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

