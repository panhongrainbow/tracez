package bpTree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Test_Check_inode_ackUpgradeIndexNode is to test the functionality of ackUpgradeIndexNode.
// When the second index node under the inode's node needs an upgrade,
// the content under the second index node is upgraded and then overwritten in the location of the inode's second index node.
// (不删除，用覆盖的方式)
func Test_Check_inode_ackUpgradeIndexNode(t *testing.T) {
	// Set up the total length and splitting length for B Plus Tree.
	BpWidth = 3
	BpHalfWidth = 2

	// Set up a top-level index node.
	inode := &BpIndex{
		Index: []int64{40},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{10, 30},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{4},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 1}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 4}, {Key: 9}},
							},
						},
					},
					{
						Index: []int64{19},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 10}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 19}, {Key: 29}},
							},
						},
					},
					{
						Index: []int64{35, 38},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 39}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 35}, {Key: 37}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 38}},
							},
						},
					},
				},
			},
			{
				Index: []int64{67, 77, 89},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{49, 59},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 40}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 49}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 59}, {Key: 65}},
							},
						},
					},
					{
						Index: []int64{73},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 67}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 73}},
							},
						},
					},
					{
						Index: []int64{81},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 77}, {Key: 78}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 81}, {Key: 86}},
							},
						},
					},
					{
						Index: []int64{96, 98},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 89}, {Key: 95}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 96}, {Key: 97}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 98}, {Key: 99}},
							},
						},
					},
				},
			},
		},
		DataNodes: []*BpData{},
	}

	// sideToOverwrite is prepared to overwrite the content at the ix position in indoe.
	sideToOverwrite := &BpIndex{
		Index: []int64{77},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{67},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{49, 59},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 40}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 49}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 59}, {Key: 65}},
							},
						},
					},
					{
						Index: []int64{73},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 67}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 73}},
							},
						},
					},
				},
			},
			{
				Index: []int64{89},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{81},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 77}, {Key: 78}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 81}, {Key: 86}},
							},
						},
					},
					{
						Index: []int64{96, 98},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 89}, {Key: 95}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 96}, {Key: 97}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 98}, {Key: 99}},
							},
						},
					},
				},
			},
		},
		DataNodes: []*BpData{},
	}

	// Expect inode after being overwritten.
	expectedMiddleAfterProtruding := &BpIndex{
		Index: []int64{40, 77},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{10, 30},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{4},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 1}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 4}, {Key: 9}},
							},
						},
					},
					{
						Index: []int64{19},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 10}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 19}, {Key: 29}},
							},
						},
					},
					{
						Index: []int64{35, 38},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 39}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 35}, {Key: 37}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 38}},
							},
						},
					},
				},
			},
			{
				Index: []int64{67},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{49, 59},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 40}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 49}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 59}, {Key: 65}},
							},
						},
					},
					{
						Index: []int64{73},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 67}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 73}},
							},
						},
					},
				},
			},
			{
				Index: []int64{89},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{81},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 77}, {Key: 78}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 81}, {Key: 86}},
							},
						},
					},
					{
						Index: []int64{96, 98},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 89}, {Key: 95}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 96}, {Key: 97}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 98}, {Key: 99}},
							},
						},
					},
				},
			},
		},
		DataNodes: []*BpData{},
	}

	// Call the function to be tested.
	inode.ackUpgradeIndexNode(1, sideToOverwrite)

	// Check the inode.
	assert.True(t, reflect.DeepEqual(expectedMiddleAfterProtruding, inode), "inode mismatch")
}

// Test_Check_inode_protrudeInOddBpWidth tests the protruding of the top-level index node in a B Plus tree,
// including the splitting of the BpIndex slice.
func Test_Check_inode_protrudeInOddBpWidth(t *testing.T) {
	// Set up the total length and splitting length for B Plus Tree.
	BpWidth = 3
	BpHalfWidth = 2

	// Set up a top-level index node.
	inode := &BpIndex{
		Index: []int64{30, 40, 89},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{10},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 4}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 10}},
					},
				},
			},
			{
				Index: []int64{38},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 30}, {Key: 35}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 38}},
					},
				},
			},
			{
				Index: []int64{81},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 40}, {Key: 67}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 81}},
					},
				},
			},
			{
				Index: []int64{96},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 89}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 96}, {Key: 98}},
					},
				},
			},
		},
		DataNodes: []*BpData{},
	}

	// Expect a new node named middle after protruding.
	expectedMiddleAfterProtruding := &BpIndex{
		Index: []int64{40},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{30},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{10},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 4}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 10}},
							},
						},
					},
					{
						Index: []int64{38},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 30}, {Key: 35}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 38}},
							},
						},
					},
				},
			},
			{
				Index: []int64{89},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{81},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 40}, {Key: 67}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 81}},
							},
						},
					},
					{
						Index: []int64{96},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 89}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 96}, {Key: 98}},
							},
						},
					},
				},
			},
		},

		DataNodes: nil,
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Call the function to be tested.
	middle, err := inode.protrudeInOddBpWidth()

	// Check for errors.
	assert.NoError(t, err, "Unexpected error")

	// Check the node named middle.
	assert.True(t, reflect.DeepEqual(expectedMiddleAfterProtruding, middle), "middle mismatch")
}

