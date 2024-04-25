package bpTree

import (
	"errors"
	"fmt"
	"sort"
)

// ➡️ The functions related to direction.

// delFromRoot is responsible for deleting an item from the root of the B Plus tree. // 这是 B 加树的删除入口
func (inode *BpIndex) delFromRoot(item BpItem) (deleted, updated bool, ix int, edgeValue int64, err error) {
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
		deleted, updated, ix, edgeValue, err = inode.delAndDir(item) // 在这里加入方向性
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

func (inode *BpIndex) delAndDir(item BpItem) (deleted, updated bool, ix int, edgeValue int64, err error) {
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
	deleted, updated, edgeValue, _, ix, err = inode.deleteToRight(item) // Delete to the rightmost node ‼️ (向右砍)

	// Return the results.
	return
}

// deleteToRight is designed to delete from the rightmost side within continuous data.  (5 - 5 - 5 - 5 - 5❌ - 6 - 7 - 8)

// deleteToRight 先放前面，因为 deleteToLeft 会抄 deleteToRight 的内容
func (inode *BpIndex) deleteToRight(item BpItem) (deleted, updated bool, edgeValue int64, status int, ix int, err error) {
	// Initialize the return value first.
	status = edgeValueInit
	edgeValue = -1

	// ✈️ Process the Index Node.
	if len(inode.IndexNodes) > 0 {
		ix = sort.Search(len(inode.Index), func(i int) bool {
			// 🖍️ The `Sort` function stops when the condition is met.
			// When it equals, it meets the condition later, so it will delete the data on the far right.
			// When it is greater than or equal to, it meets the condition earlier, so it will delete the data on the far left.
			return inode.Index[i] > item.Key // 在最右边 ‼️
		})

		// Entering the Recursive Function. 🔁
		deleted, updated, edgeValue, status, _, err = inode.IndexNodes[ix].deleteToRight(item)

		// Mechanism for updating edge values.
		if ix > 0 && status == edgeValueUpload {
			// 🖍️ In this block, the edge values will be uploaded.
			// When uploaded to a location where ix is greater than 0, it becomes an index and stops uploading.
			// (边界值会变成索引并中止)

			inode.Index[ix-1] = edgeValue
			updated = false
			status = edgeValueInit

			// Interrupted, index updated, no uploading. ⚠️
			return
		} else if ix == 0 && status == edgeValueUpload {
			// 🖍️ When uploaded to a location where ix equals 0, it continues to upload immediately until the boundary value is not 0.
			// (IX 为 0 时不停上传)

			// Continuous uploading. ⚠️
			return
		}

		// 🖍️ In this block, (temporarily) decide whether you want to update the boundary values or upload the

		// 🖐️ 状态变化 [LeaveBottom] -> Any
		if status == edgeValueRenew {

			// ⚠️ 状况一 用边界值去更新任意索引

			// 🖐️ 状态变化 [LeaveBottom] -> [Init]
			// 看到 LeaveBottom 状态时，就代表准备要更新边界值，但更新的索引不一定在最左边
			if ix-1 >= 0 {

				fmt.Println(">>>>> 更新完成")

				inode.Index[ix-1] = edgeValue

				status = edgeValueInit
				return
			} else {
				fmt.Println(">>>>> 进行上传")
				status = edgeValueUpload
				return
			}

			// To make temporary corrections, mainly to identify the problems.
		} else { // if status == statusBorrowFromIndexNode || item.Key == 960 {

			if status != statusBorrowFromIndexNode && inode.IndexNodes[ix].DataNodes != nil {
				_, _, edgeValue, err, status = inode.borrowFromBottomIndexNode(ix)
				return
			}

			if status != statusBorrowFromIndexNode && len(inode.IndexNodes[ix].Index) == 0 {
				if len(inode.IndexNodes[ix].Index) == 0 {
					inode.IndexNodes[ix].Index = []int64{edgeValue}
				}

				ix, edgeValue, status, err = inode.borrowFromIndexNode(ix)
				if ix == 0 && status == edgeValueChanges {
					status = edgeValueUpload
					return
				}
				return
			}

			if status == statusBorrowFromIndexNode {
				if len(inode.IndexNodes[ix].Index) == 0 {
					inode.IndexNodes[ix].Index = []int64{edgeValue}
				}

				ix, edgeValue, status, err = inode.borrowFromIndexNode(ix)
				if ix == 0 && status == edgeValueChanges {
					status = edgeValueUpload
					return
				}
			}

			return
		}

		// If the index at position ix becomes invalid. ‼️
		// 删除导致锁引失效 ‼️
		if len(inode.IndexNodes[ix].Index) == 0 { // invalid ❌
			if len(inode.IndexNodes[ix].DataNodes) >= 2 { // DataNode 🗂️

				// 之后从这开始开发 ‼️

				var borrowed bool

				borrowed, _, edgeValue, err, status = inode.borrowFromBottomIndexNode(ix) // Will borrow part of the node (借结点). ‼️  // 🖐️ for index node 针对索引节点
				// 看看有没有向索引节点借到资料

				if err != nil && !errors.Is(err, fmt.Errorf("the index is still there; there is no need to borrow nodes")) {
					return
				}

				if borrowed == true { // 当向其他索引节点借完后，在执行 borrowFromIndexNode，重新计算边界值

					if ix == 0 && status == edgeValueChanges {
						fmt.Println(">>>>> 进行上传")
						status = edgeValueUpload
						return
					}

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

	// ✈️ Process the Data Node.
	if len(inode.DataNodes) > 0 {
		// Call the deleteBottomItem method on the current node as it is close to the bottom layer.
		// This signifies the beginning of deleting data. (接近资料层)

		// Here, adjustments may be made to IX (IX 在这里可能会被修改) ‼️
		// var edgeValue int64
		deleted, updated, ix, edgeValue, status = inode.deleteBottomItem(item) // 🖐️ for data node 针对资料节点
		if ix == 0 && status == edgeValueChangesByDelete {                     // 当 ix 为 0 时，才要处理边界值的问题 (ix == 0，是特别加入的)
			status = edgeValueRenew
		}

		// The individual data node is now empty, and
		// it is necessary to start borrowing data from neighboring nodes.
		if len(inode.DataNodes[ix].Items) == 0 { // 会有一边的资料节点没有任何资料
			var borrowed bool
			borrowed, edgeValue, err = inode.borrowFromDataNode(ix) // Will borrow part of the data node. (向资料节点借资料)

			// 先检查是否有错误
			if err != nil {
				status = statusError
				return
			}

			if edgeValue != -1 {
				status = edgeValueUpload
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

				// 有可能新增状态
				if item.Key == 960 {
					fmt.Println()
				}
				fmt.Println("有可能新增状态")

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

					status = edgeValueInit
				} else if inode.DataNodes[ix].Next == nil {
					inode.DataNodes[ix].Previous.Next = nil

					status = edgeValueInit
				} else {
					inode.DataNodes[ix].Previous.Next = inode.DataNodes[ix].Next
					inode.DataNodes[ix].Next.Previous = inode.DataNodes[ix].Previous

					status = edgeValueInit
				}

				// Reorganize nodes.
				if ix != 0 {
					inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)             // Erase the position of ix - 1.
					inode.DataNodes = append(inode.DataNodes[:ix], inode.DataNodes[ix+1:]...) // Erase the position of ix.

					status = edgeValueInit
				} else if ix == 0 { // Conditions have already been established earlier, with the index length not equal to 0. ‼️
					inode.Index = inode.Index[1:]
					inode.DataNodes = inode.DataNodes[1:]

					// 边界值要立刻进行修改
					edgeValue = inode.DataNodes[0].Items[0].Key
					status = edgeValueUpload
				}
			}
		}

	}

	// Return the results of the deletion.
	return
}

// deleteToLeft is a method of the BpIndex type that deletes the leftmost specified BpItem. (由左边删除 👈 ‼️)
func (inode *BpIndex) deleteToLeft(item BpItem) (deleted, updated bool, ix int, err error) {
	panic("Currently developing right deletion, not developing left deletion.")

	// Return the results of the deletion.
	return
}

// deleteBottomItem will remove data from the bottom layer. (只隔一个索引 ‼️)
// If the node is too small, it will clear the entire index. (索引可能失效‼️)
// 一层 BpData 资料层，加上一个索引切片，就是一个 Bottom
func (inode *BpIndex) deleteBottomItem(item BpItem) (deleted, updated bool, ix int, edgeValue int64, status int) {

	// Perform an interruption to check the B Plus tree
	if item.Key == 960 {
		fmt.Println("Perform An Interruption")
	}

	// 初始化回传值
	edgeValue = -1

	// Use binary search to find the index (ix) where the key should be inserted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] > item.Key // No equal sign ‼️
	})

	// Call the delete method on the corresponding DataNode to delete the item.
	deleted, _, edgeValue, status = inode.DataNodes[ix]._delete(item) // 总是有错误
	// _delete 函式状况会回传 (1) 边界值没改变 (2) 边界值已改变 (3) 边界值为空

	// 这里精简
	/*if status == edgeValueChangesByDelete { // (1) 边界值已改变
		status = edgeValuePassBottom // 要通知上传的递归函式，边界值已改变
	}*/

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

			inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key // Immediately update the index

			// Return status
			updated = true
			return
		}
	}

	// Return the results of the deletion.
	return
}

