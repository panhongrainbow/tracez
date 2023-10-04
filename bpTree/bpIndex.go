package bpTree

import (
	"fmt"
	"sort"
)

// >>>>> >>>>> >>>>> basic structure

// BpIndex is the index of the B plus tree.
type BpIndex struct {
	Index      []int64    // The maximum values of each group of BpData
	IndexNodes []*BpIndex // Index nodes
	DataNodes  []*BpData  // Data nodes
	Split      bool       // After splitting the nodes, mark it.
}

// >>>>> >>>>> >>>>> get method

// indexLength returns the length of index slice. (求索引结点的长度)
func (inode *BpIndex) indexLength() (length int) {
	length = len(inode.Index)
	return
}

// iNodesLength returns the length of BpIndex Node slice. (求索引结点的长度)
func (inode *BpIndex) iNodesLength() (length int) {
	length = len(inode.IndexNodes)
	return
}

// dNodesLength returns the length of BpData Node slice. (求资料结点的长度)
func (inode *BpIndex) dNodesLength() (length int) {
	length = len(inode.DataNodes)
	return
}

// >>>>> >>>>> >>>>> insert method

// insertBpIX inserts a new index at the correct position using binary search.
func (inode *BpIndex) insertBpIX(newIx int64) (err error) {
	// Use binary search to find the position where the index should be inserted.
	ix := sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] >= newIx
	})

	// Expand the slice to accommodate the new item.
	inode.Index = append(inode.Index, 0)

	// Shift the elements to the right to make space for the new item.
	copy(inode.Index[ix+1:], inode.Index[ix:])

	// Insert the new item at the correct position.
	inode.Index[ix] = newIx

	return
}

// protrude performs index upgrade; when the middle value of the index slice pops out, it gets upgraded to the upper-level index.
// (进行索引升级，当索引切片的中间值会弹出升级成上层的 Index)
func (inode *BpIndex) protrude() (popMiddleNode *BpIndex, err error) {
	// Calculate the current index lengths for splitting.
	indexLen := (len(inode.Index) - 1) / 2
	indexNodeLen := len(inode.IndexNodes) / 2
	dataNodeLen := len(inode.DataNodes) / 2

	// Pop operation

	// Create a new left node.
	leftNode := &BpIndex{}
	leftNode.Index = append(leftNode.Index, inode.Index[:indexLen]...)
	leftNode.IndexNodes = append(leftNode.IndexNodes, inode.IndexNodes[:indexNodeLen]...)
	leftNode.DataNodes = append(leftNode.DataNodes, inode.DataNodes[:dataNodeLen]...)

	// Create a new right node.
	rightNode := &BpIndex{}
	rightNode.Index = append(rightNode.Index, inode.Index[indexLen+1:]...)
	rightNode.IndexNodes = append(rightNode.IndexNodes, inode.IndexNodes[indexNodeLen:]...)
	rightNode.DataNodes = append(rightNode.DataNodes, inode.DataNodes[dataNodeLen:]...)

	// Create a new middle node.
	middleValue := inode.Index[indexLen : indexLen+1]
	popMiddleNode = &BpIndex{
		Index:      middleValue,
		IndexNodes: []*BpIndex{leftNode, rightNode},
		DataNodes:  []*BpData{},
	}

	// Make a mark, already split.
	// inode.Split = true // 不需要

	return
}

func (inode *BpIndex) protrude2() (popMiddleNode *BpIndex, err error) {
	// Calculate the current index lengths for splitting.
	indexLen := (len(inode.Index)) / 2
	indexNodeLen := (len(inode.IndexNodes) - 1) / 2
	dataNodeLen := (len(inode.DataNodes) - 1) / 2

	// Create a new left node.
	leftNode := &BpIndex{}
	leftNode.Index = append(leftNode.Index, inode.Index[:indexLen]...)
	if indexNodeLen > 0 {
		leftNode.IndexNodes = append(leftNode.IndexNodes, inode.IndexNodes[:indexNodeLen+1]...)
	}
	if dataNodeLen > 0 {
		leftNode.DataNodes = append(leftNode.DataNodes, inode.DataNodes[:dataNodeLen+1]...)
	}

	// Create a new right node.
	rightNode := &BpIndex{}
	rightNode.Index = append(rightNode.Index, inode.Index[indexLen+1:]...)
	if indexNodeLen > 0 {
		rightNode.IndexNodes = append(rightNode.IndexNodes, inode.IndexNodes[indexNodeLen+1:]...)
	}
	if dataNodeLen > 0 {
		rightNode.DataNodes = append(rightNode.DataNodes, inode.DataNodes[dataNodeLen+1:]...)
	}

	// Create a new middle node.
	middleValue := inode.Index[indexLen : indexLen+1]
	popMiddleNode = &BpIndex{
		Index:      middleValue,
		IndexNodes: []*BpIndex{leftNode, rightNode},
		DataNodes:  []*BpData{},
	}

	// Make a mark, already split.
	// inode.Split = true // 不需要

	return
}

