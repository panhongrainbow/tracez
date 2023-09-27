package bpTree

import (
	"sync"
)

// The width and half-width for B plus tree.
var (
	BpWidth     int // the width of B plus tree.
	BpHalfWidth int // the half width of B plus tree.
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
	BpHalfWidth = int((float32(BpWidth) + 0.1) / 2)

	// Create root tree instance
	tree = &BpTree{
		root: &BpIndex{
			DataNodes: make([]*BpData, 0, BpWidth+1), // The addition of 1 is because data chunks may temporarily exceed the width.
		},
	}

	// 先準備 1 個資料切片，1 個資料切片不會產生索引
	tree.root.DataNodes = append(tree.root.DataNodes, &BpData{})

	return
}

// InsertValue ensures thread safety, insert item in B plus tree index, release lock.
func (tree *BpTree) InsertValue(item BpItem) {
	// Acquire a lock to ensure thread safety.
	tree.mutex.Lock()

	// Insert the item into the B plus tree index.
	popKey, popNode, err := tree.root.insertBpIdxNewValue(nil, item)

	if err != nil {
		panic(err)
	}

	if popNode != nil && popKey == 0 {
		tree.root.cmpAndMergeIndexNode(popNode, tree.root)
	}

	if popKey != 0 {
		tree.root.cmpAndOrganizeIndexNode(popKey, popNode, tree.root)
	}

	// Release the lock to allow other threads to access the tree.
	tree.mutex.Unlock()

	return
}