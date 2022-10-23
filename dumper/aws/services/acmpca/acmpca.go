package acmpca

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := acmpca.NewFromConfig(awscfg)
    
	out(describeCertificateAuthority(client), "acmpca", "CertificateAuthority")
    
	out(describeCertificateAuthorityAuditReport(client), "acmpca", "CertificateAuthorityAuditReport")
    
}

func describeCertificateAuthority(client *acmpca.Client) map[string]interface{} {
	input := &acmpca.DescribeCertificateAuthorityInput{}
	result, err := client.DescribeCertificateAuthority(context.TODO(), input)
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

func describeCertificateAuthorityAuditReport(client *acmpca.Client) map[string]interface{} {
	input := &acmpca.DescribeCertificateAuthorityAuditReportInput{}
	result, err := client.DescribeCertificateAuthorityAuditReport(context.TODO(), input)
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

