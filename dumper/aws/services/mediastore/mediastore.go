package mediastore

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mediastore.NewFromConfig(awscfg)
    
	out(describeContainer(client), "mediastore", "Container")
    
}

func describeContainer(client *mediastore.Client) map[string]interface{} {
	input := &mediastore.DescribeContainerInput{}
	result, err := client.DescribeContainer(context.TODO(), input)
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

