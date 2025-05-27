package piscine

func ListClear(l *List) {
	// Disconnect the list from its nodes
	l.Head = nil
	l.Tail = nil
}
