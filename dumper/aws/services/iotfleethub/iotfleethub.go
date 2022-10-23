package iotfleethub

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotfleethub"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iotfleethub.NewFromConfig(awscfg)
    
	out(describeApplication(client), "iotfleethub", "Application")
    
}

func describeApplication(client *iotfleethub.Client) map[string]interface{} {
	input := &iotfleethub.DescribeApplicationInput{}
	result, err := client.DescribeApplication(context.TODO(), input)
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

