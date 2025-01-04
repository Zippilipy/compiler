package tokenizer

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type Tokenizer struct {
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (tokenizer Tokenizer) Tokenize(input string) {
	replaced := strings.ReplaceAll(input, "\r\n", " ")
	split := strings.Split(replaced, " ")
	var buffer bytes.Buffer
	for _, str := range split {
		buffer.WriteString(str + "\n")
	}
	scanner := bufio.NewScanner(&buffer)
	for scanner.Scan() {
		token := scanner.Text()
		fmt.Println(token)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
