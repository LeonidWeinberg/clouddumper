package comprehendmedical

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := comprehendmedical.NewFromConfig(awscfg)
    
	out(describeEntitiesDetectionV2Job(client), "comprehendmedical", "EntitiesDetectionV2Job")
    
	out(describeICD10CMInferenceJob(client), "comprehendmedical", "ICD10CMInferenceJob")
    
	out(describePHIDetectionJob(client), "comprehendmedical", "PHIDetectionJob")
    
	out(describeRxNormInferenceJob(client), "comprehendmedical", "RxNormInferenceJob")
    
	out(describeSNOMEDCTInferenceJob(client), "comprehendmedical", "SNOMEDCTInferenceJob")
    
}

func describeEntitiesDetectionV2Job(client *comprehendmedical.Client) map[string]interface{} {
	input := &comprehendmedical.DescribeEntitiesDetectionV2JobInput{}
	result, err := client.DescribeEntitiesDetectionV2Job(context.TODO(), input)
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

func describeICD10CMInferenceJob(client *comprehendmedical.Client) map[string]interface{} {
	input := &comprehendmedical.DescribeICD10CMInferenceJobInput{}
	result, err := client.DescribeICD10CMInferenceJob(context.TODO(), input)
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

func describePHIDetectionJob(client *comprehendmedical.Client) map[string]interface{} {
	input := &comprehendmedical.DescribePHIDetectionJobInput{}
	result, err := client.DescribePHIDetectionJob(context.TODO(), input)
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

func describeRxNormInferenceJob(client *comprehendmedical.Client) map[string]interface{} {
	input := &comprehendmedical.DescribeRxNormInferenceJobInput{}
	result, err := client.DescribeRxNormInferenceJob(context.TODO(), input)
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

func describeSNOMEDCTInferenceJob(client *comprehendmedical.Client) map[string]interface{} {
	input := &comprehendmedical.DescribeSNOMEDCTInferenceJobInput{}
	result, err := client.DescribeSNOMEDCTInferenceJob(context.TODO(), input)
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

