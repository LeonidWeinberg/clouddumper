package appmesh

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := appmesh.NewFromConfig(awscfg)
    
	out(describeGatewayRoute(client), "appmesh", "GatewayRoute")
    
	out(describeMesh(client), "appmesh", "Mesh")
    
	out(describeRoute(client), "appmesh", "Route")
    
	out(describeVirtualGateway(client), "appmesh", "VirtualGateway")
    
	out(describeVirtualNode(client), "appmesh", "VirtualNode")
    
	out(describeVirtualRouter(client), "appmesh", "VirtualRouter")
    
	out(describeVirtualService(client), "appmesh", "VirtualService")
    
}

func describeGatewayRoute(client *appmesh.Client) map[string]interface{} {
	input := &appmesh.DescribeGatewayRouteInput{}
	result, err := client.DescribeGatewayRoute(context.TODO(), input)
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

func describeMesh(client *appmesh.Client) map[string]interface{} {
	input := &appmesh.DescribeMeshInput{}
	result, err := client.DescribeMesh(context.TODO(), input)
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

func describeRoute(client *appmesh.Client) map[string]interface{} {
	input := &appmesh.DescribeRouteInput{}
	result, err := client.DescribeRoute(context.TODO(), input)
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

func describeVirtualGateway(client *appmesh.Client) map[string]interface{} {
	input := &appmesh.DescribeVirtualGatewayInput{}
	result, err := client.DescribeVirtualGateway(context.TODO(), input)
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

func describeVirtualNode(client *appmesh.Client) map[string]interface{} {
	input := &appmesh.DescribeVirtualNodeInput{}
	result, err := client.DescribeVirtualNode(context.TODO(), input)
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

func describeVirtualRouter(client *appmesh.Client) map[string]interface{} {
	input := &appmesh.DescribeVirtualRouterInput{}
	result, err := client.DescribeVirtualRouter(context.TODO(), input)
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

func describeVirtualService(client *appmesh.Client) map[string]interface{} {
	input := &appmesh.DescribeVirtualServiceInput{}
	result, err := client.DescribeVirtualService(context.TODO(), input)
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

