package robomaker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/robomaker"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := robomaker.NewFromConfig(awscfg)
    
	out(describeDeploymentJob(client), "robomaker", "DeploymentJob")
    
	out(describeFleet(client), "robomaker", "Fleet")
    
	out(describeRobot(client), "robomaker", "Robot")
    
	out(describeRobotApplication(client), "robomaker", "RobotApplication")
    
	out(describeSimulationApplication(client), "robomaker", "SimulationApplication")
    
	out(describeSimulationJob(client), "robomaker", "SimulationJob")
    
	out(describeSimulationJobBatch(client), "robomaker", "SimulationJobBatch")
    
	out(describeWorld(client), "robomaker", "World")
    
	out(describeWorldExportJob(client), "robomaker", "WorldExportJob")
    
	out(describeWorldGenerationJob(client), "robomaker", "WorldGenerationJob")
    
	out(describeWorldTemplate(client), "robomaker", "WorldTemplate")
    
}

func describeDeploymentJob(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeDeploymentJobInput{}
	result, err := client.DescribeDeploymentJob(context.TODO(), input)
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

func describeFleet(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeFleetInput{}
	result, err := client.DescribeFleet(context.TODO(), input)
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

func describeRobot(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeRobotInput{}
	result, err := client.DescribeRobot(context.TODO(), input)
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

func describeRobotApplication(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeRobotApplicationInput{}
	result, err := client.DescribeRobotApplication(context.TODO(), input)
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

func describeSimulationApplication(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeSimulationApplicationInput{}
	result, err := client.DescribeSimulationApplication(context.TODO(), input)
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

func describeSimulationJob(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeSimulationJobInput{}
	result, err := client.DescribeSimulationJob(context.TODO(), input)
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

func describeSimulationJobBatch(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeSimulationJobBatchInput{}
	result, err := client.DescribeSimulationJobBatch(context.TODO(), input)
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

func describeWorld(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeWorldInput{}
	result, err := client.DescribeWorld(context.TODO(), input)
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

func describeWorldExportJob(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeWorldExportJobInput{}
	result, err := client.DescribeWorldExportJob(context.TODO(), input)
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

func describeWorldGenerationJob(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeWorldGenerationJobInput{}
	result, err := client.DescribeWorldGenerationJob(context.TODO(), input)
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

func describeWorldTemplate(client *robomaker.Client) map[string]interface{} {
	input := &robomaker.DescribeWorldTemplateInput{}
	result, err := client.DescribeWorldTemplate(context.TODO(), input)
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

