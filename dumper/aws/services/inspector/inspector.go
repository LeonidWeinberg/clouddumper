package inspector

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := inspector.NewFromConfig(awscfg)
    
	out(describeAssessmentRuns(client), "inspector", "AssessmentRuns")
    
	out(describeAssessmentTargets(client), "inspector", "AssessmentTargets")
    
	out(describeAssessmentTemplates(client), "inspector", "AssessmentTemplates")
    
	out(describeCrossAccountAccessRole(client), "inspector", "CrossAccountAccessRole")
    
	out(describeExclusions(client), "inspector", "Exclusions")
    
	out(describeFindings(client), "inspector", "Findings")
    
	out(describeResourceGroups(client), "inspector", "ResourceGroups")
    
	out(describeRulesPackages(client), "inspector", "RulesPackages")
    
}

func describeAssessmentRuns(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeAssessmentRunsInput{}
	result, err := client.DescribeAssessmentRuns(context.TODO(), input)
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

func describeAssessmentTargets(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeAssessmentTargetsInput{}
	result, err := client.DescribeAssessmentTargets(context.TODO(), input)
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

func describeAssessmentTemplates(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeAssessmentTemplatesInput{}
	result, err := client.DescribeAssessmentTemplates(context.TODO(), input)
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

func describeCrossAccountAccessRole(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeCrossAccountAccessRoleInput{}
	result, err := client.DescribeCrossAccountAccessRole(context.TODO(), input)
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

func describeExclusions(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeExclusionsInput{}
	result, err := client.DescribeExclusions(context.TODO(), input)
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

func describeFindings(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeFindingsInput{}
	result, err := client.DescribeFindings(context.TODO(), input)
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

func describeResourceGroups(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeResourceGroupsInput{}
	result, err := client.DescribeResourceGroups(context.TODO(), input)
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

func describeRulesPackages(client *inspector.Client) map[string]interface{} {
	input := &inspector.DescribeRulesPackagesInput{}
	result, err := client.DescribeRulesPackages(context.TODO(), input)
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

