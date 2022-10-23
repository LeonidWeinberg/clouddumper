package directconnect

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := directconnect.NewFromConfig(awscfg)
    
	out(describeConnectionLoa(client), "directconnect", "ConnectionLoa")
    
	out(describeConnections(client), "directconnect", "Connections")
    
	out(describeConnectionsOnInterconnect(client), "directconnect", "ConnectionsOnInterconnect")
    
	out(describeCustomerMetadata(client), "directconnect", "CustomerMetadata")
    
	out(describeDirectConnectGatewayAssociationProposals(client), "directconnect", "DirectConnectGatewayAssociationProposals")
    
	out(describeDirectConnectGatewayAssociations(client), "directconnect", "DirectConnectGatewayAssociations")
    
	out(describeDirectConnectGatewayAttachments(client), "directconnect", "DirectConnectGatewayAttachments")
    
	out(describeDirectConnectGateways(client), "directconnect", "DirectConnectGateways")
    
	out(describeHostedConnections(client), "directconnect", "HostedConnections")
    
	out(describeInterconnectLoa(client), "directconnect", "InterconnectLoa")
    
	out(describeInterconnects(client), "directconnect", "Interconnects")
    
	out(describeLags(client), "directconnect", "Lags")
    
	out(describeLoa(client), "directconnect", "Loa")
    
	out(describeLocations(client), "directconnect", "Locations")
    
	out(describeRouterConfiguration(client), "directconnect", "RouterConfiguration")
    
	out(describeTags(client), "directconnect", "Tags")
    
	out(describeVirtualGateways(client), "directconnect", "VirtualGateways")
    
	out(describeVirtualInterfaces(client), "directconnect", "VirtualInterfaces")
    
}

func describeConnectionLoa(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeConnectionLoaInput{}
	result, err := client.DescribeConnectionLoa(context.TODO(), input)
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

func describeConnections(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeConnectionsInput{}
	result, err := client.DescribeConnections(context.TODO(), input)
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

func describeConnectionsOnInterconnect(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeConnectionsOnInterconnectInput{}
	result, err := client.DescribeConnectionsOnInterconnect(context.TODO(), input)
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

func describeCustomerMetadata(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeCustomerMetadataInput{}
	result, err := client.DescribeCustomerMetadata(context.TODO(), input)
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

func describeDirectConnectGatewayAssociationProposals(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeDirectConnectGatewayAssociationProposalsInput{}
	result, err := client.DescribeDirectConnectGatewayAssociationProposals(context.TODO(), input)
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

func describeDirectConnectGatewayAssociations(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeDirectConnectGatewayAssociationsInput{}
	result, err := client.DescribeDirectConnectGatewayAssociations(context.TODO(), input)
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

func describeDirectConnectGatewayAttachments(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeDirectConnectGatewayAttachmentsInput{}
	result, err := client.DescribeDirectConnectGatewayAttachments(context.TODO(), input)
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

func describeDirectConnectGateways(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeDirectConnectGatewaysInput{}
	result, err := client.DescribeDirectConnectGateways(context.TODO(), input)
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

func describeHostedConnections(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeHostedConnectionsInput{}
	result, err := client.DescribeHostedConnections(context.TODO(), input)
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

func describeInterconnectLoa(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeInterconnectLoaInput{}
	result, err := client.DescribeInterconnectLoa(context.TODO(), input)
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

func describeInterconnects(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeInterconnectsInput{}
	result, err := client.DescribeInterconnects(context.TODO(), input)
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

func describeLags(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeLagsInput{}
	result, err := client.DescribeLags(context.TODO(), input)
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

func describeLoa(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeLoaInput{}
	result, err := client.DescribeLoa(context.TODO(), input)
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

func describeLocations(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeLocationsInput{}
	result, err := client.DescribeLocations(context.TODO(), input)
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

func describeRouterConfiguration(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeRouterConfigurationInput{}
	result, err := client.DescribeRouterConfiguration(context.TODO(), input)
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

func describeTags(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeTagsInput{}
	result, err := client.DescribeTags(context.TODO(), input)
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

func describeVirtualGateways(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeVirtualGatewaysInput{}
	result, err := client.DescribeVirtualGateways(context.TODO(), input)
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

func describeVirtualInterfaces(client *directconnect.Client) map[string]interface{} {
	input := &directconnect.DescribeVirtualInterfacesInput{}
	result, err := client.DescribeVirtualInterfaces(context.TODO(), input)
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