// Test_Check_inode_protrudeInEvenBpWidth tests the protruding of the top-level index node in a B Plus tree,
// including the splitting of the BpIndex slice.
func Test_Check_inode_protrudeInEvenBpWidth(t *testing.T) {
	// Set up the total length and splitting length for B Plus Tree.
	BpWidth = 4
	BpHalfWidth = 2

	// Set up a top-level index node.
	inode := &BpIndex{
		Index: []int64{30, 40, 72, 81},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{5, 10, 19},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 1}, {Key: 4}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 5}, {Key: 9}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 10}, {Key: 18}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 19}, {Key: 29}},
					},
				},
			},
			{
				Index: []int64{37},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 30}, {Key: 35}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 37}, {Key: 38}},
					},
				},
			},
			{
				Index: []int64{59, 67},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 40}, {Key: 46}, {Key: 49}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 59}, {Key: 65}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 67}, {Key: 69}},
					},
				},
			},
			{
				Index: []int64{77},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 72}, {Key: 73}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 77}, {Key: 78}},
					},
				},
			},
			{
				Index: []int64{89, 96, 98},
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 89}, {Key: 95}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 96}, {Key: 97}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 98}, {Key: 99}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 81}, {Key: 86}},
					},
				},
			},
		},
		DataNodes: []*BpData{},
	}

	// Expect a new node named middle after protruding.
	expectedMiddleAfterProtruding := &BpIndex{
		Index: []int64{72},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{30, 40},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{5, 10, 19},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 1}, {Key: 4}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 5}, {Key: 9}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 10}, {Key: 18}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 19}, {Key: 29}},
							},
						},
					},
					{
						Index: []int64{37},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 30}, {Key: 35}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 37}, {Key: 38}},
							},
						},
					},
					{
						Index: []int64{59, 67},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 40}, {Key: 46}, {Key: 49}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 59}, {Key: 65}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 67}, {Key: 69}},
							},
						},
					},
				},
			},
			{
				Index: []int64{81},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{77},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 72}, {Key: 73}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 77}, {Key: 78}},
							},
						},
					},
					{
						Index: []int64{89, 96, 98},
						DataNodes: []*BpData{
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 89}, {Key: 95}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 96}, {Key: 97}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 98}, {Key: 99}},
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 81}, {Key: 86}},
							},
						},
					},
				},
			},
		},

		DataNodes: nil,
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Call the function to be tested.
	middle, err := inode.protrudeInEvenBpWidth()

	// Check for errors.
	assert.NoError(t, err, "Unexpected error")

	// Check the node named middle.
	assert.True(t, reflect.DeepEqual(expectedMiddleAfterProtruding.IndexNodes[1], middle.IndexNodes[1]), "middle mismatch")
}

// Test_Check_inode_splitWithDnode tests the splitting of the bottom-level index node in a B Plus tree,
// including the splitting of the BpData slice.
func Test_Check_inode_splitWithDnode(t *testing.T) {
	// Set up the total length and splitting length for B Plus Tree.
	BpWidth = 3
	BpHalfWidth = 2

	// Set up a bottom-level index node.
	inode := &BpIndex{
		Index: []int64{10, 20, 30},
		DataNodes: []*BpData{
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 5}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 15}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 25}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 35}},
			},
		},
	}

	// The three parts after splitting.

	// Key after splitting.
	expectedKeyAfterSplitting := int64(20)

	// Old node iNode after splitting.
	expectedInodeAfterSplitting := &BpIndex{
		Index: []int64{10},
		DataNodes: []*BpData{
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 5}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 15}},
			},
		},
	}

	// Expect a new node named side after splitting.
	expectedSideAfterSplitting := &BpIndex{
		Index: []int64{30},
		DataNodes: []*BpData{
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 25}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 35}},
			},
		},
	}

	// Call the function to be tested.
	key, side, err := inode.splitWithDnode()

	// Check for errors.
	assert.NoError(t, err, "Unexpected error")

	// Check the returned key.
	assert.Equal(t, expectedKeyAfterSplitting, key, "Key mismatch")

	// Check the origin iNode.
	assert.True(t, reflect.DeepEqual(expectedInodeAfterSplitting, inode), "Inode mismatch")

	// Check the returned side.
	assert.True(t, reflect.DeepEqual(expectedSideAfterSplitting, side), "Side mismatch")
}

