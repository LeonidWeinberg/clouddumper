package costandusagereportservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := costandusagereportservice.NewFromConfig(awscfg)
    
	out(describeReportDefinitions(client), "costandusagereportservice", "ReportDefinitions")
    
}

func describeReportDefinitions(client *costandusagereportservice.Client) map[string]interface{} {
	input := &costandusagereportservice.DescribeReportDefinitionsInput{}
	result, err := client.DescribeReportDefinitions(context.TODO(), input)
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

