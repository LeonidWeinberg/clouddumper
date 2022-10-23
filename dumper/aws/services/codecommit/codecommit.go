package codecommit

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := codecommit.NewFromConfig(awscfg)
    
	out(describeMergeConflicts(client), "codecommit", "MergeConflicts")
    
	out(describePullRequestEvents(client), "codecommit", "PullRequestEvents")
    
}

func describeMergeConflicts(client *codecommit.Client) map[string]interface{} {
	input := &codecommit.DescribeMergeConflictsInput{}
	result, err := client.DescribeMergeConflicts(context.TODO(), input)
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

func describePullRequestEvents(client *codecommit.Client) map[string]interface{} {
	input := &codecommit.DescribePullRequestEventsInput{}
	result, err := client.DescribePullRequestEvents(context.TODO(), input)
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

