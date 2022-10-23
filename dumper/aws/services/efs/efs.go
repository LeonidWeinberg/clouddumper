package efs

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := efs.NewFromConfig(awscfg)
    
	out(describeAccessPoints(client), "efs", "AccessPoints")
    
	out(describeAccountPreferences(client), "efs", "AccountPreferences")
    
	out(describeBackupPolicy(client), "efs", "BackupPolicy")
    
	out(describeFileSystemPolicy(client), "efs", "FileSystemPolicy")
    
	out(describeFileSystems(client), "efs", "FileSystems")
    
	out(describeLifecycleConfiguration(client), "efs", "LifecycleConfiguration")
    
	out(describeMountTargetSecurityGroups(client), "efs", "MountTargetSecurityGroups")
    
	out(describeMountTargets(client), "efs", "MountTargets")
    
	out(describeReplicationConfigurations(client), "efs", "ReplicationConfigurations")
    
	out(describeTags(client), "efs", "Tags")
    
}

func describeAccessPoints(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeAccessPointsInput{}
	result, err := client.DescribeAccessPoints(context.TODO(), input)
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

func describeAccountPreferences(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeAccountPreferencesInput{}
	result, err := client.DescribeAccountPreferences(context.TODO(), input)
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

func describeBackupPolicy(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeBackupPolicyInput{}
	result, err := client.DescribeBackupPolicy(context.TODO(), input)
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

func describeFileSystemPolicy(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeFileSystemPolicyInput{}
	result, err := client.DescribeFileSystemPolicy(context.TODO(), input)
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

func describeFileSystems(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeFileSystemsInput{}
	result, err := client.DescribeFileSystems(context.TODO(), input)
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

func describeLifecycleConfiguration(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeLifecycleConfigurationInput{}
	result, err := client.DescribeLifecycleConfiguration(context.TODO(), input)
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

func describeMountTargetSecurityGroups(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeMountTargetSecurityGroupsInput{}
	result, err := client.DescribeMountTargetSecurityGroups(context.TODO(), input)
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

func describeMountTargets(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeMountTargetsInput{}
	result, err := client.DescribeMountTargets(context.TODO(), input)
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

func describeReplicationConfigurations(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeReplicationConfigurationsInput{}
	result, err := client.DescribeReplicationConfigurations(context.TODO(), input)
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

func describeTags(client *efs.Client) map[string]interface{} {
	input := &efs.DescribeTagsInput{}
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

