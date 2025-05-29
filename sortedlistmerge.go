package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// If one of the lists is empty, return the other
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	var head *NodeI
	var tail *NodeI

	// Initialize head to the smaller first node
	if n1.Data < n2.Data {
		head = n1
		n1 = n1.Next
	} else {
		head = n2
		n2 = n2.Next
	}
	tail = head

	// Merge the lists
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			tail.Next = n1
			n1 = n1.Next
		} else {
			tail.Next = n2
			n2 = n2.Next
		}
		tail = tail.Next
	}

	// Attach the remaining part
	if n1 != nil {
		tail.Next = n1
	} else {
		tail.Next = n2
	}

	return head
}
