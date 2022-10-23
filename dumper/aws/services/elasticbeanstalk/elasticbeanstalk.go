package elasticbeanstalk

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := elasticbeanstalk.NewFromConfig(awscfg)
    
	out(describeAccountAttributes(client), "elasticbeanstalk", "AccountAttributes")
    
	out(describeApplicationVersions(client), "elasticbeanstalk", "ApplicationVersions")
    
	out(describeApplications(client), "elasticbeanstalk", "Applications")
    
	out(describeConfigurationOptions(client), "elasticbeanstalk", "ConfigurationOptions")
    
	out(describeConfigurationSettings(client), "elasticbeanstalk", "ConfigurationSettings")
    
	out(describeEnvironmentHealth(client), "elasticbeanstalk", "EnvironmentHealth")
    
	out(describeEnvironmentManagedActionHistory(client), "elasticbeanstalk", "EnvironmentManagedActionHistory")
    
	out(describeEnvironmentManagedActions(client), "elasticbeanstalk", "EnvironmentManagedActions")
    
	out(describeEnvironmentResources(client), "elasticbeanstalk", "EnvironmentResources")
    
	out(describeEnvironments(client), "elasticbeanstalk", "Environments")
    
	out(describeEvents(client), "elasticbeanstalk", "Events")
    
	out(describeInstancesHealth(client), "elasticbeanstalk", "InstancesHealth")
    
	out(describePlatformVersion(client), "elasticbeanstalk", "PlatformVersion")
    
}

func describeAccountAttributes(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeAccountAttributesInput{}
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

func describeApplicationVersions(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeApplicationVersionsInput{}
	result, err := client.DescribeApplicationVersions(context.TODO(), input)
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

func describeApplications(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeApplicationsInput{}
	result, err := client.DescribeApplications(context.TODO(), input)
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

func describeConfigurationOptions(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeConfigurationOptionsInput{}
	result, err := client.DescribeConfigurationOptions(context.TODO(), input)
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

func describeConfigurationSettings(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeConfigurationSettingsInput{}
	result, err := client.DescribeConfigurationSettings(context.TODO(), input)
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

func describeEnvironmentHealth(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeEnvironmentHealthInput{}
	result, err := client.DescribeEnvironmentHealth(context.TODO(), input)
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

func describeEnvironmentManagedActionHistory(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput{}
	result, err := client.DescribeEnvironmentManagedActionHistory(context.TODO(), input)
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

func describeEnvironmentManagedActions(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeEnvironmentManagedActionsInput{}
	result, err := client.DescribeEnvironmentManagedActions(context.TODO(), input)
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

func describeEnvironmentResources(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeEnvironmentResourcesInput{}
	result, err := client.DescribeEnvironmentResources(context.TODO(), input)
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

func describeEnvironments(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeEnvironmentsInput{}
	result, err := client.DescribeEnvironments(context.TODO(), input)
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

func describeEvents(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeEventsInput{}
	result, err := client.DescribeEvents(context.TODO(), input)
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

func describeInstancesHealth(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribeInstancesHealthInput{}
	result, err := client.DescribeInstancesHealth(context.TODO(), input)
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

func describePlatformVersion(client *elasticbeanstalk.Client) map[string]interface{} {
	input := &elasticbeanstalk.DescribePlatformVersionInput{}
	result, err := client.DescribePlatformVersion(context.TODO(), input)
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

