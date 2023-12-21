package bpTree

import (
	"fmt"
	"sort"
)

// ➡️ The functions related to direction.

// delAndDir performs data deletion based on automatic direction detection.
// 自动判断资料删除方向，其實會由不同方向進行刪除
func (inode *BpIndex) delAndDir(item BpItem) (deleted, updated bool, ix int, err error) {
	// Use binary search to find the index (ix) where the key should be deleted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // no equal sign ‼️ no equal sign means delete to the right ‼️
	})

	// Check if deletion should be performed by the leftmost node first.
	if ix >= 1 { // After the second index node, it's possible to borrow data from the left ⬅️ node
		// Length of the left node
		length := len(inode.IndexNodes[ix-1].Index)

		// If it is continuous data (same value) (5❌ - 5 - 5 - 5 - 5 - 6 - 7 - 8)
		if inode.IndexNodes[ix].Index[0] == inode.IndexNodes[ix-1].Index[length-1] {
			deleted, updated, ix, err = inode.deleteToLeft(item) // Delete to the leftmost node ‼️ (向左砍)
			return
		}
	}

	// If it is discontinuous data (different values) (5 - 5 - 5 - 5 - 5❌ - 6 - 7 - 8)
	deleted, updated, ix, err = inode.deleteToRight(item) // Delete to the rightmost node ‼️ (向右砍)

	// Return the results
	return
}

// deleteToLeft is a method of the BpIndex type that deletes the leftmost specified BpItem. (由左边删除 👈 ‼️)
func (inode *BpIndex) deleteToLeft(item BpItem) (deleted, updated bool, ix int, err error) {
	// Check if there are any index nodes.
	if len(inode.IndexNodes) > 0 {
		// Use binary search to find the index (ix) where the key should be deleted.
		ix = sort.Search(len(inode.Index), func(i int) bool {
			return inode.Index[i] >= item.Key // equal sign ‼️ no equal sign means delete to the left ‼️
		})

		// Recursion keeps deletion in the left direction. 递归一直向左砍 ⬅️
		deleted, updated, _, err = inode.IndexNodes[ix].deleteToLeft(item)

		// Immediately update the index of index node.
		if updated {
			if len(inode.IndexNodes[ix].Index) != 0 {
				updated, err = inode.updateIndexBetweenIndexes(ix) // Update the index between indexes
				if err != nil {
					return
				}
			}
			if len(inode.IndexNodes[ix].Index) == 0 {
				err = inode.borrowNodeSide(ix) // Will borrow part of the node (借结点).
				if err != nil {
					return
				}
			}
		}

		// Return the results of the deletion.
		return
	}

	// Check if there are any data nodes.
	if len(inode.DataNodes) > 0 {
		// Call the deleteBottomItem method on the current node as it is close to the bottom layer.
		// This signifies the beginning of deleting data. (接近资料层)  ‼️

		// Here, adjustments may be made to IX (IX 在这里可能会被修改) ‼️
		deleted, updated, ix = inode.deleteBottomItem(item)

		// The individual data node is now empty, and
		// it is necessary to start borrowing data from neighboring nodes.
		if len(inode.DataNodes[ix].Items) == 0 {
			var borrowed bool
			borrowed, err = inode.borrowFromBothSide(ix)
			if err != nil {
				return
			}
			if borrowed == true {
				updated = true
			}
			// if borrowed == false {} // If borrowing here is not possible, partial nodes will be borrowed later.
		}

		// If the data node becomes smaller, the index will be removed.
		if len(inode.DataNodes) <= 2 && len(inode.DataNodes[ix].Items) == 0 {
			inode.Index = []int64{}
		}
	}

	// Return the results of the deletion.
	return
}

// delete is a method of the BpIndex type that deletes the specified BpItem. (由右边删除 👉 ‼️)
func (inode *BpIndex) deleteToRight(item BpItem) (deleted, updated bool, ix int, err error) {
	// Check if there are any index nodes.
	if len(inode.IndexNodes) > 0 {
		// Use binary search to find the index (ix) where the key should be deleted.
		ix = sort.Search(len(inode.Index), func(i int) bool {
			return inode.Index[i] > item.Key // no equal sign ‼️ no equal sign means delete to the right ‼️
		})

		// Recursion keeps deletion in the right direction. 递归一直向右砍 ⬅️
		deleted, updated, _, err = inode.IndexNodes[ix].deleteToRight(item)

		// Immediately update the index of index node.
		if updated {
			if len(inode.IndexNodes[ix].Index) != 0 {
				updated, err = inode.updateIndexBetweenIndexes(ix) // Update the index between indexes.
				if err != nil {
					return
				}
			}
			if len(inode.IndexNodes[ix].Index) == 0 {
				err = inode.borrowNodeSide(ix) // Will borrow part of the node (借结点).
			}
		}

		// Return the results of the deletion.
		return
	}

	// Check if there are any data nodes.
	if len(inode.DataNodes) > 0 {
		// Call the deleteBottomItem method on the current node as it is close to the bottom layer.
		// This signifies the beginning of deleting data. (接近资料层)

		// Here, adjustments may be made to IX (IX 在这里可能会被修改) ‼️
		deleted, updated, ix = inode.deleteBottomItem(item)

		// The individual data node is now empty, and
		// it is necessary to start borrowing data from neighboring nodes.
		if len(inode.DataNodes[ix].Items) == 0 {
			var borrowed bool
			borrowed, err = inode.borrowFromBothSide(ix) // If you can borrow, you can maintain the integrity of the node."
			if err != nil {
				return
			}
			if borrowed == true {
				updated = true // At the same time, it also needs to be updated.
			}
			// if borrowed == false {} // If borrowing here is not possible, partial nodes will be borrowed later.
		}

		// If the data node becomes smaller, the index will be removed.
		if len(inode.DataNodes) <= 2 && len(inode.DataNodes[ix].Items) == 0 {
			inode.Index = []int64{}
		}
	}

	// Return the results of the deletion.
	return
}

