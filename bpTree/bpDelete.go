package bpTree

import (
	"errors"
	"fmt"
	"sort"
)

// ➡️ The functions related to direction.

// delFromRoot is responsible for deleting an item from the root of the B Plus tree. // 这是 B 加树的删除入口
func (inode *BpIndex) delFromRoot(item BpItem) (deleted, updated bool, ix int, err error) {
	// 这里根节点规模太小，根节点直接就是索引节点

	if len(inode.Index) == 0 &&
		len(inode.DataNodes) == 1 {
		// 以下用 inode.DataNodes 去寻找位置，这时 根结点资料过小，只剩下 单个资料节点 了

		// ▶️ 索引节点数量 0 🗂️ 资料节点数量 1 ⛷️ 层数数量 0

		// 搜寻 🔍
		ix = sort.Search(len(inode.Index), func(i int) bool {
			// 二分法直接在资料节点进行搜寻
			return inode.DataNodes[0].Items[i].Key > item.Key // no equal sign ‼️ no equal sign means delete to the right ‼️
		})

		// 删除 💢
		if inode.DataNodes[0].Items[ix].Key == item.Key {
			inode.DataNodes[0].Items = append(inode.DataNodes[0].Items[0:ix], inode.DataNodes[0].Items[ix+1:]...)
			deleted = true
			return
		}

		// 没删到时，就要立刻中止
	} else {

		// ❌ not ( ▶️ 索引节点数量 0 🗂️ 资料节点数量 1 ⛷️ 层数数量 0 )

		// Call the delAndDir method to handle deletion and direction.
		deleted, updated, ix, err = inode.delAndDir(item) // 在这里加入方向性
		if err != nil {
			return
		}
	}

	// Return the results.
	return
}

// delAndDir performs data deletion based on automatic direction detection.  // 这是 B 加树的方向性删除入口
// 自动判断资料删除方向，其實會由不同方向進行刪除

/*
 为何要先优先向左删除资料，因最左边的相同值被删除时，就会被后面相同时递补，比较不会更动到边界值 ✌️
*/

func (inode *BpIndex) delAndDir(item BpItem) (deleted, updated bool, ix int, err error) {
	// 搜寻 🔍 (最右边 ➡️)
	// Use binary search to find the index (ix) where the key should be deleted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // 一定要大于，所以会找到最右边 ‼️
	})

	// 决定 ↩️ 是否要向左
	// Check if deletion should be performed by the leftmost node first.
	if len(inode.Index) > 0 && len(inode.IndexNodes) > 0 &&
		(ix-1) >= 1 && len(inode.IndexNodes)-1 >= (ix-1) { // 如果当前节点的左边有邻居

		// If it is continuous data (same value) (5❌ - 5 - 5 - 5 - 5 - 6 - 7 - 8)
		length := len(inode.IndexNodes[ix-1].Index) // 为了左边邻居节点最后一个索引值
		if len(inode.IndexNodes) > 0 &&             // 预防 panic 的检查
			len(inode.IndexNodes[ix].Index) > 0 && len(inode.IndexNodes[ix-1].Index) > 0 && // 预防 panic 的检查
			length > 0 && inode.IndexNodes[ix].Index[0] == inode.IndexNodes[ix-1].Index[length-1] { // 最后决定，如果最接近的索引节点有相同的索引值 ‼️

			// 搜寻 🔍 (最左边 ⬅️) (一切重来，重头开始向左搜寻)
			deleted, updated, ix, err = inode.deleteToLeft(item) // Delete to the leftmost node ‼️ (向左砍)

			// 中断了，不再考虑向右搜寻 ⚠️
			return
		}
	}

	// 搜寻 🔍 (最右边 ➡️)
	// If it is discontinuous data (different values) (5 - 5 - 5 - 5 - 5❌ - 6 - 7 - 8)
	deleted, updated, _, _, ix, err = inode.deleteToRight(item) // Delete to the rightmost node ‼️ (向右砍)

	// Return the results.
	return
}

