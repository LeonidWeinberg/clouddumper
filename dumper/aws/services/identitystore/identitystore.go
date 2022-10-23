package identitystore

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := identitystore.NewFromConfig(awscfg)
    
	out(describeGroup(client), "identitystore", "Group")
    
	out(describeGroupMembership(client), "identitystore", "GroupMembership")
    
	out(describeUser(client), "identitystore", "User")
    
}

func describeGroup(client *identitystore.Client) map[string]interface{} {
	input := &identitystore.DescribeGroupInput{}
	result, err := client.DescribeGroup(context.TODO(), input)
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

func describeGroupMembership(client *identitystore.Client) map[string]interface{} {
	input := &identitystore.DescribeGroupMembershipInput{}
	result, err := client.DescribeGroupMembership(context.TODO(), input)
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

func describeUser(client *identitystore.Client) map[string]interface{} {
	input := &identitystore.DescribeUserInput{}
	result, err := client.DescribeUser(context.TODO(), input)
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

