package directoryservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := directoryservice.NewFromConfig(awscfg)
    
	out(describeCertificate(client), "directoryservice", "Certificate")
    
	out(describeClientAuthenticationSettings(client), "directoryservice", "ClientAuthenticationSettings")
    
	out(describeConditionalForwarders(client), "directoryservice", "ConditionalForwarders")
    
	out(describeDirectories(client), "directoryservice", "Directories")
    
	out(describeDomainControllers(client), "directoryservice", "DomainControllers")
    
	out(describeEventTopics(client), "directoryservice", "EventTopics")
    
	out(describeLDAPSSettings(client), "directoryservice", "LDAPSSettings")
    
	out(describeRegions(client), "directoryservice", "Regions")
    
	out(describeSettings(client), "directoryservice", "Settings")
    
	out(describeSharedDirectories(client), "directoryservice", "SharedDirectories")
    
	out(describeSnapshots(client), "directoryservice", "Snapshots")
    
	out(describeTrusts(client), "directoryservice", "Trusts")
    
	out(describeUpdateDirectory(client), "directoryservice", "UpdateDirectory")
    
}

func describeCertificate(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeCertificateInput{}
	result, err := client.DescribeCertificate(context.TODO(), input)
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

func describeClientAuthenticationSettings(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeClientAuthenticationSettingsInput{}
	result, err := client.DescribeClientAuthenticationSettings(context.TODO(), input)
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

func describeConditionalForwarders(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeConditionalForwardersInput{}
	result, err := client.DescribeConditionalForwarders(context.TODO(), input)
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

func describeDirectories(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeDirectoriesInput{}
	result, err := client.DescribeDirectories(context.TODO(), input)
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

func describeDomainControllers(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeDomainControllersInput{}
	result, err := client.DescribeDomainControllers(context.TODO(), input)
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

func describeEventTopics(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeEventTopicsInput{}
	result, err := client.DescribeEventTopics(context.TODO(), input)
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

func describeLDAPSSettings(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeLDAPSSettingsInput{}
	result, err := client.DescribeLDAPSSettings(context.TODO(), input)
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

func describeRegions(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeRegionsInput{}
	result, err := client.DescribeRegions(context.TODO(), input)
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

func describeSettings(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeSettingsInput{}
	result, err := client.DescribeSettings(context.TODO(), input)
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

func describeSharedDirectories(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeSharedDirectoriesInput{}
	result, err := client.DescribeSharedDirectories(context.TODO(), input)
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

func describeSnapshots(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeSnapshotsInput{}
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

func describeTrusts(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeTrustsInput{}
	result, err := client.DescribeTrusts(context.TODO(), input)
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

func describeUpdateDirectory(client *directoryservice.Client) map[string]interface{} {
	input := &directoryservice.DescribeUpdateDirectoryInput{}
	result, err := client.DescribeUpdateDirectory(context.TODO(), input)
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

