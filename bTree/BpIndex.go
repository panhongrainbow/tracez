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
		fmt.Println("split")
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

	fmt.Println(">>>>>", idx-1)

	return
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
