package networkfirewall

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := networkfirewall.NewFromConfig(awscfg)
    
	out(describeFirewall(client), "networkfirewall", "Firewall")
    
	out(describeFirewallPolicy(client), "networkfirewall", "FirewallPolicy")
    
	out(describeLoggingConfiguration(client), "networkfirewall", "LoggingConfiguration")
    
	out(describeResourcePolicy(client), "networkfirewall", "ResourcePolicy")
    
	out(describeRuleGroup(client), "networkfirewall", "RuleGroup")
    
	out(describeRuleGroupMetadata(client), "networkfirewall", "RuleGroupMetadata")
    
}

func describeFirewall(client *networkfirewall.Client) map[string]interface{} {
	input := &networkfirewall.DescribeFirewallInput{}
	result, err := client.DescribeFirewall(context.TODO(), input)
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

func describeFirewallPolicy(client *networkfirewall.Client) map[string]interface{} {
	input := &networkfirewall.DescribeFirewallPolicyInput{}
	result, err := client.DescribeFirewallPolicy(context.TODO(), input)
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

func describeLoggingConfiguration(client *networkfirewall.Client) map[string]interface{} {
	input := &networkfirewall.DescribeLoggingConfigurationInput{}
	result, err := client.DescribeLoggingConfiguration(context.TODO(), input)
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

func describeResourcePolicy(client *networkfirewall.Client) map[string]interface{} {
	input := &networkfirewall.DescribeResourcePolicyInput{}
	result, err := client.DescribeResourcePolicy(context.TODO(), input)
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

func describeRuleGroup(client *networkfirewall.Client) map[string]interface{} {
	input := &networkfirewall.DescribeRuleGroupInput{}
	result, err := client.DescribeRuleGroup(context.TODO(), input)
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

func describeRuleGroupMetadata(client *networkfirewall.Client) map[string]interface{} {
	input := &networkfirewall.DescribeRuleGroupMetadataInput{}
	result, err := client.DescribeRuleGroupMetadata(context.TODO(), input)
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