// borrowFromDataNode 🛠️ only borrows a portion of data from the neighbor nodes.
// As for the direction, it may be borrowing data from the left data node, but it may also be borrowing data from the right one. (向左右两方借资料)
// The whole operation is complicated, please refer to the documentation Chapter 2.3.1 Borrow from Neighbor.
func (inode *BpIndex) borrowFromDataNode(ix int) (borrowed bool, outerEdgeValue int64, err error) {
	// ⚙️ Pre-operation and inspection.

	// Initialization Outer-Edge-Value.
	outerEdgeValue = -1

	// No data borrowing is necessary as long as the node is not empty, since all indices are still in their normal state.
	if len(inode.DataNodes[ix].Items) != 0 {
		err = fmt.Errorf("not an empty node, the current data node do not need to borrow data from either side")
		return
	}

	// ⚙️ Processing of **statuses 1** and **3**, borrowing data from the right neighbor data node.

	// This is due to the fact that for most conditions, the right neighbor data node has a higher number of data.

	// Borrow data from the right neighbor data node.
	if (ix + 1) <= len(inode.DataNodes)-1 { // Right neighbor data node exists ‼️
		borrowerLength := len(inode.DataNodes[ix+1].Items)
		if borrowerLength >= 2 { // The right neighbor data node has enough data to borrow.

			// The right neighbor node is split.
			borrowedItems := inode.DataNodes[ix+1].Items[:1] // First part contains an borrowed element. (This is the first data from the right neighbor data node.)
			remainItems := inode.DataNodes[ix+1].Items[1:]   // Second part contains the remaining elements.
			// (反正就借右资料节点的第一笔资料，只借一笔)

			// Further distribution will be completed by borrowing process.
			inode.DataNodes[ix].Items = borrowedItems
			inode.DataNodes[ix+1].Items = remainItems

			// First update an Inner-Edge-Value.
			inode.Index[ix] = inode.DataNodes[ix+1].Items[0].Key

			if ix != 0 {
				// Update another Inner-Edge-Value. (Status 2 状况 2 ⬅️)
				inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key
			} else {
				// Upload the Outer-Edge-Values. (Status 1 状况 1 ⬅️)
				outerEdgeValue = inode.DataNodes[ix].Items[0].Key
			}

			// The return status indicates that the data has been borrowed.
			borrowed = true

			return
		}
	}

	// ⚙️ Processing of **statuses 2-1** and **2-2**, borrowing data from the left neighbor data node.

	// Borrow data from the left neighbor data node.
	if (ix - 1) >= 0 { // Left neighbor data node exists ‼️
		borrowerLength := len(inode.DataNodes[ix-1].Items)
		if borrowerLength >= 2 { // The left neighbor data node has enough data to borrow.
			// The left neighbor node is split.
			remainItems := inode.DataNodes[ix-1].Items[:(borrowerLength - 1)]   // First part contains the remaining elements.
			borrowedItems := inode.DataNodes[ix-1].Items[(borrowerLength - 1):] // Second part contains an borrowed element. (This is the last data from the left neighbor data node.)

			// Further distribution will be completed by borrowing process.
			inode.DataNodes[ix-1].Items = remainItems
			inode.DataNodes[ix].Items = borrowedItems

			// Update an Inner-Edge-Value.
			inode.Index[ix-1] = inode.DataNodes[ix].Items[0].Key // (Status 2-1 2-2 状况 2-1 2-2 ⬅️ ⬅️)
			// (在不符合状况1和状况3执行此行)

			// The return status indicates that the data has been borrowed.
			borrowed = true

			return
		}
	}

	// Finally, return the result
	return
}

