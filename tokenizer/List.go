package tokenizer

type Node struct {
	value string
	next  *Node
	line  int
}

type List struct {
	head *Node
	size int
}

func (List *List) Add(value string, line int) {
	List.size++
	newNode := &Node{value: value, line: line}
	if List.head == nil {
		List.head = newNode
		return
	}

	current := List.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (List *List) Peek(amount int) (string, int) {
	current := List.head
	if amount >= List.size {
		panic("Peek amount is larger than size")
	}
	for i := 0; i < amount; i++ {
		current = current.next
	}
	return current.value, current.line
}

func (List *List) Consume() {
	List.head = List.head.next
	List.size--
}

func (List *List) IsEmpty() bool {
	return List.size == 0
}
