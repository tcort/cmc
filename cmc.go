package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tcort/cmc/parser"
	"github.com/tcort/cmc/scanner"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("usage: cms [filename]")
	}

	filename := args[0]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	tokens, err := scanner.Scan(filename, string(content))
	if err != nil {
		log.Fatal(err)
	}

	program, err := parser.Parse(filename, tokens)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(program)
}
