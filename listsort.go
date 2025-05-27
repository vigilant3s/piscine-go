package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l // List is empty or has one node — already sorted
	}

	swapped := true

	// Keep looping until no swaps are needed
	for swapped {
		swapped = false
		current := l

		for current.Next != nil {
			if current.Data > current.Next.Data {
				// Swap the data values
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true
			}
			current = current.Next
		}
	}

	return l
}
