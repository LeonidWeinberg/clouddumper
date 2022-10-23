package computeoptimizer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := computeoptimizer.NewFromConfig(awscfg)
    
	out(describeRecommendationExportJobs(client), "computeoptimizer", "RecommendationExportJobs")
    
}

func describeRecommendationExportJobs(client *computeoptimizer.Client) map[string]interface{} {
	input := &computeoptimizer.DescribeRecommendationExportJobsInput{}
	result, err := client.DescribeRecommendationExportJobs(context.TODO(), input)
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

