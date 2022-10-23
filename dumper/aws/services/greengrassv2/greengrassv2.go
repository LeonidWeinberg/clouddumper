package greengrassv2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := greengrassv2.NewFromConfig(awscfg)
    
	out(describeComponent(client), "greengrassv2", "Component")
    
}

func describeComponent(client *greengrassv2.Client) map[string]interface{} {
	input := &greengrassv2.DescribeComponentInput{}
	result, err := client.DescribeComponent(context.TODO(), input)
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

