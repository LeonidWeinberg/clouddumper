package drs

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/drs"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := drs.NewFromConfig(awscfg)
    
	out(describeJobLogItems(client), "drs", "JobLogItems")
    
	out(describeJobs(client), "drs", "Jobs")
    
	out(describeRecoveryInstances(client), "drs", "RecoveryInstances")
    
	out(describeRecoverySnapshots(client), "drs", "RecoverySnapshots")
    
	out(describeReplicationConfigurationTemplates(client), "drs", "ReplicationConfigurationTemplates")
    
	out(describeSourceServers(client), "drs", "SourceServers")
    
}

func describeJobLogItems(client *drs.Client) map[string]interface{} {
	input := &drs.DescribeJobLogItemsInput{}
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

func describeJobs(client *drs.Client) map[string]interface{} {
	input := &drs.DescribeJobsInput{}
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

func describeRecoveryInstances(client *drs.Client) map[string]interface{} {
	input := &drs.DescribeRecoveryInstancesInput{}
	result, err := client.DescribeRecoveryInstances(context.TODO(), input)
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

func describeRecoverySnapshots(client *drs.Client) map[string]interface{} {
	input := &drs.DescribeRecoverySnapshotsInput{}
	result, err := client.DescribeRecoverySnapshots(context.TODO(), input)
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

func describeReplicationConfigurationTemplates(client *drs.Client) map[string]interface{} {
	input := &drs.DescribeReplicationConfigurationTemplatesInput{}
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

func describeSourceServers(client *drs.Client) map[string]interface{} {
	input := &drs.DescribeSourceServersInput{}
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

