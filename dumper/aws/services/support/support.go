package support

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := support.NewFromConfig(awscfg)
    
	out(describeAttachment(client), "support", "Attachment")
    
	out(describeCases(client), "support", "Cases")
    
	out(describeCommunications(client), "support", "Communications")
    
	out(describeServices(client), "support", "Services")
    
	out(describeSeverityLevels(client), "support", "SeverityLevels")
    
	out(describeTrustedAdvisorCheckRefreshStatuses(client), "support", "TrustedAdvisorCheckRefreshStatuses")
    
	out(describeTrustedAdvisorCheckResult(client), "support", "TrustedAdvisorCheckResult")
    
	out(describeTrustedAdvisorCheckSummaries(client), "support", "TrustedAdvisorCheckSummaries")
    
	out(describeTrustedAdvisorChecks(client), "support", "TrustedAdvisorChecks")
    
}

func describeAttachment(client *support.Client) map[string]interface{} {
	input := &support.DescribeAttachmentInput{}
	result, err := client.DescribeAttachment(context.TODO(), input)
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

func describeCases(client *support.Client) map[string]interface{} {
	input := &support.DescribeCasesInput{}
	result, err := client.DescribeCases(context.TODO(), input)
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

func describeCommunications(client *support.Client) map[string]interface{} {
	input := &support.DescribeCommunicationsInput{}
	result, err := client.DescribeCommunications(context.TODO(), input)
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

func describeServices(client *support.Client) map[string]interface{} {
	input := &support.DescribeServicesInput{}
	result, err := client.DescribeServices(context.TODO(), input)
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

func describeSeverityLevels(client *support.Client) map[string]interface{} {
	input := &support.DescribeSeverityLevelsInput{}
	result, err := client.DescribeSeverityLevels(context.TODO(), input)
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

func describeTrustedAdvisorCheckRefreshStatuses(client *support.Client) map[string]interface{} {
	input := &support.DescribeTrustedAdvisorCheckRefreshStatusesInput{}
	result, err := client.DescribeTrustedAdvisorCheckRefreshStatuses(context.TODO(), input)
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

func describeTrustedAdvisorCheckResult(client *support.Client) map[string]interface{} {
	input := &support.DescribeTrustedAdvisorCheckResultInput{}
	result, err := client.DescribeTrustedAdvisorCheckResult(context.TODO(), input)
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

func describeTrustedAdvisorCheckSummaries(client *support.Client) map[string]interface{} {
	input := &support.DescribeTrustedAdvisorCheckSummariesInput{}
	result, err := client.DescribeTrustedAdvisorCheckSummaries(context.TODO(), input)
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

func describeTrustedAdvisorChecks(client *support.Client) map[string]interface{} {
	input := &support.DescribeTrustedAdvisorChecksInput{}
	result, err := client.DescribeTrustedAdvisorChecks(context.TODO(), input)
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

