package bpTree

import (
	"sort"
)

func (inode *BpIndex) delete(item BpItem) (deleted bool, direction int, ix int) {
	// Use binary search to find the index (ix) where the key should be deleted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	if len(inode.IndexNodes) > 0 {
		deleted, direction, ix = inode.IndexNodes[ix].delete(item)
		return
	}

	if len(inode.DataNodes) > 0 {
		deleted, direction, ix = inode.deleteBottomItem(item) // 已经接近底层，开始删蛋资料
		return
	}

	return
}

// deleteBottomItem deletes the specified BpItem from the DataNodes near the bottom layer of the BpIndex.
func (inode *BpIndex) deleteBottomItem(item BpItem) (deleted bool, direction int, ix int) {
	// Use binary search to find the index (ix) where the key should be inserted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Call the delete method on the corresponding DataNode to delete the item.
	deleted, direction = inode.DataNodes[ix].delete(item)

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
