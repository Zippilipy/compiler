package main

import (
	"compiler/parser"
	"compiler/tokenizer"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := os.ReadFile("input.xd")
	check(err)
	tokenHandler := tokenizer.Tokenizer{}
	list := tokenHandler.Tokenize(string(input))
	parsing := parser.Parser{}
	output := "global _start\n_start:\n"
	output += parsing.Parse(list)
	err = os.WriteFile("output.asm", []byte(output), 0644)
	check(err)
}