// Test_Check_BpIndex_Operation verifies the merging process after calling function splitWithDnode.
// This will overwrite the receiver pointer variable, which is the original inode.
func Test_Check_inode_mergeWithDnode(t *testing.T) {
	// Set up the total length and splitting length for B Plus Tree.
	BpWidth = 3
	BpHalfWidth = 2

	// Set up a bottom-level index node after splitting.

	// Key after splitting.
	key := int64(20)

	// Old node iNode after splitting.
	inode := &BpIndex{
		Index:      []int64{10},
		IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
		DataNodes: []*BpData{
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 5}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 15}},
			},
		},
	}

	// New node side after splitting.
	side := &BpIndex{
		Index:      []int64{30},
		IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
		DataNodes: []*BpData{
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 25}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 35}},
			},
		},
	}

	// Expect the old node named iNode after merging.
	expectedMergedInode := &BpIndex{
		Index: []int64{20},
		IndexNodes: []*BpIndex{
			{
				Index:      []int64{10},
				IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 5}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 15}},
					},
				},
			},
			{
				Index:      []int64{30},
				IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 25}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 35}},
					},
				},
			},
		},
		DataNodes: nil,
	}

	// Call the function to be tested.
	err := inode.mergeWithDnode(key, side)

	// Check for errors.
	assert.NoError(t, err, "Unexpected error")

	// Check the origin iNode
	assert.True(t, reflect.DeepEqual(expectedMergedInode, inode), "Inode mismatch")
}

// Test_Check_inode_mergeUpgradedKeyNode primarily tests
// the upgrade and subsequent merging of intermediate index values (named keys) and independent index nodes
// when the number of slices in the Index Node exceeds a certain threshold.
// (主要测试索引 key 和 width)
func Test_Check_inode_mergeUpgradedKeyNode(t *testing.T) {
	// Set up the total length and splitting length for B Plus Tree.
	BpWidth = 3
	BpHalfWidth = 2

	// Set up a bottom-level index node after splitting.
	inode := &BpIndex{
		Index: []int64{40},
		IndexNodes: []*BpIndex{
			{
				Index:      []int64{10},
				IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 4}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 10}},
					},
				},
			},
			{
				Index:      []int64{81},
				IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 81}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 98}},
					},
				},
			},
		},
		DataNodes: nil,
	}

	// Merge the upgraded index (named toBeUpgradedKey) and the index node (named toBeUpgradedInode) later.
	toBeUpgradedKey := int64(30)

	toBeUpgradeInode := &BpIndex{
		Index:      []int64{38},
		IndexNodes: []*BpIndex{},
		DataNodes: []*BpData{
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 30}, {Key: 35}},
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 38}},
			},
		},
	}

	// Expect inode after merging.
	expectMergedInode := &BpIndex{
		Index: []int64{30, 40},
		IndexNodes: []*BpIndex{
			{
				Index:      []int64{10},
				IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 4}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 10}},
					},
				},
			},
			{
				Index:      []int64{38},
				IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 30}, {Key: 35}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 38}},
					},
				},
			},
			{
				Index:      []int64{81},
				IndexNodes: []*BpIndex{}, // []*BpIndex{} or nil does not match, DeepEqual will fail.
				DataNodes: []*BpData{
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 81}},
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 98}},
					},
				},
			},
		},
		DataNodes: nil,
	}

	// Call the function to be tested.
	// insertAfterPosition is at pos0(0), insert the upgraded node after pos0(0), which is at pos1(1).
	err := inode.mergeUpgradedKeyNode(0, toBeUpgradedKey, toBeUpgradeInode)

	// Check for errors.
	assert.NoError(t, err, "Unexpected error")

	// Check the original iNode after merging.
	assert.True(t, reflect.DeepEqual(expectMergedInode.IndexNodes, inode.IndexNodes), "Inode mismatch")
}

func Test_BpIndex_InsertBpIdxNewValue(t *testing.T) {
	t.Run("InsertNewValue", func(t *testing.T) {
		BpWidth = 3
		BpHalfWidth = 2

		// Create a BpIndex instance
		index := &BpIndex{
			Index: []int64{5},
		}

		// Set up some test data nodes
		index.DataNodes = []*BpData{
			{Items: []BpItem{{Key: 3}}},
			{Items: []BpItem{{Key: 5}, {Key: 6}}},
		}

		// Create a new BpItem
		newItem := BpItem{Key: 1}

		// Call the insertItem function
		_, _, _, _, _ = index.insertItem(nil, newItem)

		// Check the updated index
		expectedIndex := []int64{5}
		assert.Equal(t, expectedIndex, index.Index)

		fmt.Println(index.DataNodes[0].Items)
	})

	t.Run("InsertExistingValue", func(t *testing.T) {
		// Create a BpIndex instance
		index := &BpIndex{
			Index: []int64{1, 2, 4},
		}

		// Call the insertExistBpIdxNewValue function
		// index.insertExistBpIdxNewValue(&BpItem{Key: 3}, 2)

		// Check the updated index
		expectedIndex := []int64{1, 2, 3, 4}
		assert.Equal(t, expectedIndex, index.Index)
	})
}

