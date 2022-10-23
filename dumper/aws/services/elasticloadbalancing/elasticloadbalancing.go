package elasticloadbalancing

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := elasticloadbalancing.NewFromConfig(awscfg)
    
	out(describeAccountLimits(client), "elasticloadbalancing", "AccountLimits")
    
	out(describeInstanceHealth(client), "elasticloadbalancing", "InstanceHealth")
    
	out(describeLoadBalancerAttributes(client), "elasticloadbalancing", "LoadBalancerAttributes")
    
	out(describeLoadBalancerPolicies(client), "elasticloadbalancing", "LoadBalancerPolicies")
    
	out(describeLoadBalancerPolicyTypes(client), "elasticloadbalancing", "LoadBalancerPolicyTypes")
    
	out(describeLoadBalancers(client), "elasticloadbalancing", "LoadBalancers")
    
	out(describeTags(client), "elasticloadbalancing", "Tags")
    
}

func describeAccountLimits(client *elasticloadbalancing.Client) map[string]interface{} {
	input := &elasticloadbalancing.DescribeAccountLimitsInput{}
	result, err := client.DescribeAccountLimits(context.TODO(), input)
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

func describeInstanceHealth(client *elasticloadbalancing.Client) map[string]interface{} {
	input := &elasticloadbalancing.DescribeInstanceHealthInput{}
	result, err := client.DescribeInstanceHealth(context.TODO(), input)
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

func describeLoadBalancerAttributes(client *elasticloadbalancing.Client) map[string]interface{} {
	input := &elasticloadbalancing.DescribeLoadBalancerAttributesInput{}
	result, err := client.DescribeLoadBalancerAttributes(context.TODO(), input)
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

func describeLoadBalancerPolicies(client *elasticloadbalancing.Client) map[string]interface{} {
	input := &elasticloadbalancing.DescribeLoadBalancerPoliciesInput{}
	result, err := client.DescribeLoadBalancerPolicies(context.TODO(), input)
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

func describeLoadBalancerPolicyTypes(client *elasticloadbalancing.Client) map[string]interface{} {
	input := &elasticloadbalancing.DescribeLoadBalancerPolicyTypesInput{}
	result, err := client.DescribeLoadBalancerPolicyTypes(context.TODO(), input)
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

func describeLoadBalancers(client *elasticloadbalancing.Client) map[string]interface{} {
	input := &elasticloadbalancing.DescribeLoadBalancersInput{}
	result, err := client.DescribeLoadBalancers(context.TODO(), input)
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

func describeTags(client *elasticloadbalancing.Client) map[string]interface{} {
	input := &elasticloadbalancing.DescribeTagsInput{}
	result, err := client.DescribeTags(context.TODO(), input)
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

