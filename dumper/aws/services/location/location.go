package location

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/location"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := location.NewFromConfig(awscfg)
    
	out(describeGeofenceCollection(client), "location", "GeofenceCollection")
    
	out(describeMap(client), "location", "Map")
    
	out(describePlaceIndex(client), "location", "PlaceIndex")
    
	out(describeRouteCalculator(client), "location", "RouteCalculator")
    
	out(describeTracker(client), "location", "Tracker")
    
}

func describeGeofenceCollection(client *location.Client) map[string]interface{} {
	input := &location.DescribeGeofenceCollectionInput{}
	result, err := client.DescribeGeofenceCollection(context.TODO(), input)
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

func describeMap(client *location.Client) map[string]interface{} {
	input := &location.DescribeMapInput{}
	result, err := client.DescribeMap(context.TODO(), input)
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

func describePlaceIndex(client *location.Client) map[string]interface{} {
	input := &location.DescribePlaceIndexInput{}
	result, err := client.DescribePlaceIndex(context.TODO(), input)
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

func describeRouteCalculator(client *location.Client) map[string]interface{} {
	input := &location.DescribeRouteCalculatorInput{}
	result, err := client.DescribeRouteCalculator(context.TODO(), input)
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

func describeTracker(client *location.Client) map[string]interface{} {
	input := &location.DescribeTrackerInput{}
	result, err := client.DescribeTracker(context.TODO(), input)
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

