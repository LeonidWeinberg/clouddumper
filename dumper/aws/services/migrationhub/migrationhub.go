package migrationhub

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := migrationhub.NewFromConfig(awscfg)
    
	out(describeApplicationState(client), "migrationhub", "ApplicationState")
    
	out(describeMigrationTask(client), "migrationhub", "MigrationTask")
    
}

func describeApplicationState(client *migrationhub.Client) map[string]interface{} {
	input := &migrationhub.DescribeApplicationStateInput{}
	result, err := client.DescribeApplicationState(context.TODO(), input)
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

func describeMigrationTask(client *migrationhub.Client) map[string]interface{} {
	input := &migrationhub.DescribeMigrationTaskInput{}
	result, err := client.DescribeMigrationTask(context.TODO(), input)
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

