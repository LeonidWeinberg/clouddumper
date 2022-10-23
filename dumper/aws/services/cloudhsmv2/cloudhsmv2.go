package cloudhsmv2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudhsmv2.NewFromConfig(awscfg)
    
	out(describeBackups(client), "cloudhsmv2", "Backups")
    
	out(describeClusters(client), "cloudhsmv2", "Clusters")
    
}

func describeBackups(client *cloudhsmv2.Client) map[string]interface{} {
	input := &cloudhsmv2.DescribeBackupsInput{}
	result, err := client.DescribeBackups(context.TODO(), input)
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

func describeClusters(client *cloudhsmv2.Client) map[string]interface{} {
	input := &cloudhsmv2.DescribeClustersInput{}
	result, err := client.DescribeClusters(context.TODO(), input)
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

