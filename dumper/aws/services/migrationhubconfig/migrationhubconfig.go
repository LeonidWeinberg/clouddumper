package migrationhubconfig

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := migrationhubconfig.NewFromConfig(awscfg)
    
	out(describeHomeRegionControls(client), "migrationhubconfig", "HomeRegionControls")
    
}

func describeHomeRegionControls(client *migrationhubconfig.Client) map[string]interface{} {
	input := &migrationhubconfig.DescribeHomeRegionControlsInput{}
	result, err := client.DescribeHomeRegionControls(context.TODO(), input)
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

