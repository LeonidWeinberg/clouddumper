package workmail

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := workmail.NewFromConfig(awscfg)
    
	out(describeEmailMonitoringConfiguration(client), "workmail", "EmailMonitoringConfiguration")
    
	out(describeGroup(client), "workmail", "Group")
    
	out(describeInboundDmarcSettings(client), "workmail", "InboundDmarcSettings")
    
	out(describeMailboxExportJob(client), "workmail", "MailboxExportJob")
    
	out(describeOrganization(client), "workmail", "Organization")
    
	out(describeResource(client), "workmail", "Resource")
    
	out(describeUser(client), "workmail", "User")
    
}

func describeEmailMonitoringConfiguration(client *workmail.Client) map[string]interface{} {
	input := &workmail.DescribeEmailMonitoringConfigurationInput{}
	result, err := client.DescribeEmailMonitoringConfiguration(context.TODO(), input)
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

func describeGroup(client *workmail.Client) map[string]interface{} {
	input := &workmail.DescribeGroupInput{}
	result, err := client.DescribeGroup(context.TODO(), input)
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

func describeInboundDmarcSettings(client *workmail.Client) map[string]interface{} {
	input := &workmail.DescribeInboundDmarcSettingsInput{}
	result, err := client.DescribeInboundDmarcSettings(context.TODO(), input)
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

func describeMailboxExportJob(client *workmail.Client) map[string]interface{} {
	input := &workmail.DescribeMailboxExportJobInput{}
	result, err := client.DescribeMailboxExportJob(context.TODO(), input)
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

func describeOrganization(client *workmail.Client) map[string]interface{} {
	input := &workmail.DescribeOrganizationInput{}
	result, err := client.DescribeOrganization(context.TODO(), input)
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

func describeResource(client *workmail.Client) map[string]interface{} {
	input := &workmail.DescribeResourceInput{}
	result, err := client.DescribeResource(context.TODO(), input)
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

func describeUser(client *workmail.Client) map[string]interface{} {
	input := &workmail.DescribeUserInput{}
	result, err := client.DescribeUser(context.TODO(), input)
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

