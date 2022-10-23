package qldb

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := qldb.NewFromConfig(awscfg)
    
	out(describeJournalKinesisStream(client), "qldb", "JournalKinesisStream")
    
	out(describeJournalS3Export(client), "qldb", "JournalS3Export")
    
	out(describeLedger(client), "qldb", "Ledger")
    
}

func describeJournalKinesisStream(client *qldb.Client) map[string]interface{} {
	input := &qldb.DescribeJournalKinesisStreamInput{}
	result, err := client.DescribeJournalKinesisStream(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeJournalS3Export(client *qldb.Client) map[string]interface{} {
	input := &qldb.DescribeJournalS3ExportInput{}
	result, err := client.DescribeJournalS3Export(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func describeLedger(client *qldb.Client) map[string]interface{} {
	input := &qldb.DescribeLedgerInput{}
	result, err := client.DescribeLedger(context.TODO(), input)
	if err != nil {
        log.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

