package quicksight

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := quicksight.NewFromConfig(awscfg)
    
	out(describeAccountCustomization(client), "quicksight", "AccountCustomization")
    
	out(describeAccountSettings(client), "quicksight", "AccountSettings")
    
	out(describeAccountSubscription(client), "quicksight", "AccountSubscription")
    
	out(describeAnalysis(client), "quicksight", "Analysis")
    
	out(describeAnalysisPermissions(client), "quicksight", "AnalysisPermissions")
    
	out(describeDashboard(client), "quicksight", "Dashboard")
    
	out(describeDashboardPermissions(client), "quicksight", "DashboardPermissions")
    
	out(describeDataSet(client), "quicksight", "DataSet")
    
	out(describeDataSetPermissions(client), "quicksight", "DataSetPermissions")
    
	out(describeDataSource(client), "quicksight", "DataSource")
    
	out(describeDataSourcePermissions(client), "quicksight", "DataSourcePermissions")
    
	out(describeFolder(client), "quicksight", "Folder")
    
	out(describeFolderPermissions(client), "quicksight", "FolderPermissions")
    
	out(describeFolderResolvedPermissions(client), "quicksight", "FolderResolvedPermissions")
    
	out(describeGroup(client), "quicksight", "Group")
    
	out(describeGroupMembership(client), "quicksight", "GroupMembership")
    
	out(describeIAMPolicyAssignment(client), "quicksight", "IAMPolicyAssignment")
    
	out(describeIngestion(client), "quicksight", "Ingestion")
    
	out(describeIpRestriction(client), "quicksight", "IpRestriction")
    
	out(describeNamespace(client), "quicksight", "Namespace")
    
	out(describeTemplate(client), "quicksight", "Template")
    
	out(describeTemplateAlias(client), "quicksight", "TemplateAlias")
    
	out(describeTemplatePermissions(client), "quicksight", "TemplatePermissions")
    
	out(describeTheme(client), "quicksight", "Theme")
    
	out(describeThemeAlias(client), "quicksight", "ThemeAlias")
    
	out(describeThemePermissions(client), "quicksight", "ThemePermissions")
    
	out(describeUser(client), "quicksight", "User")
    
}

func describeAccountCustomization(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeAccountCustomizationInput{}
	result, err := client.DescribeAccountCustomization(context.TODO(), input)
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

func describeAccountSettings(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeAccountSettingsInput{}
	result, err := client.DescribeAccountSettings(context.TODO(), input)
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

func describeAccountSubscription(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeAccountSubscriptionInput{}
	result, err := client.DescribeAccountSubscription(context.TODO(), input)
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

func describeAnalysis(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeAnalysisInput{}
	result, err := client.DescribeAnalysis(context.TODO(), input)
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

func describeAnalysisPermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeAnalysisPermissionsInput{}
	result, err := client.DescribeAnalysisPermissions(context.TODO(), input)
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

func describeDashboard(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeDashboardInput{}
	result, err := client.DescribeDashboard(context.TODO(), input)
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

func describeDashboardPermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeDashboardPermissionsInput{}
	result, err := client.DescribeDashboardPermissions(context.TODO(), input)
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

func describeDataSet(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeDataSetInput{}
	result, err := client.DescribeDataSet(context.TODO(), input)
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

func describeDataSetPermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeDataSetPermissionsInput{}
	result, err := client.DescribeDataSetPermissions(context.TODO(), input)
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

func describeDataSource(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeDataSourceInput{}
	result, err := client.DescribeDataSource(context.TODO(), input)
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

func describeDataSourcePermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeDataSourcePermissionsInput{}
	result, err := client.DescribeDataSourcePermissions(context.TODO(), input)
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

func describeFolder(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeFolderInput{}
	result, err := client.DescribeFolder(context.TODO(), input)
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

func describeFolderPermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeFolderPermissionsInput{}
	result, err := client.DescribeFolderPermissions(context.TODO(), input)
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

func describeFolderResolvedPermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeFolderResolvedPermissionsInput{}
	result, err := client.DescribeFolderResolvedPermissions(context.TODO(), input)
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

func describeGroup(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeGroupInput{}
	result, err := client.DescribeGroup(context.TODO(), input)
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

func describeGroupMembership(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeGroupMembershipInput{}
	result, err := client.DescribeGroupMembership(context.TODO(), input)
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

func describeIAMPolicyAssignment(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeIAMPolicyAssignmentInput{}
	result, err := client.DescribeIAMPolicyAssignment(context.TODO(), input)
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

func describeIngestion(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeIngestionInput{}
	result, err := client.DescribeIngestion(context.TODO(), input)
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

func describeIpRestriction(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeIpRestrictionInput{}
	result, err := client.DescribeIpRestriction(context.TODO(), input)
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

func describeNamespace(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeNamespaceInput{}
	result, err := client.DescribeNamespace(context.TODO(), input)
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

func describeTemplate(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeTemplateInput{}
	result, err := client.DescribeTemplate(context.TODO(), input)
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

func describeTemplateAlias(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeTemplateAliasInput{}
	result, err := client.DescribeTemplateAlias(context.TODO(), input)
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

func describeTemplatePermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeTemplatePermissionsInput{}
	result, err := client.DescribeTemplatePermissions(context.TODO(), input)
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

func describeTheme(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeThemeInput{}
	result, err := client.DescribeTheme(context.TODO(), input)
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

func describeThemeAlias(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeThemeAliasInput{}
	result, err := client.DescribeThemeAlias(context.TODO(), input)
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

func describeThemePermissions(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeThemePermissionsInput{}
	result, err := client.DescribeThemePermissions(context.TODO(), input)
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

func describeUser(client *quicksight.Client) map[string]interface{} {
	input := &quicksight.DescribeUserInput{}
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

