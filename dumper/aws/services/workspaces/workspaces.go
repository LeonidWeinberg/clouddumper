package workspaces

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := workspaces.NewFromConfig(awscfg)
    
	out(describeAccount(client), "workspaces", "Account")
    
	out(describeAccountModifications(client), "workspaces", "AccountModifications")
    
	out(describeClientBranding(client), "workspaces", "ClientBranding")
    
	out(describeClientProperties(client), "workspaces", "ClientProperties")
    
	out(describeConnectClientAddIns(client), "workspaces", "ConnectClientAddIns")
    
	out(describeConnectionAliasPermissions(client), "workspaces", "ConnectionAliasPermissions")
    
	out(describeConnectionAliases(client), "workspaces", "ConnectionAliases")
    
	out(describeIpGroups(client), "workspaces", "IpGroups")
    
	out(describeTags(client), "workspaces", "Tags")
    
	out(describeWorkspaceBundles(client), "workspaces", "WorkspaceBundles")
    
	out(describeWorkspaceDirectories(client), "workspaces", "WorkspaceDirectories")
    
	out(describeWorkspaceImagePermissions(client), "workspaces", "WorkspaceImagePermissions")
    
	out(describeWorkspaceImages(client), "workspaces", "WorkspaceImages")
    
	out(describeWorkspaceSnapshots(client), "workspaces", "WorkspaceSnapshots")
    
	out(describeWorkspaces(client), "workspaces", "Workspaces")
    
	out(describeWorkspacesConnectionStatus(client), "workspaces", "WorkspacesConnectionStatus")
    
}

func describeAccount(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeAccountInput{}
	result, err := client.DescribeAccount(context.TODO(), input)
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

func describeAccountModifications(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeAccountModificationsInput{}
	result, err := client.DescribeAccountModifications(context.TODO(), input)
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

func describeClientBranding(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeClientBrandingInput{}
	result, err := client.DescribeClientBranding(context.TODO(), input)
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

func describeClientProperties(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeClientPropertiesInput{}
	result, err := client.DescribeClientProperties(context.TODO(), input)
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

func describeConnectClientAddIns(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeConnectClientAddInsInput{}
	result, err := client.DescribeConnectClientAddIns(context.TODO(), input)
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

func describeConnectionAliasPermissions(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeConnectionAliasPermissionsInput{}
	result, err := client.DescribeConnectionAliasPermissions(context.TODO(), input)
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

func describeConnectionAliases(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeConnectionAliasesInput{}
	result, err := client.DescribeConnectionAliases(context.TODO(), input)
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

func describeIpGroups(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeIpGroupsInput{}
	result, err := client.DescribeIpGroups(context.TODO(), input)
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

func describeTags(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeTagsInput{}
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

func describeWorkspaceBundles(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeWorkspaceBundlesInput{}
	result, err := client.DescribeWorkspaceBundles(context.TODO(), input)
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

func describeWorkspaceDirectories(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeWorkspaceDirectoriesInput{}
	result, err := client.DescribeWorkspaceDirectories(context.TODO(), input)
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

func describeWorkspaceImagePermissions(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeWorkspaceImagePermissionsInput{}
	result, err := client.DescribeWorkspaceImagePermissions(context.TODO(), input)
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

func describeWorkspaceImages(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeWorkspaceImagesInput{}
	result, err := client.DescribeWorkspaceImages(context.TODO(), input)
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

func describeWorkspaceSnapshots(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeWorkspaceSnapshotsInput{}
	result, err := client.DescribeWorkspaceSnapshots(context.TODO(), input)
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

func describeWorkspaces(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeWorkspacesInput{}
	result, err := client.DescribeWorkspaces(context.TODO(), input)
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

func describeWorkspacesConnectionStatus(client *workspaces.Client) map[string]interface{} {
	input := &workspaces.DescribeWorkspacesConnectionStatusInput{}
	result, err := client.DescribeWorkspacesConnectionStatus(context.TODO(), input)
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

