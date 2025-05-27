package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		// Case: l1 is empty, just take all of l2
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	if l2.Head == nil {
		// Case: l2 is empty, nothing to do
		return
	}

	// Case: both lists have content, link them
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
