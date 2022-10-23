package amp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := amp.NewFromConfig(awscfg)
    
	out(describeAlertManagerDefinition(client), "amp", "AlertManagerDefinition")
    
	out(describeLoggingConfiguration(client), "amp", "LoggingConfiguration")
    
	out(describeRuleGroupsNamespace(client), "amp", "RuleGroupsNamespace")
    
	out(describeWorkspace(client), "amp", "Workspace")
    
}

func describeAlertManagerDefinition(client *amp.Client) map[string]interface{} {
	input := &amp.DescribeAlertManagerDefinitionInput{}
	result, err := client.DescribeAlertManagerDefinition(context.TODO(), input)
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

func describeLoggingConfiguration(client *amp.Client) map[string]interface{} {
	input := &amp.DescribeLoggingConfigurationInput{}
	result, err := client.DescribeLoggingConfiguration(context.TODO(), input)
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

func describeRuleGroupsNamespace(client *amp.Client) map[string]interface{} {
	input := &amp.DescribeRuleGroupsNamespaceInput{}
	result, err := client.DescribeRuleGroupsNamespace(context.TODO(), input)
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

func describeWorkspace(client *amp.Client) map[string]interface{} {
	input := &amp.DescribeWorkspaceInput{}
	result, err := client.DescribeWorkspace(context.TODO(), input)
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