// delete is a method of the BpIndex type that deletes the specified BpItem. (5 - 5 - 5 - 5 - 5❌ - 6 - 7 - 8)
// deleteToRight 先放前面，因为 deleteToLeft 会抄 deleteToRight 的内容
func (inode *BpIndex) deleteToRight(item BpItem) (deleted, updated bool, edgeValue int64, status int, ix int, err error) {
	// 设定初始值
	if status == 0 {
		status = edgeValueInit // 初始状态
	}
	if edgeValue == 0 {
		edgeValue = -1 // 边界的初始值
	}

	// 🖍️ for index node 针对索引节点

	// 搜寻 🔍 (最右边 ➡️)
	// Use binary search to find the index (ix) where the key should be deleted.
	if len(inode.IndexNodes) > 0 {
		ix = sort.Search(len(inode.Index), func(i int) bool {
			return inode.Index[i] > item.Key // 一定要大于，所以会找到最右边 ‼️
		})

		// 搜寻 🔍 (最右边 ➡️)
		// If it is discontinuous data (different values) (5 - 5 - 5 - 5 - 5❌ - 6 - 7 - 8)
		deleted, updated, edgeValue, status, _, err = inode.IndexNodes[ix].deleteToRight(item)

		// 🖐️ 状态变化 [LeaveBottom] -> Any
		if status == edgeValueLeaveBottom {

			// 🖐️ 状态变化 [LeaveBottom] -> [Init]
			if ix-1 >= 0 {
				fmt.Print("🏴‍☠️ 索引(1) ", inode.Index, "->", "位置", ix-1, "修改成", edgeValue, "->")
				inode.Index[ix-1] = edgeValue
				fmt.Print("最后变成", inode.Index, "\n")
			}

			status = edgeValueInit // 重置状态
		} else if status == statusBorrowFromIndexNode {
			// 当一个分支只剩一个索引值和一个索引节点，准备要向左合拼
			if ix-1 >= 0 && ix-1 <= len(inode.IndexNodes)-1 {
				if len(inode.IndexNodes[ix-1].Index)+1 < BpWidth { // 没错，Degree 是针对 Index
					inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index...)
					inode.IndexNodes[ix-1].IndexNodes = append(inode.IndexNodes[ix-1].IndexNodes, inode.IndexNodes[ix].IndexNodes...)
					inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
					inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
					// 合拼后，ix 的值要减 1
					status = statusIXMunus
					ix = ix - 1

					fmt.Println("这里程式还没写完1-1")

					return
				} else if len(inode.IndexNodes[ix-1].Index)+1 >= BpWidth {
					// if len(inode.IndexNodes) >= 2 { // 这里要检合拼后，多个节点层数是否相同 ⁉️
					// 后来想想，这里直接去除，因为加1后除2也会维持 Degree，只要层数相同就好
					inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index...)
					inode.IndexNodes[ix-1].IndexNodes = append(inode.IndexNodes[ix-1].IndexNodes, inode.IndexNodes[ix].IndexNodes...)
					inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
					inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)

					var middle *BpIndex

					// 中断检验
					if item.Key == 123 {
						fmt.Print()
					}

					var tailNode []*BpIndex
					// if ix >= 0 && ix <= len(inode.IndexNodes)-1 {
					tailNode = append(tailNode, inode.IndexNodes[ix:]...)
					// }

					// 要分成单偶数函式处理
					if len(inode.IndexNodes[ix-1].Index)%2 == 1 { // 单数
						// 当索引为奇数时
						middle, err = inode.IndexNodes[ix-1].protrudeInOddBpWidth()
						if err != nil {
							return
						}
						// inode.IndexNodes[ix-1] = middle // 这个错误，会造成层数不相批配
					} else if len(inode.IndexNodes[ix-1].Index)%2 == 0 { // 偶数
						// 当索引为偶数时
						middle, err = inode.IndexNodes[ix-1].protrudeInEvenBpWidth()
						if err != nil {
							return
						}
						// inode.IndexNodes[ix-1] = middle // 这个错误，会造成层数不相批配
					}

					// 在这里要整个嵌入原索引节点
					inode.Index = append([]int64{middle.Index[0]}, inode.Index...) // 嵌入索引

					inode.IndexNodes = append(inode.IndexNodes[:ix-1], middle.IndexNodes...)
					inode.IndexNodes = append(inode.IndexNodes, tailNode...)

					// 中断检验
					if item.Key == 123 {
						fmt.Print()
					}

					fmt.Println("这里程式还没写完1-2")

					status = 0

					return

					// 合拼后，ix 的值要减 1 (不会有这状况)
					// status = statusIXMunus
					// ix = ix - 1
				}
				fmt.Println("这里程式还没写完1-3")
				// }
			} else if ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1 {
				// 不能合拼后再合拼，会出事，所以用 else if，只做一次 ‼️
				if len(inode.IndexNodes[ix+1].Index)+1 < BpWidth {

					// 中断检验
					if item.Key == 123 {
						fmt.Print()
					}

					inode.IndexNodes[ix].Index = append([]int64{inode.IndexNodes[ix+1].edgeValue()}, inode.IndexNodes[ix+1].Index...)
					inode.IndexNodes[ix].IndexNodes = append(inode.IndexNodes[ix].IndexNodes, inode.IndexNodes[ix+1].IndexNodes...)
					inode.Index = append(inode.Index[:ix], inode.Index[ix+1:]...)
					inode.IndexNodes = append(inode.IndexNodes[:ix+1], inode.IndexNodes[ix+2:]...)

					var middle *BpIndex

					// 要分成单偶数函式处理
					if len(inode.Index) != 0 && len(inode.IndexNodes[ix].Index)%2 == 1 { // 单数
						// 当索引为奇数时
						middle, err = inode.IndexNodes[ix].protrudeInOddBpWidth() // 🖐️ for arrangement 针对重整结构
						if err != nil {
							return
						}

						// 在这里要整个嵌入原索引节点
						inode.IndexNodes[ix] = middle
					} else if len(inode.Index) != 0 && len(inode.IndexNodes[ix].Index)%2 == 0 { // 偶数
						// 当索引为偶数时
						middle, err = inode.IndexNodes[ix].protrudeInEvenBpWidth() // 🖐️ for index node 针对重整结构
						if err != nil {
							return
						}

						// 在这里要整个嵌入原索引节点
						inode.IndexNodes[ix] = middle

						// inode.IndexNodes[ix-1] = middle // 这个错误，会造成层数不相批配
					}

					fmt.Println("这里程式还没写完2")
				} else if len(inode.IndexNodes[ix+1].Index)+1 >= BpWidth {
					fmt.Println("这里程式还没写完3")
				}
			}
		} else {
			// 🖐️ 其他状态变化写在这里
		}

		// If the index at position ix becomes invalid. ‼️
		// 删除导致锁引失效 ‼️
		if len(inode.IndexNodes[ix].Index) == 0 { // invalid ❌
			if len(inode.IndexNodes[ix].DataNodes) >= 2 { // DataNode 🗂️

				fmt.Print("borrowFromIndexNode 执行前后，🏴‍☠️ 边界值变化 ", inode.edgeValue()) // 显示边界值

				var borrowed bool
				borrowed, err = inode.borrowFromIndexNode(ix) // Will borrow part of the node (借结点). ‼️  // 🖐️ for index node 针对索引节点
				// 看看有没有向索引节点借到资料

				fmt.Println(" -> ", inode.edgeValue()) // 显示边界值

				if err != nil && !errors.Is(err, fmt.Errorf("the index is still there; there is no need to borrow nodes")) {
					return
				}

				if borrowed == true { // 当向其他索引节点借完后，在执行 borrowFromIndexNode，重新计算边界值

					if len(inode.IndexNodes) > 0 && // 预防性检查
						len(inode.IndexNodes[0].DataNodes) > 0 && // 预防性检查
						len(inode.IndexNodes[0].DataNodes[0].Items) > 0 { // 预防性检查

						edgeValue = inode.IndexNodes[0].DataNodes[0].Items[0].Key // 边界值是由 索引节点中取出，所以可以直接把边界值放入 索引  ‼️‼️

						if edgeValue != -1 && len(inode.Index) == 0 { // 如果有正确取得 边界值 后
							inode.Index = []int64{edgeValue}
							status = statusBorrowFromIndexNode
						}
					}

					// 更新其他位置的边界值
					if ix >= 0 && ix <= len(inode.IndexNodes)-1 && // 预防性检查
						ix-1 >= 0 && ix-1 <= len(inode.Index)-1 && // 预防性检查
						len(inode.IndexNodes[ix].DataNodes) > 0 { // 预防性检查

						// 这里不会有 ix 为 0 的状况

						edgeValue = inode.IndexNodes[ix].DataNodes[0].Items[0].Key

						fmt.Print("🏴‍☠️ 索引(2) ", inode.Index, "->", "位置", ix-1, "修改成", edgeValue, "->")

						inode.Index[ix-1] = edgeValue

						fmt.Print("最后变成", inode.Index, "\n")
					}

					return
				}
			} else if len(inode.IndexNodes[ix].IndexNodes) != 0 && // IndexNode ▶️
				len(inode.IndexNodes[ix].DataNodes) == 0 {
				fmt.Println("注意 ‼️，有早期 indexMove 的机制需要评估")
				updated, err = inode.indexMove(ix) // Reorganize the indexing between nodes. (更新索引)
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
		// This signifies the beginning of deleting data. (接近资料层)

		// Here, adjustments may be made to IX (IX 在这里可能会被修改) ‼️
		// var edgeValue int64
		deleted, updated, ix, edgeValue, status = inode.deleteBottomItem(item) // 🖐️ for data node 针对资料节点
		if ix == 0 && status == edgeValuePassBottom {                          // 当 ix 为 0 时，才要处理边界值的问题
			status = edgeValueLeaveBottom
		}

		// The individual data node is now empty, and
		// it is necessary to start borrowing data from neighboring nodes.
		if len(inode.DataNodes[ix].Items) == 0 {
			updated, err = inode.borrowFromDataNode(ix) // Will borrow part of the data node. (向资料节点借资料)

			// 这段删除，程式会不稳定
			if updated == true && len(inode.DataNodes) >= 2 &&
				ix >= 0 && ix <= len(inode.DataNodes)-1 &&
				ix-1 >= 0 && ix-1 <= len(inode.DataNodes)-1 &&
				len(inode.DataNodes[ix].Items) > 0 &&
				inode.Index[ix-1] != inode.DataNodes[ix].Items[0].Key {
				// fmt.Println("计算边界值 2", inode.Index[ix-1], "->", inode.DataNodes[ix].Items[0].Key)
				inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key

				status = 0 // 抹除
				return
			}

			if updated == true || err != nil {
				// Leave as soon as you've borrowed the information.

				status = 0 // 抹除
				return
			}

			// If the data node cannot be borrowed, then information should be borrowed from the index node later.
			// 资料节点借不到，之后向索引节点借

			// During the deletion process, the node's index may become invalid.
			if len(inode.DataNodes) <= 2 {
				inode.Index = []int64{}

				// Return status
				updated = true

				status = 0 // 抹除
				return
			}

			// Wipe out the empty data node at the specified 'ix' position directly.
			if len(inode.Index) != 0 {
				// Rebuild the connections between data nodes.
				if inode.DataNodes[ix].Previous == nil {
					inode.DataNodes[ix].Next.Previous = nil

					status = 0 // 抹除
				} else if inode.DataNodes[ix].Next == nil {
					inode.DataNodes[ix].Previous.Next = nil

					status = 0 // 抹除
				} else {
					inode.DataNodes[ix].Previous.Next = inode.DataNodes[ix].Next
					inode.DataNodes[ix].Next.Previous = inode.DataNodes[ix].Previous

					status = 0 // 抹除
				}

				// Reorganize nodes.
				if ix != 0 {
					inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)             // Erase the position of ix - 1.
					inode.DataNodes = append(inode.DataNodes[:ix], inode.DataNodes[ix+1:]...) // Erase the position of ix.

					status = 0 // 抹除
				} else if ix == 0 { // Conditions have already been established earlier, with the index length not equal to 0. ‼️
					inode.Index = inode.Index[1:]
					inode.DataNodes = inode.DataNodes[1:]

					status = 0 // 抹除
				}
			}
		}

	}

	// Return the results of the deletion.
	return
}

