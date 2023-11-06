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

	// 先準備 1 個資料切片，1 個資料切片不會產生索引
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

	return
}
