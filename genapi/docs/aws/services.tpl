package services

import (
	{{range .Packages}}
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/{{.}}"
    {{end}}
    "github.com/aws/aws-sdk-go-v2/aws"
)

type Service interface {
	Dump(aws.Config, func(map[string]interface{}, string, string))
}

func GetServices() []Service {
    return []Service{
        {{range .Packages}}
        &{{.}}.Service{},
        {{end}}
    }
}
