package shield

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := shield.NewFromConfig(awscfg)
    
	out(describeAttack(client), "shield", "Attack")
    
	out(describeAttackStatistics(client), "shield", "AttackStatistics")
    
	out(describeDRTAccess(client), "shield", "DRTAccess")
    
	out(describeEmergencyContactSettings(client), "shield", "EmergencyContactSettings")
    
	out(describeProtection(client), "shield", "Protection")
    
	out(describeProtectionGroup(client), "shield", "ProtectionGroup")
    
	out(describeSubscription(client), "shield", "Subscription")
    
}

func describeAttack(client *shield.Client) map[string]interface{} {
	input := &shield.DescribeAttackInput{}
	result, err := client.DescribeAttack(context.TODO(), input)
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

func describeAttackStatistics(client *shield.Client) map[string]interface{} {
	input := &shield.DescribeAttackStatisticsInput{}
	result, err := client.DescribeAttackStatistics(context.TODO(), input)
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

func describeDRTAccess(client *shield.Client) map[string]interface{} {
	input := &shield.DescribeDRTAccessInput{}
	result, err := client.DescribeDRTAccess(context.TODO(), input)
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

func describeEmergencyContactSettings(client *shield.Client) map[string]interface{} {
	input := &shield.DescribeEmergencyContactSettingsInput{}
	result, err := client.DescribeEmergencyContactSettings(context.TODO(), input)
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

func describeProtection(client *shield.Client) map[string]interface{} {
	input := &shield.DescribeProtectionInput{}
	result, err := client.DescribeProtection(context.TODO(), input)
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

func describeProtectionGroup(client *shield.Client) map[string]interface{} {
	input := &shield.DescribeProtectionGroupInput{}
	result, err := client.DescribeProtectionGroup(context.TODO(), input)
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

func describeSubscription(client *shield.Client) map[string]interface{} {
	input := &shield.DescribeSubscriptionInput{}
	result, err := client.DescribeSubscription(context.TODO(), input)
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

