package gamelift

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := gamelift.NewFromConfig(awscfg)
    
	out(describeAlias(client), "gamelift", "Alias")
    
	out(describeBuild(client), "gamelift", "Build")
    
	out(describeEC2InstanceLimits(client), "gamelift", "EC2InstanceLimits")
    
	out(describeFleetAttributes(client), "gamelift", "FleetAttributes")
    
	out(describeFleetCapacity(client), "gamelift", "FleetCapacity")
    
	out(describeFleetEvents(client), "gamelift", "FleetEvents")
    
	out(describeFleetLocationAttributes(client), "gamelift", "FleetLocationAttributes")
    
	out(describeFleetLocationCapacity(client), "gamelift", "FleetLocationCapacity")
    
	out(describeFleetLocationUtilization(client), "gamelift", "FleetLocationUtilization")
    
	out(describeFleetPortSettings(client), "gamelift", "FleetPortSettings")
    
	out(describeFleetUtilization(client), "gamelift", "FleetUtilization")
    
	out(describeGameServer(client), "gamelift", "GameServer")
    
	out(describeGameServerGroup(client), "gamelift", "GameServerGroup")
    
	out(describeGameServerInstances(client), "gamelift", "GameServerInstances")
    
	out(describeGameSessionDetails(client), "gamelift", "GameSessionDetails")
    
	out(describeGameSessionPlacement(client), "gamelift", "GameSessionPlacement")
    
	out(describeGameSessionQueues(client), "gamelift", "GameSessionQueues")
    
	out(describeGameSessions(client), "gamelift", "GameSessions")
    
	out(describeInstances(client), "gamelift", "Instances")
    
	out(describeMatchmaking(client), "gamelift", "Matchmaking")
    
	out(describeMatchmakingConfigurations(client), "gamelift", "MatchmakingConfigurations")
    
	out(describeMatchmakingRuleSets(client), "gamelift", "MatchmakingRuleSets")
    
	out(describePlayerSessions(client), "gamelift", "PlayerSessions")
    
	out(describeRuntimeConfiguration(client), "gamelift", "RuntimeConfiguration")
    
	out(describeScalingPolicies(client), "gamelift", "ScalingPolicies")
    
	out(describeScript(client), "gamelift", "Script")
    
	out(describeVpcPeeringAuthorizations(client), "gamelift", "VpcPeeringAuthorizations")
    
	out(describeVpcPeeringConnections(client), "gamelift", "VpcPeeringConnections")
    
}

func describeAlias(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeAliasInput{}
	result, err := client.DescribeAlias(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBuild(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeBuildInput{}
	result, err := client.DescribeBuild(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeEC2InstanceLimits(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeEC2InstanceLimitsInput{}
	result, err := client.DescribeEC2InstanceLimits(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetAttributes(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetAttributesInput{}
	result, err := client.DescribeFleetAttributes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetCapacity(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetCapacityInput{}
	result, err := client.DescribeFleetCapacity(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetEvents(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetEventsInput{}
	result, err := client.DescribeFleetEvents(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetLocationAttributes(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetLocationAttributesInput{}
	result, err := client.DescribeFleetLocationAttributes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetLocationCapacity(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetLocationCapacityInput{}
	result, err := client.DescribeFleetLocationCapacity(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetLocationUtilization(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetLocationUtilizationInput{}
	result, err := client.DescribeFleetLocationUtilization(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetPortSettings(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetPortSettingsInput{}
	result, err := client.DescribeFleetPortSettings(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetUtilization(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeFleetUtilizationInput{}
	result, err := client.DescribeFleetUtilization(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeGameServer(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeGameServerInput{}
	result, err := client.DescribeGameServer(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeGameServerGroup(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeGameServerGroupInput{}
	result, err := client.DescribeGameServerGroup(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeGameServerInstances(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeGameServerInstancesInput{}
	result, err := client.DescribeGameServerInstances(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeGameSessionDetails(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeGameSessionDetailsInput{}
	result, err := client.DescribeGameSessionDetails(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeGameSessionPlacement(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeGameSessionPlacementInput{}
	result, err := client.DescribeGameSessionPlacement(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeGameSessionQueues(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeGameSessionQueuesInput{}
	result, err := client.DescribeGameSessionQueues(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeGameSessions(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeGameSessionsInput{}
	result, err := client.DescribeGameSessions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInstances(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeInstancesInput{}
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

func describeMatchmaking(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeMatchmakingInput{}
	result, err := client.DescribeMatchmaking(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMatchmakingConfigurations(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeMatchmakingConfigurationsInput{}
	result, err := client.DescribeMatchmakingConfigurations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMatchmakingRuleSets(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeMatchmakingRuleSetsInput{}
	result, err := client.DescribeMatchmakingRuleSets(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePlayerSessions(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribePlayerSessionsInput{}
	result, err := client.DescribePlayerSessions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeRuntimeConfiguration(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeRuntimeConfigurationInput{}
	result, err := client.DescribeRuntimeConfiguration(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeScalingPolicies(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeScalingPoliciesInput{}
	result, err := client.DescribeScalingPolicies(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeScript(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeScriptInput{}
	result, err := client.DescribeScript(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeVpcPeeringAuthorizations(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeVpcPeeringAuthorizationsInput{}
	result, err := client.DescribeVpcPeeringAuthorizations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeVpcPeeringConnections(client *gamelift.Client) map[string]interface{} {
	input := &gamelift.DescribeVpcPeeringConnectionsInput{}
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

