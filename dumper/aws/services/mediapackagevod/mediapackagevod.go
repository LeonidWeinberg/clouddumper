package mediapackagevod

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mediapackagevod.NewFromConfig(awscfg)
    
	out(describeAsset(client), "mediapackagevod", "Asset")
    
	out(describePackagingConfiguration(client), "mediapackagevod", "PackagingConfiguration")
    
	out(describePackagingGroup(client), "mediapackagevod", "PackagingGroup")
    
}

func describeAsset(client *mediapackagevod.Client) map[string]interface{} {
	input := &mediapackagevod.DescribeAssetInput{}
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

func describePackagingConfiguration(client *mediapackagevod.Client) map[string]interface{} {
	input := &mediapackagevod.DescribePackagingConfigurationInput{}
	result, err := client.DescribePackagingConfiguration(context.TODO(), input)
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

func describePackagingGroup(client *mediapackagevod.Client) map[string]interface{} {
	input := &mediapackagevod.DescribePackagingGroupInput{}
	result, err := client.DescribePackagingGroup(context.TODO(), input)
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

