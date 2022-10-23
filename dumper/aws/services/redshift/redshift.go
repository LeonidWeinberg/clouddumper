package redshift

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := redshift.NewFromConfig(awscfg)
    
	out(describeAccountAttributes(client), "redshift", "AccountAttributes")
    
	out(describeAuthenticationProfiles(client), "redshift", "AuthenticationProfiles")
    
	out(describeClusterDbRevisions(client), "redshift", "ClusterDbRevisions")
    
	out(describeClusterParameterGroups(client), "redshift", "ClusterParameterGroups")
    
	out(describeClusterParameters(client), "redshift", "ClusterParameters")
    
	out(describeClusterSecurityGroups(client), "redshift", "ClusterSecurityGroups")
    
	out(describeClusterSnapshots(client), "redshift", "ClusterSnapshots")
    
	out(describeClusterSubnetGroups(client), "redshift", "ClusterSubnetGroups")
    
	out(describeClusterTracks(client), "redshift", "ClusterTracks")
    
	out(describeClusterVersions(client), "redshift", "ClusterVersions")
    
	out(describeClusters(client), "redshift", "Clusters")
    
	out(describeDataShares(client), "redshift", "DataShares")
    
	out(describeDataSharesForConsumer(client), "redshift", "DataSharesForConsumer")
    
	out(describeDataSharesForProducer(client), "redshift", "DataSharesForProducer")
    
	out(describeDefaultClusterParameters(client), "redshift", "DefaultClusterParameters")
    
	out(describeEndpointAccess(client), "redshift", "EndpointAccess")
    
	out(describeEndpointAuthorization(client), "redshift", "EndpointAuthorization")
    
	out(describeEventCategories(client), "redshift", "EventCategories")
    
	out(describeEventSubscriptions(client), "redshift", "EventSubscriptions")
    
	out(describeEvents(client), "redshift", "Events")
    
	out(describeHsmClientCertificates(client), "redshift", "HsmClientCertificates")
    
	out(describeHsmConfigurations(client), "redshift", "HsmConfigurations")
    
	out(describeLoggingStatus(client), "redshift", "LoggingStatus")
    
	out(describeNodeConfigurationOptions(client), "redshift", "NodeConfigurationOptions")
    
	out(describeOrderableClusterOptions(client), "redshift", "OrderableClusterOptions")
    
	out(describePartners(client), "redshift", "Partners")
    
	out(describeReservedNodeExchangeStatus(client), "redshift", "ReservedNodeExchangeStatus")
    
	out(describeReservedNodeOfferings(client), "redshift", "ReservedNodeOfferings")
    
	out(describeReservedNodes(client), "redshift", "ReservedNodes")
    
	out(describeResize(client), "redshift", "Resize")
    
	out(describeScheduledActions(client), "redshift", "ScheduledActions")
    
	out(describeSnapshotCopyGrants(client), "redshift", "SnapshotCopyGrants")
    
	out(describeSnapshotSchedules(client), "redshift", "SnapshotSchedules")
    
	out(describeStorage(client), "redshift", "Storage")
    
	out(describeTableRestoreStatus(client), "redshift", "TableRestoreStatus")
    
	out(describeTags(client), "redshift", "Tags")
    
	out(describeUsageLimits(client), "redshift", "UsageLimits")
    
}

