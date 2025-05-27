package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data} // Step 1: Make the new node

	newNode.Next = l.Head // Step 2: Point it to the current head
	l.Head = newNode      // Step 3: Set it as the new head

	if l.Tail == nil {
		// Step 4: If the list was empty, this node is also the tail
		l.Tail = newNode
	}
}
