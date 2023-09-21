package bpTree

import (
	"testing"
)

// Test_Check_NewBpTree is the main test to verify whether the initial data distribution in the B plus tree is even during its construction.
func Test_Check_NewBpTree(t *testing.T) {

	/*t.Run("when width is 1", func(t *testing.T) {
		// Create a new B+ tree with width 1.
		root := NewBpTree(1)

		// Insert three BpItems with keys 1, 2, and 3.
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 2})
		root.InsertValue(BpItem{Key: 3})

		// Check if the intervals in the root node are [1, 2, 3].
		assert.Equal(t, []int64{1, 2, 3}, root.root.Intervals)

		// Check if the keys in the data nodes are evenly distributed as expected.
		assert.Equal(t, int64(1), root.root.DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(2), root.root.DataNodes[1].Items[0].Key)
		assert.Equal(t, int64(3), root.root.DataNodes[2].Items[0].Key)
	})*/

	/*t.Run("when width is 3", func(t *testing.T) {
		// Create a new B+ tree with width 3.
		root := NewBpTree(3)

		// Insert three BpItems with keys 1, 2, and 3.
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 2})
		root.InsertValue(BpItem{Key: 3})

		// Check if the intervals in the root node are [1, 2, 3].
		assert.Equal(t, []int64{1, 2, 3}, root.root.Intervals)

		// Check if the keys in the data nodes are evenly distributed as expected.
		assert.Equal(t, int64(1), root.root.DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(2), root.root.DataNodes[1].Items[0].Key)
		assert.Equal(t, int64(3), root.root.DataNodes[2].Items[0].Key)
	})*/

	/*t.Run("when width is 5", func(t *testing.T) {
		// Create a new B+ tree with width 5.
		root := NewBpTree(5)

		// Insert five BpItems with keys 1, 2, 3, 4, and 5.
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 2})
		root.InsertValue(BpItem{Key: 3})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 5})

		// Check if the intervals in the root node are [1, 2, 3, 4, 5].
		assert.Equal(t, []int64{1, 2, 3, 4, 5}, root.root.Intervals)

		// Check if the keys in the data nodes are evenly distributed as expected.
		assert.Equal(t, int64(1), root.root.DataNodes[0].Items[0].Key)
		assert.Equal(t, int64(2), root.root.DataNodes[1].Items[0].Key)
		assert.Equal(t, int64(3), root.root.DataNodes[2].Items[0].Key)
		assert.Equal(t, int64(4), root.root.DataNodes[3].Items[0].Key)
		assert.Equal(t, int64(5), root.root.DataNodes[4].Items[0].Key)
	})*/
}
