package rds

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := rds.NewFromConfig(awscfg)
    
	out(describeAccountAttributes(client), "rds", "AccountAttributes")
    
	out(describeCertificates(client), "rds", "Certificates")
    
	out(describeDBClusterBacktracks(client), "rds", "DBClusterBacktracks")
    
	out(describeDBClusterEndpoints(client), "rds", "DBClusterEndpoints")
    
	out(describeDBClusterParameterGroups(client), "rds", "DBClusterParameterGroups")
    
	out(describeDBClusterParameters(client), "rds", "DBClusterParameters")
    
	out(describeDBClusterSnapshotAttributes(client), "rds", "DBClusterSnapshotAttributes")
    
	out(describeDBClusterSnapshots(client), "rds", "DBClusterSnapshots")
    
	out(describeDBClusters(client), "rds", "DBClusters")
    
	out(describeDBEngineVersions(client), "rds", "DBEngineVersions")
    
	out(describeDBInstanceAutomatedBackups(client), "rds", "DBInstanceAutomatedBackups")
    
	out(describeDBInstances(client), "rds", "DBInstances")
    
	out(describeDBLogFiles(client), "rds", "DBLogFiles")
    
	out(describeDBParameterGroups(client), "rds", "DBParameterGroups")
    
	out(describeDBParameters(client), "rds", "DBParameters")
    
	out(describeDBProxies(client), "rds", "DBProxies")
    
	out(describeDBProxyEndpoints(client), "rds", "DBProxyEndpoints")
    
	out(describeDBProxyTargetGroups(client), "rds", "DBProxyTargetGroups")
    
	out(describeDBProxyTargets(client), "rds", "DBProxyTargets")
    
	out(describeDBSecurityGroups(client), "rds", "DBSecurityGroups")
    
	out(describeDBSnapshotAttributes(client), "rds", "DBSnapshotAttributes")
    
	out(describeDBSnapshots(client), "rds", "DBSnapshots")
    
	out(describeDBSubnetGroups(client), "rds", "DBSubnetGroups")
    
	out(describeEngineDefaultClusterParameters(client), "rds", "EngineDefaultClusterParameters")
    
	out(describeEngineDefaultParameters(client), "rds", "EngineDefaultParameters")
    
	out(describeEventCategories(client), "rds", "EventCategories")
    
	out(describeEventSubscriptions(client), "rds", "EventSubscriptions")
    
	out(describeEvents(client), "rds", "Events")
    
	out(describeExportTasks(client), "rds", "ExportTasks")
    
	out(describeGlobalClusters(client), "rds", "GlobalClusters")
    
	out(describeOptionGroupOptions(client), "rds", "OptionGroupOptions")
    
	out(describeOptionGroups(client), "rds", "OptionGroups")
    
	out(describeOrderableDBInstanceOptions(client), "rds", "OrderableDBInstanceOptions")
    
	out(describePendingMaintenanceActions(client), "rds", "PendingMaintenanceActions")
    
	out(describeReservedDBInstances(client), "rds", "ReservedDBInstances")
    
	out(describeReservedDBInstancesOfferings(client), "rds", "ReservedDBInstancesOfferings")
    
	out(describeSourceRegions(client), "rds", "SourceRegions")
    
	out(describeValidDBInstanceModifications(client), "rds", "ValidDBInstanceModifications")
    
}

