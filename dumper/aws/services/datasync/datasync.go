package datasync

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := datasync.NewFromConfig(awscfg)
    
	out(describeAgent(client), "datasync", "Agent")
    
	out(describeLocationEfs(client), "datasync", "LocationEfs")
    
	out(describeLocationFsxLustre(client), "datasync", "LocationFsxLustre")
    
	out(describeLocationFsxOntap(client), "datasync", "LocationFsxOntap")
    
	out(describeLocationFsxOpenZfs(client), "datasync", "LocationFsxOpenZfs")
    
	out(describeLocationFsxWindows(client), "datasync", "LocationFsxWindows")
    
	out(describeLocationHdfs(client), "datasync", "LocationHdfs")
    
	out(describeLocationNfs(client), "datasync", "LocationNfs")
    
	out(describeLocationObjectStorage(client), "datasync", "LocationObjectStorage")
    
	out(describeLocationS3(client), "datasync", "LocationS3")
    
	out(describeLocationSmb(client), "datasync", "LocationSmb")
    
	out(describeTask(client), "datasync", "Task")
    
	out(describeTaskExecution(client), "datasync", "TaskExecution")
    
}

func describeAgent(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeAgentInput{}
	result, err := client.DescribeAgent(context.TODO(), input)
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

func describeLocationEfs(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationEfsInput{}
	result, err := client.DescribeLocationEfs(context.TODO(), input)
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

func describeLocationFsxLustre(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationFsxLustreInput{}
	result, err := client.DescribeLocationFsxLustre(context.TODO(), input)
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

func describeLocationFsxOntap(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationFsxOntapInput{}
	result, err := client.DescribeLocationFsxOntap(context.TODO(), input)
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

func describeLocationFsxOpenZfs(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationFsxOpenZfsInput{}
	result, err := client.DescribeLocationFsxOpenZfs(context.TODO(), input)
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

func describeLocationFsxWindows(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationFsxWindowsInput{}
	result, err := client.DescribeLocationFsxWindows(context.TODO(), input)
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

func describeLocationHdfs(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationHdfsInput{}
	result, err := client.DescribeLocationHdfs(context.TODO(), input)
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

func describeLocationNfs(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationNfsInput{}
	result, err := client.DescribeLocationNfs(context.TODO(), input)
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

func describeLocationObjectStorage(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationObjectStorageInput{}
	result, err := client.DescribeLocationObjectStorage(context.TODO(), input)
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

func describeLocationS3(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationS3Input{}
	result, err := client.DescribeLocationS3(context.TODO(), input)
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

func describeLocationSmb(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeLocationSmbInput{}
	result, err := client.DescribeLocationSmb(context.TODO(), input)
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

func describeTask(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeTaskInput{}
	result, err := client.DescribeTask(context.TODO(), input)
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

func describeTaskExecution(client *datasync.Client) map[string]interface{} {
	input := &datasync.DescribeTaskExecutionInput{}
	result, err := client.DescribeTaskExecution(context.TODO(), input)
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

