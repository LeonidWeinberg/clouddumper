package kafkaconnect

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := kafkaconnect.NewFromConfig(awscfg)
    
	out(describeConnector(client), "kafkaconnect", "Connector")
    
	out(describeCustomPlugin(client), "kafkaconnect", "CustomPlugin")
    
	out(describeWorkerConfiguration(client), "kafkaconnect", "WorkerConfiguration")
    
}

func describeConnector(client *kafkaconnect.Client) map[string]interface{} {
	input := &kafkaconnect.DescribeConnectorInput{}
	result, err := client.DescribeConnector(context.TODO(), input)
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

func describeCustomPlugin(client *kafkaconnect.Client) map[string]interface{} {
	input := &kafkaconnect.DescribeCustomPluginInput{}
	result, err := client.DescribeCustomPlugin(context.TODO(), input)
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

func describeWorkerConfiguration(client *kafkaconnect.Client) map[string]interface{} {
	input := &kafkaconnect.DescribeWorkerConfigurationInput{}
	result, err := client.DescribeWorkerConfiguration(context.TODO(), input)
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

