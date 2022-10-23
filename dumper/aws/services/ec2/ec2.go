package ec2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ec2.NewFromConfig(awscfg)
    
	out(describeAccountAttributes(client), "ec2", "AccountAttributes")
    
	out(describeAddresses(client), "ec2", "Addresses")
    
	out(describeAddressesAttribute(client), "ec2", "AddressesAttribute")
    
	out(describeAggregateIdFormat(client), "ec2", "AggregateIdFormat")
    
	out(describeAvailabilityZones(client), "ec2", "AvailabilityZones")
    
	out(describeBundleTasks(client), "ec2", "BundleTasks")
    
	out(describeByoipCidrs(client), "ec2", "ByoipCidrs")
    
	out(describeCapacityReservationFleets(client), "ec2", "CapacityReservationFleets")
    
	out(describeCapacityReservations(client), "ec2", "CapacityReservations")
    
	out(describeCarrierGateways(client), "ec2", "CarrierGateways")
    
	out(describeClassicLinkInstances(client), "ec2", "ClassicLinkInstances")
    
	out(describeClientVpnAuthorizationRules(client), "ec2", "ClientVpnAuthorizationRules")
    
	out(describeClientVpnConnections(client), "ec2", "ClientVpnConnections")
    
	out(describeClientVpnEndpoints(client), "ec2", "ClientVpnEndpoints")
    
	out(describeClientVpnRoutes(client), "ec2", "ClientVpnRoutes")
    
	out(describeClientVpnTargetNetworks(client), "ec2", "ClientVpnTargetNetworks")
    
	out(describeCoipPools(client), "ec2", "CoipPools")
    
	out(describeConversionTasks(client), "ec2", "ConversionTasks")
    
	out(describeCustomerGateways(client), "ec2", "CustomerGateways")
    
	out(describeDhcpOptions(client), "ec2", "DhcpOptions")
    
	out(describeEgressOnlyInternetGateways(client), "ec2", "EgressOnlyInternetGateways")
    
	out(describeElasticGpus(client), "ec2", "ElasticGpus")
    
	out(describeExportImageTasks(client), "ec2", "ExportImageTasks")
    
	out(describeExportTasks(client), "ec2", "ExportTasks")
    
	out(describeFastLaunchImages(client), "ec2", "FastLaunchImages")
    
	out(describeFastSnapshotRestores(client), "ec2", "FastSnapshotRestores")
    
	out(describeFleetHistory(client), "ec2", "FleetHistory")
    
	out(describeFleetInstances(client), "ec2", "FleetInstances")
    
	out(describeFleets(client), "ec2", "Fleets")
    
	out(describeFlowLogs(client), "ec2", "FlowLogs")
    
	out(describeFpgaImageAttribute(client), "ec2", "FpgaImageAttribute")
    
	out(describeFpgaImages(client), "ec2", "FpgaImages")
    
	out(describeHostReservationOfferings(client), "ec2", "HostReservationOfferings")
    
	out(describeHostReservations(client), "ec2", "HostReservations")
    
	out(describeHosts(client), "ec2", "Hosts")
    
	out(describeIamInstanceProfileAssociations(client), "ec2", "IamInstanceProfileAssociations")
    
	out(describeIdFormat(client), "ec2", "IdFormat")
    
	out(describeIdentityIdFormat(client), "ec2", "IdentityIdFormat")
    
	out(describeImageAttribute(client), "ec2", "ImageAttribute")
    
	out(describeImages(client), "ec2", "Images")
    
	out(describeImportImageTasks(client), "ec2", "ImportImageTasks")
    
	out(describeImportSnapshotTasks(client), "ec2", "ImportSnapshotTasks")
    
	out(describeInstanceAttribute(client), "ec2", "InstanceAttribute")
    
	out(describeInstanceCreditSpecifications(client), "ec2", "InstanceCreditSpecifications")
    
	out(describeInstanceEventNotificationAttributes(client), "ec2", "InstanceEventNotificationAttributes")
    
	out(describeInstanceEventWindows(client), "ec2", "InstanceEventWindows")
    
	out(describeInstanceStatus(client), "ec2", "InstanceStatus")
    
	out(describeInstanceTypeOfferings(client), "ec2", "InstanceTypeOfferings")
    
	out(describeInstanceTypes(client), "ec2", "InstanceTypes")
    
	out(describeInstances(client), "ec2", "Instances")
    
	out(describeInternetGateways(client), "ec2", "InternetGateways")
    
	out(describeIpamPools(client), "ec2", "IpamPools")
    
	out(describeIpamScopes(client), "ec2", "IpamScopes")
    
	out(describeIpams(client), "ec2", "Ipams")
    
	out(describeIpv6Pools(client), "ec2", "Ipv6Pools")
    
	out(describeKeyPairs(client), "ec2", "KeyPairs")
    
	out(describeLaunchTemplateVersions(client), "ec2", "LaunchTemplateVersions")
    
	out(describeLaunchTemplates(client), "ec2", "LaunchTemplates")
    
	out(describeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(client), "ec2", "LocalGatewayRouteTableVirtualInterfaceGroupAssociations")
    
	out(describeLocalGatewayRouteTableVpcAssociations(client), "ec2", "LocalGatewayRouteTableVpcAssociations")
    
	out(describeLocalGatewayRouteTables(client), "ec2", "LocalGatewayRouteTables")
    
	out(describeLocalGatewayVirtualInterfaceGroups(client), "ec2", "LocalGatewayVirtualInterfaceGroups")
    
	out(describeLocalGatewayVirtualInterfaces(client), "ec2", "LocalGatewayVirtualInterfaces")
    
	out(describeLocalGateways(client), "ec2", "LocalGateways")
    
	out(describeManagedPrefixLists(client), "ec2", "ManagedPrefixLists")
    
	out(describeMovingAddresses(client), "ec2", "MovingAddresses")
    
	out(describeNatGateways(client), "ec2", "NatGateways")
    
	out(describeNetworkAcls(client), "ec2", "NetworkAcls")
    
	out(describeNetworkInsightsAccessScopeAnalyses(client), "ec2", "NetworkInsightsAccessScopeAnalyses")
    
	out(describeNetworkInsightsAccessScopes(client), "ec2", "NetworkInsightsAccessScopes")
    
	out(describeNetworkInsightsAnalyses(client), "ec2", "NetworkInsightsAnalyses")
    
	out(describeNetworkInsightsPaths(client), "ec2", "NetworkInsightsPaths")
    
	out(describeNetworkInterfaceAttribute(client), "ec2", "NetworkInterfaceAttribute")
    
	out(describeNetworkInterfacePermissions(client), "ec2", "NetworkInterfacePermissions")
    
	out(describeNetworkInterfaces(client), "ec2", "NetworkInterfaces")
    
	out(describePlacementGroups(client), "ec2", "PlacementGroups")
    
	out(describePrefixLists(client), "ec2", "PrefixLists")
    
	out(describePrincipalIdFormat(client), "ec2", "PrincipalIdFormat")
    
	out(describePublicIpv4Pools(client), "ec2", "PublicIpv4Pools")
    
	out(describeRegions(client), "ec2", "Regions")
    
	out(describeReplaceRootVolumeTasks(client), "ec2", "ReplaceRootVolumeTasks")
    
	out(describeReservedInstances(client), "ec2", "ReservedInstances")
    
	out(describeReservedInstancesListings(client), "ec2", "ReservedInstancesListings")
    
	out(describeReservedInstancesModifications(client), "ec2", "ReservedInstancesModifications")
    
	out(describeReservedInstancesOfferings(client), "ec2", "ReservedInstancesOfferings")
    
	out(describeRouteTables(client), "ec2", "RouteTables")
    
	out(describeScheduledInstanceAvailability(client), "ec2", "ScheduledInstanceAvailability")
    
	out(describeScheduledInstances(client), "ec2", "ScheduledInstances")
    
	out(describeSecurityGroupReferences(client), "ec2", "SecurityGroupReferences")
    
	out(describeSecurityGroupRules(client), "ec2", "SecurityGroupRules")
    
	out(describeSecurityGroups(client), "ec2", "SecurityGroups")
    
	out(describeSnapshotAttribute(client), "ec2", "SnapshotAttribute")
    
	out(describeSnapshotTierStatus(client), "ec2", "SnapshotTierStatus")
    
	out(describeSnapshots(client), "ec2", "Snapshots")
    
	out(describeSpotDatafeedSubscription(client), "ec2", "SpotDatafeedSubscription")
    
	out(describeSpotFleetInstances(client), "ec2", "SpotFleetInstances")
    
	out(describeSpotFleetRequestHistory(client), "ec2", "SpotFleetRequestHistory")
    
	out(describeSpotFleetRequests(client), "ec2", "SpotFleetRequests")
    
	out(describeSpotInstanceRequests(client), "ec2", "SpotInstanceRequests")
    
	out(describeSpotPriceHistory(client), "ec2", "SpotPriceHistory")
    
	out(describeStaleSecurityGroups(client), "ec2", "StaleSecurityGroups")
    
	out(describeStoreImageTasks(client), "ec2", "StoreImageTasks")
    
	out(describeSubnets(client), "ec2", "Subnets")
    
	out(describeTags(client), "ec2", "Tags")
    
	out(describeTrafficMirrorFilters(client), "ec2", "TrafficMirrorFilters")
    
	out(describeTrafficMirrorSessions(client), "ec2", "TrafficMirrorSessions")
    
	out(describeTrafficMirrorTargets(client), "ec2", "TrafficMirrorTargets")
    
	out(describeTransitGatewayAttachments(client), "ec2", "TransitGatewayAttachments")
    
	out(describeTransitGatewayConnectPeers(client), "ec2", "TransitGatewayConnectPeers")
    
	out(describeTransitGatewayConnects(client), "ec2", "TransitGatewayConnects")
    
	out(describeTransitGatewayMulticastDomains(client), "ec2", "TransitGatewayMulticastDomains")
    
	out(describeTransitGatewayPeeringAttachments(client), "ec2", "TransitGatewayPeeringAttachments")
    
	out(describeTransitGatewayPolicyTables(client), "ec2", "TransitGatewayPolicyTables")
    
	out(describeTransitGatewayRouteTableAnnouncements(client), "ec2", "TransitGatewayRouteTableAnnouncements")
    
	out(describeTransitGatewayRouteTables(client), "ec2", "TransitGatewayRouteTables")
    
	out(describeTransitGatewayVpcAttachments(client), "ec2", "TransitGatewayVpcAttachments")
    
	out(describeTransitGateways(client), "ec2", "TransitGateways")
    
	out(describeTrunkInterfaceAssociations(client), "ec2", "TrunkInterfaceAssociations")
    
	out(describeVolumeAttribute(client), "ec2", "VolumeAttribute")
    
	out(describeVolumeStatus(client), "ec2", "VolumeStatus")
    
	out(describeVolumes(client), "ec2", "Volumes")
    
	out(describeVolumesModifications(client), "ec2", "VolumesModifications")
    
	out(describeVpcAttribute(client), "ec2", "VpcAttribute")
    
	out(describeVpcClassicLink(client), "ec2", "VpcClassicLink")
    
	out(describeVpcClassicLinkDnsSupport(client), "ec2", "VpcClassicLinkDnsSupport")
    
	out(describeVpcEndpointConnectionNotifications(client), "ec2", "VpcEndpointConnectionNotifications")
    
	out(describeVpcEndpointConnections(client), "ec2", "VpcEndpointConnections")
    
	out(describeVpcEndpointServiceConfigurations(client), "ec2", "VpcEndpointServiceConfigurations")
    
	out(describeVpcEndpointServicePermissions(client), "ec2", "VpcEndpointServicePermissions")
    
	out(describeVpcEndpointServices(client), "ec2", "VpcEndpointServices")
    
	out(describeVpcEndpoints(client), "ec2", "VpcEndpoints")
    
	out(describeVpcPeeringConnections(client), "ec2", "VpcPeeringConnections")
    
	out(describeVpcs(client), "ec2", "Vpcs")
    
	out(describeVpnConnections(client), "ec2", "VpnConnections")
    
	out(describeVpnGateways(client), "ec2", "VpnGateways")
    
}

