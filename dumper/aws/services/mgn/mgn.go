package mgn

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mgn"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mgn.NewFromConfig(awscfg)
    
	out(describeJobLogItems(client), "mgn", "JobLogItems")
    
	out(describeJobs(client), "mgn", "Jobs")
    
	out(describeLaunchConfigurationTemplates(client), "mgn", "LaunchConfigurationTemplates")
    
	out(describeReplicationConfigurationTemplates(client), "mgn", "ReplicationConfigurationTemplates")
    
	out(describeSourceServers(client), "mgn", "SourceServers")
    
	out(describeVcenterClients(client), "mgn", "VcenterClients")
    
}

func describeJobLogItems(client *mgn.Client) map[string]interface{} {
	input := &mgn.DescribeJobLogItemsInput{}
	result, err := client.DescribeJobLogItems(context.TODO(), input)
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

func describeJobs(client *mgn.Client) map[string]interface{} {
	input := &mgn.DescribeJobsInput{}
	result, err := client.DescribeJobs(context.TODO(), input)
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

func describeLaunchConfigurationTemplates(client *mgn.Client) map[string]interface{} {
	input := &mgn.DescribeLaunchConfigurationTemplatesInput{}
	result, err := client.DescribeLaunchConfigurationTemplates(context.TODO(), input)
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

func describeReplicationConfigurationTemplates(client *mgn.Client) map[string]interface{} {
	input := &mgn.DescribeReplicationConfigurationTemplatesInput{}
	result, err := client.DescribeReplicationConfigurationTemplates(context.TODO(), input)
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

func describeSourceServers(client *mgn.Client) map[string]interface{} {
	input := &mgn.DescribeSourceServersInput{}
	result, err := client.DescribeSourceServers(context.TODO(), input)
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

func describeVcenterClients(client *mgn.Client) map[string]interface{} {
	input := &mgn.DescribeVcenterClientsInput{}
	result, err := client.DescribeVcenterClients(context.TODO(), input)
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

