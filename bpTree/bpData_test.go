package bpTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_BpData_getBpDataIndex includes comprehensive unit tests covering cases with item presence,
// item absence (empty slice), and side node handling, ensuring the method's correctness.
func Test_BpData_getBpDataIndex(t *testing.T) {
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
	key, err := data.getBpDataIndex()
	assert.Nil(t, err, "Expected no error when items are present")
	assert.Equal(t, int64(1), key, "Expected key to be 1")

	// Test case: Retrieve key when items are not present.
	data.Items = nil
	key, err = data.getBpDataIndex()
	assert.Error(t, err, "expect no data available")
	assert.Equal(t, int64(0), key, "Expected key to be 0 for an empty slice")
}

// Test_BpData_Check_insertBpDataValue checks the insertBpDataValue method, which inserts BpItem elements into the BpData.
func Test_BpData_Check_insertBpDataValue(t *testing.T) {
	data := &BpData{}

	item1 := BpItem{Key: 1, Val: "Value1"}
	data.insertBpDataValue(item1)
	assert.Len(t, data.Items, 1, "Expected one item in the slice")

	item2 := BpItem{Key: 2, Val: "Value2"}
	data.insertBpDataValue(item2)
	assert.Len(t, data.Items, 2, "Expected two items in the slice")
}

// Test_BpData_Check_split checks the split method, which divides a BpData node into two nodes if it contains more items than the specified width.
func Test_BpData_Check_split(t *testing.T) {
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
	key, err := data.split()
	assert.Nil(t, err, "Expected no error")

	// The index for generating a new node is 4.
	assert.Equal(t, int64(4), key)

	// Check the state of the original node.
	assert.Len(t, data.Items, 3, "Expected the original slice to have 3 items after split")

	// Check the state of the new node created by the split.
	assert.Len(t, data.Next.Items, 2, "Expected the new node slice to have 2 items after split")
}