func describeAccountAttributes(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeAccountAttributesInput{}
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

func describeAddresses(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeAddressesInput{}
	result, err := client.DescribeAddresses(context.TODO(), input)
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

func describeAddressesAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeAddressesAttributeInput{}
	result, err := client.DescribeAddressesAttribute(context.TODO(), input)
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

func describeAggregateIdFormat(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeAggregateIdFormatInput{}
	result, err := client.DescribeAggregateIdFormat(context.TODO(), input)
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

func describeAvailabilityZones(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeAvailabilityZonesInput{}
	result, err := client.DescribeAvailabilityZones(context.TODO(), input)
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

func describeBundleTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeBundleTasksInput{}
	result, err := client.DescribeBundleTasks(context.TODO(), input)
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

func describeByoipCidrs(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeByoipCidrsInput{}
	result, err := client.DescribeByoipCidrs(context.TODO(), input)
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

func describeCapacityReservationFleets(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeCapacityReservationFleetsInput{}
	result, err := client.DescribeCapacityReservationFleets(context.TODO(), input)
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

func describeCapacityReservations(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeCapacityReservationsInput{}
	result, err := client.DescribeCapacityReservations(context.TODO(), input)
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

func describeCarrierGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeCarrierGatewaysInput{}
	result, err := client.DescribeCarrierGateways(context.TODO(), input)
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

func describeClassicLinkInstances(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeClassicLinkInstancesInput{}
	result, err := client.DescribeClassicLinkInstances(context.TODO(), input)
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

func describeClientVpnAuthorizationRules(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeClientVpnAuthorizationRulesInput{}
	result, err := client.DescribeClientVpnAuthorizationRules(context.TODO(), input)
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

func describeClientVpnConnections(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeClientVpnConnectionsInput{}
	result, err := client.DescribeClientVpnConnections(context.TODO(), input)
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

func describeClientVpnEndpoints(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeClientVpnEndpointsInput{}
	result, err := client.DescribeClientVpnEndpoints(context.TODO(), input)
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

func describeClientVpnRoutes(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeClientVpnRoutesInput{}
	result, err := client.DescribeClientVpnRoutes(context.TODO(), input)
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

func describeClientVpnTargetNetworks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeClientVpnTargetNetworksInput{}
	result, err := client.DescribeClientVpnTargetNetworks(context.TODO(), input)
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

func describeCoipPools(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeCoipPoolsInput{}
	result, err := client.DescribeCoipPools(context.TODO(), input)
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

func describeConversionTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeConversionTasksInput{}
	result, err := client.DescribeConversionTasks(context.TODO(), input)
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

func describeCustomerGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeCustomerGatewaysInput{}
	result, err := client.DescribeCustomerGateways(context.TODO(), input)
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

func describeDhcpOptions(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeDhcpOptionsInput{}
	result, err := client.DescribeDhcpOptions(context.TODO(), input)
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

func describeEgressOnlyInternetGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeEgressOnlyInternetGatewaysInput{}
	result, err := client.DescribeEgressOnlyInternetGateways(context.TODO(), input)
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

func describeElasticGpus(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeElasticGpusInput{}
	result, err := client.DescribeElasticGpus(context.TODO(), input)
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

func describeExportImageTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeExportImageTasksInput{}
	result, err := client.DescribeExportImageTasks(context.TODO(), input)
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

func describeExportTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeExportTasksInput{}
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

func describeFastLaunchImages(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFastLaunchImagesInput{}
	result, err := client.DescribeFastLaunchImages(context.TODO(), input)
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

func describeFastSnapshotRestores(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFastSnapshotRestoresInput{}
	result, err := client.DescribeFastSnapshotRestores(context.TODO(), input)
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

func describeFleetHistory(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFleetHistoryInput{}
	result, err := client.DescribeFleetHistory(context.TODO(), input)
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

func describeFleetInstances(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFleetInstancesInput{}
	result, err := client.DescribeFleetInstances(context.TODO(), input)
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

func describeFleets(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFleetsInput{}
	result, err := client.DescribeFleets(context.TODO(), input)
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

func describeFlowLogs(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFlowLogsInput{}
	result, err := client.DescribeFlowLogs(context.TODO(), input)
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

func describeFpgaImageAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFpgaImageAttributeInput{}
	result, err := client.DescribeFpgaImageAttribute(context.TODO(), input)
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

func describeFpgaImages(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeFpgaImagesInput{}
	result, err := client.DescribeFpgaImages(context.TODO(), input)
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

func describeHostReservationOfferings(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeHostReservationOfferingsInput{}
	result, err := client.DescribeHostReservationOfferings(context.TODO(), input)
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

func describeHostReservations(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeHostReservationsInput{}
	result, err := client.DescribeHostReservations(context.TODO(), input)
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

func describeHosts(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeHostsInput{}
	result, err := client.DescribeHosts(context.TODO(), input)
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

func describeIamInstanceProfileAssociations(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeIamInstanceProfileAssociationsInput{}
	result, err := client.DescribeIamInstanceProfileAssociations(context.TODO(), input)
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

func describeIdFormat(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeIdFormatInput{}
	result, err := client.DescribeIdFormat(context.TODO(), input)
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

func describeIdentityIdFormat(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeIdentityIdFormatInput{}
	result, err := client.DescribeIdentityIdFormat(context.TODO(), input)
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

func describeImageAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeImageAttributeInput{}
	result, err := client.DescribeImageAttribute(context.TODO(), input)
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

func describeImages(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeImagesInput{}
	result, err := client.DescribeImages(context.TODO(), input)
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

func describeImportImageTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeImportImageTasksInput{}
	result, err := client.DescribeImportImageTasks(context.TODO(), input)
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

func describeImportSnapshotTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeImportSnapshotTasksInput{}
	result, err := client.DescribeImportSnapshotTasks(context.TODO(), input)
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

func describeInstanceAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstanceAttributeInput{}
	result, err := client.DescribeInstanceAttribute(context.TODO(), input)
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

func describeInstanceCreditSpecifications(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstanceCreditSpecificationsInput{}
	result, err := client.DescribeInstanceCreditSpecifications(context.TODO(), input)
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

func describeInstanceEventNotificationAttributes(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstanceEventNotificationAttributesInput{}
	result, err := client.DescribeInstanceEventNotificationAttributes(context.TODO(), input)
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

func describeInstanceEventWindows(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstanceEventWindowsInput{}
	result, err := client.DescribeInstanceEventWindows(context.TODO(), input)
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

func describeInstanceStatus(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstanceStatusInput{}
	result, err := client.DescribeInstanceStatus(context.TODO(), input)
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

func describeInstanceTypeOfferings(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstanceTypeOfferingsInput{}
	result, err := client.DescribeInstanceTypeOfferings(context.TODO(), input)
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

func describeInstanceTypes(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstanceTypesInput{}
	result, err := client.DescribeInstanceTypes(context.TODO(), input)
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

func describeInstances(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInstancesInput{}
	result, err := client.DescribeInstances(context.TODO(), input)
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

func describeInternetGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeInternetGatewaysInput{}
	result, err := client.DescribeInternetGateways(context.TODO(), input)
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

func describeIpamPools(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeIpamPoolsInput{}
	result, err := client.DescribeIpamPools(context.TODO(), input)
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

func describeIpamScopes(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeIpamScopesInput{}
	result, err := client.DescribeIpamScopes(context.TODO(), input)
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

func describeIpams(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeIpamsInput{}
	result, err := client.DescribeIpams(context.TODO(), input)
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

func describeIpv6Pools(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeIpv6PoolsInput{}
	result, err := client.DescribeIpv6Pools(context.TODO(), input)
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

func describeKeyPairs(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeKeyPairsInput{}
	result, err := client.DescribeKeyPairs(context.TODO(), input)
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

func describeLaunchTemplateVersions(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLaunchTemplateVersionsInput{}
	result, err := client.DescribeLaunchTemplateVersions(context.TODO(), input)
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

func describeLaunchTemplates(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLaunchTemplatesInput{}
	result, err := client.DescribeLaunchTemplates(context.TODO(), input)
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

func describeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput{}
	result, err := client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(context.TODO(), input)
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

func describeLocalGatewayRouteTableVpcAssociations(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput{}
	result, err := client.DescribeLocalGatewayRouteTableVpcAssociations(context.TODO(), input)
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

func describeLocalGatewayRouteTables(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLocalGatewayRouteTablesInput{}
	result, err := client.DescribeLocalGatewayRouteTables(context.TODO(), input)
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

func describeLocalGatewayVirtualInterfaceGroups(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput{}
	result, err := client.DescribeLocalGatewayVirtualInterfaceGroups(context.TODO(), input)
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

func describeLocalGatewayVirtualInterfaces(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLocalGatewayVirtualInterfacesInput{}
	result, err := client.DescribeLocalGatewayVirtualInterfaces(context.TODO(), input)
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

func describeLocalGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeLocalGatewaysInput{}
	result, err := client.DescribeLocalGateways(context.TODO(), input)
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

func describeManagedPrefixLists(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeManagedPrefixListsInput{}
	result, err := client.DescribeManagedPrefixLists(context.TODO(), input)
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

func describeMovingAddresses(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeMovingAddressesInput{}
	result, err := client.DescribeMovingAddresses(context.TODO(), input)
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

func describeNatGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNatGatewaysInput{}
	result, err := client.DescribeNatGateways(context.TODO(), input)
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

func describeNetworkAcls(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkAclsInput{}
	result, err := client.DescribeNetworkAcls(context.TODO(), input)
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

func describeNetworkInsightsAccessScopeAnalyses(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkInsightsAccessScopeAnalysesInput{}
	result, err := client.DescribeNetworkInsightsAccessScopeAnalyses(context.TODO(), input)
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

func describeNetworkInsightsAccessScopes(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkInsightsAccessScopesInput{}
	result, err := client.DescribeNetworkInsightsAccessScopes(context.TODO(), input)
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

func describeNetworkInsightsAnalyses(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkInsightsAnalysesInput{}
	result, err := client.DescribeNetworkInsightsAnalyses(context.TODO(), input)
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

func describeNetworkInsightsPaths(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkInsightsPathsInput{}
	result, err := client.DescribeNetworkInsightsPaths(context.TODO(), input)
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

func describeNetworkInterfaceAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkInterfaceAttributeInput{}
	result, err := client.DescribeNetworkInterfaceAttribute(context.TODO(), input)
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

func describeNetworkInterfacePermissions(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkInterfacePermissionsInput{}
	result, err := client.DescribeNetworkInterfacePermissions(context.TODO(), input)
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

func describeNetworkInterfaces(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeNetworkInterfacesInput{}
	result, err := client.DescribeNetworkInterfaces(context.TODO(), input)
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

func describePlacementGroups(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribePlacementGroupsInput{}
	result, err := client.DescribePlacementGroups(context.TODO(), input)
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

func describePrefixLists(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribePrefixListsInput{}
	result, err := client.DescribePrefixLists(context.TODO(), input)
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

func describePrincipalIdFormat(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribePrincipalIdFormatInput{}
	result, err := client.DescribePrincipalIdFormat(context.TODO(), input)
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

func describePublicIpv4Pools(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribePublicIpv4PoolsInput{}
	result, err := client.DescribePublicIpv4Pools(context.TODO(), input)
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

func describeRegions(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeRegionsInput{}
	result, err := client.DescribeRegions(context.TODO(), input)
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

func describeReplaceRootVolumeTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeReplaceRootVolumeTasksInput{}
	result, err := client.DescribeReplaceRootVolumeTasks(context.TODO(), input)
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

func describeReservedInstances(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeReservedInstancesInput{}
	result, err := client.DescribeReservedInstances(context.TODO(), input)
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

func describeReservedInstancesListings(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeReservedInstancesListingsInput{}
	result, err := client.DescribeReservedInstancesListings(context.TODO(), input)
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

func describeReservedInstancesModifications(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeReservedInstancesModificationsInput{}
	result, err := client.DescribeReservedInstancesModifications(context.TODO(), input)
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

func describeReservedInstancesOfferings(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeReservedInstancesOfferingsInput{}
	result, err := client.DescribeReservedInstancesOfferings(context.TODO(), input)
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

func describeRouteTables(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeRouteTablesInput{}
	result, err := client.DescribeRouteTables(context.TODO(), input)
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

func describeScheduledInstanceAvailability(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeScheduledInstanceAvailabilityInput{}
	result, err := client.DescribeScheduledInstanceAvailability(context.TODO(), input)
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

func describeScheduledInstances(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeScheduledInstancesInput{}
	result, err := client.DescribeScheduledInstances(context.TODO(), input)
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

func describeSecurityGroupReferences(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSecurityGroupReferencesInput{}
	result, err := client.DescribeSecurityGroupReferences(context.TODO(), input)
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

func describeSecurityGroupRules(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSecurityGroupRulesInput{}
	result, err := client.DescribeSecurityGroupRules(context.TODO(), input)
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

func describeSecurityGroups(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSecurityGroupsInput{}
	result, err := client.DescribeSecurityGroups(context.TODO(), input)
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

func describeSnapshotAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSnapshotAttributeInput{}
	result, err := client.DescribeSnapshotAttribute(context.TODO(), input)
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

func describeSnapshotTierStatus(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSnapshotTierStatusInput{}
	result, err := client.DescribeSnapshotTierStatus(context.TODO(), input)
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

func describeSnapshots(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSnapshotsInput{}
	result, err := client.DescribeSnapshots(context.TODO(), input)
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

func describeSpotDatafeedSubscription(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSpotDatafeedSubscriptionInput{}
	result, err := client.DescribeSpotDatafeedSubscription(context.TODO(), input)
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

func describeSpotFleetInstances(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSpotFleetInstancesInput{}
	result, err := client.DescribeSpotFleetInstances(context.TODO(), input)
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

func describeSpotFleetRequestHistory(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSpotFleetRequestHistoryInput{}
	result, err := client.DescribeSpotFleetRequestHistory(context.TODO(), input)
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

func describeSpotFleetRequests(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSpotFleetRequestsInput{}
	result, err := client.DescribeSpotFleetRequests(context.TODO(), input)
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

func describeSpotInstanceRequests(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSpotInstanceRequestsInput{}
	result, err := client.DescribeSpotInstanceRequests(context.TODO(), input)
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

func describeSpotPriceHistory(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSpotPriceHistoryInput{}
	result, err := client.DescribeSpotPriceHistory(context.TODO(), input)
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

func describeStaleSecurityGroups(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeStaleSecurityGroupsInput{}
	result, err := client.DescribeStaleSecurityGroups(context.TODO(), input)
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

func describeStoreImageTasks(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeStoreImageTasksInput{}
	result, err := client.DescribeStoreImageTasks(context.TODO(), input)
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

func describeSubnets(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeSubnetsInput{}
	result, err := client.DescribeSubnets(context.TODO(), input)
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

func describeTags(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTagsInput{}
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

func describeTrafficMirrorFilters(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTrafficMirrorFiltersInput{}
	result, err := client.DescribeTrafficMirrorFilters(context.TODO(), input)
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

func describeTrafficMirrorSessions(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTrafficMirrorSessionsInput{}
	result, err := client.DescribeTrafficMirrorSessions(context.TODO(), input)
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

func describeTrafficMirrorTargets(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTrafficMirrorTargetsInput{}
	result, err := client.DescribeTrafficMirrorTargets(context.TODO(), input)
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

func describeTransitGatewayAttachments(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayAttachmentsInput{}
	result, err := client.DescribeTransitGatewayAttachments(context.TODO(), input)
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

func describeTransitGatewayConnectPeers(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayConnectPeersInput{}
	result, err := client.DescribeTransitGatewayConnectPeers(context.TODO(), input)
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

func describeTransitGatewayConnects(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayConnectsInput{}
	result, err := client.DescribeTransitGatewayConnects(context.TODO(), input)
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

func describeTransitGatewayMulticastDomains(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayMulticastDomainsInput{}
	result, err := client.DescribeTransitGatewayMulticastDomains(context.TODO(), input)
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

func describeTransitGatewayPeeringAttachments(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayPeeringAttachmentsInput{}
	result, err := client.DescribeTransitGatewayPeeringAttachments(context.TODO(), input)
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

func describeTransitGatewayPolicyTables(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayPolicyTablesInput{}
	result, err := client.DescribeTransitGatewayPolicyTables(context.TODO(), input)
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

func describeTransitGatewayRouteTableAnnouncements(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayRouteTableAnnouncementsInput{}
	result, err := client.DescribeTransitGatewayRouteTableAnnouncements(context.TODO(), input)
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

func describeTransitGatewayRouteTables(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayRouteTablesInput{}
	result, err := client.DescribeTransitGatewayRouteTables(context.TODO(), input)
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

func describeTransitGatewayVpcAttachments(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewayVpcAttachmentsInput{}
	result, err := client.DescribeTransitGatewayVpcAttachments(context.TODO(), input)
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

func describeTransitGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTransitGatewaysInput{}
	result, err := client.DescribeTransitGateways(context.TODO(), input)
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

func describeTrunkInterfaceAssociations(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeTrunkInterfaceAssociationsInput{}
	result, err := client.DescribeTrunkInterfaceAssociations(context.TODO(), input)
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

func describeVolumeAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVolumeAttributeInput{}
	result, err := client.DescribeVolumeAttribute(context.TODO(), input)
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

func describeVolumeStatus(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVolumeStatusInput{}
	result, err := client.DescribeVolumeStatus(context.TODO(), input)
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

func describeVolumes(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVolumesInput{}
	result, err := client.DescribeVolumes(context.TODO(), input)
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

func describeVolumesModifications(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVolumesModificationsInput{}
	result, err := client.DescribeVolumesModifications(context.TODO(), input)
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

func describeVpcAttribute(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcAttributeInput{}
	result, err := client.DescribeVpcAttribute(context.TODO(), input)
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

func describeVpcClassicLink(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcClassicLinkInput{}
	result, err := client.DescribeVpcClassicLink(context.TODO(), input)
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

func describeVpcClassicLinkDnsSupport(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcClassicLinkDnsSupportInput{}
	result, err := client.DescribeVpcClassicLinkDnsSupport(context.TODO(), input)
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

func describeVpcEndpointConnectionNotifications(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcEndpointConnectionNotificationsInput{}
	result, err := client.DescribeVpcEndpointConnectionNotifications(context.TODO(), input)
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

func describeVpcEndpointConnections(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcEndpointConnectionsInput{}
	result, err := client.DescribeVpcEndpointConnections(context.TODO(), input)
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

func describeVpcEndpointServiceConfigurations(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcEndpointServiceConfigurationsInput{}
	result, err := client.DescribeVpcEndpointServiceConfigurations(context.TODO(), input)
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

func describeVpcEndpointServicePermissions(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcEndpointServicePermissionsInput{}
	result, err := client.DescribeVpcEndpointServicePermissions(context.TODO(), input)
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

func describeVpcEndpointServices(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcEndpointServicesInput{}
	result, err := client.DescribeVpcEndpointServices(context.TODO(), input)
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

func describeVpcEndpoints(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcEndpointsInput{}
	result, err := client.DescribeVpcEndpoints(context.TODO(), input)
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

func describeVpcPeeringConnections(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcPeeringConnectionsInput{}
	result, err := client.DescribeVpcPeeringConnections(context.TODO(), input)
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

func describeVpcs(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpcsInput{}
	result, err := client.DescribeVpcs(context.TODO(), input)
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

func describeVpnConnections(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpnConnectionsInput{}
	result, err := client.DescribeVpnConnections(context.TODO(), input)
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

func describeVpnGateways(client *ec2.Client) map[string]interface{} {
	input := &ec2.DescribeVpnGatewaysInput{}
	result, err := client.DescribeVpnGateways(context.TODO(), input)
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

