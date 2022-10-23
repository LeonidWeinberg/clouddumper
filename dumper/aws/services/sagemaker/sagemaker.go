package sagemaker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := sagemaker.NewFromConfig(awscfg)
    
	out(describeAction(client), "sagemaker", "Action")
    
	out(describeAlgorithm(client), "sagemaker", "Algorithm")
    
	out(describeApp(client), "sagemaker", "App")
    
	out(describeAppImageConfig(client), "sagemaker", "AppImageConfig")
    
	out(describeArtifact(client), "sagemaker", "Artifact")
    
	out(describeAutoMLJob(client), "sagemaker", "AutoMLJob")
    
	out(describeCodeRepository(client), "sagemaker", "CodeRepository")
    
	out(describeCompilationJob(client), "sagemaker", "CompilationJob")
    
	out(describeContext(client), "sagemaker", "Context")
    
	out(describeDataQualityJobDefinition(client), "sagemaker", "DataQualityJobDefinition")
    
	out(describeDevice(client), "sagemaker", "Device")
    
	out(describeDeviceFleet(client), "sagemaker", "DeviceFleet")
    
	out(describeDomain(client), "sagemaker", "Domain")
    
	out(describeEdgeDeploymentPlan(client), "sagemaker", "EdgeDeploymentPlan")
    
	out(describeEdgePackagingJob(client), "sagemaker", "EdgePackagingJob")
    
	out(describeEndpoint(client), "sagemaker", "Endpoint")
    
	out(describeEndpointConfig(client), "sagemaker", "EndpointConfig")
    
	out(describeExperiment(client), "sagemaker", "Experiment")
    
	out(describeFeatureGroup(client), "sagemaker", "FeatureGroup")
    
	out(describeFeatureMetadata(client), "sagemaker", "FeatureMetadata")
    
	out(describeFlowDefinition(client), "sagemaker", "FlowDefinition")
    
	out(describeHumanTaskUi(client), "sagemaker", "HumanTaskUi")
    
	out(describeHyperParameterTuningJob(client), "sagemaker", "HyperParameterTuningJob")
    
	out(describeImage(client), "sagemaker", "Image")
    
	out(describeImageVersion(client), "sagemaker", "ImageVersion")
    
	out(describeInferenceRecommendationsJob(client), "sagemaker", "InferenceRecommendationsJob")
    
	out(describeLabelingJob(client), "sagemaker", "LabelingJob")
    
	out(describeLineageGroup(client), "sagemaker", "LineageGroup")
    
	out(describeModel(client), "sagemaker", "Model")
    
	out(describeModelBiasJobDefinition(client), "sagemaker", "ModelBiasJobDefinition")
    
	out(describeModelExplainabilityJobDefinition(client), "sagemaker", "ModelExplainabilityJobDefinition")
    
	out(describeModelPackage(client), "sagemaker", "ModelPackage")
    
	out(describeModelPackageGroup(client), "sagemaker", "ModelPackageGroup")
    
	out(describeModelQualityJobDefinition(client), "sagemaker", "ModelQualityJobDefinition")
    
	out(describeMonitoringSchedule(client), "sagemaker", "MonitoringSchedule")
    
	out(describeNotebookInstance(client), "sagemaker", "NotebookInstance")
    
	out(describeNotebookInstanceLifecycleConfig(client), "sagemaker", "NotebookInstanceLifecycleConfig")
    
	out(describePipeline(client), "sagemaker", "Pipeline")
    
	out(describePipelineDefinitionForExecution(client), "sagemaker", "PipelineDefinitionForExecution")
    
	out(describePipelineExecution(client), "sagemaker", "PipelineExecution")
    
	out(describeProcessingJob(client), "sagemaker", "ProcessingJob")
    
	out(describeProject(client), "sagemaker", "Project")
    
	out(describeStudioLifecycleConfig(client), "sagemaker", "StudioLifecycleConfig")
    
	out(describeSubscribedWorkteam(client), "sagemaker", "SubscribedWorkteam")
    
	out(describeTrainingJob(client), "sagemaker", "TrainingJob")
    
	out(describeTransformJob(client), "sagemaker", "TransformJob")
    
	out(describeTrial(client), "sagemaker", "Trial")
    
	out(describeTrialComponent(client), "sagemaker", "TrialComponent")
    
	out(describeUserProfile(client), "sagemaker", "UserProfile")
    
	out(describeWorkforce(client), "sagemaker", "Workforce")
    
	out(describeWorkteam(client), "sagemaker", "Workteam")
    
}

