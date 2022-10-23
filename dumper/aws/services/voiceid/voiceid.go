package voiceid

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/voiceid"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := voiceid.NewFromConfig(awscfg)
    
	out(describeDomain(client), "voiceid", "Domain")
    
	out(describeFraudster(client), "voiceid", "Fraudster")
    
	out(describeFraudsterRegistrationJob(client), "voiceid", "FraudsterRegistrationJob")
    
	out(describeSpeaker(client), "voiceid", "Speaker")
    
	out(describeSpeakerEnrollmentJob(client), "voiceid", "SpeakerEnrollmentJob")
    
}

func describeDomain(client *voiceid.Client) map[string]interface{} {
	input := &voiceid.DescribeDomainInput{}
	result, err := client.DescribeDomain(context.TODO(), input)
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

func describeFraudster(client *voiceid.Client) map[string]interface{} {
	input := &voiceid.DescribeFraudsterInput{}
	result, err := client.DescribeFraudster(context.TODO(), input)
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

func describeFraudsterRegistrationJob(client *voiceid.Client) map[string]interface{} {
	input := &voiceid.DescribeFraudsterRegistrationJobInput{}
	result, err := client.DescribeFraudsterRegistrationJob(context.TODO(), input)
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

func describeSpeaker(client *voiceid.Client) map[string]interface{} {
	input := &voiceid.DescribeSpeakerInput{}
	result, err := client.DescribeSpeaker(context.TODO(), input)
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

func describeSpeakerEnrollmentJob(client *voiceid.Client) map[string]interface{} {
	input := &voiceid.DescribeSpeakerEnrollmentJobInput{}
	result, err := client.DescribeSpeakerEnrollmentJob(context.TODO(), input)
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

