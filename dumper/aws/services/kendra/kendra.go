package kendra

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kendra"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := kendra.NewFromConfig(awscfg)
    
	out(describeAccessControlConfiguration(client), "kendra", "AccessControlConfiguration")
    
	out(describeDataSource(client), "kendra", "DataSource")
    
	out(describeExperience(client), "kendra", "Experience")
    
	out(describeFaq(client), "kendra", "Faq")
    
	out(describeIndex(client), "kendra", "Index")
    
	out(describePrincipalMapping(client), "kendra", "PrincipalMapping")
    
	out(describeQuerySuggestionsBlockList(client), "kendra", "QuerySuggestionsBlockList")
    
	out(describeQuerySuggestionsConfig(client), "kendra", "QuerySuggestionsConfig")
    
	out(describeThesaurus(client), "kendra", "Thesaurus")
    
}

func describeAccessControlConfiguration(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeAccessControlConfigurationInput{}
	result, err := client.DescribeAccessControlConfiguration(context.TODO(), input)
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

func describeDataSource(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeDataSourceInput{}
	result, err := client.DescribeDataSource(context.TODO(), input)
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

func describeExperience(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeExperienceInput{}
	result, err := client.DescribeExperience(context.TODO(), input)
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

func describeFaq(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeFaqInput{}
	result, err := client.DescribeFaq(context.TODO(), input)
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

func describeIndex(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeIndexInput{}
	result, err := client.DescribeIndex(context.TODO(), input)
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

func describePrincipalMapping(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribePrincipalMappingInput{}
	result, err := client.DescribePrincipalMapping(context.TODO(), input)
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

func describeQuerySuggestionsBlockList(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeQuerySuggestionsBlockListInput{}
	result, err := client.DescribeQuerySuggestionsBlockList(context.TODO(), input)
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

func describeQuerySuggestionsConfig(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeQuerySuggestionsConfigInput{}
	result, err := client.DescribeQuerySuggestionsConfig(context.TODO(), input)
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

func describeThesaurus(client *kendra.Client) map[string]interface{} {
	input := &kendra.DescribeThesaurusInput{}
	result, err := client.DescribeThesaurus(context.TODO(), input)
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

