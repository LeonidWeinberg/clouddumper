package cognitoidentityprovider

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cognitoidentityprovider.NewFromConfig(awscfg)
    
	out(describeIdentityProvider(client), "cognitoidentityprovider", "IdentityProvider")
    
	out(describeResourceServer(client), "cognitoidentityprovider", "ResourceServer")
    
	out(describeRiskConfiguration(client), "cognitoidentityprovider", "RiskConfiguration")
    
	out(describeUserImportJob(client), "cognitoidentityprovider", "UserImportJob")
    
	out(describeUserPool(client), "cognitoidentityprovider", "UserPool")
    
	out(describeUserPoolClient(client), "cognitoidentityprovider", "UserPoolClient")
    
	out(describeUserPoolDomain(client), "cognitoidentityprovider", "UserPoolDomain")
    
}

func describeIdentityProvider(client *cognitoidentityprovider.Client) map[string]interface{} {
	input := &cognitoidentityprovider.DescribeIdentityProviderInput{}
	result, err := client.DescribeIdentityProvider(context.TODO(), input)
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

func describeResourceServer(client *cognitoidentityprovider.Client) map[string]interface{} {
	input := &cognitoidentityprovider.DescribeResourceServerInput{}
	result, err := client.DescribeResourceServer(context.TODO(), input)
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

func describeRiskConfiguration(client *cognitoidentityprovider.Client) map[string]interface{} {
	input := &cognitoidentityprovider.DescribeRiskConfigurationInput{}
	result, err := client.DescribeRiskConfiguration(context.TODO(), input)
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

func describeUserImportJob(client *cognitoidentityprovider.Client) map[string]interface{} {
	input := &cognitoidentityprovider.DescribeUserImportJobInput{}
	result, err := client.DescribeUserImportJob(context.TODO(), input)
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

func describeUserPool(client *cognitoidentityprovider.Client) map[string]interface{} {
	input := &cognitoidentityprovider.DescribeUserPoolInput{}
	result, err := client.DescribeUserPool(context.TODO(), input)
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

func describeUserPoolClient(client *cognitoidentityprovider.Client) map[string]interface{} {
	input := &cognitoidentityprovider.DescribeUserPoolClientInput{}
	result, err := client.DescribeUserPoolClient(context.TODO(), input)
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

func describeUserPoolDomain(client *cognitoidentityprovider.Client) map[string]interface{} {
	input := &cognitoidentityprovider.DescribeUserPoolDomainInput{}
	result, err := client.DescribeUserPoolDomain(context.TODO(), input)
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

