package forecast

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/forecast"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := forecast.NewFromConfig(awscfg)
    
	out(describeAutoPredictor(client), "forecast", "AutoPredictor")
    
	out(describeDataset(client), "forecast", "Dataset")
    
	out(describeDatasetGroup(client), "forecast", "DatasetGroup")
    
	out(describeDatasetImportJob(client), "forecast", "DatasetImportJob")
    
	out(describeExplainability(client), "forecast", "Explainability")
    
	out(describeExplainabilityExport(client), "forecast", "ExplainabilityExport")
    
	out(describeForecast(client), "forecast", "Forecast")
    
	out(describeForecastExportJob(client), "forecast", "ForecastExportJob")
    
	out(describeMonitor(client), "forecast", "Monitor")
    
	out(describePredictor(client), "forecast", "Predictor")
    
	out(describePredictorBacktestExportJob(client), "forecast", "PredictorBacktestExportJob")
    
	out(describeWhatIfAnalysis(client), "forecast", "WhatIfAnalysis")
    
	out(describeWhatIfForecast(client), "forecast", "WhatIfForecast")
    
	out(describeWhatIfForecastExport(client), "forecast", "WhatIfForecastExport")
    
}

func describeAutoPredictor(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeAutoPredictorInput{}
	result, err := client.DescribeAutoPredictor(context.TODO(), input)
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

func describeDataset(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeDatasetInput{}
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

func describeDatasetGroup(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeDatasetGroupInput{}
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

func describeDatasetImportJob(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeDatasetImportJobInput{}
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

func describeExplainability(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeExplainabilityInput{}
	result, err := client.DescribeExplainability(context.TODO(), input)
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

func describeExplainabilityExport(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeExplainabilityExportInput{}
	result, err := client.DescribeExplainabilityExport(context.TODO(), input)
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

func describeForecast(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeForecastInput{}
	result, err := client.DescribeForecast(context.TODO(), input)
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

func describeForecastExportJob(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeForecastExportJobInput{}
	result, err := client.DescribeForecastExportJob(context.TODO(), input)
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

func describeMonitor(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeMonitorInput{}
	result, err := client.DescribeMonitor(context.TODO(), input)
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

func describePredictor(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribePredictorInput{}
	result, err := client.DescribePredictor(context.TODO(), input)
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

func describePredictorBacktestExportJob(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribePredictorBacktestExportJobInput{}
	result, err := client.DescribePredictorBacktestExportJob(context.TODO(), input)
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

func describeWhatIfAnalysis(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeWhatIfAnalysisInput{}
	result, err := client.DescribeWhatIfAnalysis(context.TODO(), input)
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

func describeWhatIfForecast(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeWhatIfForecastInput{}
	result, err := client.DescribeWhatIfForecast(context.TODO(), input)
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

func describeWhatIfForecastExport(client *forecast.Client) map[string]interface{} {
	input := &forecast.DescribeWhatIfForecastExportInput{}
	result, err := client.DescribeWhatIfForecastExport(context.TODO(), input)
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

