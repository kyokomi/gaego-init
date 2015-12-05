package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	entryDirPath = "templates/"
)

//go:generate go-bindata templates/...

type Builder interface {
	Run() error
}

type builder struct {
	AppName string
}

func NewBuilder(appName string) Builder {
	return &builder{
		AppName: appName,
	}
}

func (b *builder) Run() error {
	if err := os.Mkdir(b.AppName, 0755); err != nil {
		log.Println(err)
	}

	names := AssetNames()
	for _, name := range names {
		execName := strings.Replace(name, entryDirPath, "", 1)
		execName = filepath.Join(b.AppName, execName)

		fmt.Println(name)
		fmt.Println(execName)
		fmt.Println(filepath.Dir(execName))

		// ディレクトリ作成&ファイル作成
		if err := os.MkdirAll(filepath.Dir(execName), 0755); err != nil {
			log.Println(err)
		}

		f, err := os.Create(execName)
		if err != nil {
			log.Println(err)
		}

		data, err := Asset(name)
		if err != nil {
			log.Println(err)
		}

		t := template.Must(template.New(name).Parse(string(data)))
		if err := t.Execute(f, b); err != nil {
			log.Println(err)
		}
	}

	return nil
}
