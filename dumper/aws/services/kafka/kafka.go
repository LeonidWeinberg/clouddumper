package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := kafka.NewFromConfig(awscfg)
    
	out(describeCluster(client), "kafka", "Cluster")
    
	out(describeClusterOperation(client), "kafka", "ClusterOperation")
    
	out(describeClusterV2(client), "kafka", "ClusterV2")
    
	out(describeConfiguration(client), "kafka", "Configuration")
    
	out(describeConfigurationRevision(client), "kafka", "ConfigurationRevision")
    
}

func describeCluster(client *kafka.Client) map[string]interface{} {
	input := &kafka.DescribeClusterInput{}
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

func describeClusterOperation(client *kafka.Client) map[string]interface{} {
	input := &kafka.DescribeClusterOperationInput{}
	result, err := client.DescribeClusterOperation(context.TODO(), input)
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

func describeClusterV2(client *kafka.Client) map[string]interface{} {
	input := &kafka.DescribeClusterV2Input{}
	result, err := client.DescribeClusterV2(context.TODO(), input)
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

func describeConfiguration(client *kafka.Client) map[string]interface{} {
	input := &kafka.DescribeConfigurationInput{}
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

func describeConfigurationRevision(client *kafka.Client) map[string]interface{} {
	input := &kafka.DescribeConfigurationRevisionInput{}
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

