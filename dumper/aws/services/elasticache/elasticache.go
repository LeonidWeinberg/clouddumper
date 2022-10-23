package elasticache

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := elasticache.NewFromConfig(awscfg)
    
	out(describeCacheClusters(client), "elasticache", "CacheClusters")
    
	out(describeCacheEngineVersions(client), "elasticache", "CacheEngineVersions")
    
	out(describeCacheParameterGroups(client), "elasticache", "CacheParameterGroups")
    
	out(describeCacheParameters(client), "elasticache", "CacheParameters")
    
	out(describeCacheSecurityGroups(client), "elasticache", "CacheSecurityGroups")
    
	out(describeCacheSubnetGroups(client), "elasticache", "CacheSubnetGroups")
    
	out(describeEngineDefaultParameters(client), "elasticache", "EngineDefaultParameters")
    
	out(describeEvents(client), "elasticache", "Events")
    
	out(describeGlobalReplicationGroups(client), "elasticache", "GlobalReplicationGroups")
    
	out(describeReplicationGroups(client), "elasticache", "ReplicationGroups")
    
	out(describeReservedCacheNodes(client), "elasticache", "ReservedCacheNodes")
    
	out(describeReservedCacheNodesOfferings(client), "elasticache", "ReservedCacheNodesOfferings")
    
	out(describeServiceUpdates(client), "elasticache", "ServiceUpdates")
    
	out(describeSnapshots(client), "elasticache", "Snapshots")
    
	out(describeUpdateActions(client), "elasticache", "UpdateActions")
    
	out(describeUserGroups(client), "elasticache", "UserGroups")
    
	out(describeUsers(client), "elasticache", "Users")
    
}

func describeCacheClusters(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeCacheClustersInput{}
	result, err := client.DescribeCacheClusters(context.TODO(), input)
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

func describeCacheEngineVersions(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeCacheEngineVersionsInput{}
	result, err := client.DescribeCacheEngineVersions(context.TODO(), input)
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

func describeCacheParameterGroups(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeCacheParameterGroupsInput{}
	result, err := client.DescribeCacheParameterGroups(context.TODO(), input)
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

func describeCacheParameters(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeCacheParametersInput{}
	result, err := client.DescribeCacheParameters(context.TODO(), input)
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

func describeCacheSecurityGroups(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeCacheSecurityGroupsInput{}
	result, err := client.DescribeCacheSecurityGroups(context.TODO(), input)
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

func describeCacheSubnetGroups(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeCacheSubnetGroupsInput{}
	result, err := client.DescribeCacheSubnetGroups(context.TODO(), input)
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

func describeEngineDefaultParameters(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeEngineDefaultParametersInput{}
	result, err := client.DescribeEngineDefaultParameters(context.TODO(), input)
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

func describeEvents(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeEventsInput{}
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

func describeGlobalReplicationGroups(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeGlobalReplicationGroupsInput{}
	result, err := client.DescribeGlobalReplicationGroups(context.TODO(), input)
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

func describeReplicationGroups(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeReplicationGroupsInput{}
	result, err := client.DescribeReplicationGroups(context.TODO(), input)
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

func describeReservedCacheNodes(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeReservedCacheNodesInput{}
	result, err := client.DescribeReservedCacheNodes(context.TODO(), input)
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

func describeReservedCacheNodesOfferings(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeReservedCacheNodesOfferingsInput{}
	result, err := client.DescribeReservedCacheNodesOfferings(context.TODO(), input)
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

func describeServiceUpdates(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeServiceUpdatesInput{}
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

func describeSnapshots(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeSnapshotsInput{}
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

func describeUpdateActions(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeUpdateActionsInput{}
	result, err := client.DescribeUpdateActions(context.TODO(), input)
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

func describeUserGroups(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeUserGroupsInput{}
	result, err := client.DescribeUserGroups(context.TODO(), input)
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

func describeUsers(client *elasticache.Client) map[string]interface{} {
	input := &elasticache.DescribeUsersInput{}
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