// deleteToLeft is a method of the BpIndex type that deletes the leftmost specified BpItem. (由左边删除 👈 ‼️)
func (inode *BpIndex) deleteToLeft(item BpItem) (deleted, updated bool, ix int, err error) {
	// ⬇️⬇️⬇️ for index node 针对索引节点

	// Check if there are any index nodes.
	if len(inode.IndexNodes) > 0 {
		// Use binary search to find the index (ix) where the key should be deleted.
		ix = sort.Search(len(inode.Index), func(i int) bool {
			return inode.Index[i] >= item.Key // equal sign ‼️ no equal sign means delete to the left ‼️
			// (符合条件就停)
		})

		// Recursion keeps deletion in the left direction. 递归一直向左砍 ⬅️
		deleted, updated, _, err = inode.IndexNodes[ix].deleteToLeft(item)

		// Immediately update the index of index node.
		if updated && len(inode.IndexNodes[ix].Index) == 0 {
			updated, err = inode.borrowFromIndexNode(ix) // Will borrow part of the index node (向索引节点借资料).
			if err != nil {
				return
			}
		}

		// Return the results of the deletion.
		return
	}

	// ⬇️⬇️⬇️ for data node 针对资料节点

	// Check if there are any data nodes.
	if len(inode.DataNodes) > 0 {
		// Call the deleteBottomItem method on the current node as it is close to the bottom layer.
		// This signifies the beginning of deleting data. (接近资料层) ‼️

		// Here, this is very close to the data, just one index away. (和真实资料只隔一个索引) ‼️
		deleted, updated, ix, _, _ = inode.deleteBottomItem(item)

		// The individual data node is now empty, and
		// it is necessary to start borrowing data from neighboring nodes.
		if len(inode.DataNodes[ix].Items) == 0 {
			updated, err = inode.borrowFromDataNode(ix) // Will borrow part of the data node. (向资料节点借资料)
			// If update is true, it means that data has been borrowed from the adjacent information node. ‼️
			// 如果 update 为 true，那就代表有向邻近的资料节点借到资料 ‼️
			if updated == true || err != nil {
				// Leave as soon as you've borrowed the information.
				return
			}

			// If the data node cannot be borrowed, then information should be borrowed from the index node later.
			// 资料节点借不到，之后向索引节点借

			// During the deletion process, the node's index may become invalid.
			if len(inode.DataNodes) <= 2 {
				inode.Index = []int64{}

				// Return status
				updated = true
				return
			}

			// Wipe out the empty data node at the specified 'ix' position directly.
			if len(inode.Index) != 0 {
				// Recreate links.
				inode.DataNodes[ix].Previous.Next = inode.DataNodes[ix].Next
				inode.DataNodes[ix].Next.Previous = inode.DataNodes[ix].Previous

				// Reorganize nodes.
				inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
				inode.DataNodes = append(inode.DataNodes[:ix], inode.DataNodes[ix+1:]...)

				// Return status
				updated = true
				return
			}
		}
	}

	// Return the results of the deletion.
	return
}

