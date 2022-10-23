package mq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mq.NewFromConfig(awscfg)
    
	out(describeBroker(client), "mq", "Broker")
    
	out(describeBrokerEngineTypes(client), "mq", "BrokerEngineTypes")
    
	out(describeBrokerInstanceOptions(client), "mq", "BrokerInstanceOptions")
    
	out(describeConfiguration(client), "mq", "Configuration")
    
	out(describeConfigurationRevision(client), "mq", "ConfigurationRevision")
    
	out(describeUser(client), "mq", "User")
    
}

func describeBroker(client *mq.Client) map[string]interface{} {
	input := &mq.DescribeBrokerInput{}
	result, err := client.DescribeBroker(context.TODO(), input)
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

func describeBrokerEngineTypes(client *mq.Client) map[string]interface{} {
	input := &mq.DescribeBrokerEngineTypesInput{}
	result, err := client.DescribeBrokerEngineTypes(context.TODO(), input)
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

func describeBrokerInstanceOptions(client *mq.Client) map[string]interface{} {
	input := &mq.DescribeBrokerInstanceOptionsInput{}
	result, err := client.DescribeBrokerInstanceOptions(context.TODO(), input)
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

func describeConfiguration(client *mq.Client) map[string]interface{} {
	input := &mq.DescribeConfigurationInput{}
	result, err := client.DescribeConfiguration(context.TODO(), input)
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

func describeConfigurationRevision(client *mq.Client) map[string]interface{} {
	input := &mq.DescribeConfigurationRevisionInput{}
	result, err := client.DescribeConfigurationRevision(context.TODO(), input)
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

func describeUser(client *mq.Client) map[string]interface{} {
	input := &mq.DescribeUserInput{}
	result, err := client.DescribeUser(context.TODO(), input)
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

