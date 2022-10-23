package cognitoidentity

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cognitoidentity.NewFromConfig(awscfg)
    
	out(describeIdentity(client), "cognitoidentity", "Identity")
    
	out(describeIdentityPool(client), "cognitoidentity", "IdentityPool")
    
}

func describeIdentity(client *cognitoidentity.Client) map[string]interface{} {
	input := &cognitoidentity.DescribeIdentityInput{}
	result, err := client.DescribeIdentity(context.TODO(), input)
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

func describeIdentityPool(client *cognitoidentity.Client) map[string]interface{} {
	input := &cognitoidentity.DescribeIdentityPoolInput{}
	result, err := client.DescribeIdentityPool(context.TODO(), input)
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

