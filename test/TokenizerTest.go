package main

import (
	"compiler/tokenizer"
	"errors"
	"fmt"
)

func main() {
	TestLinkedList()
}

func TestLinkedList() {
	t := tokenizer.List{}

	// Test adding to the List
	value1 := "test1"
	value2 := "test2"
	t.Add(&value1)
	t.Add(&value2)

	if t.Peek(0) != "test1" {
		panic(errors.New(fmt.Sprintf("Expected 'test1', got '%s'", t.Peek(0))))
	}

	if t.Peek(1) != "test2" {
		panic(errors.New(fmt.Sprintf("Expected 'test2', got '%s'", t.Peek(1))))
	}

	// Test consuming elements
	t.Consume()

	if t.Peek(0) != "test2" {
		panic(errors.New(fmt.Sprintf("Expected 'test2' after Consume, got '%s'", t.Peek(0))))
	}

	if t.IsEmpty() {
		panic(errors.New(fmt.Sprintf("List should not be empty")))
	}

	t.Consume()

	if !t.IsEmpty() {
		panic(errors.New(fmt.Sprintf("List should be empty after consuming all elements")))
	}
}
