package bpTree

import (
	"sort"
)

// delete is a method of the BpIndex type that deletes the specified BpItem.
func (inode *BpIndex) delete(item BpItem) (deleted, updated bool, direction int, ix int, err error) {
	// Use binary search to find the index (ix) where the key should be deleted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Check if there are any index nodes.
	if len(inode.IndexNodes) > 0 {
		// Recursive call to delete method on the corresponding IndexNode.
		deleted, updated, direction, ix, err = inode.IndexNodes[ix].delete(item)
	}

	// Check if there are any data nodes.
	if len(inode.DataNodes) > 0 {
		// Call the deleteBottomItem method on the current node as it is close to the bottom layer.
		// This signifies the beginning of deleting data.
		deleted, direction, ix = inode.deleteBottomItem(item)
	}

	// Return the results of the deletion.
	return
}

// deleteBottomItem deletes the specified BpItem from the DataNodes near the bottom layer of the BpIndex.
func (inode *BpIndex) deleteBottomItem(item BpItem) (deleted bool, direction int, ix int) {
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

	// Return the results of the deletion.
	return
}
