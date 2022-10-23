package resiliencehub

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := resiliencehub.NewFromConfig(awscfg)
    
	out(describeApp(client), "resiliencehub", "App")
    
	out(describeAppAssessment(client), "resiliencehub", "AppAssessment")
    
	out(describeAppVersionResourcesResolutionStatus(client), "resiliencehub", "AppVersionResourcesResolutionStatus")
    
	out(describeAppVersionTemplate(client), "resiliencehub", "AppVersionTemplate")
    
	out(describeDraftAppVersionResourcesImportStatus(client), "resiliencehub", "DraftAppVersionResourcesImportStatus")
    
	out(describeResiliencyPolicy(client), "resiliencehub", "ResiliencyPolicy")
    
}

func describeApp(client *resiliencehub.Client) map[string]interface{} {
	input := &resiliencehub.DescribeAppInput{}
	result, err := client.DescribeApp(context.TODO(), input)
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

func describeAppAssessment(client *resiliencehub.Client) map[string]interface{} {
	input := &resiliencehub.DescribeAppAssessmentInput{}
	result, err := client.DescribeAppAssessment(context.TODO(), input)
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

func describeAppVersionResourcesResolutionStatus(client *resiliencehub.Client) map[string]interface{} {
	input := &resiliencehub.DescribeAppVersionResourcesResolutionStatusInput{}
	result, err := client.DescribeAppVersionResourcesResolutionStatus(context.TODO(), input)
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

func describeAppVersionTemplate(client *resiliencehub.Client) map[string]interface{} {
	input := &resiliencehub.DescribeAppVersionTemplateInput{}
	result, err := client.DescribeAppVersionTemplate(context.TODO(), input)
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

func describeDraftAppVersionResourcesImportStatus(client *resiliencehub.Client) map[string]interface{} {
	input := &resiliencehub.DescribeDraftAppVersionResourcesImportStatusInput{}
	result, err := client.DescribeDraftAppVersionResourcesImportStatus(context.TODO(), input)
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

func describeResiliencyPolicy(client *resiliencehub.Client) map[string]interface{} {
	input := &resiliencehub.DescribeResiliencyPolicyInput{}
	result, err := client.DescribeResiliencyPolicy(context.TODO(), input)
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

