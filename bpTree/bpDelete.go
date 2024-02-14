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
		ix = sort.Search(len(inode.DataNodes[0].Items), func(i int) bool {
			// 二分法直接在资料节点进行搜寻
			return inode.DataNodes[0].Items[i].Key >= item.Key // no equal sign ‼️ no equal sign means delete to the right ‼️
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

		// 🖍️ 在这个区块，会上传边界值，当上传到 ix 大于 0 的地方时，会变成索引，停止上传
		// 当上传到 ix 等于 0 的地方时，就立刻持续上传，到边界值要更新的地方

		// 搜寻 🔍 (最右边 ➡️)
		// If it is discontinuous data (different values) (5 - 5 - 5 - 5 - 5❌ - 6 - 7 - 8)
		deleted, updated, edgeValue, status, _, err = inode.IndexNodes[ix].deleteToRight(item)
		if ix > 0 && status == edgeValueUpload {
			fmt.Print("🏴‍☠️ 索引(4) ", inode.Index, " -> ", " 位置 ", ix-1, " 修改成 ", edgeValue, "->")
			inode.Index[ix-1] = edgeValue
			fmt.Print("最后变成", inode.Index, " 上传中断", "\n")
			updated = false
			status = edgeValueInit
		} else if ix == 0 && status == edgeValueUpload {
			fmt.Print("🏴‍☠️ 索引(5) ", inode.Index, " -> ", " 位置 ", ix, " 边界值为 ", edgeValue, " 再上传")
		} else {
			fmt.Print("🏴‍☠️ 索引(6) ", " 位置 ", ix, " 边界值为 ", edgeValue, " 状态 ", status, " 不更新", "\n")
		}

		// 🖍️ 在这个区块，(暂时) 决定要更新边界值，还是要上传

		// 🖐️ 状态变化 [LeaveBottom] -> Any
		if status == edgeValueLeaveBottom {

			// ⚠️ 状况一 用边界值去更新任意索引

			// 🖐️ 状态变化 [LeaveBottom] -> [Init]
			// 看到 LeaveBottom 状态时，就代表准备要更新边界值，但更新的索引不一定在最左边
			if ix-1 >= 0 {
				fmt.Print("🏴‍☠️ 索引(1) ", inode.Index, "->", "位置", ix-1, "修改成", edgeValue, "->")
				inode.Index[ix-1] = edgeValue
				fmt.Print("最后变成", inode.Index, "\n")

				status = edgeValueInit // 暂时重置状态，之后可能会被改
			} else {
				status = edgeValueUpload // 暂时重置状态，之后可能会被改
			}
		} else if status == statusBorrowFromIndexNode {
			ix, edgeValue, err, status = inode.borrowFromIndexNode(ix)
			return
		}

		// If the index at position ix becomes invalid. ‼️
		// 删除导致锁引失效 ‼️
		if len(inode.IndexNodes[ix].Index) == 0 { // invalid ❌
			if len(inode.IndexNodes[ix].DataNodes) >= 2 { // DataNode 🗂️

				fmt.Print("borrowFromIndexNode 执行前后，🏴‍☠️ 边界值变化 ", inode.edgeValue()) // 显示边界值

				// 之后从这开始开发 ‼️

				var borrowed bool
				borrowed, _, edgeValue, err, status = inode.borrowFromBottomIndexNode(ix) // Will borrow part of the node (借结点). ‼️  // 🖐️ for index node 针对索引节点
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
		if ix == 0 && status == edgeValuePassBottom {                          // 当 ix 为 0 时，才要处理边界值的问题 (ix == 0，是特别加入的)
			status = edgeValueLeaveBottom
		}

		// The individual data node is now empty, and
		// it is necessary to start borrowing data from neighboring nodes.
		if len(inode.DataNodes[ix].Items) == 0 { // 会有一边的资料节点没有任何资料
			var borrowed bool
			borrowed, edgeValue, err, status = inode.borrowFromDataNode(ix) // Will borrow part of the data node. (向资料节点借资料)

			// 看之前的 if 判断式，len(inode.DataNodes) > 0 条件满足后，才会来这里
			// 由这条件可以知，目前是在底层，不是修改边界值的时机，边界值要到上层去修改
			// 在这里的工作是观察边界值是否要往上传
			if ix == 0 && status == edgeValueChanges {
				fmt.Println("上传边界值 ", edgeValue)
				status = edgeValueUpload
				return
			}

			// 先检查是否有错误
			if err != nil {
				status = statusError
				return
			}

			// If the data node cannot be borrowed, then information should be borrowed from the index node later.
			// 如果资料节点借到，就不需后续处理
			if borrowed == true {
				updated = true
				return
			}

			// 如果使用 borrowFromDataNode 没有借到资料，就要进行以下处理 ‼️ ‼️

			// ⚠️ 状况一 索引节点资料过少，整个节点失效
			// During the deletion process, the node's index may become invalid.
			// 如果资料节点数量过少
			if len(inode.DataNodes) <= 2 { // 资料节点数量过少
				inode.Index = []int64{}

				// 状况更新
				updated = true

				// 直接中断
				return
			}

			// ⚠️ 状况二 索引节点有一定数量的资料，删除部份资料后，还能维持为一个节点
			// Wipe out the empty data node at the specified 'ix' position directly.
			// 如果资料节点删除资料后，还是维持为一个节点的定义，就要进行抹除部份 ix 位置上的资料 ‼️
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
			updated, _, _, err, _ = inode.borrowFromBottomIndexNode(ix) // Will borrow part of the index node (向索引节点借资料).
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
			updated, _, err, _ = inode.borrowFromDataNode(ix) // Will borrow part of the data node. (向资料节点借资料)
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

	if item.Key == 621 {
		fmt.Println()
	}

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
func (inode *BpIndex) borrowFromDataNode(ix int) (borrowed bool, edgeValue int64, err error, status int) {
	// No data borrowing is necessary as long as the node is not empty, since all indices are still in their normal state.
	if len(inode.DataNodes[ix].Items) != 0 {
		err = fmt.Errorf("not an empty node, do not need to borrow")
		return
	}

	// 以下会向临近节点借资料，但是邻近节点会被切成 2 半 ‼️

	// Borrow from the left side first
	if (ix - 1) >= 0 { // Left neighbor exists ‼️

		// 初始化回传值
		edgeValue = inode.DataNodes[0].Items[0].Key
		status = edgeValueNoChanges

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

			// 向左借应不会有边界值的变化，到时再考虑是否要去除这段程式码 🔥
			// 检查边界值是否有变化
			if edgeValue != inode.DataNodes[0].Items[0].Key {
				edgeValue = inode.DataNodes[0].Items[0].Key
				status = edgeValueChanges
			}

			return
		}
	}

	// Borrow from the right side next.
	if (ix + 1) <= len(inode.DataNodes)-1 { // Right neighbor exists ‼️
		length := len(inode.DataNodes[ix+1].Items)
		if length >= 2 { // The right neighbor node has enough data to borrow

			// 初始化回传值
			if ix != 0 {
				edgeValue = inode.DataNodes[0].Items[0].Key
			} else if ix == 0 {
				edgeValue = -1
			}

			status = edgeValueNoChanges

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

			// 检查边界值是否有变化
			if edgeValue != inode.DataNodes[0].Items[0].Key {
				edgeValue = inode.DataNodes[0].Items[0].Key
				status = edgeValueChanges
			}
			return
		}
	}

	// Finally, return the result
	return
}

// borrowFromIndexNode will borrow more data from neighboring index nodes, including indexes.
func (inode *BpIndex) borrowFromBottomIndexNode(ix int) (borrowed bool, newIx int, edgeValue int64, err error, status int) {
	// 先初始化回传值
	newIx = -1
	edgeValue = -1
	if len(inode.IndexNodes) > 0 && len(inode.IndexNodes[0].DataNodes) > 0 && len(inode.IndexNodes[0].DataNodes[0].Items) > 0 {
		edgeValue = inode.IndexNodes[0].DataNodes[0].Items[0].Key
	}
	status = edgeValueInit

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
					// return
				}

				// inode 下的第 ix 索引节点剩 2 个资料节点，ix 索引节点 的资料被移到最左方资料
				// 如果 ix 为 0 ，就会造成边界值上传的问题，最后会处理，现在不用管
				// 如果 ix 大于 0，就不需要上传，在 inode 内进行更新
				if ix > 0 {
					inode.Index[ix-1] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key
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

					// inode 下的第 ix 索引节点剩 2 个资料节点，
					// "之前" ix 索引节点 的资料被移到最左方资料，"现在" 向右边的 邻居索引节点 借资料
					// 这个影响右边索引节点的边界值
					// 在这里进行修正
					inode.Index[ix] = inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key

					// 更新状态
					borrowed = true

					// return
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

					// inode 下的第 ix 索引节点剩 2 个资料节点，
					// "之前" ix 索引节点 的资料被移到最左方资料，"现在" 向右边的 邻居索引节点 借资料，
					// 在这里 向右边的 邻居索引节点 的资料节点数量为会减少
					// 影响到右方的邻居索引节点，要同步邻居索引节点的边界值，在这里进行修正
					inode.Index[ix] = inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key

					// 更新状态
					borrowed = true
					// return
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
						inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...) // 边界值在这里修正
						inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
					} else if ix == 0 {
						inode.Index = inode.Index[1:]
						inode.IndexNodes = inode.IndexNodes[1:]
					}

					// ix 索引节点资料先复制到 ix + 1 索引节点那，再移除 ix 索引节点
					// ix + 1 索引节点 会到 ix 位置，ix + 1 索引节点又有之前 ix 节点的资料
					// 所以新节点足够代表之前 ix 位置的索引节点
					// 也就是 ix 值不用修正
					// ix 等于 0 时，要把边界值上，这理不用管，之后会处理
					// ix 大于 0 时，在这段代码有进行修正

					// 更新状态
					borrowed = true
					// return
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

				// inode 下的第 ix 索引节点剩 2 个资料节点，ix 索引节点 的资料被移到最右方资料，就是要先形成中空
				// 如果 ix 为 0 ，就会造成边界值上传的问题，最后会处理，现在不用管，而且这里 ix 也不会为 0，因为 前面有条件 ix-1 >= 0
				// 如果 ix 大于 0，就不需要上传，在 inode 内进行更新
				if len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 {
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}
					// return
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

					// inode 下的第 ix 索引节点剩 2 个资料节点，
					// "之前" ix 索引节点 的资料被移到最右方资料，"现在" 向左边的 邻居索引节点 借资料
					// 因为是向 最左边的索引节点借的是尾部资料，这不 个会 影响右边索引节点的边界值
					// 在这里 不需要 进行修正
					// 同样，上传边界值的问题，最后会处理

					// 更新状态
					borrowed = true
					// return
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

					// inode 下的第 ix 索引节点剩 2 个资料节点，
					// "之前" ix 索引节点 的资料被移到最右方资料，"现在" 向左边的 邻居索引节点 借资料，
					// 在这里 向左边的 邻居索引节点 借尾部资料，所以不必更新索引节点的边界值
					// 但是 ix 的索引节点有向左边的邻居节点借到值，所以边界值要进行更新，进行以下修正
					inode.Index[(ix)-1] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key

					// 更新状态
					borrowed = true
					// return
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

					// ix 索引节点资料先复制到 ix - 1 索引节点那，再移除 ix 索引节点
					// ix - 1 索引节点有之前 ix 节点的资料，所以在位置 ix - 1 的索引节点能代表之前的 ix 的
					newIx = ix - 1

					// 更新状态
					borrowed = true
					// return
				} else if len(inode.IndexNodes[ix-1].DataNodes[length0-1].Items) == 0 {
					err = fmt.Errorf("节点未及时整理完成2")
					return
				}
			}
		}
	}

	if edgeValue != inode.IndexNodes[0].DataNodes[0].Items[0].Key {
		edgeValue = inode.IndexNodes[0].DataNodes[0].Items[0].Key
		status = edgeValueChanges
	}

	// Finally, return
	return
}

