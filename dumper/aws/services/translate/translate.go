package translate

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/translate"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := translate.NewFromConfig(awscfg)
    
	out(describeTextTranslationJob(client), "translate", "TextTranslationJob")
    
}

func describeTextTranslationJob(client *translate.Client) map[string]interface{} {
	input := &translate.DescribeTextTranslationJobInput{}
	result, err := client.DescribeTextTranslationJob(context.TODO(), input)
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

