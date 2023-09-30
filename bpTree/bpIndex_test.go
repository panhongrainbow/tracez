package bpTree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

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
		iSide, err := root.IndexNodes[0].IndexNodes[0].protrude()
		require.NoError(t, err, "protrude should not return an error")

		// Remove the first index node from the parent's index nodes.
		root.IndexNodes[0].IndexNodes = root.IndexNodes[0].IndexNodes[1:]

		// Create a new index node instance.
		newNode := &BpIndex{}

		// >>>>> Take apart and reassemble.

		// Take apart the index node and reassemble it using iSide.
		newNode.TakeApartReassemble(iSide, root.IndexNodes[0])

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
		iSide, err := root.IndexNodes[0].protrude()
		require.NoError(t, err, "protrude should not return an error")

		// Remove the first index node from the parent's index nodes.
		root.IndexNodes = root.IndexNodes[1:]

		// Create a new index node instance.
		newNode := &BpIndex{}

		// >>>>> Take apart and reassemble.

		// Take apart the index node and reassemble it using iSide.
		newNode.TakeApartReassemble(iSide, root)

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

		// Call the insertBpIdxNewValue function
		_, _, _ = index.insertBpIdxNewValue(nil, newItem)

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
