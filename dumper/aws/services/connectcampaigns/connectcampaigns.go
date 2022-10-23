package connectcampaigns

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connectcampaigns"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := connectcampaigns.NewFromConfig(awscfg)
    
	out(describeCampaign(client), "connectcampaigns", "Campaign")
    
}

func describeCampaign(client *connectcampaigns.Client) map[string]interface{} {
	input := &connectcampaigns.DescribeCampaignInput{}
	result, err := client.DescribeCampaign(context.TODO(), input)
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