func describeAccountAttributes(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeAccountAttributesInput{}
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

func describeCertificates(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeCertificatesInput{}
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

func describeDBClusterBacktracks(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBClusterBacktracksInput{}
	result, err := client.DescribeDBClusterBacktracks(context.TODO(), input)
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

func describeDBClusterEndpoints(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBClusterEndpointsInput{}
	result, err := client.DescribeDBClusterEndpoints(context.TODO(), input)
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

func describeDBClusterParameterGroups(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBClusterParameterGroupsInput{}
	result, err := client.DescribeDBClusterParameterGroups(context.TODO(), input)
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

func describeDBClusterParameters(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBClusterParametersInput{}
	result, err := client.DescribeDBClusterParameters(context.TODO(), input)
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

func describeDBClusterSnapshotAttributes(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBClusterSnapshotAttributesInput{}
	result, err := client.DescribeDBClusterSnapshotAttributes(context.TODO(), input)
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

func describeDBClusterSnapshots(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBClusterSnapshotsInput{}
	result, err := client.DescribeDBClusterSnapshots(context.TODO(), input)
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

func describeDBClusters(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBClustersInput{}
	result, err := client.DescribeDBClusters(context.TODO(), input)
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

func describeDBEngineVersions(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBEngineVersionsInput{}
	result, err := client.DescribeDBEngineVersions(context.TODO(), input)
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

func describeDBInstanceAutomatedBackups(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBInstanceAutomatedBackupsInput{}
	result, err := client.DescribeDBInstanceAutomatedBackups(context.TODO(), input)
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

func describeDBInstances(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBInstancesInput{}
	result, err := client.DescribeDBInstances(context.TODO(), input)
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

func describeDBLogFiles(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBLogFilesInput{}
	result, err := client.DescribeDBLogFiles(context.TODO(), input)
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

func describeDBParameterGroups(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBParameterGroupsInput{}
	result, err := client.DescribeDBParameterGroups(context.TODO(), input)
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

func describeDBParameters(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBParametersInput{}
	result, err := client.DescribeDBParameters(context.TODO(), input)
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

func describeDBProxies(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBProxiesInput{}
	result, err := client.DescribeDBProxies(context.TODO(), input)
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

func describeDBProxyEndpoints(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBProxyEndpointsInput{}
	result, err := client.DescribeDBProxyEndpoints(context.TODO(), input)
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

func describeDBProxyTargetGroups(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBProxyTargetGroupsInput{}
	result, err := client.DescribeDBProxyTargetGroups(context.TODO(), input)
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

func describeDBProxyTargets(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBProxyTargetsInput{}
	result, err := client.DescribeDBProxyTargets(context.TODO(), input)
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

func describeDBSecurityGroups(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBSecurityGroupsInput{}
	result, err := client.DescribeDBSecurityGroups(context.TODO(), input)
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

func describeDBSnapshotAttributes(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBSnapshotAttributesInput{}
	result, err := client.DescribeDBSnapshotAttributes(context.TODO(), input)
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

func describeDBSnapshots(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBSnapshotsInput{}
	result, err := client.DescribeDBSnapshots(context.TODO(), input)
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

func describeDBSubnetGroups(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeDBSubnetGroupsInput{}
	result, err := client.DescribeDBSubnetGroups(context.TODO(), input)
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

func describeEngineDefaultClusterParameters(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeEngineDefaultClusterParametersInput{}
	result, err := client.DescribeEngineDefaultClusterParameters(context.TODO(), input)
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

func describeEngineDefaultParameters(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeEngineDefaultParametersInput{}
	result, err := client.DescribeEngineDefaultParameters(context.TODO(), input)
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

func describeEventCategories(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeEventCategoriesInput{}
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

func describeEventSubscriptions(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeEventSubscriptionsInput{}
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

func describeEvents(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeEventsInput{}
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

func describeExportTasks(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeExportTasksInput{}
	result, err := client.DescribeExportTasks(context.TODO(), input)
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

func describeGlobalClusters(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeGlobalClustersInput{}
	result, err := client.DescribeGlobalClusters(context.TODO(), input)
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

func describeOptionGroupOptions(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeOptionGroupOptionsInput{}
	result, err := client.DescribeOptionGroupOptions(context.TODO(), input)
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

func describeOptionGroups(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeOptionGroupsInput{}
	result, err := client.DescribeOptionGroups(context.TODO(), input)
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

func describeOrderableDBInstanceOptions(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeOrderableDBInstanceOptionsInput{}
	result, err := client.DescribeOrderableDBInstanceOptions(context.TODO(), input)
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

func describePendingMaintenanceActions(client *rds.Client) map[string]interface{} {
	input := &rds.DescribePendingMaintenanceActionsInput{}
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

func describeReservedDBInstances(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeReservedDBInstancesInput{}
	result, err := client.DescribeReservedDBInstances(context.TODO(), input)
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

func describeReservedDBInstancesOfferings(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeReservedDBInstancesOfferingsInput{}
	result, err := client.DescribeReservedDBInstancesOfferings(context.TODO(), input)
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

func describeSourceRegions(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeSourceRegionsInput{}
	result, err := client.DescribeSourceRegions(context.TODO(), input)
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

func describeValidDBInstanceModifications(client *rds.Client) map[string]interface{} {
	input := &rds.DescribeValidDBInstanceModificationsInput{}
	result, err := client.DescribeValidDBInstanceModifications(context.TODO(), input)
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

