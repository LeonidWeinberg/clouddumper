package databrew

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databrew"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := databrew.NewFromConfig(awscfg)
    
	out(describeDataset(client), "databrew", "Dataset")
    
	out(describeJob(client), "databrew", "Job")
    
	out(describeJobRun(client), "databrew", "JobRun")
    
	out(describeProject(client), "databrew", "Project")
    
	out(describeRecipe(client), "databrew", "Recipe")
    
	out(describeRuleset(client), "databrew", "Ruleset")
    
	out(describeSchedule(client), "databrew", "Schedule")
    
}

func describeDataset(client *databrew.Client) map[string]interface{} {
	input := &databrew.DescribeDatasetInput{}
	result, err := client.DescribeDataset(context.TODO(), input)
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

func describeJob(client *databrew.Client) map[string]interface{} {
	input := &databrew.DescribeJobInput{}
	result, err := client.DescribeJob(context.TODO(), input)
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

func describeJobRun(client *databrew.Client) map[string]interface{} {
	input := &databrew.DescribeJobRunInput{}
	result, err := client.DescribeJobRun(context.TODO(), input)
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

func describeProject(client *databrew.Client) map[string]interface{} {
	input := &databrew.DescribeProjectInput{}
	result, err := client.DescribeProject(context.TODO(), input)
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

func describeRecipe(client *databrew.Client) map[string]interface{} {
	input := &databrew.DescribeRecipeInput{}
	result, err := client.DescribeRecipe(context.TODO(), input)
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

func describeRuleset(client *databrew.Client) map[string]interface{} {
	input := &databrew.DescribeRulesetInput{}
	result, err := client.DescribeRuleset(context.TODO(), input)
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

func describeSchedule(client *databrew.Client) map[string]interface{} {
	input := &databrew.DescribeScheduleInput{}
	result, err := client.DescribeSchedule(context.TODO(), input)
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