// deleteBottomItem will remove data from the bottom layer.
// If the node is too small, it will clear the entire index.
func (inode *BpIndex) deleteBottomItem(item BpItem) (deleted, updated bool, ix int) {
	// Use binary search to find the index (ix) where the key should be inserted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Call the delete method on the corresponding DataNode to delete the item.
	deleted, _ = inode.DataNodes[ix]._delete(item)

	// The Bpdatdataode is too small to form an index.
	if deleted == true && len(inode.DataNodes) < 2 {
		inode.Index = []int64{} // Wipe out the whole index.
	}

	// Return the results of the deletion.
	return
}

// ➡️ The functions related to updated indexes.

// updateIndexBetweenIndexes is for updating non-bottom-level indices. (更新非底层的索引)
func (inode *BpIndex) updateIndexBetweenIndexes(ix int) (updated bool, err error) {
	if ix > 0 && // 條件1 ix 要大於 0
		len(inode.IndexNodes[ix].IndexNodes) >= 2 && // 條件2 下層索引節點數量要大於等於 2
		(inode.Index[ix-1] != inode.IndexNodes[ix].Index[0]) { // 條件3 和原索引不同

		// 進行更新
		inode.Index[ix-1] = inode.IndexNodes[ix].Index[0]
		updated = true
	}

	// Finally, perform the return.
	return
}

// ➡️ The functions related to borrowed data.

// borrowFromBothSide only borrows a portion of data from the neighboring nodes.
func (inode *BpIndex) borrowFromBothSide(ix int) (borrowed bool, err error) {
	// Not an empty node, no need to borrow
	if len(inode.DataNodes[ix].Items) != 0 {
		err = fmt.Errorf("not an empty node, do not need to borrow")
	}

	// Borrow from the left side first
	if (ix - 1) >= 0 { // Left neighbor
		length := len(inode.DataNodes[ix-1].Items)
		if length >= 2 { // Neighbor has enough data to borrow
			firstItems := inode.DataNodes[ix-1].Items[:(length - 1)]    // First part contains the first element
			borrowedItems := inode.DataNodes[ix-1].Items[(length - 1):] // Second part contains the remaining elements

			inode.DataNodes[ix-1].Items = firstItems
			inode.DataNodes[ix].Items = borrowedItems

			inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key

			borrowed = true
		}
	}

	// Borrow from the right side next
	/*if (ix + 1) <= len(inode.DataNodes[ix].Items) { // Right neighbor
		length := len(inode.DataNodes[ix+1].Items)
		if length >= 2 { // Neighbor has enough data to borrow
			borrowedItems := inode.DataNodes[ix+1].Items[:1] // First part contains the first element
			secondItems := inode.DataNodes[ix+1].Items[1:]   // Second part contains the remaining elements

			inode.DataNodes[ix].Items = borrowedItems
			inode.DataNodes[ix+1].Items = secondItems

			inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key
			inode.Index[ix] = inode.DataNodes[ix+1].Items[0].Key

			borrowed = true
		}
	}*/

	// Finally, return the result
	return
}

// borrowNodeSide will borrow more data from neighboring nodes, including indexes.
func (inode *BpIndex) borrowNodeSide(ix int) (err error) {
	// Anyway, as the index nodes keep shrinking, eventually leaving only two DataNodes,
	// one of which may have no data. So here, we check whether the number of DataNodes is 2.
	if len(inode.IndexNodes[ix].DataNodes) != 2 {
		err = fmt.Errorf("the index is still there; there is no need to borrow nodes")
		return
	}

	// When the index of inode is 1
	if len(inode.Index) == 1 {
		// Additional code to be written here
	}
	// When the index of inode is 2
	if len(inode.Index) == 2 {
		smaller := []int64{inode.Index[0]}
		bigger := []int64{inode.Index[1]}

		if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 {
			// Borrowing a portion of nodes to the right, including some data and index.

			// Adjusting indexes
			inode.Index = smaller
			inode.IndexNodes[ix].Index = bigger
			inode.Index = append(inode.Index, inode.IndexNodes[ix+1].Index[0])

			// loading out data
			inode.IndexNodes[ix+1].Index = inode.IndexNodes[ix+1].Index[1:]

			// Receiving data
			inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix].DataNodes[1].Next.Items[0])
			inode.IndexNodes[ix].DataNodes[1].Next.Items = inode.IndexNodes[ix].DataNodes[1].Next.Items[1:]
		} else if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 {
			// Borrowing a portion of nodes to the left, including some data and index.

			// Adjusting indexes
			inode.Index = bigger
			inode.IndexNodes[ix].Index = smaller
			indexLength := len(inode.IndexNodes[ix-1].Index)
			inode.Index = append([]int64{inode.IndexNodes[ix-1].Index[indexLength-1]}, inode.Index...)

			// loading out data
			inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index[:indexLength])

			// Receiving data
			nodeLength := len(inode.IndexNodes[ix].DataNodes[0].Previous.Items)
			inode.IndexNodes[ix].DataNodes[0].Items = append([]BpItem{inode.IndexNodes[ix].DataNodes[0].Previous.Items[nodeLength-1]}, inode.IndexNodes[ix].DataNodes[0].Items...)
			inode.IndexNodes[ix].DataNodes[0].Previous.Items = inode.IndexNodes[ix].DataNodes[0].Previous.Items[:nodeLength]
		}
	}

	// Finally, return
	return
}
