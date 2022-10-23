package chime

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chime"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := chime.NewFromConfig(awscfg)
    
	out(describeAppInstance(client), "chime", "AppInstance")
    
	out(describeAppInstanceAdmin(client), "chime", "AppInstanceAdmin")
    
	out(describeAppInstanceUser(client), "chime", "AppInstanceUser")
    
	out(describeChannel(client), "chime", "Channel")
    
	out(describeChannelBan(client), "chime", "ChannelBan")
    
	out(describeChannelMembership(client), "chime", "ChannelMembership")
    
	out(describeChannelMembershipForAppInstanceUser(client), "chime", "ChannelMembershipForAppInstanceUser")
    
	out(describeChannelModeratedByAppInstanceUser(client), "chime", "ChannelModeratedByAppInstanceUser")
    
	out(describeChannelModerator(client), "chime", "ChannelModerator")
    
}

func describeAppInstance(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeAppInstanceInput{}
	result, err := client.DescribeAppInstance(context.TODO(), input)
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

func describeAppInstanceAdmin(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeAppInstanceAdminInput{}
	result, err := client.DescribeAppInstanceAdmin(context.TODO(), input)
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

func describeAppInstanceUser(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeAppInstanceUserInput{}
	result, err := client.DescribeAppInstanceUser(context.TODO(), input)
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

func describeChannel(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeChannelInput{}
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

func describeChannelBan(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeChannelBanInput{}
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

func describeChannelMembership(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeChannelMembershipInput{}
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

func describeChannelMembershipForAppInstanceUser(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeChannelMembershipForAppInstanceUserInput{}
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

func describeChannelModeratedByAppInstanceUser(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeChannelModeratedByAppInstanceUserInput{}
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

func describeChannelModerator(client *chime.Client) map[string]interface{} {
	input := &chime.DescribeChannelModeratorInput{}
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

