package resourcegroupstaggingapi

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := resourcegroupstaggingapi.NewFromConfig(awscfg)
    
	out(describeReportCreation(client), "resourcegroupstaggingapi", "ReportCreation")
    
}

func describeReportCreation(client *resourcegroupstaggingapi.Client) map[string]interface{} {
	input := &resourcegroupstaggingapi.DescribeReportCreationInput{}
	result, err := client.DescribeReportCreation(context.TODO(), input)
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

