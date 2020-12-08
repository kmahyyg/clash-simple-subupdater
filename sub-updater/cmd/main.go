package main

import (
	"flag"
	"fmt"
	"errors"
)

func main() {
	confFilePtr := flag.String("conf", "updater.yaml", "Sub-Updater Config File, Default: updater.yaml")
	currentOperation := flag.String("op", "download", "Operation You wanna do. Default: download")
	inputFilePtr := flag.String("input", "ori-config.yaml", "The Original File to be processed. Default: ori-config.yaml")
	outputFilePtr := flag.String("output", "config.yaml", "Output File Path. Default: config.yaml")
	flag.Parse()
}
