package ssm

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := ssm.NewFromConfig(awscfg)
    
	out(describeActivations(client), "ssm", "Activations")
    
	out(describeAssociation(client), "ssm", "Association")
    
	out(describeAssociationExecutionTargets(client), "ssm", "AssociationExecutionTargets")
    
	out(describeAssociationExecutions(client), "ssm", "AssociationExecutions")
    
	out(describeAutomationExecutions(client), "ssm", "AutomationExecutions")
    
	out(describeAutomationStepExecutions(client), "ssm", "AutomationStepExecutions")
    
	out(describeAvailablePatches(client), "ssm", "AvailablePatches")
    
	out(describeDocument(client), "ssm", "Document")
    
	out(describeDocumentPermission(client), "ssm", "DocumentPermission")
    
	out(describeEffectiveInstanceAssociations(client), "ssm", "EffectiveInstanceAssociations")
    
	out(describeEffectivePatchesForPatchBaseline(client), "ssm", "EffectivePatchesForPatchBaseline")
    
	out(describeInstanceAssociationsStatus(client), "ssm", "InstanceAssociationsStatus")
    
	out(describeInstanceInformation(client), "ssm", "InstanceInformation")
    
	out(describeInstancePatchStates(client), "ssm", "InstancePatchStates")
    
	out(describeInstancePatchStatesForPatchGroup(client), "ssm", "InstancePatchStatesForPatchGroup")
    
	out(describeInstancePatches(client), "ssm", "InstancePatches")
    
	out(describeInventoryDeletions(client), "ssm", "InventoryDeletions")
    
	out(describeMaintenanceWindowExecutionTaskInvocations(client), "ssm", "MaintenanceWindowExecutionTaskInvocations")
    
	out(describeMaintenanceWindowExecutionTasks(client), "ssm", "MaintenanceWindowExecutionTasks")
    
	out(describeMaintenanceWindowExecutions(client), "ssm", "MaintenanceWindowExecutions")
    
	out(describeMaintenanceWindowSchedule(client), "ssm", "MaintenanceWindowSchedule")
    
	out(describeMaintenanceWindowTargets(client), "ssm", "MaintenanceWindowTargets")
    
	out(describeMaintenanceWindowTasks(client), "ssm", "MaintenanceWindowTasks")
    
	out(describeMaintenanceWindows(client), "ssm", "MaintenanceWindows")
    
	out(describeMaintenanceWindowsForTarget(client), "ssm", "MaintenanceWindowsForTarget")
    
	out(describeOpsItems(client), "ssm", "OpsItems")
    
	out(describeParameters(client), "ssm", "Parameters")
    
	out(describePatchBaselines(client), "ssm", "PatchBaselines")
    
	out(describePatchGroupState(client), "ssm", "PatchGroupState")
    
	out(describePatchGroups(client), "ssm", "PatchGroups")
    
	out(describePatchProperties(client), "ssm", "PatchProperties")
    
	out(describeSessions(client), "ssm", "Sessions")
    
}

func describeActivations(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeActivationsInput{}
	result, err := client.DescribeActivations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAssociation(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeAssociationInput{}
	result, err := client.DescribeAssociation(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAssociationExecutionTargets(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeAssociationExecutionTargetsInput{}
	result, err := client.DescribeAssociationExecutionTargets(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAssociationExecutions(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeAssociationExecutionsInput{}
	result, err := client.DescribeAssociationExecutions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAutomationExecutions(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeAutomationExecutionsInput{}
	result, err := client.DescribeAutomationExecutions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAutomationStepExecutions(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeAutomationStepExecutionsInput{}
	result, err := client.DescribeAutomationStepExecutions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeAvailablePatches(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeAvailablePatchesInput{}
	result, err := client.DescribeAvailablePatches(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDocument(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeDocumentInput{}
	result, err := client.DescribeDocument(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeDocumentPermission(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeDocumentPermissionInput{}
	result, err := client.DescribeDocumentPermission(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeEffectiveInstanceAssociations(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeEffectiveInstanceAssociationsInput{}
	result, err := client.DescribeEffectiveInstanceAssociations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeEffectivePatchesForPatchBaseline(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeEffectivePatchesForPatchBaselineInput{}
	result, err := client.DescribeEffectivePatchesForPatchBaseline(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInstanceAssociationsStatus(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeInstanceAssociationsStatusInput{}
	result, err := client.DescribeInstanceAssociationsStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInstanceInformation(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeInstanceInformationInput{}
	result, err := client.DescribeInstanceInformation(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInstancePatchStates(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeInstancePatchStatesInput{}
	result, err := client.DescribeInstancePatchStates(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInstancePatchStatesForPatchGroup(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeInstancePatchStatesForPatchGroupInput{}
	result, err := client.DescribeInstancePatchStatesForPatchGroup(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInstancePatches(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeInstancePatchesInput{}
	result, err := client.DescribeInstancePatches(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeInventoryDeletions(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeInventoryDeletionsInput{}
	result, err := client.DescribeInventoryDeletions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindowExecutionTaskInvocations(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput{}
	result, err := client.DescribeMaintenanceWindowExecutionTaskInvocations(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindowExecutionTasks(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowExecutionTasksInput{}
	result, err := client.DescribeMaintenanceWindowExecutionTasks(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindowExecutions(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowExecutionsInput{}
	result, err := client.DescribeMaintenanceWindowExecutions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindowSchedule(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowScheduleInput{}
	result, err := client.DescribeMaintenanceWindowSchedule(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindowTargets(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowTargetsInput{}
	result, err := client.DescribeMaintenanceWindowTargets(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindowTasks(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowTasksInput{}
	result, err := client.DescribeMaintenanceWindowTasks(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindows(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowsInput{}
	result, err := client.DescribeMaintenanceWindows(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeMaintenanceWindowsForTarget(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeMaintenanceWindowsForTargetInput{}
	result, err := client.DescribeMaintenanceWindowsForTarget(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeOpsItems(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeOpsItemsInput{}
	result, err := client.DescribeOpsItems(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeParameters(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeParametersInput{}
	result, err := client.DescribeParameters(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePatchBaselines(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribePatchBaselinesInput{}
	result, err := client.DescribePatchBaselines(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePatchGroupState(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribePatchGroupStateInput{}
	result, err := client.DescribePatchGroupState(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePatchGroups(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribePatchGroupsInput{}
	result, err := client.DescribePatchGroups(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePatchProperties(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribePatchPropertiesInput{}
	result, err := client.DescribePatchProperties(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSessions(client *ssm.Client) map[string]interface{} {
	input := &ssm.DescribeSessionsInput{}
	result, err := client.DescribeSessions(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

