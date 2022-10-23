package personalize

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalize"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := personalize.NewFromConfig(awscfg)
    
	out(describeAlgorithm(client), "personalize", "Algorithm")
    
	out(describeBatchInferenceJob(client), "personalize", "BatchInferenceJob")
    
	out(describeBatchSegmentJob(client), "personalize", "BatchSegmentJob")
    
	out(describeCampaign(client), "personalize", "Campaign")
    
	out(describeDataset(client), "personalize", "Dataset")
    
	out(describeDatasetExportJob(client), "personalize", "DatasetExportJob")
    
	out(describeDatasetGroup(client), "personalize", "DatasetGroup")
    
	out(describeDatasetImportJob(client), "personalize", "DatasetImportJob")
    
	out(describeEventTracker(client), "personalize", "EventTracker")
    
	out(describeFeatureTransformation(client), "personalize", "FeatureTransformation")
    
	out(describeFilter(client), "personalize", "Filter")
    
	out(describeRecipe(client), "personalize", "Recipe")
    
	out(describeRecommender(client), "personalize", "Recommender")
    
	out(describeSchema(client), "personalize", "Schema")
    
	out(describeSolution(client), "personalize", "Solution")
    
	out(describeSolutionVersion(client), "personalize", "SolutionVersion")
    
}

func describeAlgorithm(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeAlgorithmInput{}
	result, err := client.DescribeAlgorithm(context.TODO(), input)
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

func describeBatchInferenceJob(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeBatchInferenceJobInput{}
	result, err := client.DescribeBatchInferenceJob(context.TODO(), input)
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

func describeBatchSegmentJob(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeBatchSegmentJobInput{}
	result, err := client.DescribeBatchSegmentJob(context.TODO(), input)
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

func describeCampaign(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeCampaignInput{}
	result, err := client.DescribeCampaign(context.TODO(), input)
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

func describeDataset(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeDatasetInput{}
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

func describeDatasetExportJob(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeDatasetExportJobInput{}
	result, err := client.DescribeDatasetExportJob(context.TODO(), input)
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

func describeDatasetGroup(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeDatasetGroupInput{}
	result, err := client.DescribeDatasetGroup(context.TODO(), input)
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

func describeDatasetImportJob(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeDatasetImportJobInput{}
	result, err := client.DescribeDatasetImportJob(context.TODO(), input)
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

func describeEventTracker(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeEventTrackerInput{}
	result, err := client.DescribeEventTracker(context.TODO(), input)
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

func describeFeatureTransformation(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeFeatureTransformationInput{}
	result, err := client.DescribeFeatureTransformation(context.TODO(), input)
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

func describeFilter(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeFilterInput{}
	result, err := client.DescribeFilter(context.TODO(), input)
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

func describeRecipe(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeRecipeInput{}
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

func describeRecommender(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeRecommenderInput{}
	result, err := client.DescribeRecommender(context.TODO(), input)
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

func describeSchema(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeSchemaInput{}
	result, err := client.DescribeSchema(context.TODO(), input)
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

func describeSolution(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeSolutionInput{}
	result, err := client.DescribeSolution(context.TODO(), input)
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

func describeSolutionVersion(client *personalize.Client) map[string]interface{} {
	input := &personalize.DescribeSolutionVersionInput{}
	result, err := client.DescribeSolutionVersion(context.TODO(), input)
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

