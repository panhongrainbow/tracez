package bpTree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Test_Check_BpData_delete is a test function for the delete method of the BpIndex type.
func Test_Check_BpIndex_delete(t *testing.T) {
	// Delete data at the closest bottom BpIndex. When entering the BpData node,
	// it may also be possible to search the surrounding.
	t.Run("Delete data at the closest bottom BpIndex.", func(t *testing.T) {
		// Set up the total width and half-width for the B Plus Tree.
		BpWidth = 3
		BpHalfWidth = 2

		// Create two BpData nodes representing nodes with overlapping keys.
		bpData0 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 1}, {Key: 2}},
			Split:    false,
		}

		bpData1 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 2}, {Key: 3}},
			Split:    false,
		}

		// Establish the link between the two nodes.
		bpData0.Next = &bpData1
		bpData1.Previous = &bpData0

		// To test data deletion, create a BpIndex node.
		inode := &BpIndex{
			Index:      []int64{2},
			IndexNodes: []*BpIndex{},
			DataNodes:  []*BpData{&bpData0, &bpData1},
		}

		// Execute the delete command for the first time.
		deleted, direction, ix := inode.deleteBottomItem(BpItem{Key: 2})
		require.True(t, deleted)                     // It will always delete data, just in different BpData nodes.
		require.Equal(t, deleteMiddleOne, direction) // Delete directly at the specified node without removing data from neighbor nodes.
		require.Equal(t, 0, ix)                      // Delete data on the first BpData Node.

		// Execute the delete command for the second time.
		deleted, direction, ix = inode.deleteBottomItem(BpItem{Key: 2})
		require.True(t, deleted)                    // It will always delete data, just in different BpData nodes.
		require.Equal(t, deleteRightOne, direction) // Delete at the neighbor nodes.
		require.Equal(t, 1, ix)                     // Delete data on the second BpData Node.
	})
	t.Run("Delete data at the BpIndex.", func(t *testing.T) {
		// Set up the total width and half-width for the B Plus Tree.
		BpWidth = 3
		BpHalfWidth = 2

		// Create two BpData nodes representing nodes with overlapping keys.
		bpData0 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 1}},
			Split:    false,
		}

		bpData1 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 5}, {Key: 5}},
			Split:    false,
		}

		bpData2 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 5}, {Key: 5}},
			Split:    false,
		}

		bpData3 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 5}, {Key: 10}},
			Split:    false,
		}

		// Establish the link among the nodes.
		bpData0.Previous = nil
		bpData1.Previous = &bpData0
		bpData2.Previous = &bpData1
		bpData3.Previous = &bpData2

		bpData0.Next = &bpData1
		bpData1.Next = &bpData2
		bpData2.Next = &bpData3
		bpData3.Next = nil

		// Set up a top-level index node.
		inode := &BpIndex{
			Index: []int64{5},
			IndexNodes: []*BpIndex{
				{
					Index: []int64{5},
					DataNodes: []*BpData{
						&bpData0,
						&bpData1,
					},
				},
				{
					Index: []int64{5},
					DataNodes: []*BpData{
						&bpData2,
						&bpData3,
					},
				},
			},
			DataNodes: []*BpData{},
		}

		// Execute the delete command for the first time.
		deleted, direction, ix := inode.delete(BpItem{Key: 5})
		require.True(t, deleted)                    // It will always delete data, just in different BpData nodes.
		require.Equal(t, deleteRightOne, direction) // Delete at the neighbor nodes.
		require.Equal(t, 1, ix)                     // Delete data on the first BpData Node.

		// Execute the delete command for the second time.
		deleted, direction, ix = inode.delete(BpItem{Key: 5})
		require.True(t, deleted)                    // It will always delete data, just in different BpData nodes.
		require.Equal(t, deleteRightOne, direction) // Delete at the neighbor nodes.
		require.Equal(t, 1, ix)                     // Delete data on the first BpData Node.
	})
}
