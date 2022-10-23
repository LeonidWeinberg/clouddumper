package health

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/health"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := health.NewFromConfig(awscfg)
    
	out(describeAffectedAccountsForOrganization(client), "health", "AffectedAccountsForOrganization")
    
	out(describeAffectedEntities(client), "health", "AffectedEntities")
    
	out(describeAffectedEntitiesForOrganization(client), "health", "AffectedEntitiesForOrganization")
    
	out(describeEntityAggregates(client), "health", "EntityAggregates")
    
	out(describeEventAggregates(client), "health", "EventAggregates")
    
	out(describeEventDetails(client), "health", "EventDetails")
    
	out(describeEventDetailsForOrganization(client), "health", "EventDetailsForOrganization")
    
	out(describeEventTypes(client), "health", "EventTypes")
    
	out(describeEvents(client), "health", "Events")
    
	out(describeEventsForOrganization(client), "health", "EventsForOrganization")
    
	out(describeHealthServiceStatusForOrganization(client), "health", "HealthServiceStatusForOrganization")
    
}

func describeAffectedAccountsForOrganization(client *health.Client) map[string]interface{} {
	input := &health.DescribeAffectedAccountsForOrganizationInput{}
	result, err := client.DescribeAffectedAccountsForOrganization(context.TODO(), input)
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

func describeAffectedEntities(client *health.Client) map[string]interface{} {
	input := &health.DescribeAffectedEntitiesInput{}
	result, err := client.DescribeAffectedEntities(context.TODO(), input)
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

func describeAffectedEntitiesForOrganization(client *health.Client) map[string]interface{} {
	input := &health.DescribeAffectedEntitiesForOrganizationInput{}
	result, err := client.DescribeAffectedEntitiesForOrganization(context.TODO(), input)
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

func describeEntityAggregates(client *health.Client) map[string]interface{} {
	input := &health.DescribeEntityAggregatesInput{}
	result, err := client.DescribeEntityAggregates(context.TODO(), input)
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

func describeEventAggregates(client *health.Client) map[string]interface{} {
	input := &health.DescribeEventAggregatesInput{}
	result, err := client.DescribeEventAggregates(context.TODO(), input)
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

func describeEventDetails(client *health.Client) map[string]interface{} {
	input := &health.DescribeEventDetailsInput{}
	result, err := client.DescribeEventDetails(context.TODO(), input)
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

func describeEventDetailsForOrganization(client *health.Client) map[string]interface{} {
	input := &health.DescribeEventDetailsForOrganizationInput{}
	result, err := client.DescribeEventDetailsForOrganization(context.TODO(), input)
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

func describeEventTypes(client *health.Client) map[string]interface{} {
	input := &health.DescribeEventTypesInput{}
	result, err := client.DescribeEventTypes(context.TODO(), input)
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

func describeEvents(client *health.Client) map[string]interface{} {
	input := &health.DescribeEventsInput{}
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

func describeEventsForOrganization(client *health.Client) map[string]interface{} {
	input := &health.DescribeEventsForOrganizationInput{}
	result, err := client.DescribeEventsForOrganization(context.TODO(), input)
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

func describeHealthServiceStatusForOrganization(client *health.Client) map[string]interface{} {
	input := &health.DescribeHealthServiceStatusForOrganizationInput{}
	result, err := client.DescribeHealthServiceStatusForOrganization(context.TODO(), input)
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

