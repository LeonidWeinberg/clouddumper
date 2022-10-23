package opsworks

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := opsworks.NewFromConfig(awscfg)
    
	out(describeAgentVersions(client), "opsworks", "AgentVersions")
    
	out(describeApps(client), "opsworks", "Apps")
    
	out(describeCommands(client), "opsworks", "Commands")
    
	out(describeDeployments(client), "opsworks", "Deployments")
    
	out(describeEcsClusters(client), "opsworks", "EcsClusters")
    
	out(describeElasticIps(client), "opsworks", "ElasticIps")
    
	out(describeElasticLoadBalancers(client), "opsworks", "ElasticLoadBalancers")
    
	out(describeInstances(client), "opsworks", "Instances")
    
	out(describeLayers(client), "opsworks", "Layers")
    
	out(describeLoadBasedAutoScaling(client), "opsworks", "LoadBasedAutoScaling")
    
	out(describeMyUserProfile(client), "opsworks", "MyUserProfile")
    
	out(describeOperatingSystems(client), "opsworks", "OperatingSystems")
    
	out(describePermissions(client), "opsworks", "Permissions")
    
	out(describeRaidArrays(client), "opsworks", "RaidArrays")
    
	out(describeRdsDbInstances(client), "opsworks", "RdsDbInstances")
    
	out(describeServiceErrors(client), "opsworks", "ServiceErrors")
    
	out(describeStackProvisioningParameters(client), "opsworks", "StackProvisioningParameters")
    
	out(describeStackSummary(client), "opsworks", "StackSummary")
    
	out(describeStacks(client), "opsworks", "Stacks")
    
	out(describeTimeBasedAutoScaling(client), "opsworks", "TimeBasedAutoScaling")
    
	out(describeUserProfiles(client), "opsworks", "UserProfiles")
    
	out(describeVolumes(client), "opsworks", "Volumes")
    
}

func describeAgentVersions(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeAgentVersionsInput{}
	result, err := client.DescribeAgentVersions(context.TODO(), input)
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

func describeApps(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeAppsInput{}
	result, err := client.DescribeApps(context.TODO(), input)
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

func describeCommands(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeCommandsInput{}
	result, err := client.DescribeCommands(context.TODO(), input)
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

func describeDeployments(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeDeploymentsInput{}
	result, err := client.DescribeDeployments(context.TODO(), input)
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

func describeEcsClusters(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeEcsClustersInput{}
	result, err := client.DescribeEcsClusters(context.TODO(), input)
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

func describeElasticIps(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeElasticIpsInput{}
	result, err := client.DescribeElasticIps(context.TODO(), input)
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

func describeElasticLoadBalancers(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeElasticLoadBalancersInput{}
	result, err := client.DescribeElasticLoadBalancers(context.TODO(), input)
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

func describeInstances(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeInstancesInput{}
	result, err := client.DescribeInstances(context.TODO(), input)
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

func describeLayers(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeLayersInput{}
	result, err := client.DescribeLayers(context.TODO(), input)
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

func describeLoadBasedAutoScaling(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeLoadBasedAutoScalingInput{}
	result, err := client.DescribeLoadBasedAutoScaling(context.TODO(), input)
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

func describeMyUserProfile(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeMyUserProfileInput{}
	result, err := client.DescribeMyUserProfile(context.TODO(), input)
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

func describeOperatingSystems(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeOperatingSystemsInput{}
	result, err := client.DescribeOperatingSystems(context.TODO(), input)
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

func describePermissions(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribePermissionsInput{}
	result, err := client.DescribePermissions(context.TODO(), input)
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

func describeRaidArrays(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeRaidArraysInput{}
	result, err := client.DescribeRaidArrays(context.TODO(), input)
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

func describeRdsDbInstances(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeRdsDbInstancesInput{}
	result, err := client.DescribeRdsDbInstances(context.TODO(), input)
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

func describeServiceErrors(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeServiceErrorsInput{}
	result, err := client.DescribeServiceErrors(context.TODO(), input)
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

func describeStackProvisioningParameters(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeStackProvisioningParametersInput{}
	result, err := client.DescribeStackProvisioningParameters(context.TODO(), input)
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

func describeStackSummary(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeStackSummaryInput{}
	result, err := client.DescribeStackSummary(context.TODO(), input)
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

func describeStacks(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeStacksInput{}
	result, err := client.DescribeStacks(context.TODO(), input)
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

func describeTimeBasedAutoScaling(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeTimeBasedAutoScalingInput{}
	result, err := client.DescribeTimeBasedAutoScaling(context.TODO(), input)
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

func describeUserProfiles(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeUserProfilesInput{}
	result, err := client.DescribeUserProfiles(context.TODO(), input)
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

func describeVolumes(client *opsworks.Client) map[string]interface{} {
	input := &opsworks.DescribeVolumesInput{}
	result, err := client.DescribeVolumes(context.TODO(), input)
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

