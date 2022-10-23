package applicationinsights

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := applicationinsights.NewFromConfig(awscfg)
    
	out(describeApplication(client), "applicationinsights", "Application")
    
	out(describeComponent(client), "applicationinsights", "Component")
    
	out(describeComponentConfiguration(client), "applicationinsights", "ComponentConfiguration")
    
	out(describeComponentConfigurationRecommendation(client), "applicationinsights", "ComponentConfigurationRecommendation")
    
	out(describeLogPattern(client), "applicationinsights", "LogPattern")
    
	out(describeObservation(client), "applicationinsights", "Observation")
    
	out(describeProblem(client), "applicationinsights", "Problem")
    
	out(describeProblemObservations(client), "applicationinsights", "ProblemObservations")
    
}

func describeApplication(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeApplicationInput{}
	result, err := client.DescribeApplication(context.TODO(), input)
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

func describeComponent(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeComponentInput{}
	result, err := client.DescribeComponent(context.TODO(), input)
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

func describeComponentConfiguration(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeComponentConfigurationInput{}
	result, err := client.DescribeComponentConfiguration(context.TODO(), input)
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

func describeComponentConfigurationRecommendation(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeComponentConfigurationRecommendationInput{}
	result, err := client.DescribeComponentConfigurationRecommendation(context.TODO(), input)
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

func describeLogPattern(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeLogPatternInput{}
	result, err := client.DescribeLogPattern(context.TODO(), input)
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

func describeObservation(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeObservationInput{}
	result, err := client.DescribeObservation(context.TODO(), input)
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

func describeProblem(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeProblemInput{}
	result, err := client.DescribeProblem(context.TODO(), input)
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

func describeProblemObservations(client *applicationinsights.Client) map[string]interface{} {
	input := &applicationinsights.DescribeProblemObservationsInput{}
	result, err := client.DescribeProblemObservations(context.TODO(), input)
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

