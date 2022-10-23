package guardduty

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := guardduty.NewFromConfig(awscfg)
    
	out(describeMalwareScans(client), "guardduty", "MalwareScans")
    
	out(describeOrganizationConfiguration(client), "guardduty", "OrganizationConfiguration")
    
	out(describePublishingDestination(client), "guardduty", "PublishingDestination")
    
}

func describeMalwareScans(client *guardduty.Client) map[string]interface{} {
	input := &guardduty.DescribeMalwareScansInput{}
	result, err := client.DescribeMalwareScans(context.TODO(), input)
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

func describeOrganizationConfiguration(client *guardduty.Client) map[string]interface{} {
	input := &guardduty.DescribeOrganizationConfigurationInput{}
	result, err := client.DescribeOrganizationConfiguration(context.TODO(), input)
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

func describePublishingDestination(client *guardduty.Client) map[string]interface{} {
	input := &guardduty.DescribePublishingDestinationInput{}
	result, err := client.DescribePublishingDestination(context.TODO(), input)
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

