package synthetics

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/synthetics"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := synthetics.NewFromConfig(awscfg)
    
	out(describeCanaries(client), "synthetics", "Canaries")
    
	out(describeCanariesLastRun(client), "synthetics", "CanariesLastRun")
    
	out(describeRuntimeVersions(client), "synthetics", "RuntimeVersions")
    
}

func describeCanaries(client *synthetics.Client) map[string]interface{} {
	input := &synthetics.DescribeCanariesInput{}
	result, err := client.DescribeCanaries(context.TODO(), input)
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

func describeCanariesLastRun(client *synthetics.Client) map[string]interface{} {
	input := &synthetics.DescribeCanariesLastRunInput{}
	result, err := client.DescribeCanariesLastRun(context.TODO(), input)
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

func describeRuntimeVersions(client *synthetics.Client) map[string]interface{} {
	input := &synthetics.DescribeRuntimeVersionsInput{}
	result, err := client.DescribeRuntimeVersions(context.TODO(), input)
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

