package chimesdkmessaging

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := chimesdkmessaging.NewFromConfig(awscfg)
    
	out(describeChannel(client), "chimesdkmessaging", "Channel")
    
	out(describeChannelBan(client), "chimesdkmessaging", "ChannelBan")
    
	out(describeChannelFlow(client), "chimesdkmessaging", "ChannelFlow")
    
	out(describeChannelMembership(client), "chimesdkmessaging", "ChannelMembership")
    
	out(describeChannelMembershipForAppInstanceUser(client), "chimesdkmessaging", "ChannelMembershipForAppInstanceUser")
    
	out(describeChannelModeratedByAppInstanceUser(client), "chimesdkmessaging", "ChannelModeratedByAppInstanceUser")
    
	out(describeChannelModerator(client), "chimesdkmessaging", "ChannelModerator")
    
}

func describeChannel(client *chimesdkmessaging.Client) map[string]interface{} {
	input := &chimesdkmessaging.DescribeChannelInput{}
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

func describeChannelBan(client *chimesdkmessaging.Client) map[string]interface{} {
	input := &chimesdkmessaging.DescribeChannelBanInput{}
	result, err := client.DescribeChannelBan(context.TODO(), input)
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

func describeChannelFlow(client *chimesdkmessaging.Client) map[string]interface{} {
	input := &chimesdkmessaging.DescribeChannelFlowInput{}
	result, err := client.DescribeChannelFlow(context.TODO(), input)
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

func describeChannelMembership(client *chimesdkmessaging.Client) map[string]interface{} {
	input := &chimesdkmessaging.DescribeChannelMembershipInput{}
	result, err := client.DescribeChannelMembership(context.TODO(), input)
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

func describeChannelMembershipForAppInstanceUser(client *chimesdkmessaging.Client) map[string]interface{} {
	input := &chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput{}
	result, err := client.DescribeChannelMembershipForAppInstanceUser(context.TODO(), input)
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

func describeChannelModeratedByAppInstanceUser(client *chimesdkmessaging.Client) map[string]interface{} {
	input := &chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput{}
	result, err := client.DescribeChannelModeratedByAppInstanceUser(context.TODO(), input)
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

func describeChannelModerator(client *chimesdkmessaging.Client) map[string]interface{} {
	input := &chimesdkmessaging.DescribeChannelModeratorInput{}
	result, err := client.DescribeChannelModerator(context.TODO(), input)
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

