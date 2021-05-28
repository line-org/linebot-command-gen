package main

import (
	"github.com/phuslu/log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal().Msg("Target directory is required")
	}
	cmds := GetCommandLists(os.Args[len(os.Args)-1])
	Validate(cmds)
	GenerateCommandHandler(cmds)
}

func Validate(cmds []Command) {
	var ids []string
	for _, cmd := range cmds {
		if isContain(ids, cmd.ID) {
			log.Fatal().Msg("cant use same id: " + cmd.ID)
		}
		ids = append(ids, cmd.ID)
	}
}

func isContain(base []string, target string) bool {
	for _, val := range base {
		if target == val {
			return true
		}
	}
	return false
}
