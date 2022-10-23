package grafana

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/grafana"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := grafana.NewFromConfig(awscfg)
    
	out(describeWorkspace(client), "grafana", "Workspace")
    
	out(describeWorkspaceAuthentication(client), "grafana", "WorkspaceAuthentication")
    
}

func describeWorkspace(client *grafana.Client) map[string]interface{} {
	input := &grafana.DescribeWorkspaceInput{}
	result, err := client.DescribeWorkspace(context.TODO(), input)
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

func describeWorkspaceAuthentication(client *grafana.Client) map[string]interface{} {
	input := &grafana.DescribeWorkspaceAuthenticationInput{}
	result, err := client.DescribeWorkspaceAuthentication(context.TODO(), input)
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

