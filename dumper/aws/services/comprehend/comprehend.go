package comprehend

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := comprehend.NewFromConfig(awscfg)
    
	out(describeDocumentClassificationJob(client), "comprehend", "DocumentClassificationJob")
    
	out(describeDocumentClassifier(client), "comprehend", "DocumentClassifier")
    
	out(describeDominantLanguageDetectionJob(client), "comprehend", "DominantLanguageDetectionJob")
    
	out(describeEndpoint(client), "comprehend", "Endpoint")
    
	out(describeEntitiesDetectionJob(client), "comprehend", "EntitiesDetectionJob")
    
	out(describeEntityRecognizer(client), "comprehend", "EntityRecognizer")
    
	out(describeEventsDetectionJob(client), "comprehend", "EventsDetectionJob")
    
	out(describeKeyPhrasesDetectionJob(client), "comprehend", "KeyPhrasesDetectionJob")
    
	out(describePiiEntitiesDetectionJob(client), "comprehend", "PiiEntitiesDetectionJob")
    
	out(describeResourcePolicy(client), "comprehend", "ResourcePolicy")
    
	out(describeSentimentDetectionJob(client), "comprehend", "SentimentDetectionJob")
    
	out(describeTargetedSentimentDetectionJob(client), "comprehend", "TargetedSentimentDetectionJob")
    
	out(describeTopicsDetectionJob(client), "comprehend", "TopicsDetectionJob")
    
}

func describeDocumentClassificationJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeDocumentClassificationJobInput{}
	result, err := client.DescribeDocumentClassificationJob(context.TODO(), input)
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

func describeDocumentClassifier(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeDocumentClassifierInput{}
	result, err := client.DescribeDocumentClassifier(context.TODO(), input)
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

func describeDominantLanguageDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeDominantLanguageDetectionJobInput{}
	result, err := client.DescribeDominantLanguageDetectionJob(context.TODO(), input)
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

func describeEndpoint(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeEndpointInput{}
	result, err := client.DescribeEndpoint(context.TODO(), input)
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

func describeEntitiesDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeEntitiesDetectionJobInput{}
	result, err := client.DescribeEntitiesDetectionJob(context.TODO(), input)
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

func describeEntityRecognizer(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeEntityRecognizerInput{}
	result, err := client.DescribeEntityRecognizer(context.TODO(), input)
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

func describeEventsDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeEventsDetectionJobInput{}
	result, err := client.DescribeEventsDetectionJob(context.TODO(), input)
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

func describeKeyPhrasesDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeKeyPhrasesDetectionJobInput{}
	result, err := client.DescribeKeyPhrasesDetectionJob(context.TODO(), input)
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

func describePiiEntitiesDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribePiiEntitiesDetectionJobInput{}
	result, err := client.DescribePiiEntitiesDetectionJob(context.TODO(), input)
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

func describeResourcePolicy(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeResourcePolicyInput{}
	result, err := client.DescribeResourcePolicy(context.TODO(), input)
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

func describeSentimentDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeSentimentDetectionJobInput{}
	result, err := client.DescribeSentimentDetectionJob(context.TODO(), input)
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

func describeTargetedSentimentDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeTargetedSentimentDetectionJobInput{}
	result, err := client.DescribeTargetedSentimentDetectionJob(context.TODO(), input)
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

func describeTopicsDetectionJob(client *comprehend.Client) map[string]interface{} {
	input := &comprehend.DescribeTopicsDetectionJobInput{}
	result, err := client.DescribeTopicsDetectionJob(context.TODO(), input)
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

