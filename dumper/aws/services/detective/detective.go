package detective

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/detective"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := detective.NewFromConfig(awscfg)
    
	out(describeOrganizationConfiguration(client), "detective", "OrganizationConfiguration")
    
}

func describeOrganizationConfiguration(client *detective.Client) map[string]interface{} {
	input := &detective.DescribeOrganizationConfigurationInput{}
	result, err := client.DescribeOrganizationConfiguration(context.TODO(), input)
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

