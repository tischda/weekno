// +build ignore

//go:generate echo Setting up your files...
//go:generate go get golang.org/x/mod/modfile
//go:generate go run template.go
//go:generate rm template.go
//go:generate go mod tidy

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	modfile "golang.org/x/mod/modfile"
)

func main() {

	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	moduleName := modfile.ModulePath(goModBytes)
	projectName := filepath.Base(moduleName)

	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	readmeTemplate.Execute(f, struct {
		Project    string
		Repository string
	}{
		Project:    projectName,
		Repository: moduleName,
	})
}

var readmeTemplate = template.Must(template.New("").Parse(`# {{ .Project }} [![Test](https://{{ .Repository }}/actions/workflows/test.yml/badge.svg)](https://{{ .Repository }}/actions/workflows/test.yml)

Utility written in [Go](https://www.golang.org).

### Install

~~~
go install {{ .Repository }}@latest
~~~

### Usage

~~~

~~~

Example:

~~~

~~~
`))