func describeAccountAttributes(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeAccountAttributesInput{}
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

func describeAuthenticationProfiles(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeAuthenticationProfilesInput{}
	result, err := client.DescribeAuthenticationProfiles(context.TODO(), input)
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

func describeClusterDbRevisions(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterDbRevisionsInput{}
	result, err := client.DescribeClusterDbRevisions(context.TODO(), input)
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

func describeClusterParameterGroups(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterParameterGroupsInput{}
	result, err := client.DescribeClusterParameterGroups(context.TODO(), input)
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

func describeClusterParameters(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterParametersInput{}
	result, err := client.DescribeClusterParameters(context.TODO(), input)
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

func describeClusterSecurityGroups(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterSecurityGroupsInput{}
	result, err := client.DescribeClusterSecurityGroups(context.TODO(), input)
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

func describeClusterSnapshots(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterSnapshotsInput{}
	result, err := client.DescribeClusterSnapshots(context.TODO(), input)
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

func describeClusterSubnetGroups(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterSubnetGroupsInput{}
	result, err := client.DescribeClusterSubnetGroups(context.TODO(), input)
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

func describeClusterTracks(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterTracksInput{}
	result, err := client.DescribeClusterTracks(context.TODO(), input)
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

func describeClusterVersions(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClusterVersionsInput{}
	result, err := client.DescribeClusterVersions(context.TODO(), input)
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

func describeClusters(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeClustersInput{}
	result, err := client.DescribeClusters(context.TODO(), input)
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

func describeDataShares(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeDataSharesInput{}
	result, err := client.DescribeDataShares(context.TODO(), input)
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

func describeDataSharesForConsumer(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeDataSharesForConsumerInput{}
	result, err := client.DescribeDataSharesForConsumer(context.TODO(), input)
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

func describeDataSharesForProducer(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeDataSharesForProducerInput{}
	result, err := client.DescribeDataSharesForProducer(context.TODO(), input)
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

func describeDefaultClusterParameters(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeDefaultClusterParametersInput{}
	result, err := client.DescribeDefaultClusterParameters(context.TODO(), input)
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

func describeEndpointAccess(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeEndpointAccessInput{}
	result, err := client.DescribeEndpointAccess(context.TODO(), input)
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

func describeEndpointAuthorization(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeEndpointAuthorizationInput{}
	result, err := client.DescribeEndpointAuthorization(context.TODO(), input)
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

func describeEventCategories(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeEventCategoriesInput{}
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

func describeEventSubscriptions(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeEventSubscriptionsInput{}
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

func describeEvents(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeEventsInput{}
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

func describeHsmClientCertificates(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeHsmClientCertificatesInput{}
	result, err := client.DescribeHsmClientCertificates(context.TODO(), input)
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

func describeHsmConfigurations(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeHsmConfigurationsInput{}
	result, err := client.DescribeHsmConfigurations(context.TODO(), input)
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

func describeLoggingStatus(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeLoggingStatusInput{}
	result, err := client.DescribeLoggingStatus(context.TODO(), input)
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

func describeNodeConfigurationOptions(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeNodeConfigurationOptionsInput{}
	result, err := client.DescribeNodeConfigurationOptions(context.TODO(), input)
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

func describeOrderableClusterOptions(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeOrderableClusterOptionsInput{}
	result, err := client.DescribeOrderableClusterOptions(context.TODO(), input)
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

func describePartners(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribePartnersInput{}
	result, err := client.DescribePartners(context.TODO(), input)
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

func describeReservedNodeExchangeStatus(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeReservedNodeExchangeStatusInput{}
	result, err := client.DescribeReservedNodeExchangeStatus(context.TODO(), input)
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

func describeReservedNodeOfferings(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeReservedNodeOfferingsInput{}
	result, err := client.DescribeReservedNodeOfferings(context.TODO(), input)
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

func describeReservedNodes(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeReservedNodesInput{}
	result, err := client.DescribeReservedNodes(context.TODO(), input)
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

func describeResize(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeResizeInput{}
	result, err := client.DescribeResize(context.TODO(), input)
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

func describeScheduledActions(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeScheduledActionsInput{}
	result, err := client.DescribeScheduledActions(context.TODO(), input)
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

func describeSnapshotCopyGrants(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeSnapshotCopyGrantsInput{}
	result, err := client.DescribeSnapshotCopyGrants(context.TODO(), input)
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

func describeSnapshotSchedules(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeSnapshotSchedulesInput{}
	result, err := client.DescribeSnapshotSchedules(context.TODO(), input)
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

func describeStorage(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeStorageInput{}
	result, err := client.DescribeStorage(context.TODO(), input)
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

func describeTableRestoreStatus(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeTableRestoreStatusInput{}
	result, err := client.DescribeTableRestoreStatus(context.TODO(), input)
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

func describeTags(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeTagsInput{}
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

func describeUsageLimits(client *redshift.Client) map[string]interface{} {
	input := &redshift.DescribeUsageLimitsInput{}
	result, err := client.DescribeUsageLimits(context.TODO(), input)
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

