package mediatailor

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := mediatailor.NewFromConfig(awscfg)
    
	out(describeChannel(client), "mediatailor", "Channel")
    
	out(describeLiveSource(client), "mediatailor", "LiveSource")
    
	out(describeProgram(client), "mediatailor", "Program")
    
	out(describeSourceLocation(client), "mediatailor", "SourceLocation")
    
	out(describeVodSource(client), "mediatailor", "VodSource")
    
}

func describeChannel(client *mediatailor.Client) map[string]interface{} {
	input := &mediatailor.DescribeChannelInput{}
	result, err := client.DescribeChannel(context.TODO(), input)
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

func describeLiveSource(client *mediatailor.Client) map[string]interface{} {
	input := &mediatailor.DescribeLiveSourceInput{}
	result, err := client.DescribeLiveSource(context.TODO(), input)
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

func describeProgram(client *mediatailor.Client) map[string]interface{} {
	input := &mediatailor.DescribeProgramInput{}
	result, err := client.DescribeProgram(context.TODO(), input)
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

func describeSourceLocation(client *mediatailor.Client) map[string]interface{} {
	input := &mediatailor.DescribeSourceLocationInput{}
	result, err := client.DescribeSourceLocation(context.TODO(), input)
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

func describeVodSource(client *mediatailor.Client) map[string]interface{} {
	input := &mediatailor.DescribeVodSourceInput{}
	result, err := client.DescribeVodSource(context.TODO(), input)
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