// deleteBottomItem will remove data from the bottom layer. (只隔一个索引 ‼️)
// If the node is too small, it will clear the entire index. (索引可能失效‼️)
// 一层 BpData 资料层，加上一个索引切片，就是一个 Bottom
func (inode *BpIndex) deleteBottomItem(item BpItem) (deleted, updated bool, ix int, edgeValue int64, status int) {
	// 初始化回传值
	edgeValue = -1

	// Use binary search to find the index (ix) where the key should be inserted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Call the delete method on the corresponding DataNode to delete the item.
	deleted, _, edgeValue, status = inode.DataNodes[ix]._delete(item)
	// _delete 函式状况会回传 (1) 边界值没改变 (2) 边界值已改变 (3) 边界值为空
	if status == edgeValueChanges { // (1) 边界值已改变
		status = edgeValuePassBottom // 要通知上传的递归函式，边界值已改变
	}

	if deleted == true { // 如果资料真的删除的反应
		// The BpDatda node is too small then the index is invalid.
		if len(inode.DataNodes) < 2 {
			fmt.Println("这里注意，我觉得用到的机会不多 !")
			inode.Index = []int64{} // Wipe out the whole index. (索引在此失效) ‼️
			// 索引失效也是一种状态的表达方式，当索引为空时，这将再也不是结点了

			// Return status
			updated = true
			return
		} else if len(inode.DataNodes[ix].Items) > 0 && ix > 0 && // 预防性检查
			inode.Index[ix-1] != inode.DataNodes[ix].Items[0].Key { // 检查索引是不是有变化

			// Updating within the data node is considered safer, preventing damage in the entire B plus tree index.
			// 在资料节点内更新应是比较安全，不会造成整个 B 加树的索引错乱

			fmt.Print("🏴‍☠️ 索引(3) ", inode.Index, "->", "位置", ix-1, "修改成", inode.DataNodes[ix].Items[0].Key, "->")

			inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key // Immediately update the index

			fmt.Print("最后变成", inode.Index, "\n")

			// Return status
			updated = true
			return
		}
	}

	// Return the results of the deletion.
	return
}

