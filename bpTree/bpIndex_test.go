package bpTree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

// Test_Check_BpIndex_getBpIdxIndex includes subtests that explore key retrieval from the Index slice,
// handling of side nodes within the linked DataNode, and successful key retrieval from the DataNode.
func Test_Check_BpIndex_getBpIdxIndex(t *testing.T) {
	t.Run("RetrieveFromIndexSlice", func(t *testing.T) {
		// Test case: Retrieve key from the Index slice.
		index := &BpIndex{
			Index: []int64{10, 20, 30},
		}

		key, err := index.getBpIdxIndex()
		assert.Nil(t, err, "Expected no error when retrieving from the Index slice")
		assert.Equal(t, int64(10), key, "Expected key to be 10")
	})

	t.Run("RetrieveFromIndexNode", func(t *testing.T) {
		// Test case: Retrieve key from the associated DataNode (BpData),
		// However, it is the side node and return no key message.
		index := &BpIndex{}
		dataNode1 := &BpData{
			Items: []BpItem{
				{Key: 5, Val: "Value1"},
			},
		}
		dataNode2 := &BpData{
			Items: []BpItem{
				{Key: 6, Val: "Value2"},
			},
		}
		index.DataNodes = append(index.DataNodes, dataNode1, dataNode2)

		key, err := index.getBpIdxIndex()
		assert.Equal(t, int64(0), key, "Expected key to be 0")
		// index 切片没资料就是错
		assert.EqualError(t, err, "no key available", "Expected specific error message because of no key in index")
	})

	t.Run("RetrieveFromDataNode", func(t *testing.T) {
		// Test case: Retrieve key from the associated DataNode (BpData).
		index := &BpIndex{}
		dataNode1 := &BpData{
			Items: []BpItem{
				{Key: 5, Val: "Value1"},
			},
		}
		dataNode2 := &BpData{
			Items: []BpItem{
				{Key: 6, Val: "Value2"},
			},
		}

		// Now the data node has data.
		index.DataNodes = append(index.DataNodes, dataNode1, dataNode2)

		key, err := index.getBpIdxIndex()
		assert.Equal(t, int64(0), key, "Expected key to be 0")
		// index 切片没资料就是错
		assert.EqualError(t, err, "no key available", "Expected specific error message because of no key in index")
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
		err := idx.insertBpIdxNewIndex(test.newIndex)
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
