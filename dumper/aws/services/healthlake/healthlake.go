package healthlake

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := healthlake.NewFromConfig(awscfg)
    
	out(describeFHIRDatastore(client), "healthlake", "FHIRDatastore")
    
	out(describeFHIRExportJob(client), "healthlake", "FHIRExportJob")
    
	out(describeFHIRImportJob(client), "healthlake", "FHIRImportJob")
    
}

func describeFHIRDatastore(client *healthlake.Client) map[string]interface{} {
	input := &healthlake.DescribeFHIRDatastoreInput{}
	result, err := client.DescribeFHIRDatastore(context.TODO(), input)
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

func describeFHIRExportJob(client *healthlake.Client) map[string]interface{} {
	input := &healthlake.DescribeFHIRExportJobInput{}
	result, err := client.DescribeFHIRExportJob(context.TODO(), input)
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

func describeFHIRImportJob(client *healthlake.Client) map[string]interface{} {
	input := &healthlake.DescribeFHIRImportJobInput{}
	result, err := client.DescribeFHIRImportJob(context.TODO(), input)
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

