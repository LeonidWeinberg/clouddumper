package route53recoverycontrolconfig

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := route53recoverycontrolconfig.NewFromConfig(awscfg)
    
	out(describeCluster(client), "route53recoverycontrolconfig", "Cluster")
    
	out(describeControlPanel(client), "route53recoverycontrolconfig", "ControlPanel")
    
	out(describeRoutingControl(client), "route53recoverycontrolconfig", "RoutingControl")
    
	out(describeSafetyRule(client), "route53recoverycontrolconfig", "SafetyRule")
    
}

func describeCluster(client *route53recoverycontrolconfig.Client) map[string]interface{} {
	input := &route53recoverycontrolconfig.DescribeClusterInput{}
	result, err := client.DescribeCluster(context.TODO(), input)
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

func describeControlPanel(client *route53recoverycontrolconfig.Client) map[string]interface{} {
	input := &route53recoverycontrolconfig.DescribeControlPanelInput{}
	result, err := client.DescribeControlPanel(context.TODO(), input)
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

func describeRoutingControl(client *route53recoverycontrolconfig.Client) map[string]interface{} {
	input := &route53recoverycontrolconfig.DescribeRoutingControlInput{}
	result, err := client.DescribeRoutingControl(context.TODO(), input)
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

func describeSafetyRule(client *route53recoverycontrolconfig.Client) map[string]interface{} {
	input := &route53recoverycontrolconfig.DescribeSafetyRuleInput{}
	result, err := client.DescribeSafetyRule(context.TODO(), input)
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

