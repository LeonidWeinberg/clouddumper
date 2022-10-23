package snowball

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/snowball"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := snowball.NewFromConfig(awscfg)
    
	out(describeAddress(client), "snowball", "Address")
    
	out(describeAddresses(client), "snowball", "Addresses")
    
	out(describeCluster(client), "snowball", "Cluster")
    
	out(describeJob(client), "snowball", "Job")
    
	out(describeReturnShippingLabel(client), "snowball", "ReturnShippingLabel")
    
}

func describeAddress(client *snowball.Client) map[string]interface{} {
	input := &snowball.DescribeAddressInput{}
	result, err := client.DescribeAddress(context.TODO(), input)
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

func describeAddresses(client *snowball.Client) map[string]interface{} {
	input := &snowball.DescribeAddressesInput{}
	result, err := client.DescribeAddresses(context.TODO(), input)
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

func describeCluster(client *snowball.Client) map[string]interface{} {
	input := &snowball.DescribeClusterInput{}
	result, err := client.DescribeCluster(context.TODO(), input)
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

func describeJob(client *snowball.Client) map[string]interface{} {
	input := &snowball.DescribeJobInput{}
	result, err := client.DescribeJob(context.TODO(), input)
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

func describeReturnShippingLabel(client *snowball.Client) map[string]interface{} {
	input := &snowball.DescribeReturnShippingLabelInput{}
	result, err := client.DescribeReturnShippingLabel(context.TODO(), input)
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

