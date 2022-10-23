package acm

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := acm.NewFromConfig(awscfg)
    
	out(describeCertificate(client), "acm", "Certificate")
    
}

func describeCertificate(client *acm.Client) map[string]interface{} {
	input := &acm.DescribeCertificateInput{}
	result, err := client.DescribeCertificate(context.TODO(), input)
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

