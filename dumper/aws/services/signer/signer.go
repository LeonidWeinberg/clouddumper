package signer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/signer"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := signer.NewFromConfig(awscfg)
    
	out(describeSigningJob(client), "signer", "SigningJob")
    
}

func describeSigningJob(client *signer.Client) map[string]interface{} {
	input := &signer.DescribeSigningJobInput{}
	result, err := client.DescribeSigningJob(context.TODO(), input)
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

