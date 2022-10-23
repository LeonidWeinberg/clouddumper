package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func main() {
	docsdir := os.Args[1]
	dumperdir := os.Args[2]
	sdkdir := os.Args[3]

	CleanAndMake(path.Join(dumperdir, "./aws/services"))
	files := LoadSDKFiles(sdkdir)
	apis := parseSDKFiles(files)

	generateAwsApiData(apis, docsdir)
	generateAwsapiCode(path.Join(docsdir, "awsapi.json"), path.Join(docsdir, "awsapi.tpl"), path.Join(dumperdir, "./aws/services"))

	generateAwsServicesData(apis, docsdir)
	generateServicesCode(path.Join(docsdir, "services.json"), path.Join(docsdir, "services.tpl"), path.Join(dumperdir, "./aws/services"))
}

func CleanAndMake(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

type SDKFile struct {
	Filename string
	Content  string
}

func LoadSDKFiles(sdkdir string) []SDKFile {
	var files []SDKFile

	err := filepath.Walk(sdkdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		// dirname := strings.Split(path, "/")[len(strings.Split(path, "/"))-2]
		files = append(files, SDKFile{Filename: path, Content: string(content)})

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

type AWSApi struct {
	Service    string
	Describers []string
}

var rgx = regexp.MustCompile(`func \([\w\s\*]+\) Describe(\w+)`)

func parseSDKFiles(files []SDKFile) []AWSApi {
	mapper := map[string][]string{}

	for _, file := range files {
		res := rgx.FindAllStringSubmatch(file.Content, -1)
		if res == nil {
			continue
		}

		service := strings.Split(file.Filename, "/")[len(strings.Split(file.Filename, "/"))-2]

		for i := 1; i < len(res[0]); i++ {
			mapper[service] = append(mapper[service], res[0][i])
		}
	}

	ret := []AWSApi{}
	for key, val := range mapper {
		ret = append(ret, AWSApi{Service: key, Describers: val})
	}

	return ret
}

func generateAwsApiData(awsApis []AWSApi, docsdir string) {
	intarr := []interface{}{}

	for _, awsApi := range awsApis {
		describers := []interface{}{}
		for _, method := range awsApi.Describers {
			describer := map[string]interface{}{
				"Service": awsApi.Service,
				"Method":  method,
			}

			describers = append(describers, describer)
		}

		intarr = append(intarr, map[string]interface{}{
			"Package":    awsApi.Service,
			"Describers": describers,
		})
	}

	data, err := json.MarshalIndent(intarr, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path.Join(docsdir, "awsapi.json"), data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func generateAwsServicesData(awsApis []AWSApi, docsdir string) {
	intarr := []interface{}{}

	for _, awsApi := range awsApis {
		intarr = append(intarr, awsApi.Service)
	}

	mapper := map[string]interface{}{
		"Packages": intarr,
	}

	data, err := json.MarshalIndent(mapper, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path.Join(docsdir, "services.json"), data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func generateAwsapiCode(jsonFilename, templateFilename, gendir string) {
	tpl, err := template.New("awsapi.tpl").ParseFiles(templateFilename)
	if err != nil {
		log.Fatal(err)
	}

	input, err := os.ReadFile(jsonFilename)
	if err != nil {
		log.Fatal(err)
	}

	awsapis := []map[string]interface{}{}
	err = json.Unmarshal(input, &awsapis)
	if err != nil {
		log.Fatal(err)
	}

	for _, awsapi := range awsapis {
		pkgName := awsapi["Package"].(string)

		err := os.Mkdir(path.Join(gendir, pkgName), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(path.Join(gendir, pkgName, pkgName+".go"))
		if err != nil {
			log.Fatal(err)
		}

		defer out.Close()

		log.Println("generating: " + path.Join(gendir, pkgName+".go"))

		err = tpl.Execute(out, awsapi)
		if err != nil {
			log.Fatal(err)
		}

		out.Sync()
	}
}

func generateServicesCode(jsonFilename, templateFilename, gendir string) {
	tpl, err := template.New("services.tpl").ParseFiles(templateFilename)
	if err != nil {
		log.Fatal(err)
	}

	input, err := os.ReadFile(jsonFilename)
	if err != nil {
		log.Fatal(err)
	}

	services := map[string]interface{}{}
	err = json.Unmarshal(input, &services)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(path.Join(gendir, "aws.go"))
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	log.Println("generating: " + path.Join(gendir, "aws.go"))

	err = tpl.Execute(out, services)
	if err != nil {
		log.Fatal(err)
	}

	out.Sync()
}
