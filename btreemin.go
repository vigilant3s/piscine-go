package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	// An empty tree has no minimum
	if root == nil {
		return nil
	}

	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}
