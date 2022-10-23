package mediaconvert

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mediaconvert.NewFromConfig(awscfg)
    
	out(describeEndpoints(client), "mediaconvert", "Endpoints")
    
}

func describeEndpoints(client *mediaconvert.Client) map[string]interface{} {
	input := &mediaconvert.DescribeEndpointsInput{}
	result, err := client.DescribeEndpoints(context.TODO(), input)
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

