package applicationdiscoveryservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := applicationdiscoveryservice.NewFromConfig(awscfg)
    
	out(describeAgents(client), "applicationdiscoveryservice", "Agents")
    
	out(describeConfigurations(client), "applicationdiscoveryservice", "Configurations")
    
	out(describeContinuousExports(client), "applicationdiscoveryservice", "ContinuousExports")
    
	out(describeExportConfigurations(client), "applicationdiscoveryservice", "ExportConfigurations")
    
	out(describeExportTasks(client), "applicationdiscoveryservice", "ExportTasks")
    
	out(describeImportTasks(client), "applicationdiscoveryservice", "ImportTasks")
    
	out(describeTags(client), "applicationdiscoveryservice", "Tags")
    
}

func describeAgents(client *applicationdiscoveryservice.Client) map[string]interface{} {
	input := &applicationdiscoveryservice.DescribeAgentsInput{}
	result, err := client.DescribeAgents(context.TODO(), input)
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

func describeConfigurations(client *applicationdiscoveryservice.Client) map[string]interface{} {
	input := &applicationdiscoveryservice.DescribeConfigurationsInput{}
	result, err := client.DescribeConfigurations(context.TODO(), input)
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

func describeContinuousExports(client *applicationdiscoveryservice.Client) map[string]interface{} {
	input := &applicationdiscoveryservice.DescribeContinuousExportsInput{}
	result, err := client.DescribeContinuousExports(context.TODO(), input)
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

func describeExportConfigurations(client *applicationdiscoveryservice.Client) map[string]interface{} {
	input := &applicationdiscoveryservice.DescribeExportConfigurationsInput{}
	result, err := client.DescribeExportConfigurations(context.TODO(), input)
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

func describeExportTasks(client *applicationdiscoveryservice.Client) map[string]interface{} {
	input := &applicationdiscoveryservice.DescribeExportTasksInput{}
	result, err := client.DescribeExportTasks(context.TODO(), input)
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

func describeImportTasks(client *applicationdiscoveryservice.Client) map[string]interface{} {
	input := &applicationdiscoveryservice.DescribeImportTasksInput{}
	result, err := client.DescribeImportTasks(context.TODO(), input)
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

func describeTags(client *applicationdiscoveryservice.Client) map[string]interface{} {
	input := &applicationdiscoveryservice.DescribeTagsInput{}
	result, err := client.DescribeTags(context.TODO(), input)
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

