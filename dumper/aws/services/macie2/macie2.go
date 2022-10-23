package macie2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/macie2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := macie2.NewFromConfig(awscfg)
    
	out(describeBuckets(client), "macie2", "Buckets")
    
	out(describeClassificationJob(client), "macie2", "ClassificationJob")
    
	out(describeOrganizationConfiguration(client), "macie2", "OrganizationConfiguration")
    
}

func describeBuckets(client *macie2.Client) map[string]interface{} {
	input := &macie2.DescribeBucketsInput{}
	result, err := client.DescribeBuckets(context.TODO(), input)
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

func describeClassificationJob(client *macie2.Client) map[string]interface{} {
	input := &macie2.DescribeClassificationJobInput{}
	result, err := client.DescribeClassificationJob(context.TODO(), input)
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

func describeOrganizationConfiguration(client *macie2.Client) map[string]interface{} {
	input := &macie2.DescribeOrganizationConfigurationInput{}
	result, err := client.DescribeOrganizationConfiguration(context.TODO(), input)
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

