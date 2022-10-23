package databasemigrationservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := databasemigrationservice.NewFromConfig(awscfg)
    
	out(describeAccountAttributes(client), "databasemigrationservice", "AccountAttributes")
    
	out(describeApplicableIndividualAssessments(client), "databasemigrationservice", "ApplicableIndividualAssessments")
    
	out(describeCertificates(client), "databasemigrationservice", "Certificates")
    
	out(describeConnections(client), "databasemigrationservice", "Connections")
    
	out(describeEndpointSettings(client), "databasemigrationservice", "EndpointSettings")
    
	out(describeEndpointTypes(client), "databasemigrationservice", "EndpointTypes")
    
	out(describeEndpoints(client), "databasemigrationservice", "Endpoints")
    
	out(describeEventCategories(client), "databasemigrationservice", "EventCategories")
    
	out(describeEventSubscriptions(client), "databasemigrationservice", "EventSubscriptions")
    
	out(describeEvents(client), "databasemigrationservice", "Events")
    
	out(describeFleetAdvisorCollectors(client), "databasemigrationservice", "FleetAdvisorCollectors")
    
	out(describeFleetAdvisorDatabases(client), "databasemigrationservice", "FleetAdvisorDatabases")
    
	out(describeFleetAdvisorLsaAnalysis(client), "databasemigrationservice", "FleetAdvisorLsaAnalysis")
    
	out(describeFleetAdvisorSchemaObjectSummary(client), "databasemigrationservice", "FleetAdvisorSchemaObjectSummary")
    
	out(describeFleetAdvisorSchemas(client), "databasemigrationservice", "FleetAdvisorSchemas")
    
	out(describeOrderableReplicationInstances(client), "databasemigrationservice", "OrderableReplicationInstances")
    
	out(describePendingMaintenanceActions(client), "databasemigrationservice", "PendingMaintenanceActions")
    
	out(describeRefreshSchemasStatus(client), "databasemigrationservice", "RefreshSchemasStatus")
    
	out(describeReplicationInstanceTaskLogs(client), "databasemigrationservice", "ReplicationInstanceTaskLogs")
    
	out(describeReplicationInstances(client), "databasemigrationservice", "ReplicationInstances")
    
	out(describeReplicationSubnetGroups(client), "databasemigrationservice", "ReplicationSubnetGroups")
    
	out(describeReplicationTaskAssessmentResults(client), "databasemigrationservice", "ReplicationTaskAssessmentResults")
    
	out(describeReplicationTaskAssessmentRuns(client), "databasemigrationservice", "ReplicationTaskAssessmentRuns")
    
	out(describeReplicationTaskIndividualAssessments(client), "databasemigrationservice", "ReplicationTaskIndividualAssessments")
    
	out(describeReplicationTasks(client), "databasemigrationservice", "ReplicationTasks")
    
	out(describeSchemas(client), "databasemigrationservice", "Schemas")
    
	out(describeTableStatistics(client), "databasemigrationservice", "TableStatistics")
    
}

func describeAccountAttributes(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeAccountAttributesInput{}
	result, err := client.DescribeAccountAttributes(context.TODO(), input)
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

func describeApplicableIndividualAssessments(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeApplicableIndividualAssessmentsInput{}
	result, err := client.DescribeApplicableIndividualAssessments(context.TODO(), input)
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

func describeCertificates(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeCertificatesInput{}
	result, err := client.DescribeCertificates(context.TODO(), input)
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

func describeConnections(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeConnectionsInput{}
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

func describeEndpointSettings(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeEndpointSettingsInput{}
	result, err := client.DescribeEndpointSettings(context.TODO(), input)
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

func describeEndpointTypes(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeEndpointTypesInput{}
	result, err := client.DescribeEndpointTypes(context.TODO(), input)
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

func describeEndpoints(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeEndpointsInput{}
	result, err := client.DescribeEndpoints(context.TODO(), input)
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

func describeEventCategories(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeEventCategoriesInput{}
	result, err := client.DescribeEventCategories(context.TODO(), input)
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

func describeEventSubscriptions(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeEventSubscriptionsInput{}
	result, err := client.DescribeEventSubscriptions(context.TODO(), input)
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

func describeEvents(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeEventsInput{}
	result, err := client.DescribeEvents(context.TODO(), input)
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

func describeFleetAdvisorCollectors(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeFleetAdvisorCollectorsInput{}
	result, err := client.DescribeFleetAdvisorCollectors(context.TODO(), input)
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

func describeFleetAdvisorDatabases(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeFleetAdvisorDatabasesInput{}
	result, err := client.DescribeFleetAdvisorDatabases(context.TODO(), input)
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

func describeFleetAdvisorLsaAnalysis(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput{}
	result, err := client.DescribeFleetAdvisorLsaAnalysis(context.TODO(), input)
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

func describeFleetAdvisorSchemaObjectSummary(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput{}
	result, err := client.DescribeFleetAdvisorSchemaObjectSummary(context.TODO(), input)
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

func describeFleetAdvisorSchemas(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeFleetAdvisorSchemasInput{}
	result, err := client.DescribeFleetAdvisorSchemas(context.TODO(), input)
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

func describeOrderableReplicationInstances(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeOrderableReplicationInstancesInput{}
	result, err := client.DescribeOrderableReplicationInstances(context.TODO(), input)
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

func describePendingMaintenanceActions(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribePendingMaintenanceActionsInput{}
	result, err := client.DescribePendingMaintenanceActions(context.TODO(), input)
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

func describeRefreshSchemasStatus(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeRefreshSchemasStatusInput{}
	result, err := client.DescribeRefreshSchemasStatus(context.TODO(), input)
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

func describeReplicationInstanceTaskLogs(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeReplicationInstanceTaskLogsInput{}
	result, err := client.DescribeReplicationInstanceTaskLogs(context.TODO(), input)
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

func describeReplicationInstances(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeReplicationInstancesInput{}
	result, err := client.DescribeReplicationInstances(context.TODO(), input)
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

func describeReplicationSubnetGroups(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeReplicationSubnetGroupsInput{}
	result, err := client.DescribeReplicationSubnetGroups(context.TODO(), input)
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

func describeReplicationTaskAssessmentResults(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput{}
	result, err := client.DescribeReplicationTaskAssessmentResults(context.TODO(), input)
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

func describeReplicationTaskAssessmentRuns(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput{}
	result, err := client.DescribeReplicationTaskAssessmentRuns(context.TODO(), input)
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

func describeReplicationTaskIndividualAssessments(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput{}
	result, err := client.DescribeReplicationTaskIndividualAssessments(context.TODO(), input)
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

func describeReplicationTasks(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeReplicationTasksInput{}
	result, err := client.DescribeReplicationTasks(context.TODO(), input)
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

func describeSchemas(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeSchemasInput{}
	result, err := client.DescribeSchemas(context.TODO(), input)
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

func describeTableStatistics(client *databasemigrationservice.Client) map[string]interface{} {
	input := &databasemigrationservice.DescribeTableStatisticsInput{}
	result, err := client.DescribeTableStatistics(context.TODO(), input)
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

