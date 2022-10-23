package cloudsearch

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudsearch.NewFromConfig(awscfg)
    
	out(describeAnalysisSchemes(client), "cloudsearch", "AnalysisSchemes")
    
	out(describeAvailabilityOptions(client), "cloudsearch", "AvailabilityOptions")
    
	out(describeDomainEndpointOptions(client), "cloudsearch", "DomainEndpointOptions")
    
	out(describeDomains(client), "cloudsearch", "Domains")
    
	out(describeExpressions(client), "cloudsearch", "Expressions")
    
	out(describeIndexFields(client), "cloudsearch", "IndexFields")
    
	out(describeScalingParameters(client), "cloudsearch", "ScalingParameters")
    
	out(describeServiceAccessPolicies(client), "cloudsearch", "ServiceAccessPolicies")
    
	out(describeSuggesters(client), "cloudsearch", "Suggesters")
    
}

func describeAnalysisSchemes(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeAnalysisSchemesInput{}
	result, err := client.DescribeAnalysisSchemes(context.TODO(), input)
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

func describeAvailabilityOptions(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeAvailabilityOptionsInput{}
	result, err := client.DescribeAvailabilityOptions(context.TODO(), input)
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

func describeDomainEndpointOptions(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeDomainEndpointOptionsInput{}
	result, err := client.DescribeDomainEndpointOptions(context.TODO(), input)
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

func describeDomains(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeDomainsInput{}
	result, err := client.DescribeDomains(context.TODO(), input)
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

func describeExpressions(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeExpressionsInput{}
	result, err := client.DescribeExpressions(context.TODO(), input)
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

func describeIndexFields(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeIndexFieldsInput{}
	result, err := client.DescribeIndexFields(context.TODO(), input)
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

func describeScalingParameters(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeScalingParametersInput{}
	result, err := client.DescribeScalingParameters(context.TODO(), input)
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

func describeServiceAccessPolicies(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeServiceAccessPoliciesInput{}
	result, err := client.DescribeServiceAccessPolicies(context.TODO(), input)
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

func describeSuggesters(client *cloudsearch.Client) map[string]interface{} {
	input := &cloudsearch.DescribeSuggestersInput{}
	result, err := client.DescribeSuggesters(context.TODO(), input)
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

