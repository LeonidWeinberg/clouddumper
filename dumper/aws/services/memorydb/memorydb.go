package memorydb

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/memorydb"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := memorydb.NewFromConfig(awscfg)
    
	out(describeACLs(client), "memorydb", "ACLs")
    
	out(describeClusters(client), "memorydb", "Clusters")
    
	out(describeEngineVersions(client), "memorydb", "EngineVersions")
    
	out(describeEvents(client), "memorydb", "Events")
    
	out(describeParameterGroups(client), "memorydb", "ParameterGroups")
    
	out(describeParameters(client), "memorydb", "Parameters")
    
	out(describeServiceUpdates(client), "memorydb", "ServiceUpdates")
    
	out(describeSnapshots(client), "memorydb", "Snapshots")
    
	out(describeSubnetGroups(client), "memorydb", "SubnetGroups")
    
	out(describeUsers(client), "memorydb", "Users")
    
}

func describeACLs(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeACLsInput{}
	result, err := client.DescribeACLs(context.TODO(), input)
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

func describeClusters(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeClustersInput{}
	result, err := client.DescribeClusters(context.TODO(), input)
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

func describeEngineVersions(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeEngineVersionsInput{}
	result, err := client.DescribeEngineVersions(context.TODO(), input)
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

func describeEvents(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeEventsInput{}
	result, err := client.DescribeEvents(context.TODO(), input)
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

func describeParameterGroups(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeParameterGroupsInput{}
	result, err := client.DescribeParameterGroups(context.TODO(), input)
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

func describeParameters(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeParametersInput{}
	result, err := client.DescribeParameters(context.TODO(), input)
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

func describeServiceUpdates(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeServiceUpdatesInput{}
	result, err := client.DescribeServiceUpdates(context.TODO(), input)
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

func describeSnapshots(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeSnapshotsInput{}
	result, err := client.DescribeSnapshots(context.TODO(), input)
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

func describeSubnetGroups(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeSubnetGroupsInput{}
	result, err := client.DescribeSubnetGroups(context.TODO(), input)
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

func describeUsers(client *memorydb.Client) map[string]interface{} {
	input := &memorydb.DescribeUsersInput{}
	result, err := client.DescribeUsers(context.TODO(), input)
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

