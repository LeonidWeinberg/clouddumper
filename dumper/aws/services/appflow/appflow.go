package appflow

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := appflow.NewFromConfig(awscfg)
    
	out(describeConnector(client), "appflow", "Connector")
    
	out(describeConnectorEntity(client), "appflow", "ConnectorEntity")
    
	out(describeConnectorProfiles(client), "appflow", "ConnectorProfiles")
    
	out(describeConnectors(client), "appflow", "Connectors")
    
	out(describeFlow(client), "appflow", "Flow")
    
	out(describeFlowExecutionRecords(client), "appflow", "FlowExecutionRecords")
    
}

func describeConnector(client *appflow.Client) map[string]interface{} {
	input := &appflow.DescribeConnectorInput{}
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

func describeConnectorEntity(client *appflow.Client) map[string]interface{} {
	input := &appflow.DescribeConnectorEntityInput{}
	result, err := client.DescribeConnectorEntity(context.TODO(), input)
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

func describeConnectorProfiles(client *appflow.Client) map[string]interface{} {
	input := &appflow.DescribeConnectorProfilesInput{}
	result, err := client.DescribeConnectorProfiles(context.TODO(), input)
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

func describeConnectors(client *appflow.Client) map[string]interface{} {
	input := &appflow.DescribeConnectorsInput{}
	result, err := client.DescribeConnectors(context.TODO(), input)
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

func describeFlow(client *appflow.Client) map[string]interface{} {
	input := &appflow.DescribeFlowInput{}
	result, err := client.DescribeFlow(context.TODO(), input)
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

func describeFlowExecutionRecords(client *appflow.Client) map[string]interface{} {
	input := &appflow.DescribeFlowExecutionRecordsInput{}
	result, err := client.DescribeFlowExecutionRecords(context.TODO(), input)
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

