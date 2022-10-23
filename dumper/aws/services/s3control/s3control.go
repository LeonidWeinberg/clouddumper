package s3control

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := s3control.NewFromConfig(awscfg)
    
	out(describeJob(client), "s3control", "Job")
    
	out(describeMultiRegionAccessPointOperation(client), "s3control", "MultiRegionAccessPointOperation")
    
}

func describeJob(client *s3control.Client) map[string]interface{} {
	input := &s3control.DescribeJobInput{}
	result, err := client.DescribeJob(context.TODO(), input)
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

func describeMultiRegionAccessPointOperation(client *s3control.Client) map[string]interface{} {
	input := &s3control.DescribeMultiRegionAccessPointOperationInput{}
	result, err := client.DescribeMultiRegionAccessPointOperation(context.TODO(), input)
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