// The differences between the borrowFromBottomIndexNode function ⚙️ and borrowFromIndexNode are as follows:
// `borrowFromBottomIndexNode` performs borrowing operations from the bottom-level index node, while also handling index nodes and data nodes.
// On the other hand, `borrowFromIndexNode` only deals with index nodes.
func (inode *BpIndex) borrowFromBottomIndexNode(ix int) (borrowed bool, newIx int, edgeValue int64, err error, status int) {
	// The return value is initialized to a negative value first, because the indices in the database are all positive and there won't be any negative values.
	// (初始化为负值，有更改易发现)
	newIx = -1
	edgeValue = -1

	// 🖍️ The return value is initially initialized to a negative value because the indices in the database are all positive, and there are no negative values.
	// This makes it easier to detect if there have been any modifications. (初始化为负值，有变化才容易发现)
	if len(inode.IndexNodes) > 0 && len(inode.IndexNodes[0].DataNodes) > 0 && len(inode.IndexNodes[0].DataNodes[0].Items) > 0 {
		edgeValue = inode.IndexNodes[0].DataNodes[0].Items[0].Key
	}
	status = edgeValueInit

	// 🖍️ As long as (1) the index node contains data, // 含资料的索引节点
	// but (2) becomes invalid due to an empty index, // 失效
	// and (3) has neighboring nodes, borrowing data can take place. // 有邻居
	// (符合这三条件可借资料)

	// 🖍️ However, could there be a situation where there are no neighbors?
	// No, because after merging into a single node in borrowFromBottomIndexNode, borrowing from borrowFromIndexNode will occur,
	// so there won't be no neighbors.
	// 会不会有没邻居？不，就算 borrowFromBottomIndexNode 合拼成 1 节点，borrowFromIndexNode 会去借资料，不会没邻居

	if inode.IndexNodes[ix].DataNodes != nil && len(inode.IndexNodes[ix].Index) == 0 && len(inode.IndexNodes) >= 2 {

		// 🖍️ When merging, merge with the neighbor node on the left because it may have fewer data.
		// When borrowing data, borrow from the neighbor node on the right because it may have more data.
		// (合拼向左，借资料向右)

		// 🖍️ When the right neighbor node has sufficient data and the data node has two or more elements.
		// If borrowing from the neighbor node results in its invalidation, it will be merged.
		// (2个以上足够，就算邻居节点失效，就合拼)
		if (ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1) && len(inode.IndexNodes[ix+1].DataNodes) >= 2 {

			// The following can be explained conveniently with the diagram below:
			// [] represents data nodes
			// () represents index nodes
			// <-link-> represents links

			// 🖍️ As shown below, a hollow forms between the final origin index node and the neighbor index node.
			// ( [0] <-link-> [1] )origin <-link-> ( [unknown] <-link-> [unknown] )neighbor
			// ( [1] <-link-> [0] )origin <-link-> ( [unknown] <-link-> [unknown] )neighbor
			// (形成中空)

			// 🖍️ As shown below, a solid forms between the final origin index node and the neighbor index node.
			// ( [0] <-link-> [2] )origin <-link-> ( [unknown] <-link-> [unknown] )neighbor
			// ( [1] <-link-> [1] )origin <-link-> ( [unknown] <-link-> [unknown] )neighbor
			// (形成实心)

			// 🖍️ Not considering boundary values for now, will handle them later.

			// [ Perform Simple Aid ]
			// To prepare for becoming hollow or solid.
			if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 {
				// Borrow data in the same index node from the data node first.
				inode.IndexNodes[ix].DataNodes[0].Items = append(inode.IndexNodes[ix].DataNodes[0].Items, inode.IndexNodes[ix].DataNodes[1].Items[0])
				inode.IndexNodes[ix].DataNodes[1].Items = inode.IndexNodes[ix].DataNodes[1].Items[1:]

				// Update the index of the original index node.
				if len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 {
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}
				}

				// Update inode's index.
				if ix > 0 {
					inode.Index[ix-1] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key
				}
			}

			// If the following hollow state does indeed form, we need to borrow a node from the neighbor node. (中空形成)
			if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 {

				// If the neighbor node has sufficient data, although it does not damage the neighbor, the index of the inode will be modified. (非破坏)
				// Although the neighbor node is damaged, it does not cause the neighbor node to be valid.
				if len(inode.IndexNodes[ix+1].DataNodes[0].Items) >= 2 {
					// Borrow data from the neighbor node first.
					inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix+1].DataNodes[0].Items[0])
					inode.IndexNodes[ix+1].DataNodes[0].Items = inode.IndexNodes[ix+1].DataNodes[0].Items[1:]

					// Update the index of the original index node. (ix 节点更新索引)
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// Update inode's index. (ix+1 节点边界值)
					inode.Index[ix] = inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key

					// Update the status.
					borrowed = true

					// If the neighbor node does not have sufficient data, borrowing data will result in the destruction of neighboring nodes. (被破坏)
				} else if len(inode.IndexNodes[ix+1].DataNodes[0].Items) == 1 && len(inode.IndexNodes[ix+1].DataNodes) >= 3 {
					// Borrow data from the neighbor node first.
					inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix+1].DataNodes[0].Items[0])
					inode.IndexNodes[ix+1].DataNodes[0].Items = inode.IndexNodes[ix+1].DataNodes[0].Items[1:]

					// Update the index of the original index node.
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// Rebuild the connection; inode.IndexNodes[ix+1].DataNodes[0] will transfer all links.
					inode.IndexNodes[ix+1].DataNodes[1].Previous = inode.IndexNodes[ix+1].DataNodes[0].Previous
					inode.IndexNodes[ix].DataNodes[1].Next = inode.IndexNodes[ix+1].DataNodes[0].Next

					// Remove empty node that is inode.IndexNodes[ix+1].DataNodes[0]
					inode.IndexNodes[ix+1].Index = inode.IndexNodes[ix+1].Index[1:]
					inode.IndexNodes[ix+1].DataNodes = inode.IndexNodes[ix+1].DataNodes[1:]

					// Update inode's index.
					inode.Index[ix] = inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key

					// Update the status.
					borrowed = true

					// If the neighbor node does not have sufficient data and does not have sufficient neighbors, borrowing data will result in being merged. (被合拼)
				} else if len(inode.IndexNodes[ix+1].DataNodes[0].Items) == 1 && len(inode.IndexNodes[ix+1].DataNodes) == 2 {
					// The node at position ix is going to be erased, and before erasing, its connections will be reconstructed. (被抹 ix 索引，重建)
					previousData := inode.IndexNodes[ix].DataNodes[0].Previous
					nextData := inode.IndexNodes[ix].DataNodes[0].Next

					inode.IndexNodes[ix+1].DataNodes[0].Previous = previousData
					if previousData != nil {
						previousData.Next = nextData
					}

					// All data centralized to position ix + 1.
					inode.IndexNodes[ix+1].Index = append([]int64{inode.IndexNodes[ix+1].DataNodes[0].Items[0].Key}, inode.IndexNodes[ix+1].Index...)

					// The data at ix + 1 contains that of ix, therefore the index at position ix also needs to be corrected to ix - 1.
					// ix+1 的资料内含 ix 的，之后 ix 位置的索引也要修正成 ix-1 的 (索引和索引节点只差个单位)
					inode.IndexNodes[ix+1].DataNodes = append([]*BpData{inode.IndexNodes[ix].DataNodes[0]}, inode.IndexNodes[ix+1].DataNodes...)

					// Erase the indexed node at position ix.
					if ix > 0 {
						// The index at position ix also needs to be corrected to ix-1.
						// ix 位置的索引也要修正成 ix-1 的
						inode.Index[ix] = inode.Index[ix-1]

						// Erase the indexed node at position ix.
						inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
						inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
					} else if ix == 0 {
						// Erase the indexed node at position ix.
						inode.Index = inode.Index[1:]
						inode.IndexNodes = inode.IndexNodes[1:]
					}

					// Adjust ix to the original data position after merging.
					// original data moved to ix+1, delete ix, original data moved from ix+1 to ix
					// newIX = ix

					// Update the status.
					borrowed = true
				}
			}

			// Here is the latter part discussing borrowing materials from the neighbor on the right. (现在才要讨论向右借资料) ‼️

			// The following can be explained conveniently with the diagram below:
			// [] represents data nodes
			// () represents index nodes
			// <-link-> represents links

			// 🖍️ As shown below, a hollow forms between the final origin index node and the neighbor index node.

			// ( [unknown] <-link-> [unknown] )neighbor <-link-> ( [1] <-link-> [0] )origin
			// ( [unknown] <-link-> [unknown] )neighbor <-link-> ( [0] <-link-> [1] )origin
			// (形成中空)

			// 🖍️ As shown below, a solid forms between the final origin index node and the neighbor index node.

			// ( [unknown] <-link-> [unknown] )neighbor <-link-> ( [2] <-link-> [0] )origin
			// ( [unknown] <-link-> [unknown] )neighbor <-link-> ( [1] <-link-> [1] )origin
			// (形成实心)

			// 🖍️ Not considering boundary values for now, will handle them later.

			// To prepare for becoming hollow or solid.
		} else if (ix-1 >= 0 && ix-1 <= len(inode.IndexNodes)-1) && len(inode.IndexNodes[ix-1].DataNodes) >= 2 {

			if len(inode.IndexNodes[ix].DataNodes[1].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[0].Items) > 0 {
				// Borrow data in the same index node from the data node first.
				length0 := len(inode.IndexNodes[ix].DataNodes[0].Items)
				inode.IndexNodes[ix].DataNodes[1].Items = append(inode.IndexNodes[ix].DataNodes[1].Items, inode.IndexNodes[ix].DataNodes[0].Items[length0-1])
				inode.IndexNodes[ix].DataNodes[0].Items = inode.IndexNodes[ix].DataNodes[0].Items[:length0-1] // 不包含最后一个

				// ( [unknown] <-link-> [unknown] )neighbor <-link-> ( [1] <-link-> [0] )origin
				// ( [unknown] <-link-> [unknown] )neighbor <-link-> ( [0] <-link-> [1] )origin
				// neighbor node and origin node result a phenomenon of hollow.
				// At this point, the index might still be in a invalid state, so I'll just update the index directly.
				// (在中间状态，origin 失效，但还是先更新索引)
				inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}
			}

			// If the following hollow state does indeed form, we need to borrow a node from the neighbor node. (中空形成)
			if len(inode.IndexNodes[ix].DataNodes[0].Items) == 0 && len(inode.IndexNodes[ix].DataNodes[1].Items) > 0 {

				// Knowing the number of items in the nearest data node.
				numDataNodeInNeighbor := len(inode.IndexNodes[ix-1].DataNodes)                                 // The number of data nodes in neighbor nodes.
				numItemClosestDataNode := len(inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items) // The number of items in the closest Data Node.

				// If the neighbor node has sufficient data, although it does not damage the neighbor, the index of the inode will be modified. (非破坏)
				if len(inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items) >= 2 && numDataNodeInNeighbor > 0 && numItemClosestDataNode > 0 {
					// Knowing the number of items in the nearest data node.
					inode.IndexNodes[ix].DataNodes[0].Items = append(inode.IndexNodes[ix].DataNodes[0].Items, inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items[numItemClosestDataNode-1])
					inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items = inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items[:(numItemClosestDataNode - 1)] // "Wipe out the last item."

					// After borrowing data, the index of the index node at position ix-1 will not change. ‼️
					// (ix - 1 那的索引节点都不会变 ‼️)

					// The index has already been updated, so this line of code is not executed. (更新索引)
					// inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// Update inode's index. (ix 节点边界值)
					inode.Index[ix-1] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key

					// Update the status.
					borrowed = true

					// If the neighbor node does not have sufficient data, borrowing data will result in the destruction of neighboring nodes. (被破坏)
					// Although the neighbor node is damaged, it does not cause the neighbor node to be valid.
				} else if len(inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items) == 1 && len(inode.IndexNodes[ix-1].DataNodes) >= 3 && numDataNodeInNeighbor > 0 && numItemClosestDataNode > 0 {
					// Borrow data from the neighbor node first.
					inode.IndexNodes[ix].DataNodes[0].Items = append(inode.IndexNodes[ix].DataNodes[0].Items, inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items[numItemClosestDataNode-1])
					// >>> The moved data does not need to be wiped in the original location, because the neighboring data nodes will be removed afterwards.
					// >>> (不抹除搬移资料，将删除资料节点)

					// The index has already been updated, so this line of code is not executed. (更新索引)
					inode.IndexNodes[ix].Index = []int64{inode.IndexNodes[ix].DataNodes[1].Items[0].Key}

					// Rebuild the connection; inode.IndexNodes[ix-1].DataNodes[LastOne] will transfer all links.
					inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-2].Next = inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Next
					inode.IndexNodes[ix].DataNodes[0].Previous = inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Previous

					// Remove empty node that is inode.IndexNodes[ix-1].DataNodes[LastOne]
					inode.IndexNodes[ix-1].Index = inode.IndexNodes[ix-1].Index[:(numDataNodeInNeighbor - 2)]
					inode.IndexNodes[ix-1].DataNodes = inode.IndexNodes[ix-1].DataNodes[:(numDataNodeInNeighbor - 1)] // Will not contain numDataNodeInNeighbor-1

					// Update inode's index.
					inode.Index[(ix)-1] = inode.IndexNodes[ix].DataNodes[0].Items[0].Key

					// Update the status.
					borrowed = true

					// If the neighbor node does not have sufficient data and does not have sufficient neighbors, borrowing data will result in being merged. (被合拼)
				} else if len(inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Items) == 1 && len(inode.IndexNodes[ix-1].DataNodes) == 2 && numDataNodeInNeighbor > 0 { // 邻点太小，将会被合拼，进入 [状况1-3]
					// The node at position ix is going to be erased, and before erasing, its connections will be reconstructed. (被抹 ix 索引，重建)
					previousData := inode.IndexNodes[ix].DataNodes[0].Previous
					nextData := inode.IndexNodes[ix].DataNodes[0].Next

					inode.IndexNodes[ix-1].DataNodes[numDataNodeInNeighbor-1].Next = nextData
					if nextData != nil {
						nextData.Previous = previousData
					}

					// All data centralized to position ix - 1.
					inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].DataNodes[1].Items[0].Key)

					// Instead of using borrowed data, the original data nodes and neighboring nodes are first directly merged.
					inode.IndexNodes[ix-1].DataNodes = append(inode.IndexNodes[ix-1].DataNodes, inode.IndexNodes[ix].DataNodes[1])

					// The situation here is that there is a left node at position ix-1, so the following ix-1 must not be an error
					// while being careful that ix+1 has a non-existent problem.
					if ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1 {
						inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
						inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
					} else {
						inode.Index = inode.Index[:ix-1]
						inode.IndexNodes = inode.IndexNodes[:ix]
					}

					// The data is concentrated on ix - 1 and the position is corrected.
					newIx = ix - 1

					// Update the status.
					borrowed = true
				}
			}
		}
	}

	// Finally check that the edge values have been updated.
	if len(inode.IndexNodes) > 0 && len(inode.IndexNodes[0].DataNodes) > 0 && len(inode.IndexNodes[0].DataNodes[0].Items) > 0 &&
		edgeValue != inode.IndexNodes[0].DataNodes[0].Items[0].Key {
		edgeValue = inode.IndexNodes[0].DataNodes[0].Items[0].Key
		status = edgeValueChanges
	}

	// Finally, return
	return
}

