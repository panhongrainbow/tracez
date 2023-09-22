package bpTree

import (
	"fmt"
	"sort"
)

// BpIndex is the index of the B plus tree.
type BpIndex struct {
	Index      []int64    // The maximum values of each group of BpData
	IndexNodes []*BpIndex // Index nodes
	DataNodes  []*BpData  // Data nodes
}

// getBpDataLength returns the length of BpData's items slice.
func (idx *BpIndex) getBpIndexNodesLength() (length int) {
	length = len(idx.IndexNodes)
	return
}

// getBpIdxIndex retrieves the key from the BpIndex structure.
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

// checkBpIdxIndex checks and retrieves index from nodes, handle errors appropriately.
func (idx *BpIndex) checkBpIdxIndex() (err error) {
	// Check if both IndexNodes and DataNodes have data, which is incorrect as we don't know where to retrieve the index.
	if (len(idx.IndexNodes) != 0) && (len(idx.DataNodes) != 0) {
		err = fmt.Errorf("both IndexNodes and DataNodes have data, cannot determine which one is the index source")
		return
	}

	// If IndexNodes have data, retrieve index from them.
	if len(idx.IndexNodes) != 0 {
		if len(idx.IndexNodes) != (len(idx.Index) + 1) {
			idx.Index = nil
			// abandon i = 0
			for i := 1; i < len(idx.IndexNodes); i++ {
				var ix int64
				// Retrieve index from IndexNodes
				ix, err = idx.IndexNodes[i].getBpIdxIndex()
				if err != nil {
					return
				}
				idx.Index = append(idx.Index, ix)
			}
		}
	}

	// If DataNodes have data, retrieve index from them.
	if len(idx.DataNodes) != 0 {
		if len(idx.DataNodes) != (len(idx.Index) + 1) {
			idx.Index = nil
			// abandon i = 0
			for i := 1; i < len(idx.DataNodes); i++ {
				var ix int64
				// Retrieve index from DataNodes
				ix, err = idx.DataNodes[i].getBpDataIndex()
				if err != nil {
					return
				}
				idx.Index = append(idx.Index, ix)
			}
		}
	}

	return
}

// insertBpDataValue inserts a new index into the BpIndex.
// 经由 BpIndex 直接在新增
func (idx *BpIndex) insertBpIdxNewValue(newNode *BpIndex, item BpItem) (node *BpIndex, err error) {

	var newIndex int64
	var newDataNode *BpData

	// If there are existing items, insert the new item among them.
	if newNode == nil && len(idx.Index) > 0 {
		// Verify if the index for IndexNodes is correct?
		// (先检查索吊数量是否正确)
		if len(idx.IndexNodes) > 0 &&
			(len(idx.IndexNodes) != (len(idx.Index) + 1)) {
			err = fmt.Errorf("the number of indexes is incorrect, %v", idx.Index)
			return
		}

		// Verify if the index for IndexNodes is correct?
		if len(idx.DataNodes) > 0 &&
			(len(idx.DataNodes) != (len(idx.Index) + 1)) {
			err = fmt.Errorf("the number of indexes is incorrect, %v", idx.Index)
			return
		}

		// Use binary search to find the index(i) where the key should be inserted.
		ix := sort.Search(len(idx.Index), func(i int) bool {
			return idx.Index[i] >= item.Key
		})

		// If there are index nodes, recursively insert the item into the appropriate node.
		// (这里有递回去找到接近资料切片的地方)
		if len(idx.IndexNodes) > 0 {
			_, err = idx.IndexNodes[ix].insertBpIdxNewValue(nil, item)
			return // Break here to avoid inserting elsewhere. (立刻中断)
		}

		// If there are data nodes, insert the new item at the determined index.
		if len(idx.DataNodes) > 0 {
			idx.DataNodes[ix].insertBpDataValue(item) // Insert item at index ix. // >>>>> (add to DataNodes)
			if len(idx.DataNodes[ix].Items) >= BpWidth {
				newDataNode, err = idx.DataNodes[ix].split() // newIndex
				if err != nil {
					return
				}

				// idx.DataNodes = append(idx.DataNodes, newDataNode)

				// Expand the slice to accommodate the new item.
				idx.DataNodes = append(idx.DataNodes, &BpData{})

				// Shift the elements to the right to make space for the new item.
				copy(idx.DataNodes[(ix+1)+1:], idx.DataNodes[(ix+1):])

				// Insert the new item at the correct position.
				idx.DataNodes[ix+1] = newDataNode

				newIndex = newDataNode.Items[0].Key

				// DataNode 转成 IndexNode
				// ...
			}
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
			newDataNode, err = idx.DataNodes[0].split() // newIndex
			if err != nil {
				return
			}

			idx.DataNodes = append(idx.DataNodes, newDataNode)
			newIndex = newDataNode.Items[0].Key
		}
	}

	if newDataNode != nil {
		err = idx.insertBpIdxNewIndex(newIndex)
		if err != nil {
			return
		}

		if idx.getBpIndexNodesLength() >= BpWidth {
			node, err = idx.split(BpWidth) // 这个新节点要由上层去处理
		}

		if len(idx.Index) >= BpWidth && len(idx.Index)%2 != 0 { // 进行 pop 和奇数
			indexLen := (len(idx.Index) - 1) / 2
			dataLen := len(idx.DataNodes) / 2

			leftNode := &BpIndex{
				/*Index:      idx.Index[:indexLen],
				IndexNodes: []*BpIndex{},
				DataNodes:  idx.DataNodes[:dataLen],*/
			}
			leftNode.Index = append(leftNode.Index, idx.Index[:indexLen]...)
			leftNode.DataNodes = append(leftNode.DataNodes, idx.DataNodes[:dataLen]...)

			rightNode := &BpIndex{
				/*Index:      idx.Index[indexLen+1:],
				IndexNodes: []*BpIndex{},
				DataNodes:  idx.DataNodes[dataLen:],*/
			}
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
			node, err = idx.split(BpWidth) // 这个新节点要由上层去处理
		}
	}

	return
}

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

// split divides the BpIndex's index into two parts if it contains more items than the specified width.
func (idx *BpIndex) split(width int) (node *BpIndex, err error) {
	// Check if the number of index in the BpData is less than or equal to the specified width.
	if len(idx.IndexNodes) <= width {
		// If it's not greater than the width, return an error.
		err = fmt.Errorf("cannot split IndexNodes with less than or equal to %d items", width)
		return
	}

	// Create a new index node to store the items that will be moved.
	node = &BpIndex{}
	node.IndexNodes = append(node.IndexNodes, idx.IndexNodes[width:]...)

	// Check and repair the new index node. (对新节点进行检查和修复)
	err = node.checkBpIdxIndex()
	if err != nil {
		return
	}

	// Update the current index node to retain the first width items.
	idx.IndexNodes = idx.IndexNodes[0:width]

	// Handling the indexes of two nodes. (处理两个节点的索引)
	length := len(idx.IndexNodes) - 1
	idx.Index = idx.Index[0:length]

	err = node.checkBpIdxIndex()
	if err != nil {
		return
	}

	// No error occurred during the split, so return nil to indicate success.
	return
}
