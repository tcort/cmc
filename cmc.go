package main

import (
	"github.com/tcort/cmc/scanner"
	"io/ioutil"
	"log"
	"os"
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

	for _, token := range tokens {
		log.Printf("%s:%s@%d", token.Type, token.Text, token.LineNumber)
	}
}
