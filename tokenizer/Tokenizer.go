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
	listThatCanBeParsed := List{}
	for !list.IsEmpty() {
		token, index := list.Peek(0)
		if token == "(" {
			listThatCanBeParsed.Add(token, index)
			list.Consume()
			result := ""
			left := 1
			right := 0
			for left-right != 0 {
				token, index = list.Peek(0)
				if token == "(" {
					left++
					result += token
				} else if token == ")" {
					right++
					if left-right != 0 {
						result += token
					} else {
						listThatCanBeParsed.Add(result, index)
						continue
					}
				} else {
					result += token
				}
				list.Consume()
			}
		}
		listThatCanBeParsed.Add(token, index)
		list.Consume()
	}
	return listThatCanBeParsed
}
