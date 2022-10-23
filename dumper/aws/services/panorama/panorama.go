package panorama

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/panorama"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := panorama.NewFromConfig(awscfg)
    
	out(describeApplicationInstance(client), "panorama", "ApplicationInstance")
    
	out(describeApplicationInstanceDetails(client), "panorama", "ApplicationInstanceDetails")
    
	out(describeDevice(client), "panorama", "Device")
    
	out(describeDeviceJob(client), "panorama", "DeviceJob")
    
	out(describeNode(client), "panorama", "Node")
    
	out(describeNodeFromTemplateJob(client), "panorama", "NodeFromTemplateJob")
    
	out(describePackage(client), "panorama", "Package")
    
	out(describePackageImportJob(client), "panorama", "PackageImportJob")
    
	out(describePackageVersion(client), "panorama", "PackageVersion")
    
}

func describeApplicationInstance(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribeApplicationInstanceInput{}
	result, err := client.DescribeApplicationInstance(context.TODO(), input)
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

func describeApplicationInstanceDetails(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribeApplicationInstanceDetailsInput{}
	result, err := client.DescribeApplicationInstanceDetails(context.TODO(), input)
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

func describeDevice(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribeDeviceInput{}
	result, err := client.DescribeDevice(context.TODO(), input)
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

func describeDeviceJob(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribeDeviceJobInput{}
	result, err := client.DescribeDeviceJob(context.TODO(), input)
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

func describeNode(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribeNodeInput{}
	result, err := client.DescribeNode(context.TODO(), input)
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

func describeNodeFromTemplateJob(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribeNodeFromTemplateJobInput{}
	result, err := client.DescribeNodeFromTemplateJob(context.TODO(), input)
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

func describePackage(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribePackageInput{}
	result, err := client.DescribePackage(context.TODO(), input)
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

func describePackageImportJob(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribePackageImportJobInput{}
	result, err := client.DescribePackageImportJob(context.TODO(), input)
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

func describePackageVersion(client *panorama.Client) map[string]interface{} {
	input := &panorama.DescribePackageVersionInput{}
	result, err := client.DescribePackageVersion(context.TODO(), input)
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

