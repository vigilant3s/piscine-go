package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

// Helper function with min/max boundaries
func isBST(node *TreeNode, min string, max string) bool {
	if node == nil {
		return true
	}

	// If min is set and node.Data <= min → invalid
	if min != "" && node.Data <= min {
		return false
	}
	// If max is set and node.Data > max → invalid
	if max != "" && node.Data > max {
		return false
	}

	// Check subtrees with updated bounds
	return isBST(node.Left, min, node.Data) &&
		isBST(node.Right, node.Data, max)
}
