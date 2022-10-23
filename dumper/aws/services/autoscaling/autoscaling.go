package autoscaling

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := autoscaling.NewFromConfig(awscfg)
    
	out(describeAccountLimits(client), "autoscaling", "AccountLimits")
    
	out(describeAdjustmentTypes(client), "autoscaling", "AdjustmentTypes")
    
	out(describeAutoScalingGroups(client), "autoscaling", "AutoScalingGroups")
    
	out(describeAutoScalingInstances(client), "autoscaling", "AutoScalingInstances")
    
	out(describeAutoScalingNotificationTypes(client), "autoscaling", "AutoScalingNotificationTypes")
    
	out(describeInstanceRefreshes(client), "autoscaling", "InstanceRefreshes")
    
	out(describeLaunchConfigurations(client), "autoscaling", "LaunchConfigurations")
    
	out(describeLifecycleHookTypes(client), "autoscaling", "LifecycleHookTypes")
    
	out(describeLifecycleHooks(client), "autoscaling", "LifecycleHooks")
    
	out(describeLoadBalancerTargetGroups(client), "autoscaling", "LoadBalancerTargetGroups")
    
	out(describeLoadBalancers(client), "autoscaling", "LoadBalancers")
    
	out(describeMetricCollectionTypes(client), "autoscaling", "MetricCollectionTypes")
    
	out(describeNotificationConfigurations(client), "autoscaling", "NotificationConfigurations")
    
	out(describePolicies(client), "autoscaling", "Policies")
    
	out(describeScalingActivities(client), "autoscaling", "ScalingActivities")
    
	out(describeScalingProcessTypes(client), "autoscaling", "ScalingProcessTypes")
    
	out(describeScheduledActions(client), "autoscaling", "ScheduledActions")
    
	out(describeTags(client), "autoscaling", "Tags")
    
	out(describeTerminationPolicyTypes(client), "autoscaling", "TerminationPolicyTypes")
    
	out(describeWarmPool(client), "autoscaling", "WarmPool")
    
}

func describeAccountLimits(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeAccountLimitsInput{}
	result, err := client.DescribeAccountLimits(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAdjustmentTypes(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeAdjustmentTypesInput{}
	result, err := client.DescribeAdjustmentTypes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAutoScalingGroups(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeAutoScalingGroupsInput{}
	result, err := client.DescribeAutoScalingGroups(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAutoScalingInstances(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeAutoScalingInstancesInput{}
	result, err := client.DescribeAutoScalingInstances(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAutoScalingNotificationTypes(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeAutoScalingNotificationTypesInput{}
	result, err := client.DescribeAutoScalingNotificationTypes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInstanceRefreshes(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeInstanceRefreshesInput{}
	result, err := client.DescribeInstanceRefreshes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeLaunchConfigurations(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeLaunchConfigurationsInput{}
	result, err := client.DescribeLaunchConfigurations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeLifecycleHookTypes(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeLifecycleHookTypesInput{}
	result, err := client.DescribeLifecycleHookTypes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeLifecycleHooks(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeLifecycleHooksInput{}
	result, err := client.DescribeLifecycleHooks(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeLoadBalancerTargetGroups(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeLoadBalancerTargetGroupsInput{}
	result, err := client.DescribeLoadBalancerTargetGroups(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeLoadBalancers(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeLoadBalancersInput{}
	result, err := client.DescribeLoadBalancers(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMetricCollectionTypes(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeMetricCollectionTypesInput{}
	result, err := client.DescribeMetricCollectionTypes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeNotificationConfigurations(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeNotificationConfigurationsInput{}
	result, err := client.DescribeNotificationConfigurations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePolicies(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribePoliciesInput{}
	result, err := client.DescribePolicies(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeScalingActivities(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeScalingActivitiesInput{}
	result, err := client.DescribeScalingActivities(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeScalingProcessTypes(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeScalingProcessTypesInput{}
	result, err := client.DescribeScalingProcessTypes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeScheduledActions(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeScheduledActionsInput{}
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

func describeTags(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeTagsInput{}
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

func describeTerminationPolicyTypes(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeTerminationPolicyTypesInput{}
	result, err := client.DescribeTerminationPolicyTypes(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeWarmPool(client *autoscaling.Client) map[string]interface{} {
	input := &autoscaling.DescribeWarmPoolInput{}
	result, err := client.DescribeWarmPool(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

