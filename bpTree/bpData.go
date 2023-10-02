package bpTree

import (
	"fmt"
	"sort"
)

// BpData represents the data structure for a B+ tree node.
type BpData struct {
	Previous *BpData  // Pointer to the previous BpData node.
	Next     *BpData  // Pointer to the next BpData node.
	Items    []BpItem // Slice to store BpItem elements.
	Split    bool     // After splitting the nodes, mark it.
}

// BpItem is used to record key-value pairs.
type BpItem struct {
	Key int64       `json:"key"` // The key used for indexing.
	Val interface{} `json:"val"` // The associated value.
}

// dataLength returns the length of BpData's items slice.
func (data *BpData) dataLength() (length int) {
	length = len(data.Items)
	return
}

// index retrieves the key from the first BpItem in the BpData, if available.
func (data *BpData) index() (key int64, err error) {
	// If there are items in the BpData, retrieve the key from the first item.
	if len(data.Items) > 0 {
		key = data.Items[0].Key
	}

	// If there are no items in the BpData, set an error indicating no data.
	if len(data.Items) == 0 {
		err = fmt.Errorf("there is no available index for bpdata")
	}

	return
}

// >>>>> >>>>> >>>>> insert

// insertBpDataValue inserts a BpItem into the BpData.
func (data *BpData) insert(item BpItem) {
	// If there are existing items, insert the new item among them.
	if len(data.Items) > 0 {
		data.insertAmong(item)
	}

	// If there are no existing items, simply append the new item.
	if len(data.Items) == 0 {
		data.Items = append(data.Items, item)
	}

	return
}

// insertAmong inserts a BpItem into the existing sorted BpData.
func (data *BpData) insertAmong(item BpItem) {
	// Use binary search to find the index where the item should be inserted.
	idx := sort.Search(len(data.Items), func(i int) bool {
		return data.Items[i].Key >= item.Key
	})

	// Expand the slice to accommodate the new item.
	data.Items = append(data.Items, BpItem{})

	// Shift the elements to the right to make space for the new item.
	copy(data.Items[idx+1:], data.Items[idx:])

	// Insert the new item at the correct position.
	data.Items[idx] = item
}

// split divides the BpData node into two nodes if it contains more items than the specified width.
func (data *BpData) split() (side *BpData, err error) {
	// Create a new BpData node to store the items that will be moved.
	side = &BpData{}
	length := len(data.Items)
	side.Items = append(side.Items, data.Items[(length-BpHalfWidth):length]...) // data.Items[length:length] 为空，在最后面往前 BpHalfWidth
	side.Previous = data
	side.Next = data.Next

	// Reduce the original node.
	data.Items = data.Items[:(length - BpHalfWidth)] // 上面一行切到 length-BpHalfWidth
	data.Next = side

	// Make a mark, already split.
	data.Split = true

	// No error
	return
}

// >>>>> >>>>> >>>>> delete

func (data *BpData) delete(item BpItem) (deleted bool) {
	// Use binary search to find the index where the item should be deleted.
	ix := sort.Search(len(data.Items), func(i int) bool {
		return data.Items[i].Key >= item.Key
	})

	if ix < len(data.Items) && data.Items[ix].Key == item.Key {
		copy(data.Items[ix:], data.Items[ix+1:])
		data.Items = data.Items[:len(data.Items)-1]
		deleted = true
	}

	return
}
