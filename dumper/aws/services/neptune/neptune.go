package neptune

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := neptune.NewFromConfig(awscfg)
    
	out(describeDBClusterEndpoints(client), "neptune", "DBClusterEndpoints")
    
	out(describeDBClusterParameterGroups(client), "neptune", "DBClusterParameterGroups")
    
	out(describeDBClusterParameters(client), "neptune", "DBClusterParameters")
    
	out(describeDBClusterSnapshotAttributes(client), "neptune", "DBClusterSnapshotAttributes")
    
	out(describeDBClusterSnapshots(client), "neptune", "DBClusterSnapshots")
    
	out(describeDBClusters(client), "neptune", "DBClusters")
    
	out(describeDBEngineVersions(client), "neptune", "DBEngineVersions")
    
	out(describeDBInstances(client), "neptune", "DBInstances")
    
	out(describeDBParameterGroups(client), "neptune", "DBParameterGroups")
    
	out(describeDBParameters(client), "neptune", "DBParameters")
    
	out(describeDBSubnetGroups(client), "neptune", "DBSubnetGroups")
    
	out(describeEngineDefaultClusterParameters(client), "neptune", "EngineDefaultClusterParameters")
    
	out(describeEngineDefaultParameters(client), "neptune", "EngineDefaultParameters")
    
	out(describeEventCategories(client), "neptune", "EventCategories")
    
	out(describeEventSubscriptions(client), "neptune", "EventSubscriptions")
    
	out(describeEvents(client), "neptune", "Events")
    
	out(describeGlobalClusters(client), "neptune", "GlobalClusters")
    
	out(describeOrderableDBInstanceOptions(client), "neptune", "OrderableDBInstanceOptions")
    
	out(describePendingMaintenanceActions(client), "neptune", "PendingMaintenanceActions")
    
	out(describeValidDBInstanceModifications(client), "neptune", "ValidDBInstanceModifications")
    
}

func describeDBClusterEndpoints(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBClusterEndpointsInput{}
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

func describeDBClusterParameterGroups(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBClusterParameterGroupsInput{}
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

func describeDBClusterParameters(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBClusterParametersInput{}
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

func describeDBClusterSnapshotAttributes(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBClusterSnapshotAttributesInput{}
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

func describeDBClusterSnapshots(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBClusterSnapshotsInput{}
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

func describeDBClusters(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBClustersInput{}
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

func describeDBEngineVersions(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBEngineVersionsInput{}
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

func describeDBInstances(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBInstancesInput{}
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

func describeDBParameterGroups(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBParameterGroupsInput{}
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

func describeDBParameters(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBParametersInput{}
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

func describeDBSubnetGroups(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeDBSubnetGroupsInput{}
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

func describeEngineDefaultClusterParameters(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeEngineDefaultClusterParametersInput{}
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

func describeEngineDefaultParameters(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeEngineDefaultParametersInput{}
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

func describeEventCategories(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeEventCategoriesInput{}
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

func describeEventSubscriptions(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeEventSubscriptionsInput{}
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

func describeEvents(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeEventsInput{}
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

func describeGlobalClusters(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeGlobalClustersInput{}
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

func describeOrderableDBInstanceOptions(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeOrderableDBInstanceOptionsInput{}
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

func describePendingMaintenanceActions(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribePendingMaintenanceActionsInput{}
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

func describeValidDBInstanceModifications(client *neptune.Client) map[string]interface{} {
	input := &neptune.DescribeValidDBInstanceModificationsInput{}
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

