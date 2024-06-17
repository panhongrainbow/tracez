package bpTree

// BpDataHead traverses the B Plus tree from the given index node to find and return the head (first) data node.
func (inode *BpIndex) BpDataHead() (head *BpData) {
	current := inode
	for {
		if len(current.DataNodes) == 0 {
			current = current.IndexNodes[0] // Move to the first index node at the next level.
		} else {
			return current.DataNodes[0] // Return the first data node.
		}
	}
}

// BpDataTail traverses the B Plus tree from the given index node to find and return the tail (last) data node
func (inode *BpIndex) BpDataTail() (head *BpData) {
	current := inode
	for {
		if len(current.DataNodes) == 0 {
			length := len(current.IndexNodes)
			current = current.IndexNodes[length-1] // Move to the last index node at the next level.
		} else {
			length := len(current.DataNodes)
			return current.DataNodes[length-1] // Return the last data node.
		}
	}
}

// NextNode returns the next data node in the linked list.
func (data *BpData) NextNode() (node *BpData) {
	return data.Next // Return the pointer to the next data node.
}

// Keys returns a slice of keys stored in the current data node.
func (data *BpData) Keys() (keys []int64) {
	// Iterate over the items and collect their keys.
	for _, item := range data.Items {
		keys = append(keys, item.Key)
	}

	// Return the slice of keys
	return keys
}
