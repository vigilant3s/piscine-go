package piscine

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return nil // List is empty
	}
	return l.Tail.Data // Return the data in the last node
}
