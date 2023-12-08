package bpTree

import (
	"fmt"
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
		}

		bpData1 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 2}, {Key: 3}},
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
		deleted, updated, direction, ix, err := inode.deleteBottomItemDeprecated(BpItem{Key: 2})
		require.True(t, deleted)
		require.True(t, updated)                  // Updated the index ‼️
		require.Equal(t, []int64{3}, inode.Index) // The index has been updated (2->3) ‼️
		require.Equal(t, deleteMiddleOne, direction)
		require.Equal(t, 1, ix) // Perform deletion in the 2nd node
		require.NoError(t, err)

		// Here is a key point: when data is deleted in the 2nd bpData node, the index is immediately updated.
		// The index changes from []int64{2} to []int64{3}, and the new index, in the next deletion, guides the operation to the 1st node.
		// (删除第2个bpData节点后，立即更新索引为[]int64{3}，下次删除操作将指向第1个节点)

		// Execute the delete command for the second time.
		deleted, updated, direction, ix, err = inode.deleteBottomItemDeprecated(BpItem{Key: 2})
		require.True(t, deleted)
		require.False(t, updated)                 // Not updated to the index ‼️
		require.Equal(t, []int64{3}, inode.Index) // No update to the index  (3->3) ‼️
		require.Equal(t, deleteMiddleOne, direction)
		require.Equal(t, 0, ix) // Guide the deletion operation to the 1st node. (引到第一节点)
		require.NoError(t, err)
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
		}

		bpData1 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 5}, {Key: 5}},
		}

		bpData2 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 5}, {Key: 5}},
		}

		bpData3 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 5}, {Key: 10}},
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
		deleted, updated, ix, err := inode.delete(BpItem{Key: 5})
		require.True(t, deleted)
		require.True(t, updated)                   // Updated the index ‼️
		require.Equal(t, []int64{10}, inode.Index) // The index has been updated (5->10) ‼️
		require.Equal(t, 1, ix)                    // Delete data on the second BpIndex Node. (删除第二个分支里的资料) ‼️
		require.NoError(t, err)

		fmt.Println("-------------------")

		// Execute the delete command for the second time.
		deleted, updated, ix, err = inode.delete(BpItem{Key: 5})
		require.True(t, deleted)
		require.False(t, updated)                  // Not updated to the index ‼️
		require.Equal(t, []int64{10}, inode.Index) // No update to the index  (10->10) ‼️
		require.Equal(t, 0, ix)                    // Delete data on the first BpIndex Node. (删除第一个分支里的资料) ‼️
		require.NoError(t, err)

		fmt.Println("-------------------")

		// Execute the delete command for the third time.
		deleted, updated, ix, err = inode.delete(BpItem{Key: 5})
	})
}