func (inode *BpIndex) digDigKey() (key int64) {
	node := inode
	for {
		if len(node.DataNodes) == 0 {
			node = node.IndexNodes[0]
		} else {
			key = node.DataNodes[0].Items[0].Key
			break
		}
	}
	return
}

func (inode *BpIndex) mergePopDnode(side *BpData) (err error) {
	var newIx int64

	if len(inode.IndexNodes) > 0 {
		// Cannot directly insert a pure index node.
		// When it is popped up, it is merged directly by the parent node.
		// (POP上来直接被父节点合拼)
		return fmt.Errorf("data cannot be inserted directly into wrong index nodes")
	}

	if len(inode.DataNodes) > 0 {
		// Use binary search to find the position where the index should be inserted.
		ix := sort.Search(len(inode.DataNodes), func(i int) bool {
			return inode.DataNodes[i].Items[0].Key >= side.Items[0].Key
		})

		for i := ix; i < len(inode.DataNodes); i++ {
			if inode.DataNodes[ix].Split == true {
				ix = ix + 1
				break
			}
		}

		inode.DataNodes[ix].Split = false

		// >>>>> 失处理节点

		// Expand the slice to accommodate the new item.
		inode.DataNodes = append(inode.DataNodes, &BpData{})

		// Shift the elements to the right to make space for the new item.
		copy(inode.DataNodes[ix+1:], inode.DataNodes[ix:])

		// Insert the new item at the correct position.
		inode.DataNodes[ix] = side

		// >>>>> 再处理索引

		newIx, err = side.index()
		err = inode.insertBpIX(newIx)
	}

	if len(inode.DataNodes) == 0 {
		inode.DataNodes = append(inode.DataNodes, side)
	}

	return
}

const (
	status_protrude_inode = iota + 1
	status_protrude_dnode
	status_delete_item
	status_delete_Non
	status_de_protrude
	status_delete_protrude
)

