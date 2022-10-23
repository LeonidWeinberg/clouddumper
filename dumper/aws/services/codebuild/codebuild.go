package codebuild

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := codebuild.NewFromConfig(awscfg)
    
	out(describeCodeCoverages(client), "codebuild", "CodeCoverages")
    
	out(describeTestCases(client), "codebuild", "TestCases")
    
}

func describeCodeCoverages(client *codebuild.Client) map[string]interface{} {
	input := &codebuild.DescribeCodeCoveragesInput{}
	result, err := client.DescribeCodeCoverages(context.TODO(), input)
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

func describeTestCases(client *codebuild.Client) map[string]interface{} {
	input := &codebuild.DescribeTestCasesInput{}
	result, err := client.DescribeTestCases(context.TODO(), input)
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

