package networkmanager

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := networkmanager.NewFromConfig(awscfg)
    
	out(describeGlobalNetworks(client), "networkmanager", "GlobalNetworks")
    
}

func describeGlobalNetworks(client *networkmanager.Client) map[string]interface{} {
	input := &networkmanager.DescribeGlobalNetworksInput{}
	result, err := client.DescribeGlobalNetworks(context.TODO(), input)
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

