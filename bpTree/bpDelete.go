package bpTree

import (
	"fmt"
	"reflect"
	"sort"
)

// delete is a method of the BpIndex type that deletes the specified BpItem.
func (inode *BpIndex) delete(item BpItem) (deleted, updated bool, ix int, err error) {
	// Use binary search to find the index (ix) where the key should be deleted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Check if there are any index nodes.
	if len(inode.IndexNodes) > 0 {
		// Recursive call to delete method on the corresponding IndexNode.
		deleted, updated, _, err = inode.IndexNodes[ix].delete(item)

		if updated {
			updated, err = inode.updateIndex(ix)
		}

		// Here, testing is being conducted (测试用).
		fmt.Println("not in Bottom", ix)
	}

	// Check if there are any data nodes.
	if len(inode.DataNodes) > 0 {
		// Call the deleteBottomItem method on the current node as it is close to the bottom layer.
		// This signifies the beginning of deleting data.

		// Here, adjustments may be made to IX (IX 在这里可能会被修改) ‼️
		deleted, updated, ix, err = inode.deleteBottomItem(item) // Possible index update ‼️

		// Here, testing is being conducted (测试用).
		fmt.Println("in Bottom", ix)
	}

	// Return the results of the deletion.
	return
}

// delete is a method of the BpIndex type that deletes the specified BpItem.
func (inode *BpIndex) deleteDeprecated(item BpItem) (deleted, updated bool, direction int, ix int, err error) {
	// Use binary search to find the index (ix) where the key should be deleted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Check if there are any index nodes.
	if len(inode.IndexNodes) > 0 {
		// Recursive call to delete method on the corresponding IndexNode.
		deleted, updated, direction, _, err = inode.IndexNodes[ix].deleteDeprecated(item)

		if updated {
			updated, err = inode.updateIndex(ix)
		}

		// Here, testing is being conducted (测试用).
		fmt.Println("not in Bottom", ix)
	}

	// Check if there are any data nodes.
	if len(inode.DataNodes) > 0 {
		// Call the deleteBottomItem method on the current node as it is close to the bottom layer.
		// This signifies the beginning of deleting data.

		// Here, adjustments may be made to IX (IX 在这里可能会被修改) ‼️
		deleted, updated, direction, ix, err = inode.deleteBottomItemDeprecated(item) // Possible index update ‼️

		// Here, testing is being conducted (测试用).
		fmt.Println("in Bottom", ix)
	}

	// Return the results of the deletion.
	return
}

func (inode *BpIndex) deleteBottomItem(item BpItem) (deleted, updated bool, ix int, err error) {
	// Use binary search to find the index (ix) where the key should be inserted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Call the delete method on the corresponding DataNode to delete the item.
	deleted, _ = inode.DataNodes[ix]._delete(item)

	// The following are operations for updating the index (更新索引) ‼️
	if deleted == true && len(inode.DataNodes[ix].Items) > 0 {
		updated, err = inode.updateBottomIndex(ix)
	}

	// Return the results of the deletion.
	return
}

// 准备考虑废除 mark 功能 🔥

// deleteBottomItem deletes the specified BpItem from the DataNodes near the bottom layer of the BpIndex.
func (inode *BpIndex) deleteBottomItemDeprecated(item BpItem) (deleted, updated bool, direction int, ix int, err error) {
	// ➡️ Executing the process of data deletion to remove item.

	// Use binary search to find the index (ix) where the key should be inserted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// If it is possible to delete data that extends to neighboring nodes, the index cannot be updated on time.
	// In such cases, a mask must be used to temporarily maintain the old index. ‼️
	// 如果可能会刖除到邻近结点的资料，就无法及时更新索引，要用 mask，暂时维持旧的索引 ‼️
	var mark bool
	if !(ix > 0 && ix < len(inode.Index)-1) {
		mark = true
	}

	// Call the delete method on the corresponding DataNode to delete the item.
	deleted, direction = inode.DataNodes[ix].delete(item, mark)

	// Adjust the index based on the direction of deletion.
	if deleted == true {
		if direction == deleteRightOne {
			ix = ix + 1
		} else if direction == deleteLeftOne {
			ix = ix - 1
		}
	}

	// The following are operations for updating the index (更新索引) ‼️
	updated, err = inode.updateBottomIndex(ix)

	// Return the results of the deletion.
	return
}

// This function is for updating non-bottom-level indices. (更新非底层的索引)
func (inode *BpIndex) updateIndex(ix int) (updated bool, err error) {
	if inode.Index[ix-1] != inode.IndexNodes[ix].Index[0] {
		inode.Index[ix-1] = inode.IndexNodes[ix].Index[0]
		updated = true
	}
	return
}

// updateBottomIndex cleans the data at the bottom level and updates the index. (清理底层资料并更新索引)
func (inode *BpIndex) updateBottomIndex(ix int) (updated bool, err error) {
	// Create a new index.
	newIndex := make([]int64, 0)

	// First, check whether it is a data node or an index node.
	if len(inode.DataNodes) > 0 {
		// Clean the current data index.
		err = inode.CleanMark(ix) // 清理
		if err != nil {
			return
		}
		// Reconstruct the new index.
		for i := 1; i < len(inode.DataNodes); i++ {
			newIndex = append(newIndex, inode.DataNodes[i].Items[0].Key)
		}
	} else {
		// Handle empty data nodes separately.
		err = fmt.Errorf("this is an empty node")
		return
	}

	// Compare the old and new indices; if different, update the index.
	if reflect.DeepEqual(inode.Index, newIndex) == false {
		inode.Index = newIndex
		updated = true
	}

	// Finally, perform the return.
	return
}

// CleanMark is used to maintain the BpIndex node to an appropriate size.
func (inode *BpIndex) CleanMark(ix int) (err error) {
	// ➡️ Here, addition and deletion will have different impacts.
	// Updating the index to ensure the latest and most accurate representation of the BpData.
	// When len(BpData) > 0, the quantity of BpData will be one more than the index.
	//             index[0]  index[1]  index[2]  index[3]
	// ┌─────────┬─────────┬─────────┬─────────┬─────────┐
	// │ bpdata0 │ bpdata1 │ bpdata2 │ bpdata3 │ bpdata4 │
	// └─────────┴─────────┴─────────┴─────────┴─────────┘
	//   ix=0      ix=1      ix=2      ix=3      ix=4

	// When len(BpData) = 1, the quantity of BpData is equal to the index.
	//   index[0]
	// ┌─────────┐
	// │ bpdata0 │
	// └─────────┘
	//   ix=0

	// Preventing issues that may arise when adding new data.
	//	index[0]
	// ┌─────────┬─────────┐
	// │ bpdata0 │ bpdata1 │
	// │         │ (mark)  │
	// └─────────┴─────────┘
	//	ix=0      ix=1
LOOP:
	data := inode.DataNodes[ix]
	if len(inode.DataNodes) == 2 {
		// The node is too small; data cannot be deleted anymore.
		return
	}
	if len(inode.DataNodes) <= 1 {
		// An error occurred here.
		err = fmt.Errorf("the index node has too small data to become a node")
		return
	}

	// Organizing data nodes has begun.
	for i := 0; i < len(inode.DataNodes[ix].Items); i++ {
		if data.Items[i].Mask == true {
			copy(data.Items[i:], data.Items[i+1:])
			data.Items = data.Items[:len(data.Items)-1]
			goto LOOP
		}
	}

	// Cleanup complete, returning.
	return
}
