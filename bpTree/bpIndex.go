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
}

// >>>>> >>>>> >>>>> get length and index

// getBpIdxIndex retrieves the key from the BpIndex structure. (求 索引节点的 索引)
// If the Index slice is empty, it attempts to retrieve the key from the associated DataNode.
func (idx *BpIndex) getBpIdxIndex() (key int64, err error) {
	// Check if the Index slice has values.
	if len(idx.Index) > 0 {
		key = idx.Index[0]
	}

	// If there is no index in the BpIndex, set an error indicating no key.
	if len(idx.Index) == 0 {
		err = fmt.Errorf("no key available")
	}

	return
}

// getBpDataLength returns the length of BpData's items slice. (求索引结点的长度)
func (idx *BpIndex) getBpIndexNodesLength() (length int) {
	length = len(idx.IndexNodes)
	return
}

func (idx *BpIndex) getBpDataNodesLength() (length int) {
	length = len(idx.DataNodes)
	return
}

// >>>>> >>>>> >>>>> insert index, indexNode, dataNode

// insertBpIdxNewIndex inserts a new index at the correct position using binary search.
func (idx *BpIndex) insertBpIdxNewIndex(newIx int64) (err error) {
	// Use binary search to find the position where the index should be inserted.
	ix := sort.Search(len(idx.Index), func(i int) bool {
		return idx.Index[i] >= newIx
	})

	// Expand the slice to accommodate the new item.
	idx.Index = append(idx.Index, 0)

	// Shift the elements to the right to make space for the new item.
	copy(idx.Index[ix+1:], idx.Index[ix:])

	// Insert the new item at the correct position.
	idx.Index[ix] = newIx

	return
}

// insertBpIdxNewIndexNode inserts a new index node at the correct position using binary search.
func (idx *BpIndex) insertBpIdxNewIndexNode(newIdx *BpIndex) (err error) {
	// Use binary search to find the position where the index should be inserted.
	ix := sort.Search(len(idx.Index), func(i int) bool {
		return idx.Index[i] >= newIdx.Index[0]
	})

	// >>>>> 失处理节点

	// Expand the slice to accommodate the new item.
	idx.IndexNodes = append(idx.IndexNodes, &BpIndex{})

	// Shift the elements to the right to make space for the new item.
	copy(idx.IndexNodes[ix+1:], idx.IndexNodes[ix:])

	// Insert the new item at the correct position.
	idx.IndexNodes[ix] = newIdx.IndexNodes[0]

	// >>>>> 再处理索引

	// Expand the slice to accommodate the new item.
	idx.Index = append(idx.Index, 0)

	// Shift the elements to the right to make space for the new item.
	copy(idx.Index[ix+1:], idx.Index[ix:])

	// Insert the new item at the correct position.
	idx.Index[ix] = newIdx.Index[0]

	return
}

func (idx *BpIndex) insertBpIdxNewDataNode(sideNode *BpData) (err error) {

	if len(idx.DataNodes) > 0 {
		// Use binary search to find the position where the index should be inserted.
		ix := sort.Search(len(idx.Index), func(i int) bool {
			return idx.Index[i] >= sideNode.Items[0].Key
		})

		// >>>>> 失处理节点

		// Expand the slice to accommodate the new item.
		idx.DataNodes = append(idx.DataNodes, &BpData{})

		// Shift the elements to the right to make space for the new item.
		copy(idx.DataNodes[ix+1:], idx.DataNodes[ix:])

		// Insert the new item at the correct position.
		idx.DataNodes[ix] = sideNode

		// >>>>> 再处理索引

		// Expand the slice to accommodate the new item.
		idx.Index = append(idx.Index, 0)

		// Shift the elements to the right to make space for the new item.
		copy(idx.Index[ix+1:], idx.Index[ix:])

		// Insert the new item at the correct position.
		idx.Index[ix] = sideNode.Items[0].Key
	}

	// 才一个资料切片，不会有索引
	if len(idx.DataNodes) == 0 {
		idx.DataNodes = append(idx.DataNodes, sideNode)
	}

	return
}

