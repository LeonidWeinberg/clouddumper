package sagemakera2iruntime

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := sagemakera2iruntime.NewFromConfig(awscfg)
    
	out(describeHumanLoop(client), "sagemakera2iruntime", "HumanLoop")
    
}

func describeHumanLoop(client *sagemakera2iruntime.Client) map[string]interface{} {
	input := &sagemakera2iruntime.DescribeHumanLoopInput{}
	result, err := client.DescribeHumanLoop(context.TODO(), input)
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

