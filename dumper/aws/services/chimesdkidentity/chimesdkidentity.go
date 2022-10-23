package chimesdkidentity

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := chimesdkidentity.NewFromConfig(awscfg)
    
	out(describeAppInstance(client), "chimesdkidentity", "AppInstance")
    
	out(describeAppInstanceAdmin(client), "chimesdkidentity", "AppInstanceAdmin")
    
	out(describeAppInstanceUser(client), "chimesdkidentity", "AppInstanceUser")
    
	out(describeAppInstanceUserEndpoint(client), "chimesdkidentity", "AppInstanceUserEndpoint")
    
}

func describeAppInstance(client *chimesdkidentity.Client) map[string]interface{} {
	input := &chimesdkidentity.DescribeAppInstanceInput{}
	result, err := client.DescribeAppInstance(context.TODO(), input)
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

func describeAppInstanceAdmin(client *chimesdkidentity.Client) map[string]interface{} {
	input := &chimesdkidentity.DescribeAppInstanceAdminInput{}
	result, err := client.DescribeAppInstanceAdmin(context.TODO(), input)
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

func describeAppInstanceUser(client *chimesdkidentity.Client) map[string]interface{} {
	input := &chimesdkidentity.DescribeAppInstanceUserInput{}
	result, err := client.DescribeAppInstanceUser(context.TODO(), input)
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

func describeAppInstanceUserEndpoint(client *chimesdkidentity.Client) map[string]interface{} {
	input := &chimesdkidentity.DescribeAppInstanceUserEndpointInput{}
	result, err := client.DescribeAppInstanceUserEndpoint(context.TODO(), input)
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

