package iot

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := iot.NewFromConfig(awscfg)
    
	out(describeAccountAuditConfiguration(client), "iot", "AccountAuditConfiguration")
    
	out(describeAuditFinding(client), "iot", "AuditFinding")
    
	out(describeAuditMitigationActionsTask(client), "iot", "AuditMitigationActionsTask")
    
	out(describeAuditSuppression(client), "iot", "AuditSuppression")
    
	out(describeAuditTask(client), "iot", "AuditTask")
    
	out(describeAuthorizer(client), "iot", "Authorizer")
    
	out(describeBillingGroup(client), "iot", "BillingGroup")
    
	out(describeCACertificate(client), "iot", "CACertificate")
    
	out(describeCertificate(client), "iot", "Certificate")
    
	out(describeCustomMetric(client), "iot", "CustomMetric")
    
	out(describeDefaultAuthorizer(client), "iot", "DefaultAuthorizer")
    
	out(describeDetectMitigationActionsTask(client), "iot", "DetectMitigationActionsTask")
    
	out(describeDimension(client), "iot", "Dimension")
    
	out(describeDomainConfiguration(client), "iot", "DomainConfiguration")
    
	out(describeEndpoint(client), "iot", "Endpoint")
    
	out(describeEventConfigurations(client), "iot", "EventConfigurations")
    
	out(describeFleetMetric(client), "iot", "FleetMetric")
    
	out(describeIndex(client), "iot", "Index")
    
	out(describeJob(client), "iot", "Job")
    
	out(describeJobExecution(client), "iot", "JobExecution")
    
	out(describeJobTemplate(client), "iot", "JobTemplate")
    
	out(describeManagedJobTemplate(client), "iot", "ManagedJobTemplate")
    
	out(describeMitigationAction(client), "iot", "MitigationAction")
    
	out(describeProvisioningTemplate(client), "iot", "ProvisioningTemplate")
    
	out(describeProvisioningTemplateVersion(client), "iot", "ProvisioningTemplateVersion")
    
	out(describeRoleAlias(client), "iot", "RoleAlias")
    
	out(describeScheduledAudit(client), "iot", "ScheduledAudit")
    
	out(describeSecurityProfile(client), "iot", "SecurityProfile")
    
	out(describeStream(client), "iot", "Stream")
    
	out(describeThing(client), "iot", "Thing")
    
	out(describeThingGroup(client), "iot", "ThingGroup")
    
	out(describeThingRegistrationTask(client), "iot", "ThingRegistrationTask")
    
	out(describeThingType(client), "iot", "ThingType")
    
}

func describeAccountAuditConfiguration(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeAccountAuditConfigurationInput{}
	result, err := client.DescribeAccountAuditConfiguration(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAuditFinding(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeAuditFindingInput{}
	result, err := client.DescribeAuditFinding(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAuditMitigationActionsTask(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeAuditMitigationActionsTaskInput{}
	result, err := client.DescribeAuditMitigationActionsTask(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAuditSuppression(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeAuditSuppressionInput{}
	result, err := client.DescribeAuditSuppression(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAuditTask(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeAuditTaskInput{}
	result, err := client.DescribeAuditTask(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAuthorizer(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeAuthorizerInput{}
	result, err := client.DescribeAuthorizer(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBillingGroup(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeBillingGroupInput{}
	result, err := client.DescribeBillingGroup(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeCACertificate(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeCACertificateInput{}
	result, err := client.DescribeCACertificate(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeCertificate(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeCertificateInput{}
	result, err := client.DescribeCertificate(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeCustomMetric(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeCustomMetricInput{}
	result, err := client.DescribeCustomMetric(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDefaultAuthorizer(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeDefaultAuthorizerInput{}
	result, err := client.DescribeDefaultAuthorizer(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDetectMitigationActionsTask(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeDetectMitigationActionsTaskInput{}
	result, err := client.DescribeDetectMitigationActionsTask(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDimension(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeDimensionInput{}
	result, err := client.DescribeDimension(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDomainConfiguration(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeDomainConfigurationInput{}
	result, err := client.DescribeDomainConfiguration(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeEndpoint(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeEndpointInput{}
	result, err := client.DescribeEndpoint(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeEventConfigurations(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeEventConfigurationsInput{}
	result, err := client.DescribeEventConfigurations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeFleetMetric(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeFleetMetricInput{}
	result, err := client.DescribeFleetMetric(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeIndex(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeIndexInput{}
	result, err := client.DescribeIndex(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeJob(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeJobInput{}
	result, err := client.DescribeJob(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeJobExecution(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeJobExecutionInput{}
	result, err := client.DescribeJobExecution(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeJobTemplate(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeJobTemplateInput{}
	result, err := client.DescribeJobTemplate(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeManagedJobTemplate(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeManagedJobTemplateInput{}
	result, err := client.DescribeManagedJobTemplate(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMitigationAction(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeMitigationActionInput{}
	result, err := client.DescribeMitigationAction(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProvisioningTemplate(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeProvisioningTemplateInput{}
	result, err := client.DescribeProvisioningTemplate(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProvisioningTemplateVersion(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeProvisioningTemplateVersionInput{}
	result, err := client.DescribeProvisioningTemplateVersion(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeRoleAlias(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeRoleAliasInput{}
	result, err := client.DescribeRoleAlias(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeScheduledAudit(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeScheduledAuditInput{}
	result, err := client.DescribeScheduledAudit(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSecurityProfile(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeSecurityProfileInput{}
	result, err := client.DescribeSecurityProfile(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeStream(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeStreamInput{}
	result, err := client.DescribeStream(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeThing(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeThingInput{}
	result, err := client.DescribeThing(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeThingGroup(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeThingGroupInput{}
	result, err := client.DescribeThingGroup(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeThingRegistrationTask(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeThingRegistrationTaskInput{}
	result, err := client.DescribeThingRegistrationTask(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeThingType(client *iot.Client) map[string]interface{} {
	input := &iot.DescribeThingTypeInput{}
	result, err := client.DescribeThingType(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

