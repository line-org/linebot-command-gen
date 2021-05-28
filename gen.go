package main

import (
	"fmt"
	"github.com/phuslu/log"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

// ExecCommand Copyright (c)  2020 go-generalize
//Released under the MIT license
//https://opensource.org/licenses/mit-license.php
// ExecCommand - コマンドを実行して出力結果とエラーを返す
func ExecCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	if exitCode := cmd.ProcessState.ExitCode(); exitCode != 0 {
		return "", fmt.Errorf("failed to exec git command: (exit code: %d, output: %s)", exitCode, string(b))
	}

	return string(b), nil
}

func GenerateCommandHandler(cmds []Command) {
	tmpFile, err := ioutil.ReadFile("template/handler.go.tmplate")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load tempalte file")
	}
	os.MkdirAll("gen", os.ModePerm)
	path := "gen/commands.go"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open target file")
	}
	defer file.Close()
	err = ExecuteTemplate(string(tmpFile), file, map[string]interface{}{"Cmds": cmds}, nil, path)
	if err != nil {
		log.Fatal().Err(err).Msg("failed execute template")
	}
}

func ExecuteTemplate(txt string, file io.Writer, data interface{}, funcMap template.FuncMap, path string) error {
	if funcMap == nil {
		funcMap = make(map[string]interface{})
	}

	newTmp := template.New("mongo-db-gen")
	tmp := template.Must(newTmp.Funcs(funcMap).Parse(txt))
	err := tmp.Execute(file, data)
	if err != nil {
		log.Fatal().Err(err).Msg("failed execute template")
	}
	_, err = ExecCommand("goimports", "-w", path)
	return nil
}
