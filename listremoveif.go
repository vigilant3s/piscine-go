package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	// Step 1: Remove from the front
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If list is now empty, reset tail
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Step 2: Remove from the rest
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	// Step 3: Update tail
	l.Tail = l.Head
	for l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}
}
