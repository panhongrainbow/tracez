package bpTree

import (
	"fmt"
	"sort"
)

func (inode *BpIndex) deleteBottomItem(item BpItem) (deleted bool, direction int, ix int) {
	// Use binary search to find the index(i) where the key should be inserted.
	ix = sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] >= item.Key
	})

	deleted, direction = inode.DataNodes[ix].delete(item)

	if deleted == true {
		if direction == deleteRightOne {
			ix = ix + 1
		}

		if direction == deleteLeftOne {
			ix = ix - 1
		}
	}

	return
}

func (inode *BpIndex) deleteItem2(newNode *BpIndex, item BpItem) (popIx int, popKey int64, popNode *BpIndex, status int, err error) {

	// var newIndex int64
	// var sideDataNode *BpData

	// >>>>> 进入索引结点

	// If there are existing items, insert the new item among them.
	if newNode == nil && len(inode.Index) > 0 {
		// (当索引大于 0，就可以直接开始找位置)

		// Use binary search to find the index(i) where the key should be inserted.
		ix := sort.Search(len(inode.Index), func(i int) bool {
			return inode.Index[i] >= item.Key
		})

		// >>>>> >>>>> >>>>> 进入递归

		if len(inode.IndexNodes) > 0 {
			if len(inode.IndexNodes) != (len(inode.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", inode.Index)
				return
			}

			// If there are index nodes, recursively insert the item into the appropriate node.
			// (这里有递回去找到接近资料切片的地方)
			popIx, popKey, popNode, status, err = inode.IndexNodes[ix].deleteItem2(nil, item)

			if status == statusDeProtrude && ix == 0 {
				node := &BpIndex{}
				node.Index = append(node.Index, inode.Index[0])
				inode.Index = removeElement(inode.Index, inode.Index[0])

				node.Index = append(node.Index, inode.IndexNodes[ix].Index...)
				node.Index = append(node.Index, inode.IndexNodes[ix+1].Index...)

				node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[ix].IndexNodes...)
				node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[ix+1].IndexNodes...)

				node.DataNodes = append(node.DataNodes, inode.IndexNodes[ix].DataNodes...)
				node.DataNodes = append(node.DataNodes, inode.IndexNodes[ix+1].DataNodes...)

				if len(node.Index) >= BpWidth {
					popNode, _ = node.protrudeInOddBpWidth()
					status = statusDeleteProtrude
					return
				} else {
					inode.IndexNodes[ix+1] = node

					inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
				}

				status = 0
			}

			if status == statusDeleteProtrude {
				inode.IndexNodes[ix] = popNode
				status = 0
				return
			}

			if status == statusDeProtrude && (ix-1) >= 0 {
				//
				inode.IndexNodes[ix].Index = []int64{inode.Index[ix-1]}

				inode.Index = removeElement(inode.Index, inode.Index[ix-1])

				// indexToCheck >= 0 && indexToCheck < len(originalSlice)

				if (ix+1) >= 0 && (ix+1) < len(inode.IndexNodes) {
					// if (len(inode.IndexNodes) - 1) <= (ix + 1) { // 优先向 右节点合并
					// if inode.IndexNodes[ix+1].dNodesLength() < BpHalfWidth {
					node := &BpIndex{}
					node.Index = append(node.Index, inode.IndexNodes[ix].Index...)
					node.Index = append(node.Index, inode.IndexNodes[ix+1].Index...)

					node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[ix].IndexNodes...)
					node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[ix+1].IndexNodes...)

					node.DataNodes = append(node.DataNodes, inode.IndexNodes[ix].DataNodes...)
					node.DataNodes = append(node.DataNodes, inode.IndexNodes[ix+1].DataNodes...)

					inode.IndexNodes[ix+1] = node

					inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
					// }
				} else if (ix - 1) >= 0 {
					node := &BpIndex{}
					node.Index = append(node.Index, inode.IndexNodes[ix-1].Index...)
					node.Index = append(node.Index, inode.IndexNodes[ix].Index...)

					node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[ix-1].IndexNodes...)
					node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[ix].IndexNodes...)

					node.DataNodes = append(node.DataNodes, inode.IndexNodes[ix-1].DataNodes...)
					node.DataNodes = append(node.DataNodes, inode.IndexNodes[ix].DataNodes...)

					inode.IndexNodes[ix-1] = node

					inode.IndexNodes = append(inode.IndexNodes[:ix], inode.IndexNodes[ix+1:]...)
				}

				status = 0
			}

			/*
				// 先注解 ... ... ... ... ... ... ...
				status = status_protrude_inode
				if popKey != 0 {
					err = inode.mergePopIx(ix, popKey, popNode)
					popKey = 0
					popNode = nil
				}

				if popNode != nil && popKey == 0 {
					// >>>>> >>>>> >>>>> index 分裂

					inode.insertBpIX(popNode.Index[0])
					left := inode.IndexNodes[:ix]
					right := inode.IndexNodes[ix+1:]
					node := &BpIndex{}
					node.IndexNodes = append(node.IndexNodes, left...)
					node.IndexNodes = append(node.IndexNodes, popNode.IndexNodes...)
					node.IndexNodes = append(node.IndexNodes, right...)
					inode.IndexNodes = node.IndexNodes

					popNode = nil
				}

				if len(inode.Index) >= BpWidth && len(inode.Index)%2 != 0 { // 进行 pop 和奇数
					popNode, err = inode.protrudeInOddBpWidth()
					return
				}
			*/

			return
		}

		// If there are data nodes, insert the new item at the determined index.
		if len(inode.DataNodes) > 0 {
			if len(inode.DataNodes) != (len(inode.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", inode.Index)
				return
			}

			// >>>>> 进入第 1 个资料结点入口

			deleted, _ := inode.DataNodes[ix]._delete(item) // Insert item at index ix.

			if deleted == true {
				status = statusDeleteItem
			}

			if deleted == false && len(inode.DataNodes) >= ix+1+1 {
				deleted, _ = inode.DataNodes[ix+1]._delete(item)
				if deleted == true {
					status = statusDeleteItem
					// return
				}
			}

			if deleted == false {
				deleted, _ = inode.DataNodes[ix].Next._delete(item)
				status = statusDeleteItem
				// return
			}

			if deleted == false {
				status = statusDeleteNon
				// return // 检查一下好了
			}

			if inode.DataNodes[ix].dataLength() == 0 {
				// (1) 先向其他结点求援

				// (2) 如果 IX 为 0，就删第一个索引和第一个元素

				/*if ix == 0 && len(inode.Index) == 0 {
					// 索引长度为零，会被合拼
				}*/

				if ix == 0 && len(inode.Index) == 1 {
					inode.Index = []int64{}
					inode.DataNodes = inode.DataNodes[1:]
				}

				if ix == 0 && len(inode.Index) >= 2 {
					inode.Index = inode.Index[1:]
					inode.DataNodes = inode.DataNodes[1:]
				}

				// (3) 如果 IX 不为 0, 进行以下处理
				if ix != 0 {
					inode.Index = append(inode.Index[:ix-1], inode.Index[ix:]...)             // 这会有问题
					inode.DataNodes = append(inode.DataNodes[:ix], inode.DataNodes[ix+1:]...) // 这比较不会出错
				}
			}

			if len(inode.Index) == 0 {
				status = statusDeProtrude
				return
			}

			/*
				// 先注解 ... ... ... ... ... ... ...
				if len(inode.DataNodes[ix].Items) >= BpWidth {
					sideDataNode, err = inode.DataNodes[ix].split()
					if err != nil {
						return
					}

					inode.DataNodes = append(inode.DataNodes, &BpData{})
					copy(inode.DataNodes[(ix+1)+1:], inode.DataNodes[(ix+1):])
					inode.DataNodes[ix+1] = sideDataNode

					err = inode.insertBpIX(sideDataNode.Items[0].Key)
					if err != nil {
						return
					}
				}

				if len(inode.Index) >= BpWidth {
					popKey, popNode, err = inode.splitWithDnode()
					status = status_protrude_dnode
					popIx = ix
					if err != nil {
						return
					}
				}
			*/

			return
		}
	}

	// >>>>> 进入第 2 个资料结点入口

	// The length of idx.Index is 0, which only occurs in one scenario where there is only one DataNodesDataNodes.
	// (Idx.Index 的长度为 0，只有在一个状况才会发生，资料分片只有一份)
	if newNode == nil && len(inode.Index) == 0 {
		if len(inode.DataNodes) != 1 {
			// 资料大于1，就会有索引，就不会进入这里
			err = fmt.Errorf("the number of indexes is incorrect initially")
			return
		}
		inode.DataNodes[0]._delete(item) // >>>>> (add to DataNodes)

		/*
			// 先注解 ... ... ... ... ... ... ...
			if inode.DataNodes[0].dataLength() >= BpWidth {
				sideDataNode, err = inode.DataNodes[0].split() // newIndex
				if err != nil {
					return
				}

				inode.DataNodes = append(inode.DataNodes, sideDataNode)
				newIndex = sideDataNode.Items[0].Key
			}
		*/
	}

	/*
		// 先注解 ... ... ... ... ... ... ...
		if sideDataNode != nil {
			err = inode.insertBpIX(newIndex)
			if err != nil {
				return
			}

			if len(inode.Index) >= BpWidth && len(inode.Index)%2 != 0 { // 进行 pop 和奇数 (可能没在使用)
				var node *BpIndex
				node, err = inode.protrudeInOddBpWidth()
				*inode = *node
				return
			}

			return
		}
	*/

	return
}

func removeElement(nums []int64, val int64) []int64 {
	var result []int64

	for _, num := range nums {
		if num != val {
			result = append(result, num)
		}
	}

	return result
}

func (inode *BpIndex) deProtrude(popMiddleNode *BpIndex) (err error) {
	for i := 0; i < len(popMiddleNode.IndexNodes); i++ {
		inode.Index = append(inode.Index, popMiddleNode.IndexNodes[i].Index...)
		inode.IndexNodes = append(inode.IndexNodes, popMiddleNode.IndexNodes[i].IndexNodes...)
	}

	for i := 0; i < len(popMiddleNode.Index); i++ {
		inode.Index = insertSorted(inode.Index, popMiddleNode.Index[i])
	}

	return
}

func insertSorted(arr []int64, target int64) []int64 {
	insertIndex := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	arr = append(arr[:insertIndex], append([]int64{target}, arr[insertIndex:]...)...)

	return arr
}
