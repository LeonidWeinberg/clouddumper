package opsworkscm

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := opsworkscm.NewFromConfig(awscfg)
    
	out(describeAccountAttributes(client), "opsworkscm", "AccountAttributes")
    
	out(describeBackups(client), "opsworkscm", "Backups")
    
	out(describeEvents(client), "opsworkscm", "Events")
    
	out(describeNodeAssociationStatus(client), "opsworkscm", "NodeAssociationStatus")
    
	out(describeServers(client), "opsworkscm", "Servers")
    
}

func describeAccountAttributes(client *opsworkscm.Client) map[string]interface{} {
	input := &opsworkscm.DescribeAccountAttributesInput{}
	result, err := client.DescribeAccountAttributes(context.TODO(), input)
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

func describeBackups(client *opsworkscm.Client) map[string]interface{} {
	input := &opsworkscm.DescribeBackupsInput{}
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

func describeEvents(client *opsworkscm.Client) map[string]interface{} {
	input := &opsworkscm.DescribeEventsInput{}
	result, err := client.DescribeEvents(context.TODO(), input)
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

func describeNodeAssociationStatus(client *opsworkscm.Client) map[string]interface{} {
	input := &opsworkscm.DescribeNodeAssociationStatusInput{}
	result, err := client.DescribeNodeAssociationStatus(context.TODO(), input)
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

func describeServers(client *opsworkscm.Client) map[string]interface{} {
	input := &opsworkscm.DescribeServersInput{}
	result, err := client.DescribeServers(context.TODO(), input)
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

