package backup

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := backup.NewFromConfig(awscfg)
    
	out(describeBackupJob(client), "backup", "BackupJob")
    
	out(describeBackupVault(client), "backup", "BackupVault")
    
	out(describeCopyJob(client), "backup", "CopyJob")
    
	out(describeFramework(client), "backup", "Framework")
    
	out(describeGlobalSettings(client), "backup", "GlobalSettings")
    
	out(describeProtectedResource(client), "backup", "ProtectedResource")
    
	out(describeRecoveryPoint(client), "backup", "RecoveryPoint")
    
	out(describeRegionSettings(client), "backup", "RegionSettings")
    
	out(describeReportJob(client), "backup", "ReportJob")
    
	out(describeReportPlan(client), "backup", "ReportPlan")
    
	out(describeRestoreJob(client), "backup", "RestoreJob")
    
}

func describeBackupJob(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeBackupJobInput{}
	result, err := client.DescribeBackupJob(context.TODO(), input)
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

func describeBackupVault(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeBackupVaultInput{}
	result, err := client.DescribeBackupVault(context.TODO(), input)
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

func describeCopyJob(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeCopyJobInput{}
	result, err := client.DescribeCopyJob(context.TODO(), input)
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

func describeFramework(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeFrameworkInput{}
	result, err := client.DescribeFramework(context.TODO(), input)
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

func describeGlobalSettings(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeGlobalSettingsInput{}
	result, err := client.DescribeGlobalSettings(context.TODO(), input)
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

func describeProtectedResource(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeProtectedResourceInput{}
	result, err := client.DescribeProtectedResource(context.TODO(), input)
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

func describeRecoveryPoint(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeRecoveryPointInput{}
	result, err := client.DescribeRecoveryPoint(context.TODO(), input)
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

func describeRegionSettings(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeRegionSettingsInput{}
	result, err := client.DescribeRegionSettings(context.TODO(), input)
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

func describeReportJob(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeReportJobInput{}
	result, err := client.DescribeReportJob(context.TODO(), input)
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

func describeReportPlan(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeReportPlanInput{}
	result, err := client.DescribeReportPlan(context.TODO(), input)
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

func describeRestoreJob(client *backup.Client) map[string]interface{} {
	input := &backup.DescribeRestoreJobInput{}
	result, err := client.DescribeRestoreJob(context.TODO(), input)
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

