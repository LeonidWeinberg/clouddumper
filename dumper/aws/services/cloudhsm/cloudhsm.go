package cloudhsm

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := cloudhsm.NewFromConfig(awscfg)
    
	out(describeHapg(client), "cloudhsm", "Hapg")
    
	out(describeHsm(client), "cloudhsm", "Hsm")
    
	out(describeLunaClient(client), "cloudhsm", "LunaClient")
    
}

func describeHapg(client *cloudhsm.Client) map[string]interface{} {
	input := &cloudhsm.DescribeHapgInput{}
	result, err := client.DescribeHapg(context.TODO(), input)
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

func describeHsm(client *cloudhsm.Client) map[string]interface{} {
	input := &cloudhsm.DescribeHsmInput{}
	result, err := client.DescribeHsm(context.TODO(), input)
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

func describeLunaClient(client *cloudhsm.Client) map[string]interface{} {
	input := &cloudhsm.DescribeLunaClientInput{}
	result, err := client.DescribeLunaClient(context.TODO(), input)
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

