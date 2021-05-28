package main

import (
	"fmt"
	"github.com/phuslu/log"
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

func GenerateCommandHandler(cmds []Command) {
	tmpFile, err := ioutil.ReadFile("template/handler.go.tmpl")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load tempalte file")
	}
	file, err := os.OpenFile("gen.go", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open target file")
	}
	defer file.Close()
	fmt.Printf("%#v\n", string(tmpFile))
	ExecuteTemplate(string(tmpFile), file, cmds, nil)
}

func ExecuteTemplate(txt string, file io.Writer, data interface{}, funcMap template.FuncMap) error {
	if funcMap == nil {
		funcMap = make(map[string]interface{})
	}

	newTmp := template.New("mongo-db-gen")
	newTmp.Delims("[[", "]]")
	tmp := template.Must(newTmp.Funcs(funcMap).Parse(txt))
	err := tmp.Execute(file, data)
	if err != nil {
		log.Fatal().Err(err).Msg("failed execute template")
	}
	return nil
}
