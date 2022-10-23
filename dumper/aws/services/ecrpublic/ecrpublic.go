package ecrpublic

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ecrpublic.NewFromConfig(awscfg)
    
	out(describeImageTags(client), "ecrpublic", "ImageTags")
    
	out(describeImages(client), "ecrpublic", "Images")
    
	out(describeRegistries(client), "ecrpublic", "Registries")
    
	out(describeRepositories(client), "ecrpublic", "Repositories")
    
}

func describeImageTags(client *ecrpublic.Client) map[string]interface{} {
	input := &ecrpublic.DescribeImageTagsInput{}
	result, err := client.DescribeImageTags(context.TODO(), input)
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

func describeImages(client *ecrpublic.Client) map[string]interface{} {
	input := &ecrpublic.DescribeImagesInput{}
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

func describeRegistries(client *ecrpublic.Client) map[string]interface{} {
	input := &ecrpublic.DescribeRegistriesInput{}
	result, err := client.DescribeRegistries(context.TODO(), input)
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

func describeRepositories(client *ecrpublic.Client) map[string]interface{} {
	input := &ecrpublic.DescribeRepositoriesInput{}
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

