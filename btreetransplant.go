package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	// If node has no parent, it's the root
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// Replace node as left child of its parent
		node.Parent.Left = rplc
	} else {
		// Replace node as right child of its parent
		node.Parent.Right = rplc
	}

	// Update parent pointer of replacement, if not nil
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
