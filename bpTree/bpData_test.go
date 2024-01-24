package bpTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_Check_BpData_index includes comprehensive unit tests covering cases with item presence,
// item absence (empty slice), and side node handling, ensuring the method's correctness.
func Test_Check_BpData_index(t *testing.T) {
	// Create a BpData instance with items.
	data := &BpData{
		Items: []BpItem{
			{Key: 1, Val: "Value1"},
			{Key: 2, Val: "Value2"},
		},
		Previous: &BpData{},
		Next:     &BpData{},
	}

	// Test case: Retrieve key when items are present.
	key, err := data.index()
	assert.Nil(t, err, "Expected no error when items are present")
	assert.Equal(t, int64(1), key, "Expected key to be 1")

	// Test case: Retrieve key when items are not present.
	data.Items = nil
	key, err = data.index()
	assert.Error(t, err, "expect there is no available index for BpData")
	assert.Equal(t, int64(0), key, "Expected key to be 0 for an empty slice")
}

// Test_Check_BpData_insert checks the insert method, which inserts BpItem elements into the BpData.
func Test_Check_BpData_insert(t *testing.T) {
	data := &BpData{}

	item1 := BpItem{Key: 1, Val: "Value1"}
	data.insert(item1)
	assert.Equal(t, 1, data.dataLength(), "Expected one item in the slice")

	item2 := BpItem{Key: 2, Val: "Value2"}
	data.insert(item2)
	assert.Equal(t, 2, data.dataLength(), "Expected two items in the slice")
}

// Test_Check_BpData_split checks the split method, which divides a BpData node into two nodes if it contains more items than the specified width.
func Test_Check_BpData_split(t *testing.T) {
	// Set parameters.
	BpWidth = 3
	BpHalfWidth = 2

	// Create a BpData node with five items.
	data := &BpData{
		Items: []BpItem{
			{Key: 1, Val: "Value1"},
			{Key: 2, Val: "Value2"},
			{Key: 3, Val: "Value3"},
			{Key: 4, Val: "Value4"},
			{Key: 5, Val: "Value5"},
		},
	}

	// Split the data node into two nodes with a width of 2.
	side, err := data.split()
	assert.Nil(t, err, "Expected no error")

	// The index for generating a new node is 4.
	var key int64
	key, err = side.index()
	assert.Equal(t, int64(4), key)

	// Check the state of the new node created by the split.
	assert.Len(t, data.Next.Items, 2, "Expected the new node slice to have 2 items after split")

	// Check the state of the original node.
	assert.Len(t, data.Items, 3, "Expected the original slice to have 3 items after split")
}

// Test_Check_BpData_delete is a test function for the delete method of the BpData type.
/*func Test_Check_BpData_delete(t *testing.T) {
	// Run the first subtest: Consider the same key in a single node.
	t.Run("Consider the same key in a single node.", func(t *testing.T) {
		// Set up a BpData instance with a slice of BpItems.
		data := &BpData{
			Items: []BpItem{
				{Key: 1, Val: "Value1"},
				{Key: 2, Val: "Value2"},
				{Key: 3, Val: "Value3"},
				{Key: 4, Val: "Value4"},
				{Key: 5, Val: "Value5"},
			},
		}

		// Check if the data length is as expected before deletion.
		require.Equal(t, 5, data.dataLength(), "Expected five items in the slice")

		// Create a BpItem to be deleted and perform the deletion.
		item := BpItem{Key: 3, Val: "Value3"}
		deleted, direction := data.delete(item, false)

		// Check if the deletion was successful and if the data length is as expected after deletion.
		require.True(t, deleted)
		require.Equal(t, 4, data.dataLength(), "Expected four items in the slice")
		require.Equal(t, deleteMiddleOne, direction)
	})

	// Run the second subtest: Consider the same key in two nodes.
	t.Run("Consider the same key in two nodes.", func(t *testing.T) {
		// Set up the total width and half-width for the B Plus Tree.
		BpWidth = 3
		BpHalfWidth = 2

		// Create two BpData nodes representing nodes with overlapping keys.
		bpData1 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 1}, {Key: 2}},
		}

		bpData2 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 2}, {Key: 3}},
		}

		// Establish the link between the two nodes.
		bpData1.Next = &bpData2
		bpData2.Previous = &bpData1

		// Perform the deletion of a key in the first node.
		deleted, direction := bpData1.delete(BpItem{Key: 2}, false)
		require.True(t, deleted)
		require.Len(t, bpData1.Items, 1)
		require.Len(t, bpData2.Items, 2)
		require.Equal(t, deleteMiddleOne, direction) // hit middle one

		// Perform the deletion of the same key in the first node again.
		deleted, direction = bpData1.delete(BpItem{Key: 2}, false)
		require.True(t, deleted)
		require.Len(t, bpData1.Items, 1)
		require.Len(t, bpData2.Items, 1)            // Go to the neighbor BpData node and perform the deletion.
		require.Equal(t, deleteRightOne, direction) // hit middle one
	})
}*/
