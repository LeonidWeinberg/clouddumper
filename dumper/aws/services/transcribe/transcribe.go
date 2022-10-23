package transcribe

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := transcribe.NewFromConfig(awscfg)
    
	out(describeLanguageModel(client), "transcribe", "LanguageModel")
    
}

func describeLanguageModel(client *transcribe.Client) map[string]interface{} {
	input := &transcribe.DescribeLanguageModelInput{}
	result, err := client.DescribeLanguageModel(context.TODO(), input)
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

