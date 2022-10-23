package lexmodelsv2

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := lexmodelsv2.NewFromConfig(awscfg)
    
	out(describeBot(client), "lexmodelsv2", "Bot")
    
	out(describeBotAlias(client), "lexmodelsv2", "BotAlias")
    
	out(describeBotLocale(client), "lexmodelsv2", "BotLocale")
    
	out(describeBotRecommendation(client), "lexmodelsv2", "BotRecommendation")
    
	out(describeBotVersion(client), "lexmodelsv2", "BotVersion")
    
	out(describeCustomVocabularyMetadata(client), "lexmodelsv2", "CustomVocabularyMetadata")
    
	out(describeExport(client), "lexmodelsv2", "Export")
    
	out(describeImport(client), "lexmodelsv2", "Import")
    
	out(describeIntent(client), "lexmodelsv2", "Intent")
    
	out(describeResourcePolicy(client), "lexmodelsv2", "ResourcePolicy")
    
	out(describeSlot(client), "lexmodelsv2", "Slot")
    
	out(describeSlotType(client), "lexmodelsv2", "SlotType")
    
}

func describeBot(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeBotInput{}
	result, err := client.DescribeBot(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBotAlias(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeBotAliasInput{}
	result, err := client.DescribeBotAlias(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBotLocale(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeBotLocaleInput{}
	result, err := client.DescribeBotLocale(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBotRecommendation(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeBotRecommendationInput{}
	result, err := client.DescribeBotRecommendation(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeBotVersion(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeBotVersionInput{}
	result, err := client.DescribeBotVersion(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeCustomVocabularyMetadata(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeCustomVocabularyMetadataInput{}
	result, err := client.DescribeCustomVocabularyMetadata(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeExport(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeExportInput{}
	result, err := client.DescribeExport(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeImport(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeImportInput{}
	result, err := client.DescribeImport(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeIntent(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeIntentInput{}
	result, err := client.DescribeIntent(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeResourcePolicy(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeResourcePolicyInput{}
	result, err := client.DescribeResourcePolicy(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSlot(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeSlotInput{}
	result, err := client.DescribeSlot(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeSlotType(client *lexmodelsv2.Client) map[string]interface{} {
	input := &lexmodelsv2.DescribeSlotTypeInput{}
	result, err := client.DescribeSlotType(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

