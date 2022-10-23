package ecr

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ecr.NewFromConfig(awscfg)
    
	out(describeImageReplicationStatus(client), "ecr", "ImageReplicationStatus")
    
	out(describeImageScanFindings(client), "ecr", "ImageScanFindings")
    
	out(describeImages(client), "ecr", "Images")
    
	out(describePullThroughCacheRules(client), "ecr", "PullThroughCacheRules")
    
	out(describeRegistry(client), "ecr", "Registry")
    
	out(describeRepositories(client), "ecr", "Repositories")
    
}

func describeImageReplicationStatus(client *ecr.Client) map[string]interface{} {
	input := &ecr.DescribeImageReplicationStatusInput{}
	result, err := client.DescribeImageReplicationStatus(context.TODO(), input)
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

func describeImageScanFindings(client *ecr.Client) map[string]interface{} {
	input := &ecr.DescribeImageScanFindingsInput{}
	result, err := client.DescribeImageScanFindings(context.TODO(), input)
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

func describeImages(client *ecr.Client) map[string]interface{} {
	input := &ecr.DescribeImagesInput{}
	result, err := client.DescribeImages(context.TODO(), input)
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

func describePullThroughCacheRules(client *ecr.Client) map[string]interface{} {
	input := &ecr.DescribePullThroughCacheRulesInput{}
	result, err := client.DescribePullThroughCacheRules(context.TODO(), input)
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

func describeRegistry(client *ecr.Client) map[string]interface{} {
	input := &ecr.DescribeRegistryInput{}
	result, err := client.DescribeRegistry(context.TODO(), input)
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

func describeRepositories(client *ecr.Client) map[string]interface{} {
	input := &ecr.DescribeRepositoriesInput{}
	result, err := client.DescribeRepositories(context.TODO(), input)
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

