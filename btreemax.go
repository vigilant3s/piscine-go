package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	// An empty tree has no max
	if root == nil {
		return nil
	}

	// Keep going right until there's no more
	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
