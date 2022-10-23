package codegurureviewer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := codegurureviewer.NewFromConfig(awscfg)
    
	out(describeCodeReview(client), "codegurureviewer", "CodeReview")
    
	out(describeRecommendationFeedback(client), "codegurureviewer", "RecommendationFeedback")
    
	out(describeRepositoryAssociation(client), "codegurureviewer", "RepositoryAssociation")
    
}

func describeCodeReview(client *codegurureviewer.Client) map[string]interface{} {
	input := &codegurureviewer.DescribeCodeReviewInput{}
	result, err := client.DescribeCodeReview(context.TODO(), input)
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

func describeRecommendationFeedback(client *codegurureviewer.Client) map[string]interface{} {
	input := &codegurureviewer.DescribeRecommendationFeedbackInput{}
	result, err := client.DescribeRecommendationFeedback(context.TODO(), input)
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

func describeRepositoryAssociation(client *codegurureviewer.Client) map[string]interface{} {
	input := &codegurureviewer.DescribeRepositoryAssociationInput{}
	result, err := client.DescribeRepositoryAssociation(context.TODO(), input)
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

