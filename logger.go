package main

import "github.com/phuslu/log"

func init() {
	log.DefaultLogger = log.Logger{
		TimeFormat: " ",
		Caller:     1,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}
}
