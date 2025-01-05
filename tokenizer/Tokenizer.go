package tokenizer

import (
	"fmt"
	"strings"
)

type Tokenizer struct {
}

func (tokenizer Tokenizer) Tokenize(input string) List {
	replaced := strings.ReplaceAll(input, "\r\n", "\n")
	replaced = strings.ReplaceAll(replaced, "\n", "\n ")
	replaced = strings.ReplaceAll(replaced, "(", "( ")
	replaced = strings.ReplaceAll(replaced, ")", " )")
	var keywords []string
	keywords = append(keywords, "exit")
	for _, value := range keywords {
		replaced = strings.ReplaceAll(replaced, value, fmt.Sprintf("%s ", value))
	}
	split := strings.Split(replaced, " ")
	list := List{}
	index := 1
	for _, str := range split {
		list.Add(strings.ReplaceAll(str, "\n", ""), index)
		if strings.Contains(str, "\n") {
			index++
		}
	}
	return list
}
