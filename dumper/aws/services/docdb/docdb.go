package docdb

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := docdb.NewFromConfig(awscfg)
    
	out(describeCertificates(client), "docdb", "Certificates")
    
	out(describeDBClusterParameterGroups(client), "docdb", "DBClusterParameterGroups")
    
	out(describeDBClusterParameters(client), "docdb", "DBClusterParameters")
    
	out(describeDBClusterSnapshotAttributes(client), "docdb", "DBClusterSnapshotAttributes")
    
	out(describeDBClusterSnapshots(client), "docdb", "DBClusterSnapshots")
    
	out(describeDBClusters(client), "docdb", "DBClusters")
    
	out(describeDBEngineVersions(client), "docdb", "DBEngineVersions")
    
	out(describeDBInstances(client), "docdb", "DBInstances")
    
	out(describeDBSubnetGroups(client), "docdb", "DBSubnetGroups")
    
	out(describeEngineDefaultClusterParameters(client), "docdb", "EngineDefaultClusterParameters")
    
	out(describeEventCategories(client), "docdb", "EventCategories")
    
	out(describeEventSubscriptions(client), "docdb", "EventSubscriptions")
    
	out(describeEvents(client), "docdb", "Events")
    
	out(describeGlobalClusters(client), "docdb", "GlobalClusters")
    
	out(describeOrderableDBInstanceOptions(client), "docdb", "OrderableDBInstanceOptions")
    
	out(describePendingMaintenanceActions(client), "docdb", "PendingMaintenanceActions")
    
}

func describeCertificates(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeCertificatesInput{}
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

func describeDBClusterParameterGroups(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBClusterParameterGroupsInput{}
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

func describeDBClusterParameters(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBClusterParametersInput{}
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

func describeDBClusterSnapshotAttributes(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBClusterSnapshotAttributesInput{}
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

func describeDBClusterSnapshots(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBClusterSnapshotsInput{}
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

func describeDBClusters(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBClustersInput{}
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

func describeDBEngineVersions(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBEngineVersionsInput{}
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

func describeDBInstances(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBInstancesInput{}
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

func describeDBSubnetGroups(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeDBSubnetGroupsInput{}
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

func describeEngineDefaultClusterParameters(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeEngineDefaultClusterParametersInput{}
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

func describeEventCategories(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeEventCategoriesInput{}
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

func describeEventSubscriptions(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeEventSubscriptionsInput{}
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

func describeEvents(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeEventsInput{}
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

func describeGlobalClusters(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeGlobalClustersInput{}
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

func describeOrderableDBInstanceOptions(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribeOrderableDBInstanceOptionsInput{}
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

func describePendingMaintenanceActions(client *docdb.Client) map[string]interface{} {
	input := &docdb.DescribePendingMaintenanceActionsInput{}
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

