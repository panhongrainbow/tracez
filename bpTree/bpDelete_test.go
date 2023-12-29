package bpTree

import (
	"fmt"
	"testing"
)

// Test_Check_BpData_delete is a test function for the delete method of the BpIndex type.
func Test_Check_BpIndex_delete(t *testing.T) {
	// Data deletion will take place in different directions depending on the context.
	// If the data is continuous and spans points, deletion will occur from the left.
	// However, for non-continuous data or when there are no span points, deletion will occur from the right.
	// 连续并跨节点往左边，不连续并没跨节点往右边 ‼️

	// 🧪 This test is to evaluate the deletion of consecutive data in a B Plus tree.
	t.Run("The continuous data deletion.", func(t *testing.T) {
		// Set up the total width and half-width for the B Plus Tree.
		BpWidth = 3
		BpHalfWidth = 2

		// Create seven BpData nodes.
		bpData0 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 1}, {Key: 2}},
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
			Items:    []BpItem{{Key: 5}},
		}

		bpData4 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 10}},
		}

		bpData5 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 11}},
		}

		bpData6 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 12}, {Key: 13}},
		}

		// Establish the link between the two nodes. (Next Direction)
		bpData0.Next = &bpData1
		bpData1.Next = &bpData2
		bpData2.Next = &bpData3
		bpData3.Next = &bpData4
		bpData4.Next = &bpData5
		bpData5.Next = &bpData6
		bpData6.Next = nil

		// Establish the link between the two nodes. (Previous Direction)
		bpData6.Previous = &bpData5
		bpData5.Previous = &bpData4
		bpData4.Previous = &bpData3
		bpData3.Previous = &bpData2
		bpData2.Previous = &bpData1
		bpData1.Previous = &bpData0
		bpData0.Previous = nil

		// Set up a top-level index node. (整个树建立好)
		inode := &BpIndex{
			Index: []int64{5, 10},
			IndexNodes: []*BpIndex{
				{
					Index: []int64{5},
					DataNodes: []*BpData{ // IndexNode ▶️
						&bpData0, // DataNode 🗂️
						&bpData1, // DataNode 🗂️
					},
				},
				{
					Index: []int64{5},
					DataNodes: []*BpData{ // IndexNode ▶️
						&bpData2, // DataNode 🗂️
						&bpData3, // DataNode 🗂️
					},
				},
				{
					Index: []int64{11, 12},
					DataNodes: []*BpData{ // IndexNode ▶️
						&bpData4, // DataNode 🗂️
						&bpData5, // DataNode 🗂️
						&bpData6, // DataNode 🗂️
					},
				},
			},
			DataNodes: []*BpData{},
		}

		// Execute the delete commands.
		deleted, updated, ix, err := inode.delRoot(BpItem{Key: 5})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 5})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 5})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 5})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 5})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 10})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 2})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 13})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 11})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 1})
		deleted, updated, ix, err = inode.delRoot(BpItem{Key: 12})
		fmt.Println(deleted, updated, ix, err)
	})

	// 🧪 This example is intended to test the deletion of multi-layered B plus tree data.
	t.Run("The multi layered data deletion.", func(t *testing.T) {
		// Set up the total width and half-width for the B Plus Tree.
		BpWidth = 3
		BpHalfWidth = 2

		// Create seven BpData nodes.
		bpData0 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 1}},
		}

		bpData1 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 2}},
		}

		bpData2 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 3}},
		}

		bpData3 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 4}},
		}

		bpData4 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 5}},
		}

		bpData5 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 6}},
		}

		bpData6 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 7}},
		}

		bpData7 := BpData{
			Previous: nil,
			Next:     nil,
			Items:    []BpItem{{Key: 8}, {Key: 9}},
		}

		// Establish the link between the two nodes. (Next Direction)
		bpData0.Next = &bpData1
		bpData1.Next = &bpData2
		bpData2.Next = &bpData3
		bpData3.Next = &bpData4
		bpData4.Next = &bpData5
		bpData5.Next = &bpData6
		bpData6.Next = &bpData7
		bpData7.Next = nil

		// Establish the link between the two nodes. (Previous Direction)
		bpData7.Previous = &bpData6
		bpData6.Previous = &bpData5
		bpData5.Previous = &bpData4
		bpData4.Previous = &bpData3
		bpData3.Previous = &bpData2
		bpData2.Previous = &bpData1
		bpData1.Previous = &bpData0
		bpData0.Previous = nil

		// Set up a top-level index node. (整个树建立好)
		inode := &BpIndex{
			Index: []int64{5},
			IndexNodes: []*BpIndex{ // IndexNode ▶️
				{
					Index: []int64{3},
					IndexNodes: []*BpIndex{ // IndexNode ▶️
						{
							Index: []int64{2},
							DataNodes: []*BpData{
								&bpData0, // DataNode 🗂️
								&bpData1, // DataNode 🗂️
							},
						},
						{
							Index: []int64{4},
							DataNodes: []*BpData{
								&bpData2, // DataNode 🗂️
								&bpData3, // DataNode 🗂️
							},
						},
					},
					DataNodes: []*BpData{},
				},
				{
					Index: []int64{7},
					IndexNodes: []*BpIndex{ // IndexNode ▶️
						{
							Index: []int64{6},
							DataNodes: []*BpData{
								&bpData4, // DataNode 🗂️
								&bpData5, // DataNode 🗂️
							},
						},
						{
							Index: []int64{8},
							DataNodes: []*BpData{
								&bpData6, // DataNode 🗂️
								&bpData7, // DataNode 🗂️
							},
						},
					},
					DataNodes: []*BpData{},
				},
			},
			DataNodes: []*BpData{},
		}

		// Execute the delete commands.
		deleted, updated, ix, err := inode.delRoot(BpItem{Key: 4})
		fmt.Println(deleted, updated, ix, err)

		fmt.Println("must be empty", inode.Index)
	})
}
