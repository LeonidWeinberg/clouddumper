package servicecatalog

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := servicecatalog.NewFromConfig(awscfg)
    
	out(describeConstraint(client), "servicecatalog", "Constraint")
    
	out(describeCopyProductStatus(client), "servicecatalog", "CopyProductStatus")
    
	out(describePortfolio(client), "servicecatalog", "Portfolio")
    
	out(describePortfolioShareStatus(client), "servicecatalog", "PortfolioShareStatus")
    
	out(describePortfolioShares(client), "servicecatalog", "PortfolioShares")
    
	out(describeProduct(client), "servicecatalog", "Product")
    
	out(describeProductAsAdmin(client), "servicecatalog", "ProductAsAdmin")
    
	out(describeProductView(client), "servicecatalog", "ProductView")
    
	out(describeProvisionedProduct(client), "servicecatalog", "ProvisionedProduct")
    
	out(describeProvisionedProductPlan(client), "servicecatalog", "ProvisionedProductPlan")
    
	out(describeProvisioningArtifact(client), "servicecatalog", "ProvisioningArtifact")
    
	out(describeProvisioningParameters(client), "servicecatalog", "ProvisioningParameters")
    
	out(describeRecord(client), "servicecatalog", "Record")
    
	out(describeServiceAction(client), "servicecatalog", "ServiceAction")
    
	out(describeServiceActionExecutionParameters(client), "servicecatalog", "ServiceActionExecutionParameters")
    
	out(describeTagOption(client), "servicecatalog", "TagOption")
    
}

func describeConstraint(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeConstraintInput{}
	result, err := client.DescribeConstraint(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeCopyProductStatus(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeCopyProductStatusInput{}
	result, err := client.DescribeCopyProductStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePortfolio(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribePortfolioInput{}
	result, err := client.DescribePortfolio(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePortfolioShareStatus(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribePortfolioShareStatusInput{}
	result, err := client.DescribePortfolioShareStatus(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describePortfolioShares(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribePortfolioSharesInput{}
	result, err := client.DescribePortfolioShares(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProduct(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeProductInput{}
	result, err := client.DescribeProduct(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProductAsAdmin(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeProductAsAdminInput{}
	result, err := client.DescribeProductAsAdmin(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProductView(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeProductViewInput{}
	result, err := client.DescribeProductView(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProvisionedProduct(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeProvisionedProductInput{}
	result, err := client.DescribeProvisionedProduct(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProvisionedProductPlan(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeProvisionedProductPlanInput{}
	result, err := client.DescribeProvisionedProductPlan(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProvisioningArtifact(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeProvisioningArtifactInput{}
	result, err := client.DescribeProvisioningArtifact(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeProvisioningParameters(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeProvisioningParametersInput{}
	result, err := client.DescribeProvisioningParameters(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeRecord(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeRecordInput{}
	result, err := client.DescribeRecord(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeServiceAction(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeServiceActionInput{}
	result, err := client.DescribeServiceAction(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeServiceActionExecutionParameters(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeServiceActionExecutionParametersInput{}
	result, err := client.DescribeServiceActionExecutionParameters(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeTagOption(client *servicecatalog.Client) map[string]interface{} {
	input := &servicecatalog.DescribeTagOptionInput{}
	result, err := client.DescribeTagOption(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