// insertBpDataValue inserts a new index into the BpIndex.
// 经由 BpIndex 直接在新增
func (inode *BpIndex) insertItem(newNode *BpIndex, item BpItem) (popIx int, popKey int64, popNode *BpIndex, status int, err error) {

	var newIndex int64
	var sideDataNode *BpData

	// >>>>> 进入索引结点

	// If there are existing items, insert the new item among them.
	if newNode == nil && len(inode.Index) > 0 {
		// (当索引大于 0，就可以直接开始找位置)

		// Use binary search to find the index(i) where the key should be inserted.
		ix := sort.Search(len(inode.Index), func(i int) bool {
			return inode.Index[i] >= item.Key
		})

		// >>>>> >>>>> >>>>> 进入递归

		// Verify if the index for IndexNodes is correct?
		// (先检查索吊数量是否正确)

		if len(inode.IndexNodes) > 0 {
			if len(inode.IndexNodes) != (len(inode.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", inode.Index)
				return
			}

			// If there are index nodes, recursively insert the item into the appropriate node.
			// (这里有递回去找到接近资料切片的地方)
			popIx, popKey, popNode, status, err = inode.IndexNodes[ix].insertItem(nil, item)
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
				popNode, err = inode.protrude()
				return
			} else if len(inode.Index) >= BpWidth && len(inode.Index)%2 == 0 { // 进行 pop 和奇数
				popNode, err = inode.protrude2()
				return
			}

			return // Break here to avoid inserting elsewhere. (立刻中断)
		}

		// If there are data nodes, insert the new item at the determined index.
		if len(inode.DataNodes) > 0 {
			if len(inode.DataNodes) != (len(inode.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", inode.Index)
				return
			}

			// >>>>> 进入第 1 个资料结点入口

			inode.DataNodes[ix].insert(item) // Insert item at index ix.

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
		inode.DataNodes[0].insert(item) // >>>>> (add to DataNodes)

		if inode.DataNodes[0].dataLength() >= BpWidth {
			sideDataNode, err = inode.DataNodes[0].split() // newIndex
			if err != nil {
				return
			}

			inode.DataNodes = append(inode.DataNodes, sideDataNode)
			newIndex = sideDataNode.Items[0].Key
		}
	}

	if sideDataNode != nil {
		err = inode.insertBpIX(newIndex)
		if err != nil {
			return
		}

		if len(inode.Index) >= BpWidth && len(inode.Index)%2 != 0 { // 进行 pop 和奇数 (可能没在使用)
			var node *BpIndex
			node, err = inode.protrude()
			*inode = *node
			return
		} else if len(inode.Index) >= BpWidth && len(inode.Index)%2 == 0 { // 进行 pop 和奇数 (可能没在使用)
			var node *BpIndex
			node, err = inode.protrude2()
			*inode = *node
			return
		}

		return // Break here to avoid inserting elsewhere. (立刻中断)
	}

	return
}

// >>>>> >>>>> >>>>> split and maintain

func (inode *BpIndex) splitWithDnode() (key int64, side *BpIndex, err error) {
	// Check if both IndexNodes and DataNodes have data, which is incorrect as we don't know where to retrieve the index.
	if (len(inode.IndexNodes) != 0) && (len(inode.DataNodes) != 0) {
		err = fmt.Errorf("both IndexNodes and DataNodes have data, cannot determine which one is the index source")
		return
	}

	/*if len(inode.IndexNodes) != 0 {
	// 这不要考虑，因为一定是下层的 iNode 的 index 过大
	}*/

	if len(inode.DataNodes) != 0 {
		side = &BpIndex{}
		length := len(inode.DataNodes)
		side.Index = append(side.Index, inode.Index[(length-BpHalfWidth):]...)
		side.DataNodes = append(side.DataNodes, inode.DataNodes[(length-BpHalfWidth):]...)

		key = inode.Index[(length - BpHalfWidth - 1):(length - BpHalfWidth)][0]

		inode.Index = inode.Index[0:(length - BpHalfWidth - 1)] // 减一为要少一个数量
		inode.DataNodes = inode.DataNodes[0:(length - BpHalfWidth)]
	}

	return
}

// >>>>> >>>>> >>>>> compare and merge

func (inode *BpIndex) TakeApartReassemble(indexes ...*BpIndex) {
	//
	for _, v := range indexes {
		if len(v.Index) == 0 {
			if len(v.IndexNodes) > 0 {
				v.Index = insertAtFront(v.Index, v.IndexNodes[0].Index[0])
			} else if len(v.DataNodes) > 0 {
				v.Index = insertAtFront(v.Index, v.DataNodes[0].Items[0].Key)
			}
		}
	}

	//
	sort.SliceStable(indexes, func(i, j int) bool {
		return (*indexes[i]).Index[0] < (*indexes[j]).Index[0]
	})

	for _, v := range indexes {
		if len(v.IndexNodes) > 0 && len(v.Index) != len(v.IndexNodes) {
			v.Index = insertAtFront(v.Index, v.IndexNodes[0].Index[0])
		}
		if len(v.DataNodes) > 0 && len(v.Index) != len(v.DataNodes) {
			v.Index = insertAtFront(v.Index, v.DataNodes[0].Items[0].Key)
		}
	}

	//
	inode.Index = []int64{}
	inode.IndexNodes = []*BpIndex{}
	inode.DataNodes = []*BpData{}

	//
	for _, v := range indexes {
		inode.Index = append(inode.Index, v.Index...)
		inode.IndexNodes = append(inode.IndexNodes, v.IndexNodes...)
		inode.DataNodes = append(inode.DataNodes, v.DataNodes...)
	}

	//
	inode.Index = inode.Index[1:]

	return
}

func (inode *BpIndex) prepareProtrudeDnode(podIx int, podKey int64, indexes ...*BpIndex) {
	newTree := &BpIndex{}
	newTree.Index = []int64{podKey}

	for _, v := range indexes {
		var tmp = &BpIndex{}
		tmp.Index = append(tmp.Index, v.Index...)
		tmp.IndexNodes = append(tmp.IndexNodes, v.IndexNodes...)
		tmp.DataNodes = append(tmp.DataNodes, v.DataNodes...)
		newTree.IndexNodes = append(newTree.IndexNodes, tmp)
	}

	*inode = *newTree

	return
}

func (inode *BpIndex) prepareProtrudeDnode2(podIx int, podKey int64, indexes ...*BpIndex) {
	//
	newTree := &BpIndex{}
	newTree.Index = []int64{podKey}

	var tmp = &BpIndex{}

	for _, v := range indexes {
		tmp.Index = append(tmp.Index, v.Index...)
		// tmp.IndexNodes = append(tmp.IndexNodes, v.IndexNodes...)
		tmp.DataNodes = append(tmp.DataNodes, v.DataNodes[0:podIx]...)
		tmp.DataNodes = append(tmp.DataNodes, v.DataNodes[0:podIx]...)
		newTree.IndexNodes = append(newTree.IndexNodes, tmp)
	}

	*inode = *newTree

	return
}

func (inode *BpIndex) mergePopIx(ix int, popKey int64, indexNode *BpIndex) (err error) {
	ix = ix + 1

	err = inode.insertBpIX(popKey)

	inode.IndexNodes = append(inode.IndexNodes, &BpIndex{})
	copy(inode.IndexNodes[ix+1:], inode.IndexNodes[ix:])
	inode.IndexNodes[ix] = indexNode

	return
}

func insertAtFront(slice []int64, newElement int64) []int64 {
	// 创建一个新切片，长度比原切片多1
	newSlice := make([]int64, len(slice)+1)

	// 将新元素放在新切片的第一个位置
	newSlice[0] = newElement

	// 将原切片的所有元素追加到新切片后面
	copy(newSlice[1:], slice)

	return newSlice
}
