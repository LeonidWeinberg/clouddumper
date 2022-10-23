package codeartifact

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := codeartifact.NewFromConfig(awscfg)
    
	out(describeDomain(client), "codeartifact", "Domain")
    
	out(describePackage(client), "codeartifact", "Package")
    
	out(describePackageVersion(client), "codeartifact", "PackageVersion")
    
	out(describeRepository(client), "codeartifact", "Repository")
    
}

func describeDomain(client *codeartifact.Client) map[string]interface{} {
	input := &codeartifact.DescribeDomainInput{}
	result, err := client.DescribeDomain(context.TODO(), input)
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

func describePackage(client *codeartifact.Client) map[string]interface{} {
	input := &codeartifact.DescribePackageInput{}
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

func describePackageVersion(client *codeartifact.Client) map[string]interface{} {
	input := &codeartifact.DescribePackageVersionInput{}
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

func describeRepository(client *codeartifact.Client) map[string]interface{} {
	input := &codeartifact.DescribeRepositoryInput{}
	result, err := client.DescribeRepository(context.TODO(), input)
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

