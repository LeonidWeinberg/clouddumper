package devopsguru

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := devopsguru.NewFromConfig(awscfg)
    
	out(describeAccountHealth(client), "devopsguru", "AccountHealth")
    
	out(describeAccountOverview(client), "devopsguru", "AccountOverview")
    
	out(describeAnomaly(client), "devopsguru", "Anomaly")
    
	out(describeEventSourcesConfig(client), "devopsguru", "EventSourcesConfig")
    
	out(describeFeedback(client), "devopsguru", "Feedback")
    
	out(describeInsight(client), "devopsguru", "Insight")
    
	out(describeOrganizationHealth(client), "devopsguru", "OrganizationHealth")
    
	out(describeOrganizationOverview(client), "devopsguru", "OrganizationOverview")
    
	out(describeOrganizationResourceCollectionHealth(client), "devopsguru", "OrganizationResourceCollectionHealth")
    
	out(describeResourceCollectionHealth(client), "devopsguru", "ResourceCollectionHealth")
    
	out(describeServiceIntegration(client), "devopsguru", "ServiceIntegration")
    
}

func describeAccountHealth(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeAccountHealthInput{}
	result, err := client.DescribeAccountHealth(context.TODO(), input)
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

func describeAccountOverview(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeAccountOverviewInput{}
	result, err := client.DescribeAccountOverview(context.TODO(), input)
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

func describeAnomaly(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeAnomalyInput{}
	result, err := client.DescribeAnomaly(context.TODO(), input)
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

func describeEventSourcesConfig(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeEventSourcesConfigInput{}
	result, err := client.DescribeEventSourcesConfig(context.TODO(), input)
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

func describeFeedback(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeFeedbackInput{}
	result, err := client.DescribeFeedback(context.TODO(), input)
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

func describeInsight(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeInsightInput{}
	result, err := client.DescribeInsight(context.TODO(), input)
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

func describeOrganizationHealth(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeOrganizationHealthInput{}
	result, err := client.DescribeOrganizationHealth(context.TODO(), input)
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

func describeOrganizationOverview(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeOrganizationOverviewInput{}
	result, err := client.DescribeOrganizationOverview(context.TODO(), input)
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

func describeOrganizationResourceCollectionHealth(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeOrganizationResourceCollectionHealthInput{}
	result, err := client.DescribeOrganizationResourceCollectionHealth(context.TODO(), input)
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

func describeResourceCollectionHealth(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeResourceCollectionHealthInput{}
	result, err := client.DescribeResourceCollectionHealth(context.TODO(), input)
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

func describeServiceIntegration(client *devopsguru.Client) map[string]interface{} {
	input := &devopsguru.DescribeServiceIntegrationInput{}
	result, err := client.DescribeServiceIntegration(context.TODO(), input)
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

