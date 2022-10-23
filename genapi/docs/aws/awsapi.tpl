package {{.Package}}

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/{{.Package}}"
)

type Service struct {
}

func (c *Service) Dump(awscfg aws.Config, out func(map[string]interface{}, string, string)) {
	client := {{.Package}}.NewFromConfig(awscfg)
    {{range .Describers}}
	out(describe{{.Method}}(client), "{{.Service}}", "{{.Method}}")
    {{end}}
}
{{range .Describers}}
func describe{{.Method}}(client *{{.Service}}.Client) map[string]interface{} {
	input := &{{.Service}}.Describe{{.Method}}Input{}
	result, err := client.Describe{{.Method}}(context.TODO(), input)
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
{{end}}
