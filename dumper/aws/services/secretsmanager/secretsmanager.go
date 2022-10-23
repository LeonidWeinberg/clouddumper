package secretsmanager

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := secretsmanager.NewFromConfig(awscfg)
    
	out(describeSecret(client), "secretsmanager", "Secret")
    
}

func describeSecret(client *secretsmanager.Client) map[string]interface{} {
	input := &secretsmanager.DescribeSecretInput{}
	result, err := client.DescribeSecret(context.TODO(), input)
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

