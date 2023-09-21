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
	node.IndexNodes = idx.IndexNodes[width:]

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

// >>>>> >>>>> >>>>> >>>>> >>>>>

// BpIndex2 is the index of the B+ tree.
/*type BpIndex2 struct {
	IsLeaf     bool        // Whether it is approaching the bottom data level
	Intervals  []int64     // The maximum values of each group of BpData
	IndexNodes []*BpIndex2 // Index nodes
	DataNodes  []*BpData   // Data nodes
}*/

// NewBpIdxIndexNode creates a new index node.
/*func NewBpIdxIndexNode() (index *BpIndex2) {
	index = &BpIndex2{
		DataNodes: []*BpData{},
		IsLeaf:    false,
	}
	return
}*/

// NewBpIdxDataNode creates a new data node.
/*func NewBpIdxDataNode() (index *BpIndex2) {
	index = &BpIndex2{
		DataNodes: make([]*BpData, BpWidth),
		IsLeaf:    true,
	}
	for i := 0; i < BpWidth; i++ {
		index.DataNodes[i] = &BpData{
			Items: make([]BpItem, BpWidth),
		}
	}
	return
}*/

/*func (index *BpIndex2) insertIndexValue(item BpItem) {
	if index.IsLeaf {
		if len(index.Intervals) == 0 {
			// 插入最左邊
			index.Intervals = append(index.Intervals, item.Key)
			index.DataNodes[0].Items = append(index.DataNodes[0].Items, item)
		} else {
			index.insertExistIndexValue(item)
		}
	}
	if !index.IsLeaf {
		fmt.Println()
		idx := sort.Search(len(index.Intervals), func(i int) bool {
			return index.Intervals[i] >= item.Key
		})
		index.IndexNodes[idx].insertIndexValue(item)
	}
	return
}*/

//   .   .   .
// --- --- ---

/*func (index *BpIndex2) insertExistIndexValue(item BpItem) {
idx := sort.Search(len(index.Intervals), func(i int) bool {
	return index.Intervals[i] >= item.Key
})

if idx == 0 && len(index.DataNodes[0].Items) < BpWidth {
	index.DataNodes[0].insertBpDataValue(item)
	return
}

if idx == 0 && len(index.DataNodes[0].Items) >= BpWidth {
	// >>>>> split
	index.DataNodes[0].insertBpDataValue(item)
	extra := index.SplitIndex()

	if len(index.IndexNodes) == 0 {
		// main := NewBpIndex([]BpItem{})

		/*sub := NewBpIndex([]BpItem{})
		sub.IsLeaf = true*/

/*main := NewBpIdxIndexNode()
sub := NewBpIdxDataNode()*/

/*for i := 0; i < len(extra); i++ {
	sub.insertIndexValue(extra[i])
}*/

/*backup := copyBpIndex(index)

			main.IndexNodes = append(main.IndexNodes, sub, backup)

			for i := 0; i < len(main.IndexNodes); i++ {
				length := len(main.IndexNodes[i].Intervals)

				//  .  .
				// -- --

				main.Intervals = append(main.Intervals, main.IndexNodes[i].Intervals[length-1])
			}

			*index = *main

			return
		}

		if len(index.IndexNodes) != 0 {
			//
			return
		}

		return
	}

	if idx > 0 && idx < BpWidth && len(index.Intervals) < BpWidth {
		index.DataNodes[idx].insertBpDataValue(item)
		if len(index.Intervals) < (idx + 1) { // (len(index.IndexNodes)-1) == idx
			index.Intervals = append(index.Intervals, item.Key)
			return
		}
		if len(index.Intervals) >= (idx + 1) {
			length := len(index.DataNodes[idx].Items)
			index.Intervals[idx] = index.DataNodes[idx].Items[length-1].Key
		}
	}

	return
}*/

/*func copyBpIndex(index *BpIndex2) *BpIndex2 {
	if index == nil {
		return nil
	}

	// 复制Intervals切片
	intervalsCopy := make([]int64, len(index.Intervals))
	copy(intervalsCopy, index.Intervals)

	// 递归复制Index切片
	var indexCopy []*BpIndex2
	for _, subIndex := range index.IndexNodes {
		subIndexCopy := subIndex
		indexCopy = append(indexCopy, subIndexCopy)
	}

	// 递归复制Data切片
	var dataCopy []*BpData
	for _, data := range index.DataNodes {
		dataCopy = append(dataCopy, data) // 此处假设BpData为结构体，直接复制指针
	}

	// 创建新的BpIndex结构体并复制字段
	return &BpIndex2{
		IsLeaf:     index.IsLeaf,
		Intervals:  intervalsCopy,
		IndexNodes: indexCopy,
		DataNodes:  dataCopy,
	}
}*/

/*func (index *BpIndex2) insertExistIndexValue2(item BpItem) {
	idx := sort.Search(BpWidth, func(i int) bool {
		return index.Intervals[i] >= item.Key
	})

	index.Intervals = append(index.Intervals, 0)
	copy(index.Intervals[idx+1:], index.Intervals[idx:])
	index.Intervals[idx] = item.Key

	dataIndex := idx - 1
	index.DataNodes[dataIndex].Items = append(index.DataNodes[dataIndex].Items, BpItem{})
	copy(index.DataNodes[dataIndex].Items[idx+1:], index.DataNodes[dataIndex].Items[idx:])
	index.DataNodes[dataIndex].Items[idx] = item

	return
}*/
