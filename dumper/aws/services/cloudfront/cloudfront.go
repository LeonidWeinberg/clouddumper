package cloudfront

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudfront.NewFromConfig(awscfg)
    
	out(describeFunction(client), "cloudfront", "Function")
    
}

func describeFunction(client *cloudfront.Client) map[string]interface{} {
	input := &cloudfront.DescribeFunctionInput{}
	result, err := client.DescribeFunction(context.TODO(), input)
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

