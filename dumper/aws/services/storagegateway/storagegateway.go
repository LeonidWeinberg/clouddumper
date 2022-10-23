package storagegateway

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := storagegateway.NewFromConfig(awscfg)
    
	out(describeAvailabilityMonitorTest(client), "storagegateway", "AvailabilityMonitorTest")
    
	out(describeBandwidthRateLimit(client), "storagegateway", "BandwidthRateLimit")
    
	out(describeBandwidthRateLimitSchedule(client), "storagegateway", "BandwidthRateLimitSchedule")
    
	out(describeCache(client), "storagegateway", "Cache")
    
	out(describeCachediSCSIVolumes(client), "storagegateway", "CachediSCSIVolumes")
    
	out(describeChapCredentials(client), "storagegateway", "ChapCredentials")
    
	out(describeFileSystemAssociations(client), "storagegateway", "FileSystemAssociations")
    
	out(describeGatewayInformation(client), "storagegateway", "GatewayInformation")
    
	out(describeMaintenanceStartTime(client), "storagegateway", "MaintenanceStartTime")
    
	out(describeNFSFileShares(client), "storagegateway", "NFSFileShares")
    
	out(describeSMBFileShares(client), "storagegateway", "SMBFileShares")
    
	out(describeSMBSettings(client), "storagegateway", "SMBSettings")
    
	out(describeSnapshotSchedule(client), "storagegateway", "SnapshotSchedule")
    
	out(describeStorediSCSIVolumes(client), "storagegateway", "StorediSCSIVolumes")
    
	out(describeTapeArchives(client), "storagegateway", "TapeArchives")
    
	out(describeTapeRecoveryPoints(client), "storagegateway", "TapeRecoveryPoints")
    
	out(describeTapes(client), "storagegateway", "Tapes")
    
	out(describeUploadBuffer(client), "storagegateway", "UploadBuffer")
    
	out(describeVTLDevices(client), "storagegateway", "VTLDevices")
    
	out(describeWorkingStorage(client), "storagegateway", "WorkingStorage")
    
}

func describeAvailabilityMonitorTest(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeAvailabilityMonitorTestInput{}
	result, err := client.DescribeAvailabilityMonitorTest(context.TODO(), input)
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

func describeBandwidthRateLimit(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeBandwidthRateLimitInput{}
	result, err := client.DescribeBandwidthRateLimit(context.TODO(), input)
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

func describeBandwidthRateLimitSchedule(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeBandwidthRateLimitScheduleInput{}
	result, err := client.DescribeBandwidthRateLimitSchedule(context.TODO(), input)
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

func describeCache(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeCacheInput{}
	result, err := client.DescribeCache(context.TODO(), input)
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

func describeCachediSCSIVolumes(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeCachediSCSIVolumesInput{}
	result, err := client.DescribeCachediSCSIVolumes(context.TODO(), input)
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

func describeChapCredentials(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeChapCredentialsInput{}
	result, err := client.DescribeChapCredentials(context.TODO(), input)
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

func describeFileSystemAssociations(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeFileSystemAssociationsInput{}
	result, err := client.DescribeFileSystemAssociations(context.TODO(), input)
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

func describeGatewayInformation(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeGatewayInformationInput{}
	result, err := client.DescribeGatewayInformation(context.TODO(), input)
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

func describeMaintenanceStartTime(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeMaintenanceStartTimeInput{}
	result, err := client.DescribeMaintenanceStartTime(context.TODO(), input)
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

func describeNFSFileShares(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeNFSFileSharesInput{}
	result, err := client.DescribeNFSFileShares(context.TODO(), input)
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

func describeSMBFileShares(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeSMBFileSharesInput{}
	result, err := client.DescribeSMBFileShares(context.TODO(), input)
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

func describeSMBSettings(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeSMBSettingsInput{}
	result, err := client.DescribeSMBSettings(context.TODO(), input)
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

func describeSnapshotSchedule(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeSnapshotScheduleInput{}
	result, err := client.DescribeSnapshotSchedule(context.TODO(), input)
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

func describeStorediSCSIVolumes(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeStorediSCSIVolumesInput{}
	result, err := client.DescribeStorediSCSIVolumes(context.TODO(), input)
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

func describeTapeArchives(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeTapeArchivesInput{}
	result, err := client.DescribeTapeArchives(context.TODO(), input)
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

func describeTapeRecoveryPoints(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeTapeRecoveryPointsInput{}
	result, err := client.DescribeTapeRecoveryPoints(context.TODO(), input)
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

func describeTapes(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeTapesInput{}
	result, err := client.DescribeTapes(context.TODO(), input)
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

func describeUploadBuffer(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeUploadBufferInput{}
	result, err := client.DescribeUploadBuffer(context.TODO(), input)
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

func describeVTLDevices(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeVTLDevicesInput{}
	result, err := client.DescribeVTLDevices(context.TODO(), input)
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

func describeWorkingStorage(client *storagegateway.Client) map[string]interface{} {
	input := &storagegateway.DescribeWorkingStorageInput{}
	result, err := client.DescribeWorkingStorage(context.TODO(), input)
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

