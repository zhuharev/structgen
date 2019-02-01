package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli"
	"github.com/zhuharev/structgen"
	yaml "gopkg.in/yaml.v2"
)

const (
	FlagConfig   = "config"
	FlagTemplate = "template"
	FlagOut      = "out"
)

func main() {
	app := &cli.App{
		Action: run,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  FlagConfig + ", c",
				Value: "structs.yml",
			},
			cli.StringFlag{
				Name:  FlagTemplate + ", t",
				Value: "template.tmpl",
			},
			cli.StringFlag{
				Name:  FlagOut + ", o",
				Usage: "destination folder",
				Value: "models/",
			},
		},
	}
	app.Run(os.Args)

}

func run(ctx *cli.Context) (err error) {
	t, err := template.ParseFiles(ctx.String(FlagTemplate))
	if err != nil {
		return
	}

	var s structgen.Schema

	data, err := ioutil.ReadFile(ctx.String(FlagConfig))
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &s)
	if err != nil {
		return
	}

	s.Init()

	log.Printf("%+v", s)

	b := bytes.NewBuffer(nil)

	err = t.Execute(b, map[string]interface{}{
		"structs": s.Structs,
	})
	if err != nil {
		return
	}

	err = os.MkdirAll(ctx.String(FlagOut), os.ModePerm)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filepath.Join(ctx.String(FlagOut), "models_gen.go"), b.Bytes(), os.ModePerm)
	return
}
