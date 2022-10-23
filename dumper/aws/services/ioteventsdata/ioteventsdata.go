package ioteventsdata

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ioteventsdata.NewFromConfig(awscfg)
    
	out(describeAlarm(client), "ioteventsdata", "Alarm")
    
	out(describeDetector(client), "ioteventsdata", "Detector")
    
}

func describeAlarm(client *ioteventsdata.Client) map[string]interface{} {
	input := &ioteventsdata.DescribeAlarmInput{}
	result, err := client.DescribeAlarm(context.TODO(), input)
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

func describeDetector(client *ioteventsdata.Client) map[string]interface{} {
	input := &ioteventsdata.DescribeDetectorInput{}
	result, err := client.DescribeDetector(context.TODO(), input)
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

