package tokenizer

type Node struct {
	value *string
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (LinkedList *LinkedList) Add(value *string) {
	newNode := &Node{value: value}
	if LinkedList.head == nil {
		LinkedList.head = newNode
		return
	}

	current := LinkedList.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	LinkedList.size += 1
}

func (LinkedList *LinkedList) Peek(amount int) string {
	current := LinkedList.head
	if amount > LinkedList.size {
		panic("Peek amount is larger than size")
	}
	for i := 0; i < amount; i++ {
		current = current.next
	}
	return *current.value
}

func (LinkedList *LinkedList) Consume() {
	LinkedList.head = LinkedList.head.next
}

func (LinkedList *LinkedList) IsEmpty() bool {
	return LinkedList.size == 0
}
