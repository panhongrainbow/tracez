package bpTree

import (
	"errors"
	"fmt"
	"sort"
)

// ➡️ The functions related to direction.

// The function delRoot is responsible for deleting an item from the root of the B+ tree.
func (inode *BpIndex) delRoot(item BpItem) (deleted, updated bool, ix int, err error) {
	// Check if the root node is empty and has only one data node with a matching key.
	if len(inode.Index) == 0 &&
		len(inode.DataNodes) == 1 {
		if inode.DataNodes[0].Items[0].Key == item.Key {
			// If the root node has only one data node and its key matches the target key, remove the root node.
			node := &BpIndex{
				DataNodes: make([]*BpData, 0, BpWidth+1), // The addition of 1 is because data chunks may temporarily exceed the width.
			}
			*inode = *node
			return
		}
	}

	// Call the delAndDir method to handle deletion and direction.
	deleted, updated, ix, err = inode.delAndDir(item)
	if err != nil {
		return
	}

	// If there's not much data in the root node and it has two data nodes, handle the cases.
	if len(inode.Index) == 0 &&
		len(inode.DataNodes) == 2 {
		// If the first data node is empty, replace the root node with the second data node.
		if len(inode.DataNodes[0].Items) == 0 {
			inode.Index = nil
			inode.DataNodes = []*BpData{inode.DataNodes[1]}
			return
		}
		// If the second data node is empty, replace the root node with the first data node.
		if len(inode.DataNodes[1].Items) == 0 {
			inode.Index = nil
			inode.DataNodes = []*BpData{inode.DataNodes[0]}
			return
		}
	}

	if len(inode.Index) == 0 && len(inode.IndexNodes) == 1 {
		*inode = *inode.IndexNodes[0]
	}
	if len(inode.Index) == 0 && len(inode.IndexNodes) > 0 {
		node := &BpIndex{}
		for i := 0; i < len(inode.IndexNodes); i++ {
			node.Index = append(node.Index, inode.IndexNodes[i].Index...)
			node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[i].IndexNodes...)
			node.DataNodes = append(node.DataNodes, inode.IndexNodes[i].DataNodes...)
		}

		*inode = *node
	}

	if len(inode.Index) == 0 && len(inode.DataNodes) > 0 {
		node := &BpIndex{}
		for i := 0; i < len(inode.IndexNodes); i++ {
			node.DataNodes = append(node.DataNodes, inode.IndexNodes[i].DataNodes...)
		}
		for i := 0; i < len(node.DataNodes); i++ {
			if i != 0 {
				node.Index = append(node.Index, node.DataNodes[i].Items[0].Key)
			}
		}

		*inode = *node
	}

	// Return the results
	return
}

