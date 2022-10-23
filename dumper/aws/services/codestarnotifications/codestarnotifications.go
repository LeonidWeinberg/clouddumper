package codestarnotifications

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := codestarnotifications.NewFromConfig(awscfg)
    
	out(describeNotificationRule(client), "codestarnotifications", "NotificationRule")
    
}

func describeNotificationRule(client *codestarnotifications.Client) map[string]interface{} {
	input := &codestarnotifications.DescribeNotificationRuleInput{}
	result, err := client.DescribeNotificationRule(context.TODO(), input)
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

