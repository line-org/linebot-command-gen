package main

import (
	"github.com/phuslu/log"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal().Msg("Target directory is required")
	}
	path := os.Args[len(os.Args)-1]
	fInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal().Err(err).Msg("failed check path")
	}
	var cmds []Command
	if fInfo.IsDir() {
		filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if strings.HasSuffix(path, "yaml") {
				cmds = append(cmds, GetCommandLists(path)...)
			}
			return nil
		})
	} else {
		cmds = GetCommandLists(path)
	}
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
