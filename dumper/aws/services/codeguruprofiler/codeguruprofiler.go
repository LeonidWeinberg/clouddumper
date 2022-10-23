package codeguruprofiler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := codeguruprofiler.NewFromConfig(awscfg)
    
	out(describeProfilingGroup(client), "codeguruprofiler", "ProfilingGroup")
    
}

func describeProfilingGroup(client *codeguruprofiler.Client) map[string]interface{} {
	input := &codeguruprofiler.DescribeProfilingGroupInput{}
	result, err := client.DescribeProfilingGroup(context.TODO(), input)
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

