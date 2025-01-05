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
	replaced = strings.ReplaceAll(replaced, "(", " ( ")
	replaced = strings.ReplaceAll(replaced, ")", " )")
	split := strings.Split(replaced, " ")
	var buffer bytes.Buffer
	for _, str := range split {
		buffer.WriteString(str + "\n")
	}
	list := LinkedList{}
	scanner := bufio.NewScanner(&buffer)
	for scanner.Scan() {
		token := scanner.Text()
		list.Add(&token)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for !list.IsEmpty() {
		token := list.Peek(0)
		fmt.Println(token)
		list.Consume()
	}
}
