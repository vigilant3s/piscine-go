package piscine

type NodeL struct {
	Data interface{} // Holds any type of data
	Next *NodeL      // Points to the next node
}

type List struct {
	Head *NodeL // Start of the list
	Tail *NodeL // End of the list
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data} // Step 1: Create the new car

	if l.Head == nil {
		// Step 2: If list is empty, set Head and Tail to new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Step 3: If list has stuff, link the new car to the end
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
