package ssoadmin

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ssoadmin.NewFromConfig(awscfg)
    
	out(describeAccountAssignmentCreationStatus(client), "ssoadmin", "AccountAssignmentCreationStatus")
    
	out(describeAccountAssignmentDeletionStatus(client), "ssoadmin", "AccountAssignmentDeletionStatus")
    
	out(describeInstanceAccessControlAttributeConfiguration(client), "ssoadmin", "InstanceAccessControlAttributeConfiguration")
    
	out(describePermissionSet(client), "ssoadmin", "PermissionSet")
    
	out(describePermissionSetProvisioningStatus(client), "ssoadmin", "PermissionSetProvisioningStatus")
    
}

func describeAccountAssignmentCreationStatus(client *ssoadmin.Client) map[string]interface{} {
	input := &ssoadmin.DescribeAccountAssignmentCreationStatusInput{}
	result, err := client.DescribeAccountAssignmentCreationStatus(context.TODO(), input)
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

func describeAccountAssignmentDeletionStatus(client *ssoadmin.Client) map[string]interface{} {
	input := &ssoadmin.DescribeAccountAssignmentDeletionStatusInput{}
	result, err := client.DescribeAccountAssignmentDeletionStatus(context.TODO(), input)
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

func describeInstanceAccessControlAttributeConfiguration(client *ssoadmin.Client) map[string]interface{} {
	input := &ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput{}
	result, err := client.DescribeInstanceAccessControlAttributeConfiguration(context.TODO(), input)
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

func describePermissionSet(client *ssoadmin.Client) map[string]interface{} {
	input := &ssoadmin.DescribePermissionSetInput{}
	result, err := client.DescribePermissionSet(context.TODO(), input)
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

func describePermissionSetProvisioningStatus(client *ssoadmin.Client) map[string]interface{} {
	input := &ssoadmin.DescribePermissionSetProvisioningStatusInput{}
	result, err := client.DescribePermissionSetProvisioningStatus(context.TODO(), input)
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

