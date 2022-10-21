package ec2

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func DescribeInstances(awscfg aws.Config) {
	client := ec2.NewFromConfig(awscfg)
	input := &ec2.DescribeInstancesInput{}

	result, err := client.DescribeInstances(context.TODO(), input)

	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return
	}

	/*
		for _, r := range result.Reservations {
			fmt.Println("Reservation ID: " + *r.ReservationId)
			fmt.Println("Instance IDs:")
			for _, i := range r.Instances {
				fmt.Println("   " + *i.InstanceId)
			}

			fmt.Println("")
		}
	*/

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %v\n", string(data))
}
