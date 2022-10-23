package elasticinference

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticinference"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := elasticinference.NewFromConfig(awscfg)
    
	out(describeAcceleratorOfferings(client), "elasticinference", "AcceleratorOfferings")
    
	out(describeAcceleratorTypes(client), "elasticinference", "AcceleratorTypes")
    
	out(describeAccelerators(client), "elasticinference", "Accelerators")
    
}

func describeAcceleratorOfferings(client *elasticinference.Client) map[string]interface{} {
	input := &elasticinference.DescribeAcceleratorOfferingsInput{}
	result, err := client.DescribeAcceleratorOfferings(context.TODO(), input)
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

func describeAcceleratorTypes(client *elasticinference.Client) map[string]interface{} {
	input := &elasticinference.DescribeAcceleratorTypesInput{}
	result, err := client.DescribeAcceleratorTypes(context.TODO(), input)
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

func describeAccelerators(client *elasticinference.Client) map[string]interface{} {
	input := &elasticinference.DescribeAcceleratorsInput{}
	result, err := client.DescribeAccelerators(context.TODO(), input)
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

