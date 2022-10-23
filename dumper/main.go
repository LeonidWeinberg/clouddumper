package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/leonidweinberg/clouddumper/dumper/aws/services"
	_ "github.com/lib/pq"
)

func main() {
	awscfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	_, err = sts.NewFromConfig(awscfg).GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}

	out := func(awsdata map[string]interface{}, service, method string) {
		data, err := json.MarshalIndent(awsdata, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		// filename := "./dump/" + service + "_" + method + ".json"
		err = os.WriteFile("./dump/"+service+"_"+method+".json", data, 0777)
		if err != nil {
			log.Fatal(err)
		}

		// log.Println(filename)
	}

	err = os.RemoveAll("./dump")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir("./dump", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	ser := services.GetServices()

	for _, service := range ser {
		service.Dump(awscfg, out)
	}
}

// postgres := embeddedpostgres.NewDatabase()
// err := postgres.Start()

//if err != nil {
//	log.Fatal(err)
//}

// postgres.Stop()
