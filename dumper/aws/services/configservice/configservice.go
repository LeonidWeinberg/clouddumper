package configservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := configservice.NewFromConfig(awscfg)
    
	out(describeAggregateComplianceByConfigRules(client), "configservice", "AggregateComplianceByConfigRules")
    
	out(describeAggregateComplianceByConformancePacks(client), "configservice", "AggregateComplianceByConformancePacks")
    
	out(describeAggregationAuthorizations(client), "configservice", "AggregationAuthorizations")
    
	out(describeComplianceByConfigRule(client), "configservice", "ComplianceByConfigRule")
    
	out(describeComplianceByResource(client), "configservice", "ComplianceByResource")
    
	out(describeConfigRuleEvaluationStatus(client), "configservice", "ConfigRuleEvaluationStatus")
    
	out(describeConfigRules(client), "configservice", "ConfigRules")
    
	out(describeConfigurationAggregatorSourcesStatus(client), "configservice", "ConfigurationAggregatorSourcesStatus")
    
	out(describeConfigurationAggregators(client), "configservice", "ConfigurationAggregators")
    
	out(describeConfigurationRecorderStatus(client), "configservice", "ConfigurationRecorderStatus")
    
	out(describeConfigurationRecorders(client), "configservice", "ConfigurationRecorders")
    
	out(describeConformancePackCompliance(client), "configservice", "ConformancePackCompliance")
    
	out(describeConformancePackStatus(client), "configservice", "ConformancePackStatus")
    
	out(describeConformancePacks(client), "configservice", "ConformancePacks")
    
	out(describeDeliveryChannelStatus(client), "configservice", "DeliveryChannelStatus")
    
	out(describeDeliveryChannels(client), "configservice", "DeliveryChannels")
    
	out(describeOrganizationConfigRuleStatuses(client), "configservice", "OrganizationConfigRuleStatuses")
    
	out(describeOrganizationConfigRules(client), "configservice", "OrganizationConfigRules")
    
	out(describeOrganizationConformancePackStatuses(client), "configservice", "OrganizationConformancePackStatuses")
    
	out(describeOrganizationConformancePacks(client), "configservice", "OrganizationConformancePacks")
    
	out(describePendingAggregationRequests(client), "configservice", "PendingAggregationRequests")
    
	out(describeRemediationConfigurations(client), "configservice", "RemediationConfigurations")
    
	out(describeRemediationExceptions(client), "configservice", "RemediationExceptions")
    
	out(describeRemediationExecutionStatus(client), "configservice", "RemediationExecutionStatus")
    
	out(describeRetentionConfigurations(client), "configservice", "RetentionConfigurations")
    
}

func describeAggregateComplianceByConfigRules(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeAggregateComplianceByConfigRulesInput{}
	result, err := client.DescribeAggregateComplianceByConfigRules(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAggregateComplianceByConformancePacks(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeAggregateComplianceByConformancePacksInput{}
	result, err := client.DescribeAggregateComplianceByConformancePacks(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAggregationAuthorizations(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeAggregationAuthorizationsInput{}
	result, err := client.DescribeAggregationAuthorizations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeComplianceByConfigRule(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeComplianceByConfigRuleInput{}
	result, err := client.DescribeComplianceByConfigRule(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeComplianceByResource(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeComplianceByResourceInput{}
	result, err := client.DescribeComplianceByResource(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConfigRuleEvaluationStatus(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConfigRuleEvaluationStatusInput{}
	result, err := client.DescribeConfigRuleEvaluationStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConfigRules(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConfigRulesInput{}
	result, err := client.DescribeConfigRules(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConfigurationAggregatorSourcesStatus(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConfigurationAggregatorSourcesStatusInput{}
	result, err := client.DescribeConfigurationAggregatorSourcesStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConfigurationAggregators(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConfigurationAggregatorsInput{}
	result, err := client.DescribeConfigurationAggregators(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConfigurationRecorderStatus(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConfigurationRecorderStatusInput{}
	result, err := client.DescribeConfigurationRecorderStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConfigurationRecorders(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConfigurationRecordersInput{}
	result, err := client.DescribeConfigurationRecorders(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConformancePackCompliance(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConformancePackComplianceInput{}
	result, err := client.DescribeConformancePackCompliance(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConformancePackStatus(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConformancePackStatusInput{}
	result, err := client.DescribeConformancePackStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeConformancePacks(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeConformancePacksInput{}
	result, err := client.DescribeConformancePacks(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDeliveryChannelStatus(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeDeliveryChannelStatusInput{}
	result, err := client.DescribeDeliveryChannelStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDeliveryChannels(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeDeliveryChannelsInput{}
	result, err := client.DescribeDeliveryChannels(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeOrganizationConfigRuleStatuses(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeOrganizationConfigRuleStatusesInput{}
	result, err := client.DescribeOrganizationConfigRuleStatuses(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeOrganizationConfigRules(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeOrganizationConfigRulesInput{}
	result, err := client.DescribeOrganizationConfigRules(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeOrganizationConformancePackStatuses(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeOrganizationConformancePackStatusesInput{}
	result, err := client.DescribeOrganizationConformancePackStatuses(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeOrganizationConformancePacks(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeOrganizationConformancePacksInput{}
	result, err := client.DescribeOrganizationConformancePacks(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePendingAggregationRequests(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribePendingAggregationRequestsInput{}
	result, err := client.DescribePendingAggregationRequests(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeRemediationConfigurations(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeRemediationConfigurationsInput{}
	result, err := client.DescribeRemediationConfigurations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeRemediationExceptions(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeRemediationExceptionsInput{}
	result, err := client.DescribeRemediationExceptions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeRemediationExecutionStatus(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeRemediationExecutionStatusInput{}
	result, err := client.DescribeRemediationExecutionStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeRetentionConfigurations(client *configservice.Client) map[string]interface{} {
	input := &configservice.DescribeRetentionConfigurationsInput{}
	result, err := client.DescribeRetentionConfigurations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

