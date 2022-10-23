package savingsplans

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := savingsplans.NewFromConfig(awscfg)
    
	out(describeSavingsPlanRates(client), "savingsplans", "SavingsPlanRates")
    
	out(describeSavingsPlans(client), "savingsplans", "SavingsPlans")
    
	out(describeSavingsPlansOfferingRates(client), "savingsplans", "SavingsPlansOfferingRates")
    
	out(describeSavingsPlansOfferings(client), "savingsplans", "SavingsPlansOfferings")
    
}

func describeSavingsPlanRates(client *savingsplans.Client) map[string]interface{} {
	input := &savingsplans.DescribeSavingsPlanRatesInput{}
	result, err := client.DescribeSavingsPlanRates(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSavingsPlans(client *savingsplans.Client) map[string]interface{} {
	input := &savingsplans.DescribeSavingsPlansInput{}
	result, err := client.DescribeSavingsPlans(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSavingsPlansOfferingRates(client *savingsplans.Client) map[string]interface{} {
	input := &savingsplans.DescribeSavingsPlansOfferingRatesInput{}
	result, err := client.DescribeSavingsPlansOfferingRates(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSavingsPlansOfferings(client *savingsplans.Client) map[string]interface{} {
	input := &savingsplans.DescribeSavingsPlansOfferingsInput{}
	result, err := client.DescribeSavingsPlansOfferings(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