// ➡️ The following function will make detailed adjustments for the B Plus tree.

// borrowFromDataNode only borrows a portion of data from the neighboring nodes.
func (inode *BpIndex) borrowFromDataNode(ix int) (borrowed bool, err error) {
	// No data borrowing is necessary as long as the node is not empty, since all indices are still in their normal state.
	if len(inode.DataNodes[ix].Items) != 0 {
		err = fmt.Errorf("not an empty node, do not need to borrow")
		return
	}

	// Borrow from the left side first
	if (ix - 1) >= 0 { // Left neighbor exists ‼️
		length := len(inode.DataNodes[ix-1].Items)
		if length >= 2 { // The left neighbor node has enough data to borrow
			// ⬇️ The left neighbor node is split.
			firstItems := inode.DataNodes[ix-1].Items[:(length - 1)]    // First part contains the first element
			borrowedItems := inode.DataNodes[ix-1].Items[(length - 1):] // Second part contains the remaining elements

			// ⬇️ Data reassignment
			inode.DataNodes[ix-1].Items = firstItems
			inode.DataNodes[ix].Items = borrowedItems

			// ⬇️ Index reassignment

			// This counts as a safe index update, within the internal structure of the DataNode itself. ✔️
			// 在 DataNode 内部更新索引算安全 ✔️
			inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key

			// ⬇️ Return status
			borrowed = true
			return
		}
	}

	// Borrow from the right side next.
	if (ix + 1) <= len(inode.DataNodes)-1 { // Right neighbor exists ‼️
		length := len(inode.DataNodes[ix+1].Items)
		if length >= 2 { // The right neighbor node has enough data to borrow
			// ⬇️ The right neighbor node is split.
			borrowedItems := inode.DataNodes[ix+1].Items[:1] // First part contains the first element
			secondItems := inode.DataNodes[ix+1].Items[1:]   // Second part contains the remaining elements

			// ⬇️ Data reassignment
			inode.DataNodes[ix].Items = borrowedItems
			inode.DataNodes[ix+1].Items = secondItems

			// ⬇️ Index reassignment
			if ix != 0 {
				// 最左边的 dataNode 不会产生索引
				inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key
			}

			// other conditions
			inode.Index[ix] = inode.DataNodes[ix+1].Items[0].Key

			// ⬇️ Return status
			borrowed = true
			return
		}
	}

	// Finally, return the result
	return
}

