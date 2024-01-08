package bpTree

import (
	"sync"
)

// The width and half-width for B plus tree.
var (
	BpWidth     int // the width of B plus tree.
	BpHalfWidth int // the half-width of B plus tree.
)

// BpTree is the root of Tree B plus.
type BpTree struct {
	mutex sync.Mutex // lock
	root  *BpIndex   // root tree
}

// NewBpTree initializes B plus tree structure with specified width and data entries.
func NewBpTree(width int) (tree *BpTree) {
	// Set the width and half-width for B plus tree.
	if width < 3 { // The minimum width for B plus tree is 3.
		width = 3
	}
	BpWidth = width
	BpHalfWidth = int((float32(BpWidth)-0.1)/2) + 1

	// Create root tree instance
	tree = &BpTree{
		root: &BpIndex{
			DataNodes: make([]*BpData, 0, BpWidth+1), // The addition of 1 is because data chunks may temporarily exceed the width.
		},
	}

	// Prepare one data slice first; one data slice will not generate an index.
	tree.root.DataNodes = append(tree.root.DataNodes, &BpData{})

	return
}

// InsertValue ensures thread safety, insert item in B plus tree index, release lock.
func (tree *BpTree) InsertValue(item BpItem) {
	// Acquire a lock to ensure thread safety.
	tree.mutex.Lock()

	// Insert the item into the B plus tree index.
	_, popKey, popNode, status, err := tree.root.insertItem(nil, item)

	if err != nil {
		panic(err)
	}

	if status == statusProtrudeInode && popNode != nil {
		// Here, it will increase the entire tree's depth. (层数增加)
		tree.root = popNode
		status = statusNormal
	}

	if status == statusProtrudeDnode {
		err = tree.root.mergeWithDnode(popKey, popNode)
		status = statusNormal
		if err != nil {
			return
		}
	}

	if len(tree.root.Index) >= BpWidth && len(tree.root.Index)%2 != 0 {
		popNode, _ = tree.root.protrudeInOddBpWidth()
		tree.root = popNode
	} else if len(tree.root.Index) >= BpWidth && len(tree.root.Index)%2 == 0 {
		popNode, _ = tree.root.protrudeInEvenBpWidth()
		tree.root = popNode
	}

	// Release the lock to allow other threads to access the tree.
	tree.mutex.Unlock()

	// Performing a return.
	return
}

// RemoveValue ensures thread safety, remove item in B plus tree index, release lock.
func (tree *BpTree) RemoveValue(item BpItem) (deleted, updated bool, ix int, err error) {
	// Acquire a lock to ensure thread safety.
	tree.mutex.Lock()

	// Release the lock to allow other threads to access the tree.
	defer tree.mutex.Unlock()

	// The deletion operation is currently managed by the root node to prevent issues with mismatched levels of child nodes.
	// If the levels of child nodes are not correct, the B plus tree may malfunction. ‼️
	// 删除操作由根节点管理，确保所有子节点层级相同 ‼️

	// Performing deletion operation.
	deleted, updated, ix, err = tree.root.delFromRoot(item)

	// ⚠️ The following is the B plus tree merging operation.
	// The merging criteria here do not rely on an empty node index. ‼️
	// This is done to increase the chances of merging, as the index may not be cleared on time.
	// 这里不以节点 index 为空为合拼标准，因为 index 可能没有及时清空

	// ⚠️ When there is only one remaining index child node. (索引节点的升级合拼)
	if len(tree.root.IndexNodes) == 1 && len(tree.root.DataNodes) == 0 {
		*tree.root = *tree.root.IndexNodes[0]
		return
	}

	// ⚠️ When there is only two remaining data child nodes. (资料节点的升级合拼)
	if len(tree.root.DataNodes) == 2 && len(tree.root.IndexNodes) == 0 {
		// If one of the data nodes indeed has no data.
		if len(tree.root.DataNodes[0].Items) == 0 && len(tree.root.DataNodes[1].Items) != 0 {
			// If the first data node is empty, replace the root node with the second data node.
			tree.root.Index = nil
			tree.root.DataNodes = []*BpData{tree.root.DataNodes[1]}
			return
		} else if len(tree.root.DataNodes[1].Items) == 0 && len(tree.root.DataNodes[0].Items) != 0 {
			// If the second data node is empty, replace the root node with the first data node.
			tree.root.Index = nil
			tree.root.DataNodes = []*BpData{tree.root.DataNodes[0]}
			return
		}
	}

	// ⚠️ If there are only 2 index nodes, but the data is not a lot, they can be merged.
	if len(tree.root.IndexNodes) == 2 &&
		BpWidth > (len(tree.root.IndexNodes[0].Index)+len(tree.root.IndexNodes[1].Index)) { // Combine within the range of BpWidth.

		// Begin the merger process.
		if len(tree.root.Index) == 0 && len(tree.root.IndexNodes) > 0 {
			// Create a new BpIndex node.
			node := &BpIndex{}

			// Merge all index nodes into a new node.
			for i := 0; i < len(tree.root.IndexNodes); i++ {
				node.Index = append(node.Index, tree.root.IndexNodes[i].Index...)
				node.IndexNodes = append(node.IndexNodes, tree.root.IndexNodes[i].IndexNodes...)
				node.DataNodes = append(node.DataNodes, tree.root.IndexNodes[i].DataNodes...)
			}

			// Replace the original root node with the new node.
			*tree.root = *node
		}
	}

	// ⚠️ Warning: The following code appears to perform a restructuring operation on a B+ tree.
	if len(tree.root.DataNodes) == 2 &&
		BpWidth > (len(tree.root.DataNodes[0].Items)+len(tree.root.DataNodes[1].Items)) {
		// Create a new BpIndex node.
		node := &BpIndex{}

		// Copy the DataNodes from the root's IndexNodes to the new node.
		for i := 0; i < len(tree.root.IndexNodes); i++ {
			node.DataNodes = append(node.DataNodes, tree.root.IndexNodes[i].DataNodes...)
		}

		// Extract the keys from the DataNodes and populate the Index of the new node.
		for i := 0; i < len(node.DataNodes); i++ {
			if i != 0 {
				node.Index = append(node.Index, node.DataNodes[i].Items[0].Key)
			}
		}

		// Replace the original root node with the new node.
		*tree.root = *node
	}

	// Performing a return.
	return
}
