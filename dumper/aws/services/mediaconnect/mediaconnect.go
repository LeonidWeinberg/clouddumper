package mediaconnect

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mediaconnect.NewFromConfig(awscfg)
    
	out(describeFlow(client), "mediaconnect", "Flow")
    
	out(describeOffering(client), "mediaconnect", "Offering")
    
	out(describeReservation(client), "mediaconnect", "Reservation")
    
}

func describeFlow(client *mediaconnect.Client) map[string]interface{} {
	input := &mediaconnect.DescribeFlowInput{}
	result, err := client.DescribeFlow(context.TODO(), input)
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

func describeOffering(client *mediaconnect.Client) map[string]interface{} {
	input := &mediaconnect.DescribeOfferingInput{}
	result, err := client.DescribeOffering(context.TODO(), input)
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

func describeReservation(client *mediaconnect.Client) map[string]interface{} {
	input := &mediaconnect.DescribeReservationInput{}
	result, err := client.DescribeReservation(context.TODO(), input)
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

