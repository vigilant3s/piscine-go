package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Step 1: Apply function to current node's Data
	f(root.Data)

	// Step 2: Traverse the left subtree
	BTreeApplyPreorder(root.Left, f)

	// Step 3: Traverse the right subtree
	BTreeApplyPreorder(root.Right, f)
}
