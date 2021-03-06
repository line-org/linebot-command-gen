package main

import (
	"flag"
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
	var outputDir = flag.String("o", "gen/", "output path")
	flag.Parse()

	path := os.Args[len(os.Args)-1]
	fInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal().Err(err).Msg("failed check path")
	}
	var cmds []*Command
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

	for _, c := range cmds {
		c.BaseText = strings.ReplaceAll(c.BaseText, " ", "")
		var slice []string
		for _, s := range c.SubTexts {
			slice = append(slice, strings.ReplaceAll(s, " ", ""))
		}
		c.SubTexts = slice
	}
	Validate(cmds)
	GenerateCommandHandler(cmds, *outputDir)
}

func Validate(cmds []*Command) {
	var ids []string
	var texts []string
	for _, cmd := range cmds {
		if isContain(ids, cmd.ID) {
			log.Fatal().Msg("cant use same id: " + cmd.ID)
		}
		ids = append(ids, cmd.ID)
		if isContain(texts, cmd.BaseText) {
			log.Fatal().Msg("cant use same base text: " + cmd.BaseText)
		}
		texts = append(texts, cmd.BaseText)
		for _, s := range cmd.SubTexts {
			if isContain(texts, s) {
				log.Fatal().Msg("cant use same sub text: " + s)
			}
			texts = append(texts, s)
		}
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
