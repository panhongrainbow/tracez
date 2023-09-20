package bpTree

import (
	"sync"
)

// The width and half-width for B plus tree.
var (
	BpWidth     int // the width of B plus tree.
	BpHalfWidth int // the half width of B plus tree.
)

// BpTree is the root of Tree B.
type BpTree struct {
	mutex sync.Mutex // lock
	root  *BpIndex2  // root tree
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
		root: &BpIndex2{
			DataNodes: make([]*BpData, BpWidth),
			IsLeaf:    true, // Initially, the root node will be filled with data.
		},
	}

	// Prepare a certain number of links to the bpData nodes initially.
	// root
	//             <= Links are used to record the maximum values of bpData.
	//  []  []  [] <= Represents individual data entries.
	for i := 0; i < BpWidth; i++ {
		tree.root.DataNodes[i] = &BpData{
			Items: make([]BpItem, 0, width),
		}
	}

	return
}

// InsertValue ensures thread safety, insert item in B plus tree index, release lock.

// root
//	  .  .  . <= The first-level index.
//	---  -  - <= The data is concentrated on the left side.

// root
// .               .    <= The first-level index.
// .	        .  .  . <= The second-level index.
// - [] [] ---  -  - <= Represents individual data entries, and the data is concentrated on the left side.
func (tree *BpTree) InsertValue(item BpItem) {
	// Acquire a lock to ensure thread safety.
	tree.mutex.Lock()

	// Insert the item into the B plus tree index.
	tree.root.insertIndexValue(item)

	// Release the lock to allow other threads to access the tree.
	tree.mutex.Unlock()

	return
}
