package polly

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/polly"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := polly.NewFromConfig(awscfg)
    
	out(describeVoices(client), "polly", "Voices")
    
}

func describeVoices(client *polly.Client) map[string]interface{} {
	input := &polly.DescribeVoicesInput{}
	result, err := client.DescribeVoices(context.TODO(), input)
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

