package pi

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pi"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := pi.NewFromConfig(awscfg)
    
	out(describeDimensionKeys(client), "pi", "DimensionKeys")
    
}

func describeDimensionKeys(client *pi.Client) map[string]interface{} {
	input := &pi.DescribeDimensionKeysInput{}
	result, err := client.DescribeDimensionKeys(context.TODO(), input)
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

