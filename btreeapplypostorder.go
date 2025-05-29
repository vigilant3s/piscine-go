package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Step 1: Traverse the left subtree
	BTreeApplyPostorder(root.Left, f)

	// Step 2: Traverse the right subtree
	BTreeApplyPostorder(root.Right, f)

	// Step 3: Visit the current node
	f(root.Data)
}
