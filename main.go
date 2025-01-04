package main

import (
	"fmt"
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
	fmt.Println(string(input))
	output := []byte("testing\ntesting2")
	err = os.WriteFile("output.asm", output, 0644)
	check(err)
}
