package firehose

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := firehose.NewFromConfig(awscfg)
    
	out(describeDeliveryStream(client), "firehose", "DeliveryStream")
    
}

func describeDeliveryStream(client *firehose.Client) map[string]interface{} {
	input := &firehose.DescribeDeliveryStreamInput{}
	result, err := client.DescribeDeliveryStream(context.TODO(), input)
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

