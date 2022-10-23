package connect

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connect"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := connect.NewFromConfig(awscfg)
    
	out(describeAgentStatus(client), "connect", "AgentStatus")
    
	out(describeContact(client), "connect", "Contact")
    
	out(describeContactFlow(client), "connect", "ContactFlow")
    
	out(describeContactFlowModule(client), "connect", "ContactFlowModule")
    
	out(describeHoursOfOperation(client), "connect", "HoursOfOperation")
    
	out(describeInstance(client), "connect", "Instance")
    
	out(describeInstanceAttribute(client), "connect", "InstanceAttribute")
    
	out(describeInstanceStorageConfig(client), "connect", "InstanceStorageConfig")
    
	out(describePhoneNumber(client), "connect", "PhoneNumber")
    
	out(describeQueue(client), "connect", "Queue")
    
	out(describeQuickConnect(client), "connect", "QuickConnect")
    
	out(describeRoutingProfile(client), "connect", "RoutingProfile")
    
	out(describeSecurityProfile(client), "connect", "SecurityProfile")
    
	out(describeTrafficDistributionGroup(client), "connect", "TrafficDistributionGroup")
    
	out(describeUser(client), "connect", "User")
    
	out(describeUserHierarchyGroup(client), "connect", "UserHierarchyGroup")
    
	out(describeUserHierarchyStructure(client), "connect", "UserHierarchyStructure")
    
	out(describeVocabulary(client), "connect", "Vocabulary")
    
}

func describeAgentStatus(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeAgentStatusInput{}
	result, err := client.DescribeAgentStatus(context.TODO(), input)
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

func describeContact(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeContactInput{}
	result, err := client.DescribeContact(context.TODO(), input)
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

func describeContactFlow(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeContactFlowInput{}
	result, err := client.DescribeContactFlow(context.TODO(), input)
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

func describeContactFlowModule(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeContactFlowModuleInput{}
	result, err := client.DescribeContactFlowModule(context.TODO(), input)
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

func describeHoursOfOperation(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeHoursOfOperationInput{}
	result, err := client.DescribeHoursOfOperation(context.TODO(), input)
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

func describeInstance(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeInstanceInput{}
	result, err := client.DescribeInstance(context.TODO(), input)
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

func describeInstanceAttribute(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeInstanceAttributeInput{}
	result, err := client.DescribeInstanceAttribute(context.TODO(), input)
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

func describeInstanceStorageConfig(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeInstanceStorageConfigInput{}
	result, err := client.DescribeInstanceStorageConfig(context.TODO(), input)
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

func describePhoneNumber(client *connect.Client) map[string]interface{} {
	input := &connect.DescribePhoneNumberInput{}
	result, err := client.DescribePhoneNumber(context.TODO(), input)
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

func describeQueue(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeQueueInput{}
	result, err := client.DescribeQueue(context.TODO(), input)
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

func describeQuickConnect(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeQuickConnectInput{}
	result, err := client.DescribeQuickConnect(context.TODO(), input)
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

func describeRoutingProfile(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeRoutingProfileInput{}
	result, err := client.DescribeRoutingProfile(context.TODO(), input)
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

func describeSecurityProfile(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeSecurityProfileInput{}
	result, err := client.DescribeSecurityProfile(context.TODO(), input)
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

func describeTrafficDistributionGroup(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeTrafficDistributionGroupInput{}
	result, err := client.DescribeTrafficDistributionGroup(context.TODO(), input)
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

func describeUser(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeUserInput{}
	result, err := client.DescribeUser(context.TODO(), input)
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

func describeUserHierarchyGroup(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeUserHierarchyGroupInput{}
	result, err := client.DescribeUserHierarchyGroup(context.TODO(), input)
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

func describeUserHierarchyStructure(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeUserHierarchyStructureInput{}
	result, err := client.DescribeUserHierarchyStructure(context.TODO(), input)
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

func describeVocabulary(client *connect.Client) map[string]interface{} {
	input := &connect.DescribeVocabularyInput{}
	result, err := client.DescribeVocabulary(context.TODO(), input)
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

