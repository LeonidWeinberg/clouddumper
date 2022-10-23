package marketplacecatalog

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := marketplacecatalog.NewFromConfig(awscfg)
    
	out(describeChangeSet(client), "marketplacecatalog", "ChangeSet")
    
	out(describeEntity(client), "marketplacecatalog", "Entity")
    
}

func describeChangeSet(client *marketplacecatalog.Client) map[string]interface{} {
	input := &marketplacecatalog.DescribeChangeSetInput{}
	result, err := client.DescribeChangeSet(context.TODO(), input)
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

func describeEntity(client *marketplacecatalog.Client) map[string]interface{} {
	input := &marketplacecatalog.DescribeEntityInput{}
	result, err := client.DescribeEntity(context.TODO(), input)
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

