package medialive

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := medialive.NewFromConfig(awscfg)
    
	out(describeChannel(client), "medialive", "Channel")
    
	out(describeInput(client), "medialive", "Input")
    
	out(describeInputDevice(client), "medialive", "InputDevice")
    
	out(describeInputDeviceThumbnail(client), "medialive", "InputDeviceThumbnail")
    
	out(describeInputSecurityGroup(client), "medialive", "InputSecurityGroup")
    
	out(describeMultiplex(client), "medialive", "Multiplex")
    
	out(describeMultiplexProgram(client), "medialive", "MultiplexProgram")
    
	out(describeOffering(client), "medialive", "Offering")
    
	out(describeReservation(client), "medialive", "Reservation")
    
	out(describeSchedule(client), "medialive", "Schedule")
    
}

func describeChannel(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeChannelInput{}
	result, err := client.DescribeChannel(context.TODO(), input)
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

func describeInput(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeInputInput{}
	result, err := client.DescribeInput(context.TODO(), input)
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

func describeInputDevice(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeInputDeviceInput{}
	result, err := client.DescribeInputDevice(context.TODO(), input)
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

func describeInputDeviceThumbnail(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeInputDeviceThumbnailInput{}
	result, err := client.DescribeInputDeviceThumbnail(context.TODO(), input)
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

func describeInputSecurityGroup(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeInputSecurityGroupInput{}
	result, err := client.DescribeInputSecurityGroup(context.TODO(), input)
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

func describeMultiplex(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeMultiplexInput{}
	result, err := client.DescribeMultiplex(context.TODO(), input)
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

func describeMultiplexProgram(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeMultiplexProgramInput{}
	result, err := client.DescribeMultiplexProgram(context.TODO(), input)
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

func describeOffering(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeOfferingInput{}
	result, err := client.DescribeOffering(context.TODO(), input)
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

func describeReservation(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeReservationInput{}
	result, err := client.DescribeReservation(context.TODO(), input)
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

func describeSchedule(client *medialive.Client) map[string]interface{} {
	input := &medialive.DescribeScheduleInput{}
	result, err := client.DescribeSchedule(context.TODO(), input)
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

