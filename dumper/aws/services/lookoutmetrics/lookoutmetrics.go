package lookoutmetrics

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := lookoutmetrics.NewFromConfig(awscfg)
    
	out(describeAlert(client), "lookoutmetrics", "Alert")
    
	out(describeAnomalyDetectionExecutions(client), "lookoutmetrics", "AnomalyDetectionExecutions")
    
	out(describeAnomalyDetector(client), "lookoutmetrics", "AnomalyDetector")
    
	out(describeMetricSet(client), "lookoutmetrics", "MetricSet")
    
}

func describeAlert(client *lookoutmetrics.Client) map[string]interface{} {
	input := &lookoutmetrics.DescribeAlertInput{}
	result, err := client.DescribeAlert(context.TODO(), input)
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

func describeAnomalyDetectionExecutions(client *lookoutmetrics.Client) map[string]interface{} {
	input := &lookoutmetrics.DescribeAnomalyDetectionExecutionsInput{}
	result, err := client.DescribeAnomalyDetectionExecutions(context.TODO(), input)
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

func describeAnomalyDetector(client *lookoutmetrics.Client) map[string]interface{} {
	input := &lookoutmetrics.DescribeAnomalyDetectorInput{}
	result, err := client.DescribeAnomalyDetector(context.TODO(), input)
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

func describeMetricSet(client *lookoutmetrics.Client) map[string]interface{} {
	input := &lookoutmetrics.DescribeMetricSetInput{}
	result, err := client.DescribeMetricSet(context.TODO(), input)
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

