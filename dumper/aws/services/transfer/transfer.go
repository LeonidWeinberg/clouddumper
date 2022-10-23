package transfer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := transfer.NewFromConfig(awscfg)
    
	out(describeAccess(client), "transfer", "Access")
    
	out(describeAgreement(client), "transfer", "Agreement")
    
	out(describeCertificate(client), "transfer", "Certificate")
    
	out(describeConnector(client), "transfer", "Connector")
    
	out(describeExecution(client), "transfer", "Execution")
    
	out(describeHostKey(client), "transfer", "HostKey")
    
	out(describeProfile(client), "transfer", "Profile")
    
	out(describeSecurityPolicy(client), "transfer", "SecurityPolicy")
    
	out(describeServer(client), "transfer", "Server")
    
	out(describeUser(client), "transfer", "User")
    
	out(describeWorkflow(client), "transfer", "Workflow")
    
}

func describeAccess(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeAccessInput{}
	result, err := client.DescribeAccess(context.TODO(), input)
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

func describeAgreement(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeAgreementInput{}
	result, err := client.DescribeAgreement(context.TODO(), input)
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

func describeCertificate(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeCertificateInput{}
	result, err := client.DescribeCertificate(context.TODO(), input)
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

func describeConnector(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeConnectorInput{}
	result, err := client.DescribeConnector(context.TODO(), input)
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

func describeExecution(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeExecutionInput{}
	result, err := client.DescribeExecution(context.TODO(), input)
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

func describeHostKey(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeHostKeyInput{}
	result, err := client.DescribeHostKey(context.TODO(), input)
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

func describeProfile(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeProfileInput{}
	result, err := client.DescribeProfile(context.TODO(), input)
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

func describeSecurityPolicy(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeSecurityPolicyInput{}
	result, err := client.DescribeSecurityPolicy(context.TODO(), input)
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

func describeServer(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeServerInput{}
	result, err := client.DescribeServer(context.TODO(), input)
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

func describeUser(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeUserInput{}
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

func describeWorkflow(client *transfer.Client) map[string]interface{} {
	input := &transfer.DescribeWorkflowInput{}
	result, err := client.DescribeWorkflow(context.TODO(), input)
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