// insertBpDataValue inserts a new index into the BpIndex.
// 经由 BpIndex 直接在新增
func (idx *BpIndex) insertBpIdxNewValue(newNode *BpIndex, item BpItem) (popKey int64, popNode *BpIndex, err error) {

	var newIndex int64
	var sideDataNode *BpData

	// If there are existing items, insert the new item among them.
	if newNode == nil && len(idx.Index) > 0 {
		// (当索引大于 0，就可以直接开始找位置)

		// Use binary search to find the index(i) where the key should be inserted.
		ix := sort.Search(len(idx.Index), func(i int) bool {
			return idx.Index[i] >= item.Key
		})

		// >>>>> >>>>> >>>>> 进入递归

		// Verify if the index for IndexNodes is correct?
		// (先检查索吊数量是否正确)

		if len(idx.IndexNodes) > 0 {
			if len(idx.IndexNodes) != (len(idx.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", idx.Index)
				return
			}

			// If there are index nodes, recursively insert the item into the appropriate node.
			// (这里有递回去找到接近资料切片的地方)
			popKey, popNode, err = idx.IndexNodes[ix].insertBpIdxNewValue(nil, item)
			// >>>>>>>>>>>>>> XXXXXXXXXXX
			if popKey != 0 {
				err = idx.cmpAndCombineIndexNode(popKey, popNode)
				popKey = 0
				popNode = nil
			}

			if len(idx.Index) >= BpWidth && len(idx.Index)%2 != 0 { // 进行 pop 和奇数
				indexLen := (len(idx.Index) - 1) / 2
				indexNodeLen := len(idx.IndexNodes) / 2
				dataNodeLen := len(idx.DataNodes) / 2

				leftNode := &BpIndex{}
				leftNode.Index = append(leftNode.Index, idx.Index[:indexLen]...)
				leftNode.IndexNodes = append(leftNode.IndexNodes, idx.IndexNodes[:indexNodeLen]...)
				leftNode.DataNodes = append(leftNode.DataNodes, idx.DataNodes[:dataNodeLen]...)

				rightNode := &BpIndex{}
				rightNode.Index = append(rightNode.Index, idx.Index[indexLen+1:]...)
				rightNode.IndexNodes = append(rightNode.IndexNodes, idx.IndexNodes[indexNodeLen:]...)
				rightNode.DataNodes = append(rightNode.DataNodes, idx.DataNodes[dataNodeLen:]...)

				middleValue := idx.Index[indexLen : indexLen+1]
				middleNode := &BpIndex{
					Index:      middleValue,
					IndexNodes: []*BpIndex{leftNode, rightNode},
					DataNodes:  []*BpData{},
				}

				*idx = *middleNode

				return
			}

			return // Break here to avoid inserting elsewhere. (立刻中断)
		}

		// If there are data nodes, insert the new item at the determined index.
		if len(idx.DataNodes) > 0 {
			// Verify if the index for IndexNodes is correct?
			if len(idx.DataNodes) != (len(idx.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", idx.Index)
				return
			}

			// >>>>> >>>>> >>>>> 进入底层，新增资料
			idx.DataNodes[ix].insertBpDataValue(item) // Insert item at index ix.

			if len(idx.DataNodes[ix].Items) >= BpWidth {
				sideDataNode, err = idx.DataNodes[ix].split()
				if err != nil {
					return
				}

				idx.DataNodes = append(idx.DataNodes, &BpData{})
				copy(idx.DataNodes[(ix+1)+1:], idx.DataNodes[(ix+1):])
				idx.DataNodes[ix+1] = sideDataNode

				err = idx.insertBpIdxNewIndex(sideDataNode.Items[0].Key)
				if err != nil {
					return
				}
			}

			if len(idx.Index) >= BpWidth {
				popKey, popNode, err = idx.basicSplit()
				if err != nil {
					return
				}
			}

			return
		}
	}

	// The length of idx.Index is 0, which only occurs in one scenario where there is only one DataNodesDataNodes.
	// (Idx.Index 的长度为 0，只有在一个状况才会发生，资料分片只有一份)
	if newNode == nil && len(idx.Index) == 0 {
		if len(idx.DataNodes) != 1 {
			err = fmt.Errorf("the number of indexes is incorrect initially")
			return
		}
		idx.DataNodes[0].insertBpDataValue(item) // >>>>> (add to DataNodes)

		if idx.DataNodes[0].getBpDataLength() >= BpWidth {
			sideDataNode, err = idx.DataNodes[0].split() // newIndex
			if err != nil {
				return
			}

			idx.DataNodes = append(idx.DataNodes, sideDataNode)
			newIndex = sideDataNode.Items[0].Key
		}
	}

	if sideDataNode != nil {
		err = idx.insertBpIdxNewIndex(newIndex)
		if err != nil {
			return
		}

		if idx.getBpIndexNodesLength() >= BpWidth {
			// node, err = idx.split(BpWidth) // 这个新节点要由上层去处理
			_, popNode, err = idx.basicSplit()
		}

		if len(idx.Index) >= BpWidth && len(idx.Index)%2 != 0 { // 进行 pop 和奇数
			indexLen := (len(idx.Index) - 1) / 2
			dataLen := len(idx.DataNodes) / 2

			leftNode := &BpIndex{}
			leftNode.Index = append(leftNode.Index, idx.Index[:indexLen]...)
			leftNode.DataNodes = append(leftNode.DataNodes, idx.DataNodes[:dataLen]...)

			rightNode := &BpIndex{}
			rightNode.Index = append(rightNode.Index, idx.Index[indexLen+1:]...)
			rightNode.DataNodes = append(rightNode.DataNodes, idx.DataNodes[dataLen:]...)

			middleValue := idx.Index[indexLen : indexLen+1]
			middleNode := &BpIndex{
				Index:      middleValue,
				IndexNodes: []*BpIndex{leftNode, rightNode},
				DataNodes:  []*BpData{},
			}

			*idx = *middleNode

			return
		}

		return // Break here to avoid inserting elsewhere. (立刻中断)
	}

	if newNode != nil {
		err = idx.insertBpIdxNewIndexNode(newNode)
		if err != nil {
			return
		}

		if idx.getBpIndexNodesLength() >= BpWidth {
			popKey, popNode, err = idx.basicSplit() // 这个新节点要由上层去处理
		}
	}

	return
}

// >>>>> >>>>> >>>>> split and maintain

// 每次切开会有一个 Key 弹出
func (idx *BpIndex) basicSplit() (key int64, side *BpIndex, err error) {
	BpHalfWidth = 2

	// Check if both IndexNodes and DataNodes have data, which is incorrect as we don't know where to retrieve the index.
	if (len(idx.IndexNodes) != 0) && (len(idx.DataNodes) != 0) {
		err = fmt.Errorf("both IndexNodes and DataNodes have data, cannot determine which one is the index source")
		return
	}

	if len(idx.IndexNodes) != 0 {
		side = &BpIndex{}
		side.Index = append(side.Index, idx.Index[BpHalfWidth:]...)
		side.IndexNodes = append(side.IndexNodes, idx.IndexNodes[BpHalfWidth:]...)

		key = idx.Index[BpHalfWidth-1 : BpHalfWidth][0]

		idx.Index = idx.Index[0 : BpHalfWidth-1] // 减一为要少一个数量
		idx.IndexNodes = idx.IndexNodes[0:BpHalfWidth]
	}

	if len(idx.DataNodes) != 0 {
		side = &BpIndex{}
		length := len(idx.DataNodes)
		side.Index = append(side.Index, idx.Index[(length-BpHalfWidth):]...)
		side.DataNodes = append(side.DataNodes, idx.DataNodes[(length-BpHalfWidth):]...)

		key = idx.Index[(length - BpHalfWidth - 1):(length - BpHalfWidth)][0]

		idx.Index = idx.Index[0:(length - BpHalfWidth - 1)] // 减一为要少一个数量
		idx.DataNodes = idx.DataNodes[0:(length - BpHalfWidth)]
	}

	return
}

// >>>>> >>>>> >>>>> compare and merge

func (idx *BpIndex) cmpAndMergeIndexNode(indexes ...*BpIndex) {
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
	idx.Index = []int64{}
	idx.IndexNodes = []*BpIndex{}
	idx.DataNodes = []*BpData{}

	//
	for _, v := range indexes {
		idx.Index = append(idx.Index, v.Index...)
		idx.IndexNodes = append(idx.IndexNodes, v.IndexNodes...)
		idx.DataNodes = append(idx.DataNodes, v.DataNodes...)
	}

	//
	idx.Index = idx.Index[1:]

	return
}

func (idx *BpIndex) cmpAndOrganizeIndexNode(podKey int64, indexes ...*BpIndex) {
	//
	sort.SliceStable(indexes, func(i, j int) bool {
		return (*indexes[i]).Index[0] < (*indexes[j]).Index[0]
	})

	//
	newTree := &BpIndex{}
	newTree.Index = []int64{podKey}

	//
	for _, v := range indexes {
		var tmp = &BpIndex{}
		tmp.Index = append(tmp.Index, v.Index...)
		tmp.IndexNodes = append(tmp.IndexNodes, v.IndexNodes...)
		tmp.DataNodes = append(tmp.DataNodes, v.DataNodes...)
		newTree.IndexNodes = append(newTree.IndexNodes, tmp)
	}

	*idx = *newTree

	return
}

func (idx *BpIndex) cmpAndCombineIndexNode(popKey int64, indexNode *BpIndex) (err error) {
	//
	ix := sort.Search(len(idx.IndexNodes), func(i int) bool {
		return idx.IndexNodes[i].Index[0] >= indexNode.Index[0]
	})

	idx.IndexNodes = append(idx.IndexNodes, &BpIndex{})
	copy(idx.IndexNodes[ix+1:], idx.IndexNodes[ix:])
	idx.IndexNodes[ix] = indexNode

	//
	err = idx.insertBpIdxNewIndex(popKey)

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
