package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Case 1 & 2: Node has no left child
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		// Case 2: Node has no right child
		root = BTreeTransplant(root, node, node.Left)
	} else {
		// Case 3: Node has two children
		// Find the in-order successor (min of right subtree)
		successor := BTreeMin(node.Right)

		// If successor is not the direct right child of the node
		if successor.Parent != node {
			// Replace successor with its own right child
			root = BTreeTransplant(root, successor, successor.Right)
			// Move node's right subtree to successor
			successor.Right = node.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}

		// Replace node with successor
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
	}

	return root
}
