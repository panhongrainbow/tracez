package bTree

import (
	"fmt"
	"sort"
)

// BpIndex 是B加樹的索引
type BpIndex struct {
	Isleaf    bool
	Intervals []int64 // 如果沒有 interval 長度為 0 ，就走到 Index
	Index     []*BpIndex
	Data      []*BpData
}

// BpItem 用於記錄數值
type BpItem struct {
	Key int64
	Val interface{}
}

func NewBpIndex(item []BpItem) (index *BpIndex) {
	index = &BpIndex{
		Data: make([]*BpData, BpWidth),
		// Isleaf: true,
	}
	for i := 0; i < BpWidth; i++ {
		index.Data[i] = &BpData{
			Items: make([]BpItem, 0, BpWidth),
		}
	}
	for i := 0; i < len(item); i++ {
		index.insertIndexValue(item[i])
	}

	/*index.Data[0].Items = append(index.Data[0].Items, item...)
	length := len(index.Data[0].Items)
	index.Intervals = append(index.Intervals, item[length-1].Key)*/

	return
}

func (index *BpIndex) insertIndexValue(item BpItem) {
	if index.Isleaf {
		if len(index.Intervals) == 0 {
			// 插入最左邊
			index.Intervals = append(index.Intervals, item.Key)
			index.Data[0].Items = append(index.Data[0].Items, item)
		} else {
			index.insertExistIndexValue(item)
		}
	}
	if !index.Isleaf {
		fmt.Println()
		idx := sort.Search(len(index.Intervals), func(i int) bool {
			return index.Intervals[i] >= item.Key
		})
		index.Index[idx].insertIndexValue(item)
	}
	return
}

//   .   .   .
// --- --- ---

func (index *BpIndex) insertExistIndexValue(item BpItem) {
	idx := sort.Search(len(index.Intervals), func(i int) bool {
		return index.Intervals[i] >= item.Key
	})

	if idx == 0 && len(index.Data[0].Items) < BpWidth {
		index.Data[0].insertBpDataValue(item)
		return
	}

	if idx == 0 && len(index.Data[0].Items) >= BpWidth {
		// >>>>> split
		index.Data[0].insertBpDataValue(item)
		extra := index.SplitIndex()

		if len(index.Index) == 0 {
			main := NewBpIndex([]BpItem{})

			sub := NewBpIndex([]BpItem{})
			sub.Isleaf = true

			for i := 0; i < len(extra); i++ {
				sub.insertIndexValue(extra[i])
			}

			backup := copyBpIndex(index)

			main.Index = append(main.Index, sub, backup)

			for i := 0; i < len(main.Index); i++ {
				length := len(main.Index[i].Intervals)

				//  .  .
				// -- --

				main.Intervals = append(main.Intervals, main.Index[i].Intervals[length-1])
			}

			*index = *main

			return
		}

		if len(index.Index) != 0 {
			//
			return
		}

		return
	}

	if idx > 0 && idx < BpWidth && len(index.Intervals) < BpWidth {
		index.Data[idx].insertBpDataValue(item)
		if len(index.Intervals) < (idx + 1) { // (len(index.Index)-1) == idx
			index.Intervals = append(index.Intervals, item.Key)
			return
		}
		if len(index.Intervals) >= (idx + 1) {
			length := len(index.Data[idx].Items)
			index.Intervals[idx] = index.Data[idx].Items[length-1].Key
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
	for _, subIndex := range index.Index {
		subIndexCopy := subIndex
		indexCopy = append(indexCopy, subIndexCopy)
	}

	// 递归复制Data切片
	var dataCopy []*BpData
	for _, data := range index.Data {
		dataCopy = append(dataCopy, data) // 此处假设BpData为结构体，直接复制指针
	}

	// 创建新的BpIndex结构体并复制字段
	return &BpIndex{
		Isleaf:    index.Isleaf,
		Intervals: intervalsCopy,
		Index:     indexCopy,
		Data:      dataCopy,
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
	index.Data[dataIndex].Items = append(index.Data[dataIndex].Items, BpItem{})
	copy(index.Data[dataIndex].Items[idx+1:], index.Data[dataIndex].Items[idx:])
	index.Data[dataIndex].Items[idx] = item

	return
}
