package eks

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := eks.NewFromConfig(awscfg)
    
	out(describeAddon(client), "eks", "Addon")
    
	out(describeAddonVersions(client), "eks", "AddonVersions")
    
	out(describeCluster(client), "eks", "Cluster")
    
	out(describeFargateProfile(client), "eks", "FargateProfile")
    
	out(describeIdentityProviderConfig(client), "eks", "IdentityProviderConfig")
    
	out(describeNodegroup(client), "eks", "Nodegroup")
    
	out(describeUpdate(client), "eks", "Update")
    
}

func describeAddon(client *eks.Client) map[string]interface{} {
	input := &eks.DescribeAddonInput{}
	result, err := client.DescribeAddon(context.TODO(), input)
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

func describeAddonVersions(client *eks.Client) map[string]interface{} {
	input := &eks.DescribeAddonVersionsInput{}
	result, err := client.DescribeAddonVersions(context.TODO(), input)
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

func describeCluster(client *eks.Client) map[string]interface{} {
	input := &eks.DescribeClusterInput{}
	result, err := client.DescribeCluster(context.TODO(), input)
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

func describeFargateProfile(client *eks.Client) map[string]interface{} {
	input := &eks.DescribeFargateProfileInput{}
	result, err := client.DescribeFargateProfile(context.TODO(), input)
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

func describeIdentityProviderConfig(client *eks.Client) map[string]interface{} {
	input := &eks.DescribeIdentityProviderConfigInput{}
	result, err := client.DescribeIdentityProviderConfig(context.TODO(), input)
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

func describeNodegroup(client *eks.Client) map[string]interface{} {
	input := &eks.DescribeNodegroupInput{}
	result, err := client.DescribeNodegroup(context.TODO(), input)
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

func describeUpdate(client *eks.Client) map[string]interface{} {
	input := &eks.DescribeUpdateInput{}
	result, err := client.DescribeUpdate(context.TODO(), input)
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

