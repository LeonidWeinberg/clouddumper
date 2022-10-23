package dynamodb

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := dynamodb.NewFromConfig(awscfg)
    
	out(describeBackup(client), "dynamodb", "Backup")
    
	out(describeContinuousBackups(client), "dynamodb", "ContinuousBackups")
    
	out(describeContributorInsights(client), "dynamodb", "ContributorInsights")
    
	out(describeEndpoints(client), "dynamodb", "Endpoints")
    
	out(describeExport(client), "dynamodb", "Export")
    
	out(describeGlobalTable(client), "dynamodb", "GlobalTable")
    
	out(describeGlobalTableSettings(client), "dynamodb", "GlobalTableSettings")
    
	out(describeImport(client), "dynamodb", "Import")
    
	out(describeKinesisStreamingDestination(client), "dynamodb", "KinesisStreamingDestination")
    
	out(describeLimits(client), "dynamodb", "Limits")
    
	out(describeTable(client), "dynamodb", "Table")
    
	out(describeTableReplicaAutoScaling(client), "dynamodb", "TableReplicaAutoScaling")
    
	out(describeTimeToLive(client), "dynamodb", "TimeToLive")
    
}

func describeBackup(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeBackupInput{}
	result, err := client.DescribeBackup(context.TODO(), input)
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

func describeContinuousBackups(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeContinuousBackupsInput{}
	result, err := client.DescribeContinuousBackups(context.TODO(), input)
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

func describeContributorInsights(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeContributorInsightsInput{}
	result, err := client.DescribeContributorInsights(context.TODO(), input)
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

func describeEndpoints(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeEndpointsInput{}
	result, err := client.DescribeEndpoints(context.TODO(), input)
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

func describeExport(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeExportInput{}
	result, err := client.DescribeExport(context.TODO(), input)
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

func describeGlobalTable(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeGlobalTableInput{}
	result, err := client.DescribeGlobalTable(context.TODO(), input)
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

func describeGlobalTableSettings(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeGlobalTableSettingsInput{}
	result, err := client.DescribeGlobalTableSettings(context.TODO(), input)
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

func describeImport(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeImportInput{}
	result, err := client.DescribeImport(context.TODO(), input)
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

func describeKinesisStreamingDestination(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeKinesisStreamingDestinationInput{}
	result, err := client.DescribeKinesisStreamingDestination(context.TODO(), input)
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

func describeLimits(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeLimitsInput{}
	result, err := client.DescribeLimits(context.TODO(), input)
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

func describeTable(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeTableInput{}
	result, err := client.DescribeTable(context.TODO(), input)
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

func describeTableReplicaAutoScaling(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeTableReplicaAutoScalingInput{}
	result, err := client.DescribeTableReplicaAutoScaling(context.TODO(), input)
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

func describeTimeToLive(client *dynamodb.Client) map[string]interface{} {
	input := &dynamodb.DescribeTimeToLiveInput{}
	result, err := client.DescribeTimeToLive(context.TODO(), input)
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

