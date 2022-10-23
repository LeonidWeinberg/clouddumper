package securityhub

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := securityhub.NewFromConfig(awscfg)
    
	out(describeActionTargets(client), "securityhub", "ActionTargets")
    
	out(describeHub(client), "securityhub", "Hub")
    
	out(describeOrganizationConfiguration(client), "securityhub", "OrganizationConfiguration")
    
	out(describeProducts(client), "securityhub", "Products")
    
	out(describeStandards(client), "securityhub", "Standards")
    
	out(describeStandardsControls(client), "securityhub", "StandardsControls")
    
}

func describeActionTargets(client *securityhub.Client) map[string]interface{} {
	input := &securityhub.DescribeActionTargetsInput{}
	result, err := client.DescribeActionTargets(context.TODO(), input)
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

func describeHub(client *securityhub.Client) map[string]interface{} {
	input := &securityhub.DescribeHubInput{}
	result, err := client.DescribeHub(context.TODO(), input)
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

func describeOrganizationConfiguration(client *securityhub.Client) map[string]interface{} {
	input := &securityhub.DescribeOrganizationConfigurationInput{}
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

func describeProducts(client *securityhub.Client) map[string]interface{} {
	input := &securityhub.DescribeProductsInput{}
	result, err := client.DescribeProducts(context.TODO(), input)
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

func describeStandards(client *securityhub.Client) map[string]interface{} {
	input := &securityhub.DescribeStandardsInput{}
	result, err := client.DescribeStandards(context.TODO(), input)
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

func describeStandardsControls(client *securityhub.Client) map[string]interface{} {
	input := &securityhub.DescribeStandardsControlsInput{}
	result, err := client.DescribeStandardsControls(context.TODO(), input)
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

