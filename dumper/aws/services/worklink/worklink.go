package worklink

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := worklink.NewFromConfig(awscfg)
    
	out(describeAuditStreamConfiguration(client), "worklink", "AuditStreamConfiguration")
    
	out(describeCompanyNetworkConfiguration(client), "worklink", "CompanyNetworkConfiguration")
    
	out(describeDevice(client), "worklink", "Device")
    
	out(describeDevicePolicyConfiguration(client), "worklink", "DevicePolicyConfiguration")
    
	out(describeDomain(client), "worklink", "Domain")
    
	out(describeFleetMetadata(client), "worklink", "FleetMetadata")
    
	out(describeIdentityProviderConfiguration(client), "worklink", "IdentityProviderConfiguration")
    
	out(describeWebsiteCertificateAuthority(client), "worklink", "WebsiteCertificateAuthority")
    
}

func describeAuditStreamConfiguration(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeAuditStreamConfigurationInput{}
	result, err := client.DescribeAuditStreamConfiguration(context.TODO(), input)
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

func describeCompanyNetworkConfiguration(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeCompanyNetworkConfigurationInput{}
	result, err := client.DescribeCompanyNetworkConfiguration(context.TODO(), input)
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

func describeDevice(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeDeviceInput{}
	result, err := client.DescribeDevice(context.TODO(), input)
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

func describeDevicePolicyConfiguration(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeDevicePolicyConfigurationInput{}
	result, err := client.DescribeDevicePolicyConfiguration(context.TODO(), input)
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

func describeDomain(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeDomainInput{}
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

func describeFleetMetadata(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeFleetMetadataInput{}
	result, err := client.DescribeFleetMetadata(context.TODO(), input)
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

func describeIdentityProviderConfiguration(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeIdentityProviderConfigurationInput{}
	result, err := client.DescribeIdentityProviderConfiguration(context.TODO(), input)
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

func describeWebsiteCertificateAuthority(client *worklink.Client) map[string]interface{} {
	input := &worklink.DescribeWebsiteCertificateAuthorityInput{}
	result, err := client.DescribeWebsiteCertificateAuthority(context.TODO(), input)
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