func (inode *BpIndex) borrowFromRootIndexNode(ix int, edgeValue int64) (err error) {
	if len(inode.IndexNodes[ix].Index) == 0 {
		inode.IndexNodes[ix].Index = []int64{edgeValue}
	}
	_, _, _, err = inode.borrowFromIndexNode(ix)
	return
}

// borrowFromIndexNode function ⚙️ is used to borrow data when it is not a root node or a data node, to quickly maintain the operation of the B Plus Tree.
// (在 非根节点 和 非资料节点)
// When a B Plus Tree deletes data, the index nodes may need to borrow data.
// The reason B Plus Tree borrows data is to quickly adjust its index to ensure the normal operation of the B Plus Tree.
// Scanning the entire B Plus tree and making large-scale adjustments is impractical and may cause performance bottlenecks. (借资料维持整个树的运作)
// Therefore, I believe that the operations of deleting data in a B P Tree may be slower than adding new data's. (我认为 B 加树删除操作会比新增较慢)
func (inode *BpIndex) borrowFromIndexNode(ix int) (newIx int, edgeValue int64, status int, err error) {

	// 🩻 The index at position ix must be set first, otherwise the number of indexes and nodes won't match up later.
	if len(inode.IndexNodes[ix].Index) == 0 {
		err = fmt.Errorf("the index at position ix must be set first")
		return
	}

	// There is a neighbor node on the left.
	if ix-1 >= 0 && ix-1 <= len(inode.IndexNodes)-1 {

		// 🖍️ The index node may not be able to borrow data, this is when the neighboring node has too little data,
		// then the index node and the neighboring node will be merged to one index node. (借不到就合拼)
		//
		// 🖍️ If only one index node remains after merging in inode, (借资枓失败，上层再处理)
		// the upper-level node will continue to borrow data to maintain the operation of the entire tree.

		// 🖍️ it's better to merge to the left neighbor node because the data nodes on the left side usually have fewer data,
		// which makes the merging less likely to be too large and thus safer. (优先向左合拼)

		// There is a neighbor node on the left.
		if len(inode.IndexNodes[ix-1].Index)+1 < BpWidth { // That's right, "Degree" is for the index. ‼️

			// Merge into the left neighbor node first.
			inode.combineToLeftNeighborNode(ix)

			// ⚠️ Here, because the node is too small after merging, the data borrowing might fail, leading the upper-level node to continue borrowing data. (合并后太小了)

			// 🖍️ [IX] ix-1 indicates the position of the newly merged index node. (ix-1 为新的位置)
			newIx = ix - 1

			// 🖍️ [Link] Here, there's no need to reconstruct data node links as there are no operations involving data nodes. (不重建连结)
			// nothing

			// 🖍️ Because the original data in position ix is being merged to the left, the edge value of the leftmost index node will not change. (边界值不变)
			status = edgeValueInit

			return

		} else if len(inode.IndexNodes[ix-1].Index)+1 >= BpWidth {

			// Merge into the left neighbor node first.
			inode.combineToLeftNeighborNode(ix)

			// 🦺 The index of the merged node becomes excessively large, requiring reallocation using either protrudeInOddBpWidth or protrudeInEvenBpWidth.

			// The original data is located at ix-1. Subsequently, backing up the data of the index nodes occurs after position ix (inclusive 包含).
			var embedNode *BpIndex
			var tailIndexNodes []*BpIndex
			tailIndexNodes = append(tailIndexNodes, inode.IndexNodes[ix:]...) // 原资料在 ix-1，那备份 ix 之后的索引节点的资料
			// The position difference between the index and the index node is one.
			// 备份 ix 之后的索引节点的资料，那索引就是备份 ix 之后的位置
			tailIndex := make([]int64, len(inode.Index[ix-1:])) // Deep copying to prevent value changes
			copy(tailIndex, inode.Index[ix-1:])

			// The merged nodes are subjected to reallocation.
			if len(inode.IndexNodes[ix-1].Index)%2 == 1 { // For odd quantity of index, reallocate using the odd function.
				if embedNode, err = inode.IndexNodes[ix-1].protrudeInOddBpWidth(); err != nil {
					return
				}
			} else if len(inode.IndexNodes[ix-1].Index)%2 == 0 { // For even quantity of index, reallocate using the even function.
				if embedNode, err = inode.IndexNodes[ix-1].protrudeInEvenBpWidth(); err != nil {
					return
				}
			}

			// 🖍️ The data to be merged should be divided into three segments:
			// Front Segment (inode.IndexNodes[:ix-1]): The segment before ix-1 (exclusive 不含)
			// Middle Segment (embedNode) : The data at ix-1
			// Back Segment (tailIndexNodes) : The segment after ix (inclusive)
			inode.IndexNodes = append(inode.IndexNodes[:ix-1], embedNode.IndexNodes...)
			inode.IndexNodes = append(inode.IndexNodes, tailIndexNodes...)

			// Let's adjust the index.

			// The original data is at ix-1. Using this position as a boundary, if ix-2 >= 0, it indicates the presence of the Front Segment.
			if ix-2 >= 0 { // 原始数据位于 ix-1，如果 ix-2 >= 0，则表示存在前半部分
				// 🖍️ After merging with the left node, the data is redistributed and split into two nodes again, with only one index value changes, which is at the position of index node ix.
				// 合拼后再重分配后，只有一个索引值会变，就在索引节点的位置为 ix 的地方
				inode.Index = append(inode.Index[:ix-1], embedNode.Index[0]) // 但是要转换到索引位置时，要减1，为ix-1，也就是 inode.Index[:ix-1]
				inode.Index = append(inode.Index, tailIndex...)
			} else {
				// 🖍️ If ix is not 0, it is 1, there must be a neighbor node on the left side, so ix is 1.
				// The original data is merged into the position of ix-1, which is also 0, and then redistributed.
				// So, it's fine to directly use embedNode.Index to form the new index.

				// ix 不是 0，就是 1，一定有左边的邻居节点，所以 ix 就是 1
				// 原始数据合并到 ix-1 的位置，也是 0，再重新分配
				// 所以直接用 embedNode.Index 去组成新索引就好了
				inode.Index = append(embedNode.Index, tailIndex...)
			}

			// 🖍️ [IX] After merging with the left node, it is redistributed and split into two nodes again, so the position of ix remains unchanged.
			// (合拼到左节点后，再重新分配并分割成两个节点，所以 ix 位置不变)

			// 🖍️ [Link] Here, there's no need to reconstruct data node links as there are no operations involving data nodes. (不重建连结)
			// nothing

			// 🖍️ [Status] Because the entire index position is being merged to the left and be split into two nodes again,
			// the edge value of the leftmost index node will not change. (边界值不变)

			status = edgeValueInit

			return
		}

		// 🖍️ When unable to borrow data from the left neighbor node, start borrowing data from the right neighbor node.
		// Here we don't simplify the code by changing `ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1` to `ix == 0`,
		// because even if `ix == 0`, when `inode` has only one index node left, there may be no neighbor nodes at all, and borrowing data may still not be possible.
		// (只剩一个索引节点时，没邻居，会有都借不到的问题，条件不能精简成 ix == 1)

		// 🖍️ Borrowing data repeatedly is not allowed; It can only be done once.
		// Therefore, it is crucial to use 'else if' here.
	} else if ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1 { // 不能连续借资料，必用 else if ⚠️

		if len(inode.IndexNodes[ix+1].Index)+1 < BpWidth { // 没错，Degree 是针对 Index

			// Merge into the right neighbor node first.
			inode.combineToRightNeighborNode(ix)

			// ⚠️ Here, because the node is too small after merging, the data borrowing might fail, leading the upper-level node to continue borrowing data. (合并后太小了)

			// 🖍️ [IX] The IX position remains unchanged, as mentioned earlier. (ix 位置不变)
			// empty

			// 🖍️ [Link] Here, there's no need to reconstruct data node links as there are no operations involving data nodes. (不重建连结)
			// nothing

			// 🖍️ [Status] Because the original data in position ix is being merged to the right, the edge value of the leftmost index node will not change. (边界值不变)
			status = edgeValueInit

			return

		} else if len(inode.IndexNodes[ix+1].Index)+1 >= BpWidth {

			// Merge into the right neighbor node first.
			inode.combineToRightNeighborNode(ix)

			// 🦺 The index of the merged node becomes excessively large, requiring reallocation using either protrudeInOddBpWidth or protrudeInEvenBpWidth.

			// The original data is located at ix. Subsequently, backing up the data of the index nodes occurs after position ix+1 (inclusive 包含).
			var embedNode *BpIndex
			var tailIndexNodes []*BpIndex
			tailIndex := make([]int64, len(inode.Index[ix:])) // Deep copying to prevent value changes

			// 🖍️ [Check] The index node under the inode has been previously merged, so now we need to check if the index node at position ix+1 exists.
			// 再检查一次 ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1
			if ix+1 >= 0 && ix+1 <= len(inode.IndexNodes)-1 {
				tailIndexNodes = append(tailIndexNodes, inode.IndexNodes[ix+1:]...) // 原资料在 ix，那备份 ix+1 之后的索引节点的资料
				// The position difference between the index and the index node is one.
				// 备份 ix+1 之后的索引节点的资料，那索引就是备份 ix 之后的位置
				copy(tailIndex, inode.Index[ix:]) // Deep copying to prevent value changes
			}

			// The merged nodes are subjected to reallocation.
			if len(inode.IndexNodes[ix].Index)%2 == 1 { // For odd quantity of index, reallocate using the odd function.
				// 当索引为奇数时
				if embedNode, err = inode.IndexNodes[ix].protrudeInOddBpWidth(); err != nil {
					return
				}
			} else if len(inode.IndexNodes[ix].Index)%2 == 0 { // For even quantity of index, reallocate using the even function.
				// 当索引为偶数时
				if embedNode, err = inode.IndexNodes[ix].protrudeInEvenBpWidth(); err != nil {
					return
				}
			}

			// 🖍️ The data to be merged should be divided into three segments:
			// Front Segment (inode.IndexNodes[:ix]): The segment before ix (exclusive 不含)
			// Middle Segment (embedNode) : The data at ix
			// Back Segment (tailIndexNodes) : The segment after ix+1 (inclusive)
			inode.IndexNodes = append(inode.IndexNodes[:ix], embedNode.IndexNodes...)
			inode.IndexNodes = append(inode.IndexNodes, tailIndexNodes...)

			// Let's adjust the index.

			// The original data is at ix. Using this position as a boundary, if ix-1 >= 0, it indicates the presence of the Front Segment.
			if ix-1 >= 0 { // 原始数据位于 ix，如果 ix-1 >= 0，则表示存在前半部分
				// 🖍️ After merging with the right node, the data is redistributed and split into two nodes again, with only one index value changes, which is at the position of index node ix+1.
				// 合拼后再重分配后，只有一个索引值会变，就在索引节点的位置为 ix+1 的地方
				inode.Index = append(inode.Index[:ix], embedNode.Index[0]) // 但是要转换到索引位置时，要减1，为ix，也就是 inode.Index[:ix]
				inode.Index = append(inode.Index, tailIndex...)
			} else {
				// If there is no the Front Segment.
				inode.Index = append(embedNode.Index, tailIndex...)
			}

			// 🖍️ [IX] After merging with the right node, it is redistributed and split into two nodes again, so the position of ix remains unchanged.
			// (合拼到右节点后，再重新分配并分割成两个节点，所以 ix 位置不变)

			// 🖍️ [Link] Here, there's no need to reconstruct data node links as there are no operations involving data nodes. (不重建连结)
			// nothing

			// 🖍️ [Status] Because the entire index position is being merged to the left and be split into two nodes again,
			// the edge value of the leftmost index node will not change. (边界值不变)

			status = edgeValueInit

			return
		}
	}
	return
}

