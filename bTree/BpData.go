package bTree

import "sort"

// BpData 是B加樹的資料
type BpData struct {
	// MaxKey int64
	Items []BpItem
}

// not assign index

func (data *BpData) insertBpDataValue(item BpItem) {
	if len(data.Items) == 0 {
		data.Items = append(data.Items, item)
	} else {
		data.insertExistBpDataValue(item)
	}
	return
}

func (data *BpData) insertExistBpDataValue(item BpItem) {
	idx := sort.Search(len(data.Items), func(i int) bool {
		return data.Items[i].Key >= item.Key
	})

	data.Items = append(data.Items, BpItem{})
	copy(data.Items[idx+1:], data.Items[idx:])
	data.Items[idx] = item
}

// assign index

func (data *BpData) insertBpDataValue2(idx int, item BpItem) {
	if len(data.Items) == 0 {
		data.Items = append(data.Items, item)
	} else {
		data.insertExistBpDataValue2(idx, item)
	}
	return
}

func (data *BpData) insertExistBpDataValue2(idx int, item BpItem) {
	data.Items = append(data.Items, BpItem{})
	copy(data.Items[idx+1:], data.Items[idx:])
	data.Items[idx] = item
}

func arrangeInterval(interval []int64, idx int, key int64) (ret []int64) {
	if len(interval) == 0 {
		ret = append(ret, key)
	} else {
		ret = interval
		ret = append(ret, 0)
		copy(ret[idx+1:], ret[idx:])
		ret[idx] = key
	}
	return
}
