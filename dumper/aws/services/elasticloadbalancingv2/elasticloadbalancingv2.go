package elasticloadbalancingv2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := elasticloadbalancingv2.NewFromConfig(awscfg)
    
	out(describeAccountLimits(client), "elasticloadbalancingv2", "AccountLimits")
    
	out(describeListenerCertificates(client), "elasticloadbalancingv2", "ListenerCertificates")
    
	out(describeListeners(client), "elasticloadbalancingv2", "Listeners")
    
	out(describeLoadBalancerAttributes(client), "elasticloadbalancingv2", "LoadBalancerAttributes")
    
	out(describeLoadBalancers(client), "elasticloadbalancingv2", "LoadBalancers")
    
	out(describeRules(client), "elasticloadbalancingv2", "Rules")
    
	out(describeSSLPolicies(client), "elasticloadbalancingv2", "SSLPolicies")
    
	out(describeTags(client), "elasticloadbalancingv2", "Tags")
    
	out(describeTargetGroupAttributes(client), "elasticloadbalancingv2", "TargetGroupAttributes")
    
	out(describeTargetGroups(client), "elasticloadbalancingv2", "TargetGroups")
    
	out(describeTargetHealth(client), "elasticloadbalancingv2", "TargetHealth")
    
}

func describeAccountLimits(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeAccountLimitsInput{}
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

func describeListenerCertificates(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeListenerCertificatesInput{}
	result, err := client.DescribeListenerCertificates(context.TODO(), input)
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

func describeListeners(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeListenersInput{}
	result, err := client.DescribeListeners(context.TODO(), input)
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

func describeLoadBalancerAttributes(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeLoadBalancerAttributesInput{}
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

func describeLoadBalancers(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeLoadBalancersInput{}
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

func describeRules(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeRulesInput{}
	result, err := client.DescribeRules(context.TODO(), input)
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

func describeSSLPolicies(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeSSLPoliciesInput{}
	result, err := client.DescribeSSLPolicies(context.TODO(), input)
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

func describeTags(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeTagsInput{}
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

func describeTargetGroupAttributes(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeTargetGroupAttributesInput{}
	result, err := client.DescribeTargetGroupAttributes(context.TODO(), input)
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

func describeTargetGroups(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeTargetGroupsInput{}
	result, err := client.DescribeTargetGroups(context.TODO(), input)
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

func describeTargetHealth(client *elasticloadbalancingv2.Client) map[string]interface{} {
	input := &elasticloadbalancingv2.DescribeTargetHealthInput{}
	result, err := client.DescribeTargetHealth(context.TODO(), input)
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