func describeAction(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeActionInput{}
	result, err := client.DescribeAction(context.TODO(), input)
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

func describeAlgorithm(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeAlgorithmInput{}
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

func describeApp(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeAppInput{}
	result, err := client.DescribeApp(context.TODO(), input)
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

func describeAppImageConfig(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeAppImageConfigInput{}
	result, err := client.DescribeAppImageConfig(context.TODO(), input)
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

func describeArtifact(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeArtifactInput{}
	result, err := client.DescribeArtifact(context.TODO(), input)
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

func describeAutoMLJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeAutoMLJobInput{}
	result, err := client.DescribeAutoMLJob(context.TODO(), input)
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

func describeCodeRepository(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeCodeRepositoryInput{}
	result, err := client.DescribeCodeRepository(context.TODO(), input)
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

func describeCompilationJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeCompilationJobInput{}
	result, err := client.DescribeCompilationJob(context.TODO(), input)
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

func describeContext(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeContextInput{}
	result, err := client.DescribeContext(context.TODO(), input)
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

func describeDataQualityJobDefinition(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeDataQualityJobDefinitionInput{}
	result, err := client.DescribeDataQualityJobDefinition(context.TODO(), input)
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

func describeDevice(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeDeviceInput{}
	result, err := client.DescribeDevice(context.TODO(), input)
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

func describeDeviceFleet(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeDeviceFleetInput{}
	result, err := client.DescribeDeviceFleet(context.TODO(), input)
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

func describeDomain(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeDomainInput{}
	result, err := client.DescribeDomain(context.TODO(), input)
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

func describeEdgeDeploymentPlan(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeEdgeDeploymentPlanInput{}
	result, err := client.DescribeEdgeDeploymentPlan(context.TODO(), input)
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

func describeEdgePackagingJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeEdgePackagingJobInput{}
	result, err := client.DescribeEdgePackagingJob(context.TODO(), input)
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

func describeEndpoint(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeEndpointInput{}
	result, err := client.DescribeEndpoint(context.TODO(), input)
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

func describeEndpointConfig(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeEndpointConfigInput{}
	result, err := client.DescribeEndpointConfig(context.TODO(), input)
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

func describeExperiment(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeExperimentInput{}
	result, err := client.DescribeExperiment(context.TODO(), input)
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

func describeFeatureGroup(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeFeatureGroupInput{}
	result, err := client.DescribeFeatureGroup(context.TODO(), input)
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

func describeFeatureMetadata(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeFeatureMetadataInput{}
	result, err := client.DescribeFeatureMetadata(context.TODO(), input)
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

func describeFlowDefinition(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeFlowDefinitionInput{}
	result, err := client.DescribeFlowDefinition(context.TODO(), input)
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

func describeHumanTaskUi(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeHumanTaskUiInput{}
	result, err := client.DescribeHumanTaskUi(context.TODO(), input)
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

func describeHyperParameterTuningJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeHyperParameterTuningJobInput{}
	result, err := client.DescribeHyperParameterTuningJob(context.TODO(), input)
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

func describeImage(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeImageInput{}
	result, err := client.DescribeImage(context.TODO(), input)
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

func describeImageVersion(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeImageVersionInput{}
	result, err := client.DescribeImageVersion(context.TODO(), input)
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

func describeInferenceRecommendationsJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeInferenceRecommendationsJobInput{}
	result, err := client.DescribeInferenceRecommendationsJob(context.TODO(), input)
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

func describeLabelingJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeLabelingJobInput{}
	result, err := client.DescribeLabelingJob(context.TODO(), input)
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

func describeLineageGroup(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeLineageGroupInput{}
	result, err := client.DescribeLineageGroup(context.TODO(), input)
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

func describeModel(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeModelInput{}
	result, err := client.DescribeModel(context.TODO(), input)
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

func describeModelBiasJobDefinition(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeModelBiasJobDefinitionInput{}
	result, err := client.DescribeModelBiasJobDefinition(context.TODO(), input)
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

func describeModelExplainabilityJobDefinition(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeModelExplainabilityJobDefinitionInput{}
	result, err := client.DescribeModelExplainabilityJobDefinition(context.TODO(), input)
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

func describeModelPackage(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeModelPackageInput{}
	result, err := client.DescribeModelPackage(context.TODO(), input)
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

func describeModelPackageGroup(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeModelPackageGroupInput{}
	result, err := client.DescribeModelPackageGroup(context.TODO(), input)
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

func describeModelQualityJobDefinition(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeModelQualityJobDefinitionInput{}
	result, err := client.DescribeModelQualityJobDefinition(context.TODO(), input)
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

func describeMonitoringSchedule(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeMonitoringScheduleInput{}
	result, err := client.DescribeMonitoringSchedule(context.TODO(), input)
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

func describeNotebookInstance(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeNotebookInstanceInput{}
	result, err := client.DescribeNotebookInstance(context.TODO(), input)
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

func describeNotebookInstanceLifecycleConfig(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeNotebookInstanceLifecycleConfigInput{}
	result, err := client.DescribeNotebookInstanceLifecycleConfig(context.TODO(), input)
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

func describePipeline(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribePipelineInput{}
	result, err := client.DescribePipeline(context.TODO(), input)
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

func describePipelineDefinitionForExecution(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribePipelineDefinitionForExecutionInput{}
	result, err := client.DescribePipelineDefinitionForExecution(context.TODO(), input)
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

func describePipelineExecution(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribePipelineExecutionInput{}
	result, err := client.DescribePipelineExecution(context.TODO(), input)
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

func describeProcessingJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeProcessingJobInput{}
	result, err := client.DescribeProcessingJob(context.TODO(), input)
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

func describeProject(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeProjectInput{}
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

func describeStudioLifecycleConfig(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeStudioLifecycleConfigInput{}
	result, err := client.DescribeStudioLifecycleConfig(context.TODO(), input)
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

func describeSubscribedWorkteam(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeSubscribedWorkteamInput{}
	result, err := client.DescribeSubscribedWorkteam(context.TODO(), input)
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

func describeTrainingJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeTrainingJobInput{}
	result, err := client.DescribeTrainingJob(context.TODO(), input)
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

func describeTransformJob(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeTransformJobInput{}
	result, err := client.DescribeTransformJob(context.TODO(), input)
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

func describeTrial(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeTrialInput{}
	result, err := client.DescribeTrial(context.TODO(), input)
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

func describeTrialComponent(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeTrialComponentInput{}
	result, err := client.DescribeTrialComponent(context.TODO(), input)
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

func describeUserProfile(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeUserProfileInput{}
	result, err := client.DescribeUserProfile(context.TODO(), input)
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

func describeWorkforce(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeWorkforceInput{}
	result, err := client.DescribeWorkforce(context.TODO(), input)
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

func describeWorkteam(client *sagemaker.Client) map[string]interface{} {
	input := &sagemaker.DescribeWorkteamInput{}
	result, err := client.DescribeWorkteam(context.TODO(), input)
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

