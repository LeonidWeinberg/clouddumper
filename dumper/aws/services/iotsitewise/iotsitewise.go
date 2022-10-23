package iotsitewise

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iotsitewise.NewFromConfig(awscfg)
    
	out(describeAccessPolicy(client), "iotsitewise", "AccessPolicy")
    
	out(describeAsset(client), "iotsitewise", "Asset")
    
	out(describeAssetModel(client), "iotsitewise", "AssetModel")
    
	out(describeAssetProperty(client), "iotsitewise", "AssetProperty")
    
	out(describeBulkImportJob(client), "iotsitewise", "BulkImportJob")
    
	out(describeDashboard(client), "iotsitewise", "Dashboard")
    
	out(describeDefaultEncryptionConfiguration(client), "iotsitewise", "DefaultEncryptionConfiguration")
    
	out(describeGateway(client), "iotsitewise", "Gateway")
    
	out(describeGatewayCapabilityConfiguration(client), "iotsitewise", "GatewayCapabilityConfiguration")
    
	out(describeLoggingOptions(client), "iotsitewise", "LoggingOptions")
    
	out(describePortal(client), "iotsitewise", "Portal")
    
	out(describeProject(client), "iotsitewise", "Project")
    
	out(describeStorageConfiguration(client), "iotsitewise", "StorageConfiguration")
    
	out(describeTimeSeries(client), "iotsitewise", "TimeSeries")
    
}

func describeAccessPolicy(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeAccessPolicyInput{}
	result, err := client.DescribeAccessPolicy(context.TODO(), input)
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

func describeAsset(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeAssetInput{}
	result, err := client.DescribeAsset(context.TODO(), input)
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

func describeAssetModel(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeAssetModelInput{}
	result, err := client.DescribeAssetModel(context.TODO(), input)
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

func describeAssetProperty(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeAssetPropertyInput{}
	result, err := client.DescribeAssetProperty(context.TODO(), input)
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

func describeBulkImportJob(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeBulkImportJobInput{}
	result, err := client.DescribeBulkImportJob(context.TODO(), input)
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

func describeDashboard(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeDashboardInput{}
	result, err := client.DescribeDashboard(context.TODO(), input)
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

func describeDefaultEncryptionConfiguration(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeDefaultEncryptionConfigurationInput{}
	result, err := client.DescribeDefaultEncryptionConfiguration(context.TODO(), input)
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

func describeGateway(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeGatewayInput{}
	result, err := client.DescribeGateway(context.TODO(), input)
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

func describeGatewayCapabilityConfiguration(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeGatewayCapabilityConfigurationInput{}
	result, err := client.DescribeGatewayCapabilityConfiguration(context.TODO(), input)
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

func describeLoggingOptions(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeLoggingOptionsInput{}
	result, err := client.DescribeLoggingOptions(context.TODO(), input)
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

func describePortal(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribePortalInput{}
	result, err := client.DescribePortal(context.TODO(), input)
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

func describeProject(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeProjectInput{}
	result, err := client.DescribeProject(context.TODO(), input)
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

func describeStorageConfiguration(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeStorageConfigurationInput{}
	result, err := client.DescribeStorageConfiguration(context.TODO(), input)
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

func describeTimeSeries(client *iotsitewise.Client) map[string]interface{} {
	input := &iotsitewise.DescribeTimeSeriesInput{}
	result, err := client.DescribeTimeSeries(context.TODO(), input)
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

