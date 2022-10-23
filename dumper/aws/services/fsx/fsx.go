package fsx

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := fsx.NewFromConfig(awscfg)
    
	out(describeBackups(client), "fsx", "Backups")
    
	out(describeDataRepositoryAssociations(client), "fsx", "DataRepositoryAssociations")
    
	out(describeDataRepositoryTasks(client), "fsx", "DataRepositoryTasks")
    
	out(describeFileCaches(client), "fsx", "FileCaches")
    
	out(describeFileSystemAliases(client), "fsx", "FileSystemAliases")
    
	out(describeFileSystems(client), "fsx", "FileSystems")
    
	out(describeSnapshots(client), "fsx", "Snapshots")
    
	out(describeStorageVirtualMachines(client), "fsx", "StorageVirtualMachines")
    
	out(describeVolumes(client), "fsx", "Volumes")
    
}

func describeBackups(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeBackupsInput{}
	result, err := client.DescribeBackups(context.TODO(), input)
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

func describeDataRepositoryAssociations(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeDataRepositoryAssociationsInput{}
	result, err := client.DescribeDataRepositoryAssociations(context.TODO(), input)
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

func describeDataRepositoryTasks(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeDataRepositoryTasksInput{}
	result, err := client.DescribeDataRepositoryTasks(context.TODO(), input)
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

func describeFileCaches(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeFileCachesInput{}
	result, err := client.DescribeFileCaches(context.TODO(), input)
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

func describeFileSystemAliases(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeFileSystemAliasesInput{}
	result, err := client.DescribeFileSystemAliases(context.TODO(), input)
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

func describeFileSystems(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeFileSystemsInput{}
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

func describeSnapshots(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeSnapshotsInput{}
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

func describeStorageVirtualMachines(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeStorageVirtualMachinesInput{}
	result, err := client.DescribeStorageVirtualMachines(context.TODO(), input)
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

func describeVolumes(client *fsx.Client) map[string]interface{} {
	input := &fsx.DescribeVolumesInput{}
	result, err := client.DescribeVolumes(context.TODO(), input)
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

