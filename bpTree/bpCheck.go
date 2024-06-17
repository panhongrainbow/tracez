package bpTree

// CheckBpNodeAndNextContinuity is to check the continuity of the current node and the next node.
func CheckBpNodeAndNextContinuity(inode *BpIndex) bool {
	if inode == nil { // If the root is nil, the tree is empty.
		return true // An empty tree is considered continuous.
	}

	// Start with the data head node.
	currentNode := inode.BpDataHead()

	// Traverse the linked list of nodes.
	for currentNode != nil {
		currentKeys := currentNode.Keys()
		nextNode := currentNode.NextNode()

		// Check the continuity within the current node.
		if !CheckNodeBpNodeContinuity(currentKeys) {
			return false
		}

		// Check if the last key of the current node and the first key of the next node are continuous.
		if nextNode != nil { // If there is a next node.
			nextKeys := nextNode.Keys()
			if len(currentKeys) > 0 && len(nextKeys) > 0 {
				if currentKeys[len(currentKeys)-1] > nextKeys[0] {
					return false
				}
			}
		}

		// Move to the next node in the linked list.
		currentNode = nextNode
	}

	// All nodes and keys are continuous, return true.
	return true
}

// CheckNodeBpNodeContinuity is to check the continuity within the node.
func CheckNodeBpNodeContinuity(keys []int64) bool {
	// Loop through the keys to check if they are in ascending order.
	for i := 0; i < len(keys)-1; i++ {
		if keys[i] > keys[i+1] {
			return false
		}
	}

	// All keys are in order, return true.
	return true
}