func Test_BpIndex_InsertBpIdxNewIndex(t *testing.T) {
	tests := []struct {
		inputIndex []int64
		newIndex   int64
		expected   []int64
	}{
		{
			inputIndex: []int64{1, 3, 5, 7},
			newIndex:   4,
			expected:   []int64{1, 3, 4, 5, 7},
		},
		{
			inputIndex: []int64{2, 4, 6, 8},
			newIndex:   10,
			expected:   []int64{2, 4, 6, 8, 10},
		},
	}

	for _, test := range tests {
		// Create a new BpIndex
		idx := &BpIndex{
			Index: test.inputIndex,
		}

		// Call the function
		idx.insertBpIX(test.newIndex)

		// Check if the result matches the expected output
		assert.Equal(t, test.expected, idx.Index, "For inputIndex %v and newIndex %v", test.inputIndex, test.newIndex)
	}
}

// Test_Check_BpIndex_Split tests index splitting based on width and updating the index nodes accordingly.
func Test_Check_BpIndex_Split(t *testing.T) {
	// Create a BpIndex instance
	index := &BpIndex{
		Index: []int64{1, 2, 3, 4, 5},
	}

	// Define the expected result after the split
	expectedSplit := []int64{4, 5}
	expectedRemain := []int64{1, 2, 3}

	// Call the split function with width 2
	//err := index.split(2)

	// Check for any errors
	/*if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}*/

	// Check if the index has been split correctly
	if !reflect.DeepEqual(index.Index, expectedRemain) {
		t.Errorf("Unexpected index after split. Got: %v, Expected: %v", index.Index, expectedRemain)
	}

	// Check the new index node
	if !reflect.DeepEqual(index.IndexNodes[0].Index, expectedSplit) {
		t.Errorf("Unexpected index in the new node. Got: %v, Expected: %v", index.IndexNodes[0].Index, expectedSplit)
	}
}

// Test_Check_BpIndex_insertBpIdxValue examines different B plus tree index key insertion scenarios.
func Test_Check_BpIndex_insertBpIdxValue(t *testing.T) {
	t.Run("InsertNewKey", func(t *testing.T) {
		// Test case: Insert a new key into an empty BpIndex.
		idx := &BpIndex{}
		key := int64(10)

		//idx.insertBpIdxValue(key)

		assert.Len(t, idx.Index, 1, "Expected one key in the Index slice")
		assert.Equal(t, key, idx.Index[0], "Expected the inserted key to be 10")
	})

	t.Run("InsertExistingKey", func(t *testing.T) {
		// Test case: Insert an existing key into the BpIndex with other keys.
		idx := &BpIndex{
			Index: []int64{5, 15, 25},
		}
		key := int64(15)

		//idx.insertBpIdxValue(key)

		assert.Len(t, idx.Index, 4, "Expected one more key in the Index slice")
		assert.Equal(t, key, idx.Index[2], "Expected the inserted key to be at the correct position")
	})

	t.Run("InsertExistingKeyAtBeginning", func(t *testing.T) {
		// Test case: Insert an existing key at the beginning of the BpIndex.
		idx := &BpIndex{
			Index: []int64{5, 15, 25},
		}
		key := int64(5)

		//idx.insertBpIdxValue(key)

		assert.Len(t, idx.Index, 4, "Expected one more key in the Index slice")
		assert.Equal(t, key, idx.Index[0], "Expected the inserted key to be at the beginning")
	})

	t.Run("InsertExistingKeyAtEnd", func(t *testing.T) {
		// Test case: Insert an existing key at the end of the BpIndex.
		idx := &BpIndex{
			Index: []int64{5, 15, 25},
		}
		key := int64(25)

		//idx.insertBpIdxValue(key)

		assert.Len(t, idx.Index, 4, "Expected one more key in the Index slice")
		assert.Equal(t, key, idx.Index[3], "Expected the inserted key to be at the end")
	})

	t.Run("InsertExistingKeyInMiddle", func(t *testing.T) {
		// Test case: Insert an existing key into the middle of the BpIndex.
		idx := &BpIndex{
			Index: []int64{10, 20, 40},
		}
		key := int64(30)

		//idx.insertBpIdxValue(key)

		assert.Len(t, idx.Index, 4, "Expected one more key in the Index slice")
		assert.Equal(t, key, idx.Index[2], "Expected the inserted key to be in the middle")
	})
}
