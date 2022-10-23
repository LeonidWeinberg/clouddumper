package costexplorer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := costexplorer.NewFromConfig(awscfg)
    
	out(describeCostCategoryDefinition(client), "costexplorer", "CostCategoryDefinition")
    
}

func describeCostCategoryDefinition(client *costexplorer.Client) map[string]interface{} {
	input := &costexplorer.DescribeCostCategoryDefinitionInput{}
	result, err := client.DescribeCostCategoryDefinition(context.TODO(), input)
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

