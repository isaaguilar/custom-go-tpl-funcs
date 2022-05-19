package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/jmespath/go-jmespath"
	"gopkg.in/yaml.v2"
)

type Values struct {
	Adjective string
}

func main() {
	values := Values{"fine"}

	tplfileb, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	tplfile := string(tplfileb)
	tpl := template.Must(template.New("").Funcs(Custom).Parse(tplfile))
	var buf bytes.Buffer
	err = tpl.Execute(&buf, values)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(buf.String())

}

var Custom = map[string]interface{}{
	"read": Read,
	"yaml": func() string { return "yaml" },
}

func Read(kind, filename, query string) string {
	if kind == "yaml" {
		return Yaml(filename, query)
	}
	return ""
}

type Trait struct {
	Name     string
	Favorite Favorite
}

type Favorite struct {
	Color        string
	Number       int64
	Food         string
	BooleanValue bool `yaml:"booleanValue"`
}

func Yaml(filename, query string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
		return ""
	}
	traits := Trait{}
	err = yaml.Unmarshal(b, &traits)
	if err != nil {
		log.Panic(err)
		return ""
	}

	result, err := jmespath.Search(query, traits)
	if err != nil {
		return ""
	}

	switch r := result.(type) {
	case string:
		return r
	case int64:
		return fmt.Sprintf("%d", r)
	case bool:
		return fmt.Sprintf("%t", r)
	default:
		return ""
	}

}
