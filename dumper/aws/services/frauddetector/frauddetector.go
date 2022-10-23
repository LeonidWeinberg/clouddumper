package frauddetector

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := frauddetector.NewFromConfig(awscfg)
    
	out(describeDetector(client), "frauddetector", "Detector")
    
	out(describeModelVersions(client), "frauddetector", "ModelVersions")
    
}

func describeDetector(client *frauddetector.Client) map[string]interface{} {
	input := &frauddetector.DescribeDetectorInput{}
	result, err := client.DescribeDetector(context.TODO(), input)
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

func describeModelVersions(client *frauddetector.Client) map[string]interface{} {
	input := &frauddetector.DescribeModelVersionsInput{}
	result, err := client.DescribeModelVersions(context.TODO(), input)
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

