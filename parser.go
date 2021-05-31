package main

import (
	"github.com/phuslu/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetCommandLists(path string) []Command {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open commands file")
	}
	var cmds []Command
	err = yaml.Unmarshal(file, &cmds)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse commands file: " + path)
	}
	return cmds
}
