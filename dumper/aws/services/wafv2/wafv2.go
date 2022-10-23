package wafv2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := wafv2.NewFromConfig(awscfg)
    
	out(describeManagedRuleGroup(client), "wafv2", "ManagedRuleGroup")
    
}

func describeManagedRuleGroup(client *wafv2.Client) map[string]interface{} {
	input := &wafv2.DescribeManagedRuleGroupInput{}
	result, err := client.DescribeManagedRuleGroup(context.TODO(), input)
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

