package bpTree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 10}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 38}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 81}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 96}, {Key: 98}},
						Split:    false,
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
								Split:    false,
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 10}},
								Split:    false,
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
								Split:    false,
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 38}},
								Split:    false,
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
								Split:    false,
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 81}},
								Split:    false,
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
								Split:    false,
							},
							{
								Previous: nil,
								Next:     nil,
								Items:    []BpItem{{Key: 96}, {Key: 98}},
								Split:    false,
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
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 15}},
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 25}},
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 35}},
				Split:    false,
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
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 15}},
				Split:    false,
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
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 35}},
				Split:    false,
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
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 15}},
				Split:    false,
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
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 35}},
				Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 15}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 35}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 10}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 98}},
						Split:    false,
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
				Split:    false,
			},
			{
				Previous: nil,
				Next:     nil,
				Items:    []BpItem{{Key: 38}},
				Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 10}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 38}},
						Split:    false,
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
						Split:    false,
					},
					{
						Previous: nil,
						Next:     nil,
						Items:    []BpItem{{Key: 98}},
						Split:    false,
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

// Test_Check_BpIndex_Operation tests the splitting of the bottom-level index node in a B Plus tree,
// including the splitting of the BpData slice.
func Test_Check_BpIndex_Operation(t *testing.T) {
	t.Run("pop and insert dNode", func(t *testing.T) {
		// Set up Bp Parameters.
		BpWidth = 3
		BpHalfWidth = 2

		// Create a root tree with a specific structure (indexes are [7], [5], [11 to 13]).
		root := createRootTree7and5and11to13()

		// Check the updated data structure.
		assert.Equal(t, []int64{11, 13}, root.IndexNodes[1].Index)
		assert.Equal(t, []int64{10}, root.IndexNodes[1].IndexNodes[0].Index)
		assert.Equal(t, 2, root.IndexNodes[1].IndexNodes[0].DataNodes[0].dataLength())

		// >>>>> Insert a BpItem.

		// Insert a key (8) among the data items in the first data node.
		root.IndexNodes[1].IndexNodes[0].DataNodes[0].insertAmong(BpItem{Key: 8})

		// >>>>> Split a BpData.

		// Split the data node and obtain the side and any potential error.
		side, err := root.IndexNodes[1].IndexNodes[0].DataNodes[0].split()
		require.NoError(t, err, "split should not return an error")

		// >>>>> Merge a popped dNode.

		// Merge a popped data node based on the split side.
		err = root.IndexNodes[1].IndexNodes[0].mergePopDnode(side)
		require.NoError(t, err, "mergePopDnode should not return an error")

		// Check the updated data structure.
		assert.Equal(t, []int64{8, 10}, root.IndexNodes[1].IndexNodes[0].Index)
		assert.Equal(t, 3, len(root.IndexNodes[1].IndexNodes[0].DataNodes))
		assert.Equal(t, int64(7), root.IndexNodes[1].IndexNodes[0].DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(8), root.IndexNodes[1].IndexNodes[0].DataNodes[1].Items[0].Key)
		assert.Equal(t, int64(9), root.IndexNodes[1].IndexNodes[0].DataNodes[1].Items[1].Key)
		assert.Equal(t, int64(10), root.IndexNodes[1].IndexNodes[0].DataNodes[2].Items[0].Key)

		// Print the resulting tree structure.
		// root.Print()
	})
	t.Run("protrude and insert iNode", func(t *testing.T) {
		// Set up Bp Parameters.
		BpWidth = 3
		BpHalfWidth = 2

		// Create a root tree with a specific structure (indexes are [7], [5], [11 to 13]).
		root := createRootTree7and5and11to13()

		// Check the updated data structure.
		assert.Equal(t, []int64{11, 13}, root.IndexNodes[1].Index)
		assert.Equal(t, []int64{10}, root.IndexNodes[1].IndexNodes[0].Index)
		assert.Equal(t, 2, root.IndexNodes[1].IndexNodes[0].DataNodes[0].dataLength())

		// >>>>> Insert a BpItem.

		// Insert a key (2) among the data items in the first data node.
		root.IndexNodes[0].IndexNodes[0].DataNodes[0].insertAmong(BpItem{Key: 2})

		// >>>>> Split a BpData.

		// Split the data node and obtain the side and any potential error.
		dSide, err := root.IndexNodes[0].IndexNodes[0].DataNodes[0].split()
		require.NoError(t, err, "split should not return an error")

		// >>>>> Merge a popped dNode.

		// Merge a popped data node based on the split side.
		err = root.IndexNodes[0].IndexNodes[0].mergePopDnode(dSide)
		require.NoError(t, err, "mergePopDnode should not return an error")

		// Check the updated data structure.
		assert.Equal(t, []int64{1, 3, 4}, root.IndexNodes[0].IndexNodes[0].Index)
		assert.Equal(t, 4, len(root.IndexNodes[0].IndexNodes[0].DataNodes))
		assert.Equal(t, int64(1), root.IndexNodes[0].IndexNodes[0].DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(1), root.IndexNodes[0].IndexNodes[0].DataNodes[1].Items[0].Key)
		assert.Equal(t, int64(2), root.IndexNodes[0].IndexNodes[0].DataNodes[1].Items[1].Key)
		assert.Equal(t, int64(3), root.IndexNodes[0].IndexNodes[0].DataNodes[2].Items[0].Key)
		assert.Equal(t, int64(4), root.IndexNodes[0].IndexNodes[0].DataNodes[2].Items[1].Key)
		assert.Equal(t, int64(4), root.IndexNodes[0].IndexNodes[0].DataNodes[3].Items[0].Key)
		assert.Equal(t, int64(5), root.IndexNodes[0].IndexNodes[0].DataNodes[3].Items[1].Key)

		// >>>>> Protrude an iNode.

		// Protrude the index node and retrieve the side (iSide).
		_, err = root.IndexNodes[0].IndexNodes[0].protrudeInOddBpWidth()
		require.NoError(t, err, "protrude should not return an error")

		// Remove the first index node from the parent's index nodes.
		root.IndexNodes[0].IndexNodes = root.IndexNodes[0].IndexNodes[1:]

		// Create a new index node instance.
		newNode := &BpIndex{}

		// >>>>> Take apart and reassemble.

		// Take apart the index node and reassemble it using iSide.
		// newNode.TakeApartReassemble(iSide, root.IndexNodes[0])

		// Update the parent's index nodes with the reassembled node.
		root.IndexNodes[0] = newNode

		// Check the updated data structure.
		assert.Equal(t, []int64{3, 5}, root.IndexNodes[0].Index)
		assert.Equal(t, []int64{1}, root.IndexNodes[0].IndexNodes[0].Index)
		assert.Equal(t, []int64{4}, root.IndexNodes[0].IndexNodes[1].Index)
		assert.Equal(t, []int64{6}, root.IndexNodes[0].IndexNodes[2].Index)
		assert.Equal(t, []int64{11, 13}, root.IndexNodes[1].Index)
		assert.Equal(t, int64(1), root.IndexNodes[0].IndexNodes[0].DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(1), root.IndexNodes[0].IndexNodes[0].DataNodes[1].Items[0].Key)
		assert.Equal(t, int64(2), root.IndexNodes[0].IndexNodes[0].DataNodes[1].Items[1].Key)
		assert.Equal(t, int64(3), root.IndexNodes[0].IndexNodes[1].DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(4), root.IndexNodes[0].IndexNodes[1].DataNodes[1].Items[0].Key)
		assert.Equal(t, int64(5), root.IndexNodes[0].IndexNodes[2].DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(6), root.IndexNodes[0].IndexNodes[2].DataNodes[1].Items[0].Key)

		// Print the resulting tree structure.
		// newNode.Print()
	})
	t.Run("protrude 15 15to15 15", func(t *testing.T) {
		// Set up Bp Parameters.
		BpWidth = 3
		BpHalfWidth = 2

		// Create a root tree with a specific structure (indexes are [15], [15to15], [15]).
		root := createRootTree15and15to15and15()

		// >>>>> Insert a BpItem.

		// Insert a key (15) among the data items in the first data node.
		root.IndexNodes[0].DataNodes[0].insertAmong(BpItem{Key: 15})

		// >>>>> Split a BpData.

		// Split the data node and obtain the side and any potential error.
		dSide, err := root.IndexNodes[0].DataNodes[0].split()
		require.NoError(t, err, "split should not return an error")

		// >>>>> Merge a popped dNode.

		// Merge a popped data node based on the split side.
		err = root.IndexNodes[0].mergePopDnode(dSide)
		require.NoError(t, err, "mergePopDnode should not return an error")

		// >>>>> Protrude an iNode.

		// Protrude the index node and retrieve the side (iSide).
		_, err = root.IndexNodes[0].protrudeInOddBpWidth()
		require.NoError(t, err, "protrude should not return an error")

		// Remove the first index node from the parent's index nodes.
		root.IndexNodes = root.IndexNodes[1:]

		// Create a new index node instance.
		newNode := &BpIndex{}

		// >>>>> Take apart and reassemble.

		// Take apart the index node and reassemble it using iSide.
		// newNode.TakeApartReassemble(iSide, root)

		// Check the updated data structure.
		assert.Equal(t, 1, newNode.IndexNodes[0].DataNodes[0].dataLength())
		assert.Equal(t, 2, newNode.IndexNodes[0].DataNodes[1].dataLength())

		// Print the resulting tree structure.
		// newNode.Print()
	})
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
		err := idx.insertBpIX(test.newIndex)
		require.NoError(t, err)

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
