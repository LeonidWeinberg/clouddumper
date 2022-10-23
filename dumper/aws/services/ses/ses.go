package ses

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ses.NewFromConfig(awscfg)
    
	out(describeActiveReceiptRuleSet(client), "ses", "ActiveReceiptRuleSet")
    
	out(describeConfigurationSet(client), "ses", "ConfigurationSet")
    
	out(describeReceiptRule(client), "ses", "ReceiptRule")
    
	out(describeReceiptRuleSet(client), "ses", "ReceiptRuleSet")
    
}

func describeActiveReceiptRuleSet(client *ses.Client) map[string]interface{} {
	input := &ses.DescribeActiveReceiptRuleSetInput{}
	result, err := client.DescribeActiveReceiptRuleSet(context.TODO(), input)
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

func describeConfigurationSet(client *ses.Client) map[string]interface{} {
	input := &ses.DescribeConfigurationSetInput{}
	result, err := client.DescribeConfigurationSet(context.TODO(), input)
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

func describeReceiptRule(client *ses.Client) map[string]interface{} {
	input := &ses.DescribeReceiptRuleInput{}
	result, err := client.DescribeReceiptRule(context.TODO(), input)
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

func describeReceiptRuleSet(client *ses.Client) map[string]interface{} {
	input := &ses.DescribeReceiptRuleSetInput{}
	result, err := client.DescribeReceiptRuleSet(context.TODO(), input)
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

