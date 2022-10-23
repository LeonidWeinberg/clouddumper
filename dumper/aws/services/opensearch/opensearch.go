package opensearch

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearch"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := opensearch.NewFromConfig(awscfg)
    
	out(describeDomain(client), "opensearch", "Domain")
    
	out(describeDomainAutoTunes(client), "opensearch", "DomainAutoTunes")
    
	out(describeDomainChangeProgress(client), "opensearch", "DomainChangeProgress")
    
	out(describeDomainConfig(client), "opensearch", "DomainConfig")
    
	out(describeDomains(client), "opensearch", "Domains")
    
	out(describeInboundConnections(client), "opensearch", "InboundConnections")
    
	out(describeInstanceTypeLimits(client), "opensearch", "InstanceTypeLimits")
    
	out(describeOutboundConnections(client), "opensearch", "OutboundConnections")
    
	out(describePackages(client), "opensearch", "Packages")
    
	out(describeReservedInstanceOfferings(client), "opensearch", "ReservedInstanceOfferings")
    
	out(describeReservedInstances(client), "opensearch", "ReservedInstances")
    
}

func describeDomain(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeDomainInput{}
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

func describeDomainAutoTunes(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeDomainAutoTunesInput{}
	result, err := client.DescribeDomainAutoTunes(context.TODO(), input)
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

func describeDomainChangeProgress(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeDomainChangeProgressInput{}
	result, err := client.DescribeDomainChangeProgress(context.TODO(), input)
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

func describeDomainConfig(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeDomainConfigInput{}
	result, err := client.DescribeDomainConfig(context.TODO(), input)
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

func describeDomains(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeDomainsInput{}
	result, err := client.DescribeDomains(context.TODO(), input)
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

func describeInboundConnections(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeInboundConnectionsInput{}
	result, err := client.DescribeInboundConnections(context.TODO(), input)
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

func describeInstanceTypeLimits(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeInstanceTypeLimitsInput{}
	result, err := client.DescribeInstanceTypeLimits(context.TODO(), input)
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

func describeOutboundConnections(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeOutboundConnectionsInput{}
	result, err := client.DescribeOutboundConnections(context.TODO(), input)
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

func describePackages(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribePackagesInput{}
	result, err := client.DescribePackages(context.TODO(), input)
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

func describeReservedInstanceOfferings(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeReservedInstanceOfferingsInput{}
	result, err := client.DescribeReservedInstanceOfferings(context.TODO(), input)
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

func describeReservedInstances(client *opensearch.Client) map[string]interface{} {
	input := &opensearch.DescribeReservedInstancesInput{}
	result, err := client.DescribeReservedInstances(context.TODO(), input)
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

