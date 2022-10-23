package iotevents

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotevents"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iotevents.NewFromConfig(awscfg)
    
	out(describeAlarmModel(client), "iotevents", "AlarmModel")
    
	out(describeDetectorModel(client), "iotevents", "DetectorModel")
    
	out(describeDetectorModelAnalysis(client), "iotevents", "DetectorModelAnalysis")
    
	out(describeInput(client), "iotevents", "Input")
    
	out(describeLoggingOptions(client), "iotevents", "LoggingOptions")
    
}

func describeAlarmModel(client *iotevents.Client) map[string]interface{} {
	input := &iotevents.DescribeAlarmModelInput{}
	result, err := client.DescribeAlarmModel(context.TODO(), input)
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

func describeDetectorModel(client *iotevents.Client) map[string]interface{} {
	input := &iotevents.DescribeDetectorModelInput{}
	result, err := client.DescribeDetectorModel(context.TODO(), input)
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

func describeDetectorModelAnalysis(client *iotevents.Client) map[string]interface{} {
	input := &iotevents.DescribeDetectorModelAnalysisInput{}
	result, err := client.DescribeDetectorModelAnalysis(context.TODO(), input)
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

func describeInput(client *iotevents.Client) map[string]interface{} {
	input := &iotevents.DescribeInputInput{}
	result, err := client.DescribeInput(context.TODO(), input)
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

func describeLoggingOptions(client *iotevents.Client) map[string]interface{} {
	input := &iotevents.DescribeLoggingOptionsInput{}
	result, err := client.DescribeLoggingOptions(context.TODO(), input)
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

