package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	// Base case: empty tree or match found
	if root == nil || root.Data == elem {
		return root
	}

	// If elem is less, go left
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}

	// Else, go right
	return BTreeSearchItem(root.Right, elem)
}
