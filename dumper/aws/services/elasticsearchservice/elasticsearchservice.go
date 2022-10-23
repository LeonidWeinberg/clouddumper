package elasticsearchservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := elasticsearchservice.NewFromConfig(awscfg)
    
	out(describeDomainAutoTunes(client), "elasticsearchservice", "DomainAutoTunes")
    
	out(describeDomainChangeProgress(client), "elasticsearchservice", "DomainChangeProgress")
    
	out(describeElasticsearchDomain(client), "elasticsearchservice", "ElasticsearchDomain")
    
	out(describeElasticsearchDomainConfig(client), "elasticsearchservice", "ElasticsearchDomainConfig")
    
	out(describeElasticsearchDomains(client), "elasticsearchservice", "ElasticsearchDomains")
    
	out(describeElasticsearchInstanceTypeLimits(client), "elasticsearchservice", "ElasticsearchInstanceTypeLimits")
    
	out(describeInboundCrossClusterSearchConnections(client), "elasticsearchservice", "InboundCrossClusterSearchConnections")
    
	out(describeOutboundCrossClusterSearchConnections(client), "elasticsearchservice", "OutboundCrossClusterSearchConnections")
    
	out(describePackages(client), "elasticsearchservice", "Packages")
    
	out(describeReservedElasticsearchInstanceOfferings(client), "elasticsearchservice", "ReservedElasticsearchInstanceOfferings")
    
	out(describeReservedElasticsearchInstances(client), "elasticsearchservice", "ReservedElasticsearchInstances")
    
}

func describeDomainAutoTunes(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeDomainAutoTunesInput{}
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

func describeDomainChangeProgress(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeDomainChangeProgressInput{}
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

func describeElasticsearchDomain(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeElasticsearchDomainInput{}
	result, err := client.DescribeElasticsearchDomain(context.TODO(), input)
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

func describeElasticsearchDomainConfig(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeElasticsearchDomainConfigInput{}
	result, err := client.DescribeElasticsearchDomainConfig(context.TODO(), input)
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

func describeElasticsearchDomains(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeElasticsearchDomainsInput{}
	result, err := client.DescribeElasticsearchDomains(context.TODO(), input)
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

func describeElasticsearchInstanceTypeLimits(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput{}
	result, err := client.DescribeElasticsearchInstanceTypeLimits(context.TODO(), input)
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

func describeInboundCrossClusterSearchConnections(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput{}
	result, err := client.DescribeInboundCrossClusterSearchConnections(context.TODO(), input)
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

func describeOutboundCrossClusterSearchConnections(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput{}
	result, err := client.DescribeOutboundCrossClusterSearchConnections(context.TODO(), input)
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

func describePackages(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribePackagesInput{}
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

func describeReservedElasticsearchInstanceOfferings(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput{}
	result, err := client.DescribeReservedElasticsearchInstanceOfferings(context.TODO(), input)
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

func describeReservedElasticsearchInstances(client *elasticsearchservice.Client) map[string]interface{} {
	input := &elasticsearchservice.DescribeReservedElasticsearchInstancesInput{}
	result, err := client.DescribeReservedElasticsearchInstances(context.TODO(), input)
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