// combineToLeftNeighborNode is part of borrowFromIndexNode, where the current index node will be merged into the left neighbor node.
// (borrowFromIndexNode 的一部份)
func (inode *BpIndex) combineToLeftNeighborNode(ix int) {
	// The data merges with the left neighbor node.
	inode.IndexNodes[ix-1].Index = append(inode.IndexNodes[ix-1].Index, inode.IndexNodes[ix].Index...)
	inode.IndexNodes[ix-1].IndexNodes = append(inode.IndexNodes[ix-1].IndexNodes, inode.IndexNodes[ix].IndexNodes...)

	// Deleting the data node at position ix will result in the original data being at position ix - 1. (原资料就在 ix -1)
	inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)
	inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
	return
}

// combineToRightNeighborNode is part of borrowFromIndexNode, where the current index node will be merged into the right neighbor node.
// (borrowFromIndexNode 的一部份)
func (inode *BpIndex) combineToRightNeighborNode(ix int) {
	// The data merges with the right neighbor node.
	inode.IndexNodes[ix].Index = append([]int64{inode.IndexNodes[ix+1].edgeValue()}, inode.IndexNodes[ix+1].Index...)
	inode.IndexNodes[ix].IndexNodes = append(inode.IndexNodes[ix].IndexNodes, inode.IndexNodes[ix+1].IndexNodes...)

	// 🖍️ At first, the original data is located at index ix. (原始资料在 ix)
	// Next, the original data will be merged into the neighbor node on the right, shifting the original data to position ix+1. (原始资料合拼到 ix+1)
	// Then, the index node at position ix will be erased, and the original data returns to position ix. (抹除 ix 节点，原始资料又回到 ix)
	// 再来，原始资料会先合并到右方的邻居节点，原始资料移动到位置 ix+1
	// 之后，再抹除 ix 位置上的索引节点，原始料料又回到位置 ix
	inode.Index = append(inode.Index[:ix], inode.Index[ix+1:]...)
	inode.IndexNodes = append(inode.IndexNodes[:ix+1], inode.IndexNodes[ix+2:]...)
	return
}
