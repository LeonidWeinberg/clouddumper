package pinpointsmsvoicev2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := pinpointsmsvoicev2.NewFromConfig(awscfg)
    
	out(describeAccountAttributes(client), "pinpointsmsvoicev2", "AccountAttributes")
    
	out(describeAccountLimits(client), "pinpointsmsvoicev2", "AccountLimits")
    
	out(describeConfigurationSets(client), "pinpointsmsvoicev2", "ConfigurationSets")
    
	out(describeKeywords(client), "pinpointsmsvoicev2", "Keywords")
    
	out(describeOptOutLists(client), "pinpointsmsvoicev2", "OptOutLists")
    
	out(describeOptedOutNumbers(client), "pinpointsmsvoicev2", "OptedOutNumbers")
    
	out(describePhoneNumbers(client), "pinpointsmsvoicev2", "PhoneNumbers")
    
	out(describePools(client), "pinpointsmsvoicev2", "Pools")
    
	out(describeSenderIds(client), "pinpointsmsvoicev2", "SenderIds")
    
	out(describeSpendLimits(client), "pinpointsmsvoicev2", "SpendLimits")
    
}

func describeAccountAttributes(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeAccountAttributesInput{}
	result, err := client.DescribeAccountAttributes(context.TODO(), input)
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

func describeAccountLimits(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeAccountLimitsInput{}
	result, err := client.DescribeAccountLimits(context.TODO(), input)
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

func describeConfigurationSets(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeConfigurationSetsInput{}
	result, err := client.DescribeConfigurationSets(context.TODO(), input)
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

func describeKeywords(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeKeywordsInput{}
	result, err := client.DescribeKeywords(context.TODO(), input)
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

func describeOptOutLists(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeOptOutListsInput{}
	result, err := client.DescribeOptOutLists(context.TODO(), input)
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

func describeOptedOutNumbers(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeOptedOutNumbersInput{}
	result, err := client.DescribeOptedOutNumbers(context.TODO(), input)
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

func describePhoneNumbers(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribePhoneNumbersInput{}
	result, err := client.DescribePhoneNumbers(context.TODO(), input)
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

func describePools(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribePoolsInput{}
	result, err := client.DescribePools(context.TODO(), input)
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

func describeSenderIds(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeSenderIdsInput{}
	result, err := client.DescribeSenderIds(context.TODO(), input)
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

func describeSpendLimits(client *pinpointsmsvoicev2.Client) map[string]interface{} {
	input := &pinpointsmsvoicev2.DescribeSpendLimitsInput{}
	result, err := client.DescribeSpendLimits(context.TODO(), input)
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

