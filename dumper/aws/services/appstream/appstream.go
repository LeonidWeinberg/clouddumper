package appstream

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := appstream.NewFromConfig(awscfg)
    
	out(describeAppBlocks(client), "appstream", "AppBlocks")
    
	out(describeApplicationFleetAssociations(client), "appstream", "ApplicationFleetAssociations")
    
	out(describeApplications(client), "appstream", "Applications")
    
	out(describeDirectoryConfigs(client), "appstream", "DirectoryConfigs")
    
	out(describeEntitlements(client), "appstream", "Entitlements")
    
	out(describeFleets(client), "appstream", "Fleets")
    
	out(describeImageBuilders(client), "appstream", "ImageBuilders")
    
	out(describeImagePermissions(client), "appstream", "ImagePermissions")
    
	out(describeImages(client), "appstream", "Images")
    
	out(describeSessions(client), "appstream", "Sessions")
    
	out(describeStacks(client), "appstream", "Stacks")
    
	out(describeUsageReportSubscriptions(client), "appstream", "UsageReportSubscriptions")
    
	out(describeUserStackAssociations(client), "appstream", "UserStackAssociations")
    
	out(describeUsers(client), "appstream", "Users")
    
}

func describeAppBlocks(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeAppBlocksInput{}
	result, err := client.DescribeAppBlocks(context.TODO(), input)
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

func describeApplicationFleetAssociations(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeApplicationFleetAssociationsInput{}
	result, err := client.DescribeApplicationFleetAssociations(context.TODO(), input)
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

func describeApplications(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeApplicationsInput{}
	result, err := client.DescribeApplications(context.TODO(), input)
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

func describeDirectoryConfigs(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeDirectoryConfigsInput{}
	result, err := client.DescribeDirectoryConfigs(context.TODO(), input)
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

func describeEntitlements(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeEntitlementsInput{}
	result, err := client.DescribeEntitlements(context.TODO(), input)
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

func describeFleets(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeFleetsInput{}
	result, err := client.DescribeFleets(context.TODO(), input)
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

func describeImageBuilders(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeImageBuildersInput{}
	result, err := client.DescribeImageBuilders(context.TODO(), input)
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

func describeImagePermissions(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeImagePermissionsInput{}
	result, err := client.DescribeImagePermissions(context.TODO(), input)
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

func describeImages(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeImagesInput{}
	result, err := client.DescribeImages(context.TODO(), input)
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

func describeSessions(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeSessionsInput{}
	result, err := client.DescribeSessions(context.TODO(), input)
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

func describeStacks(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeStacksInput{}
	result, err := client.DescribeStacks(context.TODO(), input)
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

func describeUsageReportSubscriptions(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeUsageReportSubscriptionsInput{}
	result, err := client.DescribeUsageReportSubscriptions(context.TODO(), input)
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

func describeUserStackAssociations(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeUserStackAssociationsInput{}
	result, err := client.DescribeUserStackAssociations(context.TODO(), input)
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

func describeUsers(client *appstream.Client) map[string]interface{} {
	input := &appstream.DescribeUsersInput{}
	result, err := client.DescribeUsers(context.TODO(), input)
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

