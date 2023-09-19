package bpTree

import (
	"fmt"
	"sort"
)

// BpIndex is the index of the B+ tree.
type BpIndex struct {
	IsLeaf     bool       // Whether it is approaching the bottom data level
	Intervals  []int64    // The maximum values of each group of BpData
	IndexNodes []*BpIndex // Index nodes
	DataNodes  []*BpData  // Data nodes
}

// NewBpIdxIndexNode creates a new index node.
func NewBpIdxIndexNode() (index *BpIndex) {
	index = &BpIndex{
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
func NewBpIdxDataNode() (index *BpIndex) {
	index = &BpIndex{
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

func (index *BpIndex) insertIndexValue(item BpItem) {
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

func (index *BpIndex) insertExistIndexValue(item BpItem) {
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

func copyBpIndex(index *BpIndex) *BpIndex {
	if index == nil {
		return nil
	}

	// 复制Intervals切片
	intervalsCopy := make([]int64, len(index.Intervals))
	copy(intervalsCopy, index.Intervals)

	// 递归复制Index切片
	var indexCopy []*BpIndex
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
	return &BpIndex{
		IsLeaf:     index.IsLeaf,
		Intervals:  intervalsCopy,
		IndexNodes: indexCopy,
		DataNodes:  dataCopy,
	}
}

func (index *BpIndex) insertExistIndexValue2(item BpItem) {
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
