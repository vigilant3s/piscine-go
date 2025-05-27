package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current // Found the correct position
		}
		current = current.Next // Move to the next node
		index++
	}

	return nil // If pos is too big or list is empty
}
