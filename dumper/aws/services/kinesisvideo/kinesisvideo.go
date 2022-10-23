package kinesisvideo

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := kinesisvideo.NewFromConfig(awscfg)
    
	out(describeImageGenerationConfiguration(client), "kinesisvideo", "ImageGenerationConfiguration")
    
	out(describeNotificationConfiguration(client), "kinesisvideo", "NotificationConfiguration")
    
	out(describeSignalingChannel(client), "kinesisvideo", "SignalingChannel")
    
	out(describeStream(client), "kinesisvideo", "Stream")
    
}

func describeImageGenerationConfiguration(client *kinesisvideo.Client) map[string]interface{} {
	input := &kinesisvideo.DescribeImageGenerationConfigurationInput{}
	result, err := client.DescribeImageGenerationConfiguration(context.TODO(), input)
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

func describeNotificationConfiguration(client *kinesisvideo.Client) map[string]interface{} {
	input := &kinesisvideo.DescribeNotificationConfigurationInput{}
	result, err := client.DescribeNotificationConfiguration(context.TODO(), input)
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

func describeSignalingChannel(client *kinesisvideo.Client) map[string]interface{} {
	input := &kinesisvideo.DescribeSignalingChannelInput{}
	result, err := client.DescribeSignalingChannel(context.TODO(), input)
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

func describeStream(client *kinesisvideo.Client) map[string]interface{} {
	input := &kinesisvideo.DescribeStreamInput{}
	result, err := client.DescribeStream(context.TODO(), input)
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

