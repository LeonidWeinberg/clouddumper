package iotsecuretunneling

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iotsecuretunneling.NewFromConfig(awscfg)
    
	out(describeTunnel(client), "iotsecuretunneling", "Tunnel")
    
}

func describeTunnel(client *iotsecuretunneling.Client) map[string]interface{} {
	input := &iotsecuretunneling.DescribeTunnelInput{}
	result, err := client.DescribeTunnel(context.TODO(), input)
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

