package bpTree

import (
	"fmt"
	"sort"
)

// BpIndex is the index of the B plus tree.
type BpIndex struct {
	Index      []int64    // The maximum values of each group of BpData
	IndexNodes []*BpIndex // Index nodes
	DataNode   []*BpData  // Data nodes
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

func (idx *BpIndex) PushBpIndex(idxs ...*BpIndex) {
	return
}

// 中间拆，推上去
func (idx *BpIndex) SplitBpIndexAndPop() (key int64) {
	return
}

func (idx *BpIndex) PushBpData(datas ...*BpData) {
	return
}

func (idx *BpIndex) SplitBpDataAndPop() (newIdx *BpIndex) {
	return
}

// >>>>> >>>>> >>>>> >>>>> >>>>>

// insertBpDataValue inserts a key into the BpIndex.
func (idx *BpIndex) insertBpIdxValue(key int64) {
	// If there are existing items, insert the new item among them.
	if len(idx.Index) > 0 {
		idx.insertExistBpIdxValue(key)
	}

	// If there is no existing index, simply append the new key.
	if len(idx.Index) == 0 {
		idx.Index = append(idx.Index, key)
		return
	}

	return
}

// insertExistBpIdxValue inserts a key into the existing sorted BpIndex's index.
func (idx *BpIndex) insertExistBpIdxValue(key int64) {
	// Use binary search to find the index(i) where the key should be inserted.
	i := sort.Search(len(idx.Index), func(i int) bool {
		return idx.Index[i] >= key
	})

	// Expand the slice to accommodate the new key.
	idx.Index = append(idx.Index, 0)

	// Shift the elements to the right to make space for the new key.
	copy(idx.Index[i+1:], idx.Index[i:])

	// Insert the new key at the correct position.
	idx.Index[i] = key
}

// split divides the BpIndex's index into two parts if it contains more items than the specified width.
func (idx *BpIndex) split(width int) (err error) {
	// Check if the number of index in the BpData is less than or equal to the specified width.
	if len(idx.Index) <= width {
		// If it's not greater than the width, return an error.
		return fmt.Errorf("cannot split BpData node with less than or equal to %d items", width)
	}

	// Create a new index node to store the items that will be moved.
	node := &BpIndex{}
	node.Index = idx.Index[width:]

	//
	// node.IndexNodes

	// Update the current index node to retain the first 'width' items.
	idx.Index = idx.Index[0:width]

	// No error occurred during the split, so return nil to indicate success.
	return nil
}

// >>>>> >>>>> >>>>> >>>>> >>>>>

// BpIndex2 is the index of the B+ tree.
type BpIndex2 struct {
	IsLeaf     bool        // Whether it is approaching the bottom data level
	Intervals  []int64     // The maximum values of each group of BpData
	IndexNodes []*BpIndex2 // Index nodes
	DataNodes  []*BpData   // Data nodes
}

// NewBpIdxIndexNode creates a new index node.
func NewBpIdxIndexNode() (index *BpIndex2) {
	index = &BpIndex2{
		DataNodes: []*BpData{},
		IsLeaf:    false,
	}
	/*for i := 0; i < BpWidth; i++ {
		index.DataNodes[i] = &BpData{
			Items: make([]BpItem, BpWidth),
		}
	}*/
	return
}

// NewBpIdxDataNode creates a new data node.
func NewBpIdxDataNode() (index *BpIndex2) {
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
}

func (index *BpIndex2) insertIndexValue(item BpItem) {
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
}

//   .   .   .
// --- --- ---

func (index *BpIndex2) insertExistIndexValue(item BpItem) {
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

			main := NewBpIdxIndexNode()
			sub := NewBpIdxDataNode()

			for i := 0; i < len(extra); i++ {
				sub.insertIndexValue(extra[i])
			}

			backup := copyBpIndex(index)

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
}

func copyBpIndex(index *BpIndex2) *BpIndex2 {
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
}

func (index *BpIndex2) insertExistIndexValue2(item BpItem) {
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
}
