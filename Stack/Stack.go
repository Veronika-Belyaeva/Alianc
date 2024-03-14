package stack

type Node struct {
	data string
	next *Node
}

type Stack struct {
	head *Node
}

func (stack *Stack) push(value string) {

	node := &Node{data: value}
	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
}

func (stack *Stack) pop() (string, string) {
	if stack.head == nil {
		return "", "nil"
	} else {
		value := stack.head.data
		stack.head = stack.head.next
		return value, "nil"
	}
}