package budgets

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := budgets.NewFromConfig(awscfg)
    
	out(describeBudget(client), "budgets", "Budget")
    
	out(describeBudgetAction(client), "budgets", "BudgetAction")
    
	out(describeBudgetActionHistories(client), "budgets", "BudgetActionHistories")
    
	out(describeBudgetActionsForAccount(client), "budgets", "BudgetActionsForAccount")
    
	out(describeBudgetActionsForBudget(client), "budgets", "BudgetActionsForBudget")
    
	out(describeBudgetNotificationsForAccount(client), "budgets", "BudgetNotificationsForAccount")
    
	out(describeBudgetPerformanceHistory(client), "budgets", "BudgetPerformanceHistory")
    
	out(describeBudgets(client), "budgets", "Budgets")
    
	out(describeNotificationsForBudget(client), "budgets", "NotificationsForBudget")
    
	out(describeSubscribersForNotification(client), "budgets", "SubscribersForNotification")
    
}

func describeBudget(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetInput{}
	result, err := client.DescribeBudget(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBudgetAction(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetActionInput{}
	result, err := client.DescribeBudgetAction(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBudgetActionHistories(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetActionHistoriesInput{}
	result, err := client.DescribeBudgetActionHistories(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBudgetActionsForAccount(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetActionsForAccountInput{}
	result, err := client.DescribeBudgetActionsForAccount(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBudgetActionsForBudget(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetActionsForBudgetInput{}
	result, err := client.DescribeBudgetActionsForBudget(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBudgetNotificationsForAccount(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetNotificationsForAccountInput{}
	result, err := client.DescribeBudgetNotificationsForAccount(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBudgetPerformanceHistory(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetPerformanceHistoryInput{}
	result, err := client.DescribeBudgetPerformanceHistory(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBudgets(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeBudgetsInput{}
	result, err := client.DescribeBudgets(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeNotificationsForBudget(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeNotificationsForBudgetInput{}
	result, err := client.DescribeNotificationsForBudget(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSubscribersForNotification(client *budgets.Client) map[string]interface{} {
	input := &budgets.DescribeSubscribersForNotificationInput{}
	result, err := client.DescribeSubscribersForNotification(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