// indexMove performs index movement operations.
func (inode *BpIndex) indexMove(ix int) (updated bool, err error) {
	// If the index of a child node is empty, start index movement and push it down.
	if len(inode.IndexNodes[ix].Index) == 0 {
		if len(inode.Index) == 1 {
			// ⬇️ Scenario 1: Directly push down the only index from the upper level, making the upper-level index empty.
			inode.IndexNodes[ix].Index = []int64{inode.Index[0]}
			inode.Index = []int64{}

			// The top-level index disappears, create a new node for direct merging.
			node := &BpIndex{}

			// Merge indices
			node.Index = append(node.Index, inode.IndexNodes[0].Index...)
			node.Index = append(node.Index, inode.IndexNodes[1].Index...)

			// Merge indices
			node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[0].IndexNodes...)
			node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[1].IndexNodes...)

			// Save the modification at the end
			*inode = *node

			updated = true
		} else if len(inode.Index) > 1 && ix > 0 {
			// ⬇️ Scenario 2: Directly push down one index from the upper level, leaving others unchanged.
			inode.IndexNodes[ix].Index = []int64{inode.Index[ix-1]}
			inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)

			// Return status
			updated = true
		}
	}
	return
}

// borrowFromIndexNode will borrow more data from neighboring index nodes, including indexes.
func (inode *BpIndex) borrowFromIndexNode(ix int) (borrowed bool, err error) {
	// ⬇️ Check if there is an opportunity to borrow data from the index node. Data node with invalid index has neighbors.
	// (索引失效的资料节点 有邻居)
	if len(inode.IndexNodes[ix].Index) == 0 && // The underlying index is invalid; repair is required.
		inode.IndexNodes[ix].DataNodes != nil && // This is an issue that the index node needs to address.
		len(inode.IndexNodes) >= 2 { // There are multiple neighboring index nodes that can share data. 空资料节点有邻居 // (这是所有的状况要遵守的条件)
		// (先向右边借，因右边资料比较多)
		if (ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1) &&
			len(inode.IndexNodes[ix+1].DataNodes) >= 2 { // 邻居资料结点资料够多，可向右借; 当有 ix+1 时，不是 [状况3] 就是 [状况4] // (这是状况3和状况4要遵守的)
			// ➡️ Check if there is a chance to borrow data to the right.

			if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 { // 由 [狀況3] 發生，要先形成中间有空
				// 🔴 Case 3 Operation

				// 先向同一个 [索引节点] 下的 [资料节点] 借资料
				inode.IndexNodes[ix].DataNodes[0].Items = append(inode.IndexNodes[ix].DataNodes[0].Items, inode.IndexNodes[ix].DataNodes[1].Items[0])
				inode.IndexNodes[ix].DataNodes[1].Items = inode.IndexNodes[ix].DataNodes[1].Items[1:]

				// 如果能更新索引就进行更新
				if len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 {
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}
					return
				}
			}

			if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 { // 执行完后有可能由 [状况3] 变成 [状况4] 的状态，中间变成空的

				// 🔴 Case 4 Operation

				if len(inode.IndexNodes[ix+1].DataNodes[0].Items) >= 2 { // 如果最邻近的资料结点也有足够的资料，这时不会破坏邻近节点，进入 [状况4-1]，最好的状况
					// 🔴 Case 4-1 Operation

					// 先不让 资料 为空
					inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix+1].DataNodes[0].Items[0])
					inode.IndexNodes[ix+1].DataNodes[0].Items = inode.IndexNodes[ix+1].DataNodes[0].Items[1:]

					// 正常更新索引
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// 更新状态
					borrowed = true
					return
				} else if len(inode.IndexNodes[ix+1].DataNodes[0].Items) == 1 && len(inode.IndexNodes[ix+1].DataNodes) >= 3 { // 如果最邻近的资料结点没有足够的资料，这一借，邻居节点将会破坏，进入 [状况4-2]
					// 三个被抢一个，还有 2 个，不会对树的结构进行破坏 ✌️

					// 🔴 Case 4-2 Operation

					// 先不让 资料 为空
					inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix+1].DataNodes[0].Items[0])

					// 再 锁引 不能为空
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// 重建连结
					inode.IndexNodes[ix+1].DataNodes[1].Previous = inode.IndexNodes[ix+1].DataNodes[0].Previous
					inode.IndexNodes[ix].DataNodes[1].Next = inode.IndexNodes[ix+1].DataNodes[1]

					// 唯一值被取走，被破坏了，清空无效索引和资料节点
					inode.IndexNodes[ix+1].Index = inode.IndexNodes[ix+1].Index[1:]         // 都各退一个
					inode.IndexNodes[ix+1].DataNodes = inode.IndexNodes[ix+1].DataNodes[1:] // 都各退一个

					// ☢️ 更改上层索引，应可以，因这里接近底层资料
					inode.Index[(ix+1)-1] = inode.IndexNodes[(ix + 1)].DataNodes[0].Items[0].Key

					// 更新状态
					borrowed = true
					return
				} else if len(inode.IndexNodes[ix+1].DataNodes[0].Items) == 1 && len(inode.IndexNodes[ix+1].DataNodes) == 2 { // 邻点太小，将会被合拼，进入 [状况4-3]
					// 🔴 Case 4-3 Operation

					// 重建连结
					inode.IndexNodes[ix+1].DataNodes[0].Previous = inode.IndexNodes[ix].DataNodes[0]
					inode.IndexNodes[ix].DataNodes[0].Next = inode.IndexNodes[ix+1].DataNodes[0]

					// 不用借了，先直接合拼
					inode.IndexNodes[ix+1].Index = append([]int64{inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key}, inode.IndexNodes[ix+1].Index...)
					inode.IndexNodes[ix+1].DataNodes = append([]*BpData{inode.IndexNodes[ix].DataNodes[0]}, inode.IndexNodes[ix+1].DataNodes...)

					// 抹除 ix 位置
					if ix > 0 {
						inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
						inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
					} else if ix == 0 {
						inode.Index = inode.Index[1:]
						inode.IndexNodes = inode.IndexNodes[1:]
					}

					// 更新状态
					borrowed = true
					return
				} else if len(inode.IndexNodes[ix+1].DataNodes[0].Items) == 0 {
					err = fmt.Errorf("节点未及时整理完成1")
					return
				}
			}
		} else if (ix-1 >= 0 && ix-1 <= len(inode.IndexNodes)-1) &&
			len(inode.IndexNodes[ix-1].DataNodes) >= 2 { // 邻居资料结点资料够多，可向左借; 当有 ix-1 时，不是 [状况1] 就是 [状况2] // (这是状况1和状况2要遵守的)
			// ⬅️ Check if there is a chance to borrow data to the left.

			// (再向左边借)
			if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 { // 由 [狀況2] 發生，要先形成中间有空
				// 🔴 Case 2 Operation

				// 先向同一个 [索引节点] 下的 [资料节点] 借资料
				length0 := len(inode.IndexNodes[ix].DataNodes[0].Items)
				inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix].DataNodes[0].Items[length0-1])
				inode.IndexNodes[ix].DataNodes[0].Items = inode.IndexNodes[ix].DataNodes[0].Items[:length0-1] // 不包含最后一个

				// 如果能更新索引就进行更新
				if len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 {
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}
					return
				}
			}

			if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 && ix != 0 { // 执行完后有可能由 [状况2] 变成 [状况1] 的状态，中间变成空的

				// 先由出尾端的位置
				length0 := len(inode.IndexNodes[ix-1].DataNodes)
				length1 := len(inode.IndexNodes[ix-1].DataNodes[length0-1].Items)
				length2 := len(inode.IndexNodes[ix-1].DataNodes)

				// 🔴 Case 1 Operation
				if len(inode.IndexNodes[ix-1].DataNodes[length0-1].Items) >= 2 && length0 > 0 && length1 > 0 { // 如果最邻近的资料结点也有足够的资料，这时不会破坏邻近节点，进入 [状况4-1]，最好的状况
					// 🔴 Case 1-1 Operation

					// 先不让 资料 为空，再 锁引 不能为空
					inode.IndexNodes[ix].DataNodes[0].Items = append(inode.IndexNodes[ix].DataNodes[0].Items, inode.IndexNodes[ix-1].DataNodes[length0-1].Items[length1-1])
					inode.IndexNodes[ix-1].DataNodes[length0-1].Items = inode.IndexNodes[ix-1].DataNodes[length0-1].Items[:(length1 - 1)]

					// 正常更新索引
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// 更新状态
					borrowed = true
					return
				} else if len(inode.IndexNodes[ix-1].DataNodes[length0-1].Items) == 1 && len(inode.IndexNodes[ix-1].DataNodes) >= 3 && length0 > 0 && length1 > 0 { // 如果最邻近的资料结点没有足够的资料，这一借，邻居节点将会破坏，进入 [状况1-2]
					// 三个被抢一个，还有 2 个，不会对树的结构进行破坏 ✌️

					// 🔴 Case 1-2 Operation

					// 先不让 资料 为空，再 锁引 不能为空
					inode.IndexNodes[ix].DataNodes[0].Items = append(inode.IndexNodes[ix].DataNodes[0].Items, inode.IndexNodes[ix-1].DataNodes[length0-1].Items[length1-1])

					// 再 锁引 不能为空
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// 重建连结
					/*inode.IndexNodes[ix+1].DataNodes[length0-2].Next = inode.IndexNodes[ix+1].DataNodes[length0-1].Next
					inode.IndexNodes[ix].DataNodes[0].Previous = inode.IndexNodes[ix+1].DataNodes[length0-2]*/

					// 唯一值被取走，被破坏了，清空无效索引和资料节点
					inode.IndexNodes[ix-1].Index = inode.IndexNodes[ix-1].Index[:(length2 - 2)]
					inode.IndexNodes[ix-1].DataNodes = inode.IndexNodes[ix-1].DataNodes[:(length2 - 1)]

					// ☢️ 更改上层索引，应可以，因这里接近底层资料
					inode.Index[(ix)-1] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key

					// 更新状态
					borrowed = true
					return
				} else if len(inode.IndexNodes[ix-1].DataNodes[length0-1].Items) == 1 && len(inode.IndexNodes[ix-1].DataNodes) == 2 && length0 > 0 { // 邻点太小，将会被合拼，进入 [状况1-3]
					// 🔴 Case 1-3 Operation

					// 重建连结
					inode.IndexNodes[ix-1].DataNodes[length0-1].Next = inode.IndexNodes[ix].DataNodes[1]
					inode.IndexNodes[ix].DataNodes[1].Previous = inode.IndexNodes[ix-1].DataNodes[length0-1]

					// 不用借了，先直接合拼
					inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].DataNodes[1].Items[0].Key)
					inode.IndexNodes[ix-1].DataNodes = append(inode.IndexNodes[ix-1].DataNodes, inode.IndexNodes[ix].DataNodes[1])

					// 抹除 ix 位置
					inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
					inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)

					// 更新状态
					borrowed = true
					return
				} else if len(inode.IndexNodes[ix-1].DataNodes[length0-1].Items) == 0 {
					err = fmt.Errorf("节点未及时整理完成2")
					return
				}
			}
		}
	}

	// Finally, return
	return
}
