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
	GenerateCommandHandler(cmds)
}
