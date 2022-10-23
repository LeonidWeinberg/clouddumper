package workdocs

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workdocs"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := workdocs.NewFromConfig(awscfg)
    
	out(describeActivities(client), "workdocs", "Activities")
    
	out(describeComments(client), "workdocs", "Comments")
    
	out(describeDocumentVersions(client), "workdocs", "DocumentVersions")
    
	out(describeFolderContents(client), "workdocs", "FolderContents")
    
	out(describeGroups(client), "workdocs", "Groups")
    
	out(describeNotificationSubscriptions(client), "workdocs", "NotificationSubscriptions")
    
	out(describeResourcePermissions(client), "workdocs", "ResourcePermissions")
    
	out(describeRootFolders(client), "workdocs", "RootFolders")
    
	out(describeUsers(client), "workdocs", "Users")
    
}

func describeActivities(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeActivitiesInput{}
	result, err := client.DescribeActivities(context.TODO(), input)
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

func describeComments(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeCommentsInput{}
	result, err := client.DescribeComments(context.TODO(), input)
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

func describeDocumentVersions(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeDocumentVersionsInput{}
	result, err := client.DescribeDocumentVersions(context.TODO(), input)
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

func describeFolderContents(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeFolderContentsInput{}
	result, err := client.DescribeFolderContents(context.TODO(), input)
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

func describeGroups(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeGroupsInput{}
	result, err := client.DescribeGroups(context.TODO(), input)
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

func describeNotificationSubscriptions(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeNotificationSubscriptionsInput{}
	result, err := client.DescribeNotificationSubscriptions(context.TODO(), input)
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

func describeResourcePermissions(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeResourcePermissionsInput{}
	result, err := client.DescribeResourcePermissions(context.TODO(), input)
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

func describeRootFolders(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeRootFoldersInput{}
	result, err := client.DescribeRootFolders(context.TODO(), input)
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

func describeUsers(client *workdocs.Client) map[string]interface{} {
	input := &workdocs.DescribeUsersInput{}
	result, err := client.DescribeUsers(context.TODO(), input)
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

