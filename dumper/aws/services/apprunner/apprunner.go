package apprunner

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := apprunner.NewFromConfig(awscfg)
    
	out(describeAutoScalingConfiguration(client), "apprunner", "AutoScalingConfiguration")
    
	out(describeCustomDomains(client), "apprunner", "CustomDomains")
    
	out(describeObservabilityConfiguration(client), "apprunner", "ObservabilityConfiguration")
    
	out(describeService(client), "apprunner", "Service")
    
	out(describeVpcConnector(client), "apprunner", "VpcConnector")
    
}

func describeAutoScalingConfiguration(client *apprunner.Client) map[string]interface{} {
	input := &apprunner.DescribeAutoScalingConfigurationInput{}
	result, err := client.DescribeAutoScalingConfiguration(context.TODO(), input)
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

func describeCustomDomains(client *apprunner.Client) map[string]interface{} {
	input := &apprunner.DescribeCustomDomainsInput{}
	result, err := client.DescribeCustomDomains(context.TODO(), input)
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

func describeObservabilityConfiguration(client *apprunner.Client) map[string]interface{} {
	input := &apprunner.DescribeObservabilityConfigurationInput{}
	result, err := client.DescribeObservabilityConfiguration(context.TODO(), input)
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

func describeService(client *apprunner.Client) map[string]interface{} {
	input := &apprunner.DescribeServiceInput{}
	result, err := client.DescribeService(context.TODO(), input)
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

func describeVpcConnector(client *apprunner.Client) map[string]interface{} {
	input := &apprunner.DescribeVpcConnectorInput{}
	result, err := client.DescribeVpcConnector(context.TODO(), input)
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