// delAndDir performs data deletion based on automatic direction detection.
// 自动判断资料删除方向，其實會由不同方向進行刪除
func (inode *BpIndex) delAndDir(item BpItem) (deleted, updated bool, ix int, err error) {
	// Use binary search to find the index (ix) where the key should be deleted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // no equal sign ‼️ no equal sign means delete to the right ‼️
	})

	// Check if deletion should be performed by the leftmost node first.
	if (ix-1) >= 0 &&
		len(inode.IndexNodes)-1 >= (ix-1) && ix != 0 { // After the second index node, it's possible to borrow data from the left ⬅️ node
		// Length of the left node
		length := len(inode.IndexNodes[ix-1].Index)

		// If it is continuous data (same value) (5❌ - 5 - 5 - 5 - 5 - 6 - 7 - 8)
		if ix >= 1 && ix <= len(inode.IndexNodes)-1 && ix-1 >= 1 && ix-1 <= len(inode.IndexNodes)-1 && length >= 1 && inode.IndexNodes[ix].Index[0] == inode.IndexNodes[ix-1].Index[length-1] {
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
				updated, err = inode.borrowNodeSide(ix) // Will borrow part of the node (借结点).
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

		// Integrate the scattered nodes. (这段还没测试)
		if len(inode.DataNodes[ix].Items) == 0 && len(inode.Index) != 0 {
			// Rebuild connections.
			inode.DataNodes[ix].Previous.Next = inode.DataNodes[ix].Next
			inode.DataNodes[ix].Next.Previous = inode.DataNodes[ix].Previous

			// Reorganize nodes.
			inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
			inode.DataNodes = append(inode.DataNodes[:ix], inode.DataNodes[ix+1:]...)
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
			if len(inode.IndexNodes[ix].Index) == 0 &&
				len(inode.IndexNodes[ix].DataNodes) == 2 {
				updated, err = inode.borrowNodeSide(ix) // Will borrow part of the node (借结点).
				if err != nil && !errors.Is(err, fmt.Errorf("the index is still there; there is no need to borrow nodes")) {
					return
				}
			} else if len(inode.IndexNodes[ix].Index) == 0 && // 不是在资料节点，就是在索引节点，在这里要连起来
				len(inode.IndexNodes[ix].IndexNodes) != 0 &&
				len(inode.IndexNodes[ix].DataNodes) == 0 {
				updated, err = inode.indexesMove(ix)
				if err != nil {
					return
				}
			}
		}

		if item.Key == 282 {
			fmt.Println("由这里开发 ！")
		}

		if len(inode.Index) == 0 &&
			len(inode.IndexNodes) == 2 &&
			len(inode.IndexNodes[0].DataNodes) > 0 &&
			len(inode.IndexNodes[1].DataNodes) > 0 {
			if len(inode.IndexNodes[0].DataNodes) == 2 {
				if len(inode.IndexNodes[0].DataNodes[0].Items) == 0 {
					// 再编写
				} else if len(inode.IndexNodes[0].DataNodes[1].Items) == 0 {
					// 再编写
				}
			} else if len(inode.IndexNodes[1].DataNodes) == 2 {
				if len(inode.IndexNodes[1].DataNodes[0].Items) == 0 {
					inode.IndexNodes[0].Index = append(inode.IndexNodes[0].Index, inode.IndexNodes[1].Index[0])
					inode.IndexNodes[1].DataNodes[1].Previous = inode.IndexNodes[1].DataNodes[0].Previous
					inode.IndexNodes[0].DataNodes = append(inode.IndexNodes[0].DataNodes, inode.IndexNodes[1].DataNodes[1])
					inode.IndexNodes = []*BpIndex{inode.IndexNodes[0]}
					length := len(inode.IndexNodes[0].Index)
					if length >= BpWidth {
						var key int64
						var side *BpIndex
						key, side, err = inode.IndexNodes[0].splitWithDnode()
						inode.Index = []int64{key}
						inode.IndexNodes = append(inode.IndexNodes, side)
					}
				} else if len(inode.IndexNodes[1].DataNodes[1].Items) == 0 {
					// 再编写
				}
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
			borrowed, err = inode.borrowFromBothSide(ix) // If you can borrow, you can maintain the integrity of the node.
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
			updated = true
		}

		// Integrate the scattered nodes.
		if len(inode.DataNodes[ix].Items) == 0 && len(inode.Index) != 0 {
			// Rebuild links.
			if inode.DataNodes[ix].Previous == nil {
				// 第 1 个资料结点
				inode.DataNodes[ix].Next.Previous = nil
			} else if inode.DataNodes[ix].Next == nil {
				// 第 2 个资料结点
				inode.DataNodes[ix].Previous.Next = nil
			} else {
				inode.DataNodes[ix].Previous.Next = inode.DataNodes[ix].Next
				inode.DataNodes[ix].Next.Previous = inode.DataNodes[ix].Previous
			}

			// Reorganize nodes.
			if ix != 0 {
				inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
				inode.DataNodes = append(inode.DataNodes[:ix], inode.DataNodes[ix+1:]...)
			} else if ix == 0 {
				inode.Index = inode.Index[1:]
				inode.DataNodes = inode.DataNodes[1:]
			}
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
		updated = true
	}

	if deleted == true && ix > 0 && len(inode.DataNodes[ix].Items) > 0 {
		if inode.Index[ix-1] != inode.DataNodes[ix].Items[0].Key {
			inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key
			updated = true
		}
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

	// 以下先注解，因为无法对邻近节点的索引进行修改
	// Borrow from the right side next.
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
func (inode *BpIndex) borrowNodeSide(ix int) (updated bool, err error) {
	// 如果邻近节点资料很多，先拼左，再拼右
	if len(inode.IndexNodes[ix].Index) == 0 && inode.IndexNodes[ix].DataNodes != nil && len(inode.IndexNodes) == 2 {
		if ix-1 >= 0 && ix-1 <= len(inode.IndexNodes)-1 && len(inode.IndexNodes[ix-1].Index) >= 2 { // 可以向左借
			// 未开发
			fmt.Println()
		} else if ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1 && len(inode.IndexNodes[ix+1].Index) >= 2 {
			// 未开发
			fmt.Println()
			if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 {
				// 未开发
				fmt.Println()
			} else if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 {
				inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix+1].DataNodes[0].Items[0])
				inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

				inode.IndexNodes[ix+1].Index = inode.IndexNodes[ix+1].Index[1:]
				inode.IndexNodes[ix+1].DataNodes = inode.IndexNodes[ix+1].DataNodes[1:]

				inode.Index = []int64{inode.IndexNodes[ix+1].Index[0]}

				updated = true
				return
			}
		}
	}

	// Anyway, as the index nodes keep shrinking, eventually leaving only two DataNodes,
	// one of which may have no data. So here, we check whether the number of DataNodes is 2.
	if len(inode.IndexNodes[ix].DataNodes) != 2 {
		err = fmt.Errorf("the index is still there; there is no need to borrow nodes")
		return
	}

	// 在这里索引之间会移动
	if len(inode.Index) == 2 {
		// When the length of the neighbor node's index is 2, we perform operations related to borrowing nodes.
		smaller := []int64{inode.Index[0]}
		bigger := []int64{inode.Index[1]}

		if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 &&
			ix+1 <= len(inode.IndexNodes)-1 &&
			ix+1 >= 0 &&
			len(inode.IndexNodes[ix+1].Index) >= 2 {
			// Borrowing a portion of nodes to the right, including some data and index.

			// Adjusting indexes
			inode.Index = smaller
			inode.IndexNodes[ix].Index = bigger
			inode.Index = append(inode.Index, inode.IndexNodes[ix+1].Index[0])

			// loading out data
			inode.IndexNodes[ix+1].Index = inode.IndexNodes[ix+1].Index[1:]

			// Receiving data
			inode.IndexNodes[ix].DataNodes[1] = inode.IndexNodes[ix+1].DataNodes[0]
			inode.IndexNodes[ix+1].DataNodes = inode.IndexNodes[ix+1].DataNodes[1:]

			updated = true
		} else if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 &&
			ix-1 <= len(inode.IndexNodes)-1 &&
			ix-1 >= 0 &&
			len(inode.IndexNodes[ix-1].Index) >= 2 {
			// Borrowing a portion of nodes to the left, including some data and index.

			// Adjusting indexes
			inode.Index = bigger
			inode.IndexNodes[ix].Index = smaller
			indexLength := len(inode.IndexNodes[ix-1].Index)
			inode.Index = append([]int64{inode.IndexNodes[ix-1].Index[indexLength-1]}, inode.Index...)

			// loading out data
			inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index[:indexLength-1]) // 不含最后一个

			// Receiving data
			// nodeLength := len(inode.IndexNodes[ix].DataNodes)
			previousLength := len(inode.IndexNodes[ix-1].DataNodes)
			inode.IndexNodes[ix].DataNodes[0] = inode.IndexNodes[ix-1].DataNodes[previousLength-1]
			inode.IndexNodes[ix-1].DataNodes = inode.IndexNodes[ix-1].DataNodes[:previousLength-1]

			updated = true
		}

		// When the length of the neighbor node's index is 1, we perform operations related to merging nodes.
		if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 &&
			ix+1 <= len(inode.IndexNodes)-1 &&
			ix+1 >= 0 &&
			len(inode.IndexNodes[ix+1].Index) >= 1 {
			// 这里先不处理，因其他索引节点会变动 !
		} else if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 &&
			ix-1 <= len(inode.IndexNodes)-1 &&
			ix-1 >= 0 &&
			len(inode.IndexNodes[ix-1].Index) >= 1 {
			// 重建連結
			length := len(inode.IndexNodes[ix-1].DataNodes)
			inode.IndexNodes[ix-1].DataNodes[length-1].Next = inode.IndexNodes[ix].DataNodes[1]
			inode.IndexNodes[ix].DataNodes[1].Previous = inode.IndexNodes[ix-1].DataNodes[length-1]
			// 資料移動
			inode.IndexNodes[ix-1].DataNodes = append(inode.IndexNodes[ix-1].DataNodes, inode.IndexNodes[ix].DataNodes[1])
			inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].DataNodes[1].Items[0].Key) // 更新索引
			//
			inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
			//
			inode.Index = inode.Index[1:]
			//
			updated = true
		}
	}

	// 这里只是简单的进行简单的节点合拼，这里是在索引节点的内部
	length := len(inode.IndexNodes)
	if len(inode.Index) == 1 && ix >= 0 && ix <= length-1 && len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 && ix+1 <= len(inode.IndexNodes)-1 && ix+1 >= 0 && inode.IndexNodes[ix+1].DataNodes != nil {
		node := &BpIndex{
			Index:     []int64{inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key, inode.IndexNodes[ix+1].DataNodes[1].Items[0].Key},
			DataNodes: []*BpData{inode.IndexNodes[ix].DataNodes[0], inode.IndexNodes[ix+1].DataNodes[0], inode.IndexNodes[ix+1].DataNodes[1]},
		}
		node.DataNodes[0].Next = node.DataNodes[1]
		node.DataNodes[1].Previous = node.DataNodes[0]

		*inode = *node

		updated = true
	}

	// 开始进行层数缩
	length = len(inode.IndexNodes)
	if len(inode.IndexNodes) > 0 && ix >= 0 && ix <= length-1 && len(inode.IndexNodes[ix].Index) == 0 { // 索引失效
		// 下放索引
		// 在 ix 位罝上，IX 位置上的节点失效
		if ix == 0 { // 在第 1 个位置就直接抹除
			/*inode.Index = inode.Index[1:]
			inode.IndexNodes = inode.IndexNodes[1:]

			updated = true*/
		} else if ix != 0 {
			if len(inode.Index) == 1 { // 上层直接下放唯一的索引，上层索引直接为空
				inode.IndexNodes[ix].Index = []int64{inode.Index[0]}
				inode.Index = []int64{}

				updated = true
			} else if len(inode.Index) > 1 { // 上层直接下放其中一个索引，其他不变
				inode.IndexNodes[ix].Index = []int64{inode.Index[ix-1]}
				inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)

				updated = true
			}

			// 如果发生下放资料后，节点和索引的数量差距超过 1
			// 这时，只剩 2 个 DataNode
			if inode.IndexNodes[ix].DataNodes != nil &&
				len(inode.IndexNodes[ix].DataNodes) == 2 && (len(inode.IndexNodes)-len(inode.Index)) > 1 {
				if len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 &&
					inode.IndexNodes[ix].DataNodes[0].Items[0].Key > inode.IndexNodes[ix].Index[0] {
					inode.IndexNodes[ix].Index[0] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key
				}
				if len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 &&
					inode.IndexNodes[ix].DataNodes[1].Items[0].Key < inode.IndexNodes[ix].Index[0] {
					inode.IndexNodes[ix].Index[0] = inode.IndexNodes[ix].DataNodes[1].Items[0].Key
				}
				// 开始决定是向左，还是向右合拼
				var combined bool
				if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 &&
					ix-1 >= 0 && ix-1 <= len(inode.IndexNodes)-1 {
					inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index...)
					inode.IndexNodes[ix-1].DataNodes = append(inode.IndexNodes[ix-1].DataNodes, inode.IndexNodes[ix].DataNodes[1])
					combined = true
				} else if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 &&
					ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1 {
					inode.IndexNodes[ix+1].Index = append(inode.IndexNodes[ix].Index, inode.IndexNodes[ix+1].Index...)
					inode.IndexNodes[ix+1].DataNodes = append([]*BpData{inode.IndexNodes[ix].DataNodes[0]}, inode.IndexNodes[ix+1].DataNodes...)
					combined = true
				}
				// 如果还是没合拼
				if combined == false {
					if len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 {
						inode.IndexNodes[ix].Index[0] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key

						inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index...)
						inode.IndexNodes[ix-1].DataNodes = append(inode.IndexNodes[ix-1].DataNodes, inode.IndexNodes[ix].DataNodes[0])

						combined = true
					}
					if len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 {
						inode.IndexNodes[ix].Index[0] = inode.IndexNodes[ix].DataNodes[1].Items[0].Key

						inode.IndexNodes[ix+1].Index = append(inode.IndexNodes[ix].Index, inode.IndexNodes[ix+1].Index...)
						inode.IndexNodes[ix+1].DataNodes = append([]*BpData{inode.IndexNodes[ix].DataNodes[1]}, inode.IndexNodes[ix+1].DataNodes...)

						combined = true
					}
				}

				if combined == true {
					inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
				}
			}
		}
		if ix >= 0 && ix <= len(inode.IndexNodes)-1 &&
			len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 &&
			len(inode.IndexNodes[ix].DataNodes[1].Items) != 0 {
			if ix+1 >= 0 && len(inode.IndexNodes)-1 >= ix+1 {
				// 重建连结，在 ix 位置上的索引节点会有其中一个资料节点为空
				// 只有在 2 个资料节点，其中一个为空，就会索引失效
				if inode.IndexNodes[ix].DataNodes[0].Previous != nil {
					inode.IndexNodes[ix].DataNodes[0].Previous.Next = inode.IndexNodes[ix].DataNodes[0].Next
				}
				if inode.IndexNodes[ix].DataNodes[0].Next != nil {
					inode.IndexNodes[ix].DataNodes[0].Next.Previous = inode.IndexNodes[ix].DataNodes[0].Previous
				}
				// 合拼到右节点
				if len(inode.IndexNodes[ix].Index) == 0 {
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}
				}
				inode.IndexNodes[ix+1].Index = append([]int64{inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key}, inode.IndexNodes[ix+1].Index...) // 之前上层已经下放索引
				inode.IndexNodes[ix+1].DataNodes = append([]*BpData{inode.IndexNodes[ix].DataNodes[1]}, inode.IndexNodes[ix+1].DataNodes...)
				// 删除整个 ix 位置上的索引节点
				inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)

				inode.Index = []int64{}

				// *inode = *inode.IndexNodes[0] // 先不要
			}
		} else if ix >= 0 && ix <= len(inode.IndexNodes)-1 && len(inode.IndexNodes[ix].DataNodes[0].Items) != 0 &&
			len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 {
			if ix-1 >= 0 && len(inode.IndexNodes)-1 >= ix-1 {
				// 重建连结，在 ix 位置上的索引节点会有其中一个资料节点为空
				// 只有在 2 个资料节点，其中一个为空，就会索引失效
				if inode.IndexNodes[ix].DataNodes[1].Previous != nil {
					inode.IndexNodes[ix].DataNodes[1].Previous.Next = inode.IndexNodes[ix].DataNodes[1].Next
				}
				if inode.IndexNodes[ix].DataNodes[1].Next != nil {
					inode.IndexNodes[ix].DataNodes[1].Next.Previous = inode.IndexNodes[ix].DataNodes[1].Previous
				}
				// 合拼到左节点
				inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index[0])
				inode.IndexNodes[ix-1].DataNodes = append(inode.IndexNodes[ix-1].DataNodes, inode.IndexNodes[ix].DataNodes[0])
				// 删除整个 ix 位置上的索引节点
				inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
			}
		}
	}

	// Finally, return
	return
}

func (inode *BpIndex) indexesMove(ix int) (updated bool, err error) {
	// 底下有一个索引结点的索引为空，开始进行索引流动
	if len(inode.IndexNodes[ix].Index) == 0 {
		// 下放索引
		// 在 ix 位罝上，IX 位置上的节点失效

		if len(inode.Index) == 1 { // 上层直接下放唯一的索引，上层索引直接为空
			inode.IndexNodes[ix].Index = []int64{inode.Index[0]}
			inode.Index = []int64{}

			// 顶层索引消失，直接进行合拼
			node := &BpIndex{}
			node.Index = append(node.Index, inode.IndexNodes[0].Index...)
			node.Index = append(node.Index, inode.IndexNodes[1].Index...)

			node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[0].IndexNodes...)
			node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[1].IndexNodes...)

			// 最后储存改写
			*inode = *node

			updated = true
		} else if len(inode.Index) > 1 { // 上层直接下放其中一个索引，其他不变
			inode.IndexNodes[ix].Index = []int64{inode.Index[ix-1]}
			inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)

			updated = true
		}
	}
	return
}
