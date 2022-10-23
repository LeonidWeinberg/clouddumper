package snowdevicemanagement

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := snowdevicemanagement.NewFromConfig(awscfg)
    
	out(describeDevice(client), "snowdevicemanagement", "Device")
    
	out(describeDeviceEc2Instances(client), "snowdevicemanagement", "DeviceEc2Instances")
    
	out(describeExecution(client), "snowdevicemanagement", "Execution")
    
	out(describeTask(client), "snowdevicemanagement", "Task")
    
}

func describeDevice(client *snowdevicemanagement.Client) map[string]interface{} {
	input := &snowdevicemanagement.DescribeDeviceInput{}
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

func describeDeviceEc2Instances(client *snowdevicemanagement.Client) map[string]interface{} {
	input := &snowdevicemanagement.DescribeDeviceEc2InstancesInput{}
	result, err := client.DescribeDeviceEc2Instances(context.TODO(), input)
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

func describeExecution(client *snowdevicemanagement.Client) map[string]interface{} {
	input := &snowdevicemanagement.DescribeExecutionInput{}
	result, err := client.DescribeExecution(context.TODO(), input)
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

func describeTask(client *snowdevicemanagement.Client) map[string]interface{} {
	input := &snowdevicemanagement.DescribeTaskInput{}
	result, err := client.DescribeTask(context.TODO(), input)
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

