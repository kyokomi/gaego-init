package main

import (
	"log"
	"os"
	"text/template"
)

//go:generate go-bindata templates/*

func main() {
	f, _ := os.Create("app.yaml")
	data, err := Asset("templates/src/app.yaml")
	if err != nil {
		log.Println(err)
	}

	t := template.Must(template.New("app.yaml").Parse(string(data)))
	if err := t.Execute(f, map[string]string{"appName": "hoge-app"}); err != nil {
		log.Println(err)
	}
}
