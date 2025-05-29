package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Step 1: Go to the left subtree
	BTreeApplyInorder(root.Left, f)

	// Step 2: Apply the function to current node's data
	f(root.Data)

	// Step 3: Go to the right subtree
	BTreeApplyInorder(root.Right, f)
}
