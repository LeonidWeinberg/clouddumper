package mediapackage

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mediapackage.NewFromConfig(awscfg)
    
	out(describeChannel(client), "mediapackage", "Channel")
    
	out(describeHarvestJob(client), "mediapackage", "HarvestJob")
    
	out(describeOriginEndpoint(client), "mediapackage", "OriginEndpoint")
    
}

func describeChannel(client *mediapackage.Client) map[string]interface{} {
	input := &mediapackage.DescribeChannelInput{}
	result, err := client.DescribeChannel(context.TODO(), input)
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

func describeHarvestJob(client *mediapackage.Client) map[string]interface{} {
	input := &mediapackage.DescribeHarvestJobInput{}
	result, err := client.DescribeHarvestJob(context.TODO(), input)
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

func describeOriginEndpoint(client *mediapackage.Client) map[string]interface{} {
	input := &mediapackage.DescribeOriginEndpointInput{}
	result, err := client.DescribeOriginEndpoint(context.TODO(), input)
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

