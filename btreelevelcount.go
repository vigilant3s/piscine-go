package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Get level count for left and right subtrees
	leftCount := BTreeLevelCount(root.Left)
	rightCount := BTreeLevelCount(root.Right)

	// Return the deeper one + 1 (for the current node)
	if leftCount > rightCount {
		return leftCount + 1
	}
	return rightCount + 1
}
