package main

import (
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
	output := []byte("testing\ntesting2")
	err = os.WriteFile("output.asm", output, 0644)
	check(err)
	tokenHandler := tokenizer.Tokenizer{}
	tokenHandler.Tokenize(string(input))
}
