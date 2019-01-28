package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/zhuharev/structgen"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	t, err := template.ParseFiles("template.tmpl")
	if err != nil {
		panic(err)
	}

	var s structgen.Schema

	data, err := ioutil.ReadFile("example.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}

	s.Init()

	log.Printf("%+v", s)

	b := bytes.NewBuffer(nil)

	err = t.Execute(b, map[string]interface{}{
		"structs": s.Structs,
	})
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("out.go", b.Bytes(), os.ModePerm)

}
