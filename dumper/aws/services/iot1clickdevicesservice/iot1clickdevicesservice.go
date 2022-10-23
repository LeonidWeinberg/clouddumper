package iot1clickdevicesservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iot1clickdevicesservice.NewFromConfig(awscfg)
    
	out(describeDevice(client), "iot1clickdevicesservice", "Device")
    
}

func describeDevice(client *iot1clickdevicesservice.Client) map[string]interface{} {
	input := &iot1clickdevicesservice.DescribeDeviceInput{}
	result, err := client.DescribeDevice(context.TODO(), input)
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

