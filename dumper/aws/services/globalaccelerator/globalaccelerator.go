package globalaccelerator

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := globalaccelerator.NewFromConfig(awscfg)
    
	out(describeAccelerator(client), "globalaccelerator", "Accelerator")
    
	out(describeAcceleratorAttributes(client), "globalaccelerator", "AcceleratorAttributes")
    
	out(describeCustomRoutingAccelerator(client), "globalaccelerator", "CustomRoutingAccelerator")
    
	out(describeCustomRoutingAcceleratorAttributes(client), "globalaccelerator", "CustomRoutingAcceleratorAttributes")
    
	out(describeCustomRoutingEndpointGroup(client), "globalaccelerator", "CustomRoutingEndpointGroup")
    
	out(describeCustomRoutingListener(client), "globalaccelerator", "CustomRoutingListener")
    
	out(describeEndpointGroup(client), "globalaccelerator", "EndpointGroup")
    
	out(describeListener(client), "globalaccelerator", "Listener")
    
}

func describeAccelerator(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeAcceleratorInput{}
	result, err := client.DescribeAccelerator(context.TODO(), input)
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

func describeAcceleratorAttributes(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeAcceleratorAttributesInput{}
	result, err := client.DescribeAcceleratorAttributes(context.TODO(), input)
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

func describeCustomRoutingAccelerator(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeCustomRoutingAcceleratorInput{}
	result, err := client.DescribeCustomRoutingAccelerator(context.TODO(), input)
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

func describeCustomRoutingAcceleratorAttributes(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput{}
	result, err := client.DescribeCustomRoutingAcceleratorAttributes(context.TODO(), input)
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

func describeCustomRoutingEndpointGroup(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeCustomRoutingEndpointGroupInput{}
	result, err := client.DescribeCustomRoutingEndpointGroup(context.TODO(), input)
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

func describeCustomRoutingListener(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeCustomRoutingListenerInput{}
	result, err := client.DescribeCustomRoutingListener(context.TODO(), input)
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

func describeEndpointGroup(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeEndpointGroupInput{}
	result, err := client.DescribeEndpointGroup(context.TODO(), input)
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

func describeListener(client *globalaccelerator.Client) map[string]interface{} {
	input := &globalaccelerator.DescribeListenerInput{}
	result, err := client.DescribeListener(context.TODO(), input)
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

