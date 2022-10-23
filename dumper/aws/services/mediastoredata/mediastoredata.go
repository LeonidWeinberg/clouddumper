package mediastoredata

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mediastoredata.NewFromConfig(awscfg)
    
	out(describeObject(client), "mediastoredata", "Object")
    
}

func describeObject(client *mediastoredata.Client) map[string]interface{} {
	input := &mediastoredata.DescribeObjectInput{}
	result, err := client.DescribeObject(context.TODO(), input)
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

