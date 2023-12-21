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

/*func (inode *BpIndex) mergeWithEmptyIndex(ix int) {
	//
	for i := 0; i < len(inode.IndexNodes); i++ {
		if len(inode.IndexNodes[i].Index) == 0 {
			if len(inode.IndexNodes[i].DataNodes) > 0 {

				// 如果 inode.IndexNodes[i].Index 長度為 0，不是索引節點為空，那就是資料節點為空

				if len(inode.IndexNodes[i].IndexNodes) > 0 {
					// 這裡是 IndexNode 有資料
					// (不可以向別的節點借資料)
					// 以後再處理
				} else if len(inode.IndexNodes[i].DataNodes) > 0 {

					if i == 0 {
						//
						fmt.Println()
					}

					// 看條件是否符合能向自己借
					if len(inode.IndexNodes[i].Index) == 0 && len(inode.IndexNodes[i].DataNodes[0].Items) > 1 {
						if err := inode.IndexNodes[i].splitAndDeleteSelf(); err != nil {
							return
						}
					}
				}
			}
		}
	}
}*/

// 當 Items 為空，刪除 DpData 的部份索引
/*func (inode *BpIndex) dropIndexIfdataNodeEmpty(ix int) {
	// 如果第一個和第二個 DataNode 為空，那第一個索引就刪除
	if (ix == 0 || ix == 1) && len(inode.DataNodes[ix].Items) == 0 {
		// 删除索引
		inode.Index = inode.Index[1:]

		// 重建連結
		if ix == 0 {
			inode.DataNodes = inode.DataNodes[1:]
			inode.DataNodes[1].Previous = nil
		}
		if ix == 1 {
			inode.DataNodes[0].Next = inode.DataNodes[1].Next
			inode.DataNodes[1].Next.Previous = inode.DataNodes[0]
			inode.DataNodes[1] = nil
			inode.DataNodes = append(inode.DataNodes[:1], inode.DataNodes[2:]...) // inode.DataNodes[:1], inode.DataNodes[1], inode.DataNodes[2:]... 位置 1 資料消失
		}
	} else {
		// 檢查第三個節點以後，是否為空的 BpData
		if len(inode.DataNodes[ix].Items) == 0 {
			// 删除索引
			copy(inode.Index[0:ix-1], inode.Index[ix:]) // 位置 i 的資料不見了

			// 重建連結
			inode.DataNodes[ix-1].Next = inode.DataNodes[ix+1]
			inode.DataNodes[ix+1].Previous = inode.DataNodes[ix-1]
		}
	}

	return
}*/

/*func (inode *BpIndex) mergeWithEmptyIndex2() {
	//
	for i := 0; i < len(inode.IndexNodes); i++ {
		if len(inode.IndexNodes[i].Index) == 0 {
			if len(inode.IndexNodes[i].DataNodes) > 0 {

				// 如果 inode.IndexNodes[i].Index 長度為 0，不是索引節點為空，那就是資料節點為空

				if len(inode.IndexNodes[i].IndexNodes) > 0 {
					// 這裡是 IndexNode 有資料
					// (不可以向別的節點借資料)
					// 以後再處理
				} else if len(inode.IndexNodes[i].DataNodes) > 0 {

					if i == 0 {
						//
						fmt.Println()
					}

					// 看條件是否符合能向自己借
					if len(inode.IndexNodes[i].Index) == 0 && len(inode.IndexNodes[i].DataNodes[0].Items) > 1 {
						if err := inode.IndexNodes[i].splitAndDeleteSelf(); err != nil {
							return
						}
					}
				}
			}
		}
	}
}*/

/*func (inode *BpIndex) splitAndDeleteSelf() (err error) {
	//
	firstItems := inode.DataNodes[0].Items[:1] // 第一份包含第一个元素
	otherItems := inode.DataNodes[0].Items[1:] // 第二份包含剩余的元素

	firstBpData := BpData{Items: firstItems}
	secondBpData := BpData{Items: otherItems}

	firstBpData.Previous = nil
	firstBpData.Next = &secondBpData

	secondBpData.Previous = &firstBpData
	secondBpData.Next = inode.DataNodes[0].Next

	secondBpData.Next.Previous = &secondBpData

	inode.Index = []int64{otherItems[0].Key}
	inode.DataNodes = []*BpData{&firstBpData, &secondBpData}

	return
}*/

// delete is a method of the BpIndex type that deletes the specified BpItem.
/*func (inode *BpIndex) deleteDeprecated(item BpItem) (deleted, updated bool, direction int, ix int, err error) {
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
}*/

// 准备考虑废除 mark 功能 🔥

// deleteBottomItem deletes the specified BpItem from the DataNodes near the bottom layer of the BpIndex.
/*func (inode *BpIndex) deleteBottomItemDeprecated(item BpItem) (deleted, updated bool, direction int, ix int, err error) {
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
}*/

/*func (inode *BpIndex) updateBottomIndex(ix int) (updated bool, err error) {
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
}*/

// CleanMark is used to maintain the BpIndex node to an appropriate size.
/*func (inode *BpIndex) CleanMark(ix int) (err error) {
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
}*/
