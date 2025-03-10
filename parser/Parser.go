package parser

import (
	"compiler/tokenizer"
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"strconv"
)

type Parser struct {
}

func (parser Parser) Parse(list tokenizer.List) string {
	var result string
	for !list.IsEmpty() {
		token, line := list.Peek(0)
		switch token {
		case "exit":
			list.Consume()
			checkEmpty(list, "(", line)
			token, line = list.Peek(0)
			if token != "(" {
				panic(errors.New(fmt.Sprintf("expected (, got %s at line %s", token, strconv.Itoa(line))))
			}
			list.Consume()
			checkEmpty(list, "number", line)
			expression, line := list.Peek(0)
			number, err := Evaluate(expression)
			if err != nil {
				panic(errors.New(fmt.Sprintf("expected number, got %s at line %s", number, strconv.Itoa(line))))
			}
			list.Consume()
			checkEmpty(list, ")", line)
			token, line = list.Peek(0)
			if token != ")" {
				panic(errors.New(fmt.Sprintf("expected ), got %s at line %s", token, strconv.Itoa(line))))
			}
			list.Consume()
			output := fmt.Sprintf("    mov rax, 60\n    mov rdi, %s\n    syscall\n", number)
			result += output
			continue
		default:
			list.Consume()
			continue
		}
	}
	return result
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func checkEmpty(list tokenizer.List, expected string, line int) {
	if list.IsEmpty() {
		panic(errors.New(fmt.Sprintf("expected %s at line %s", expected, strconv.Itoa(line))))
	}
}

func Evaluate(expression string) (string, error) {
	fileSet := token.NewFileSet()
	tr, err := types.Eval(fileSet, nil, token.NoPos, expression)
	if err != nil {
		return "", err
	}
	return tr.Value.String(), nil
}
