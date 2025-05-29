package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// Create the new node
	newNode := &TreeNode{Data: data}

	// If root is nil, return newNode as the new root
	if root == nil {
		return newNode
	}

	// Start from the root
	current := root

	for {
		// If the new data is less than current node's data, go left
		if data < current.Data {
			// If there's no left child, insert here
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				break
			}
			current = current.Left
		} else {
			// If the new data is greater or equal, go right
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				break
			}
			current = current.Right
		}
	}

	// Return the original root unchanged
	return root
}
