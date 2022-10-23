package cloudwatch

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudwatch.NewFromConfig(awscfg)
    
	out(describeAlarmHistory(client), "cloudwatch", "AlarmHistory")
    
	out(describeAlarms(client), "cloudwatch", "Alarms")
    
	out(describeAlarmsForMetric(client), "cloudwatch", "AlarmsForMetric")
    
	out(describeAnomalyDetectors(client), "cloudwatch", "AnomalyDetectors")
    
	out(describeInsightRules(client), "cloudwatch", "InsightRules")
    
}

func describeAlarmHistory(client *cloudwatch.Client) map[string]interface{} {
	input := &cloudwatch.DescribeAlarmHistoryInput{}
	result, err := client.DescribeAlarmHistory(context.TODO(), input)
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

func describeAlarms(client *cloudwatch.Client) map[string]interface{} {
	input := &cloudwatch.DescribeAlarmsInput{}
	result, err := client.DescribeAlarms(context.TODO(), input)
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

func describeAlarmsForMetric(client *cloudwatch.Client) map[string]interface{} {
	input := &cloudwatch.DescribeAlarmsForMetricInput{}
	result, err := client.DescribeAlarmsForMetric(context.TODO(), input)
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

func describeAnomalyDetectors(client *cloudwatch.Client) map[string]interface{} {
	input := &cloudwatch.DescribeAnomalyDetectorsInput{}
	result, err := client.DescribeAnomalyDetectors(context.TODO(), input)
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

func describeInsightRules(client *cloudwatch.Client) map[string]interface{} {
	input := &cloudwatch.DescribeInsightRulesInput{}
	result, err := client.DescribeInsightRules(context.TODO(), input)
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

