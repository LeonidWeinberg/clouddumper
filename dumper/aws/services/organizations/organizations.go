package organizations

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := organizations.NewFromConfig(awscfg)
    
	out(describeAccount(client), "organizations", "Account")
    
	out(describeCreateAccountStatus(client), "organizations", "CreateAccountStatus")
    
	out(describeEffectivePolicy(client), "organizations", "EffectivePolicy")
    
	out(describeHandshake(client), "organizations", "Handshake")
    
	out(describeOrganization(client), "organizations", "Organization")
    
	out(describeOrganizationalUnit(client), "organizations", "OrganizationalUnit")
    
	out(describePolicy(client), "organizations", "Policy")
    
}

func describeAccount(client *organizations.Client) map[string]interface{} {
	input := &organizations.DescribeAccountInput{}
	result, err := client.DescribeAccount(context.TODO(), input)
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

func describeCreateAccountStatus(client *organizations.Client) map[string]interface{} {
	input := &organizations.DescribeCreateAccountStatusInput{}
	result, err := client.DescribeCreateAccountStatus(context.TODO(), input)
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

func describeEffectivePolicy(client *organizations.Client) map[string]interface{} {
	input := &organizations.DescribeEffectivePolicyInput{}
	result, err := client.DescribeEffectivePolicy(context.TODO(), input)
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

func describeHandshake(client *organizations.Client) map[string]interface{} {
	input := &organizations.DescribeHandshakeInput{}
	result, err := client.DescribeHandshake(context.TODO(), input)
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

func describeOrganization(client *organizations.Client) map[string]interface{} {
	input := &organizations.DescribeOrganizationInput{}
	result, err := client.DescribeOrganization(context.TODO(), input)
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

func describeOrganizationalUnit(client *organizations.Client) map[string]interface{} {
	input := &organizations.DescribeOrganizationalUnitInput{}
	result, err := client.DescribeOrganizationalUnit(context.TODO(), input)
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

func describePolicy(client *organizations.Client) map[string]interface{} {
	input := &organizations.DescribePolicyInput{}
	result, err := client.DescribePolicy(context.TODO(), input)
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

