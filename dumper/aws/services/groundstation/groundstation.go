package groundstation

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/groundstation"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := groundstation.NewFromConfig(awscfg)
    
	out(describeContact(client), "groundstation", "Contact")
    
}

func describeContact(client *groundstation.Client) map[string]interface{} {
	input := &groundstation.DescribeContactInput{}
	result, err := client.DescribeContact(context.TODO(), input)
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