func (inode *BpIndex) borrowFromIndexNode(ix int) (newIx int, edgeValue int64, err error, status int) {
	// 🖍️ 在这个区块，是在进行借完资料后处理
	// 要就全合拼，不然就先合拼再重分配

	// ⚠️ 状况二 当一个分支只剩一个索引值和一个索引节点，准备要向左合拼
	// 思考后，还是向左合拼比较好，因为左边的资料结点的资料会比较少，合并时，比较不会过大，比较安全
	if ix-1 >= 0 && ix-1 <= len(inode.IndexNodes)-1 {
		// ⚠️ 状况二之一 先向左合并
		if len(inode.IndexNodes[ix-1].Index)+1 < BpWidth { // 没错，Degree 是针对 Index
			// ⚠️ 状况二之一之一 先向左合并，合拼后底层索引节点过小，合拼成一个新节点
			inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index...)
			inode.IndexNodes[ix-1].IndexNodes = append(inode.IndexNodes[ix-1].IndexNodes, inode.IndexNodes[ix].IndexNodes...)
			inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
			inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)

			// 合拼后，ix 的值要减 1
			newIx = ix - 1

			// 在这里不需要重建连结，因为没有资料节点的操作 ‼️
			// 因为是整个 ix 位置的索引节点向左合拼，最左边索引节点的边界值是不会变的

			status = edgeValueInit

			return
		} else if len(inode.IndexNodes[ix-1].Index)+1 >= BpWidth {
			// ⚠️ 状况二之一之二 先向左合并，合拼后底层索引节点过大，要用 protrudeInOddBpWidth 或 protrudeInEvenBpWidth 重新分配

			// if len(inode.IndexNodes) >= 2 { // 这里要检合拼后，多个节点层数是否相同 ⁉️
			// 后来想想，这里直接去除，因为加1后除2也会维持 Degree，只要层数相同就好

			inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index...) // 剩1个索引和1个索引节点，所以可以直接合拼，但很容易出错

			inode.IndexNodes[ix-1].IndexNodes = append(inode.IndexNodes[ix-1].IndexNodes, inode.IndexNodes[ix].IndexNodes...)
			inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
			inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)

			// 准备要嵌入的节点
			var embed *BpIndex
			var tailIndex = inode.Index[ix-1:]
			var tailIndexNodes []*BpIndex
			tailIndexNodes = append(tailIndexNodes, inode.IndexNodes[ix:]...)

			// 要分成单偶数函式处理
			if len(inode.IndexNodes[ix-1].Index)%2 == 1 { // 针对单数数量的索引节点
				// 当索引为奇数时
				embed, err = inode.IndexNodes[ix-1].protrudeInOddBpWidth() // 进行重新分配
				if err != nil {
					return
				}
			} else if len(inode.IndexNodes[ix-1].Index)%2 == 0 { // 针对偶数数量的索引节点
				// 当索引为偶数时
				embed, err = inode.IndexNodes[ix-1].protrudeInEvenBpWidth() // 进行重新分配
				if err != nil {
					return
				}
			}

			// 在这里要整个嵌入原索引节点

			if ix-2 >= 0 { // 其实考虑可以改成 ix-2 > 0
				// 会用到原始索引的前半段
				inode.Index = append(inode.Index[:ix-2], embed.Index[0])
				inode.Index = append(inode.Index, tailIndex...)
			} else {
				// 不 会用到原始索引的前半段
				inode.Index = append(embed.Index, tailIndex...)
			}

			// 合拼后，执行 protrudeInOddBpWidth 和 protrudeInEvenBpWidth 的，
			// 索引和索引节点都会增加一个单位，另外，因是向左合拼，ix 会大于等于 1
			inode.IndexNodes = append(inode.IndexNodes[:ix-1], embed.IndexNodes...)
			inode.IndexNodes = append(inode.IndexNodes, tailIndexNodes...)

			// 在这里不需要重建连结，因为没有资料节点的操作 ‼️
			// 因为是整个 ix 位置的索引节点向左合拼，最左边索引节点的边界值是不会变的

			status = edgeValueInit

			return
		}
	} else if ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1 { // 不能合拼后再合拼，会出事，所以用 else if，只做一次 ‼️
		// ⚠️ 状况二之二 再向右合并
		if len(inode.IndexNodes[ix+1].Index)+1 < BpWidth { // 没错，Degree 是针对 Index
			// ⚠️ 状况二之二之一 先向右合并，合拼后底层索引节点过小，合拼成一个新节点
			inode.IndexNodes[ix].Index = append([]int64{inode.IndexNodes[ix+1].edgeValue()}, inode.IndexNodes[ix+1].Index...)
			inode.IndexNodes[ix].IndexNodes = append(inode.IndexNodes[ix].IndexNodes, inode.IndexNodes[ix+1].IndexNodes...)
			inode.Index = append(inode.Index[:ix], inode.Index[ix+1:]...)
			inode.IndexNodes = append(inode.IndexNodes[:ix+1], inode.IndexNodes[ix+2:]...)

			status = edgeValueInit

			return
		} else if len(inode.IndexNodes[ix+1].Index)+1 >= BpWidth {
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
		}
	}
	return
}
