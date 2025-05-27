package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	current := l.Head
	l.Tail = l.Head // After reversal, old head becomes new tail

	for current != nil {
		next := current.Next // Save the next car
		current.Next = prev  // Flip the arrow backwards
		prev = current       // Move prev forward
		current = next       // Move current forward
	}

	l.Head = prev // Set new head to the last non-nil node
}
