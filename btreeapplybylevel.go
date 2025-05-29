package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Create a queue of nodes to visit
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Take the first node in the queue
		current := queue[0]
		queue = queue[1:]

		// Apply the function to current node's data
		f(current.Data)

		// Add left and right children to the queue if they exist
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
