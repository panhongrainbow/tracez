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
}

// BpItem is used to record key-value pairs.
type BpItem struct {
	Key int64       // The key used for indexing.
	Val interface{} // The associated value.
}

// getBpDataLength returns the length of BpData's items slice.
func (data *BpData) getBpDataLength() (length int) {
	length = len(data.Items)
	return
}

// getBpDataIndex retrieves the key from the first BpItem in the BpData, if available.
func (data *BpData) getBpDataIndex() (key int64, err error) {
	// If there are items in the BpData, retrieve the key from the first item.
	if len(data.Items) > 0 {
		key = data.Items[0].Key
	}

	// If there are no items in the BpData, set an error indicating no data.
	if len(data.Items) == 0 {
		err = fmt.Errorf("no data available")
	}

	return
}

// insertBpDataValue inserts a BpItem into the BpData.
func (data *BpData) insertBpDataValue(item BpItem) {
	// If there are existing items, insert the new item among them.
	if len(data.Items) > 0 {
		data.insertExistBpDataValue(item)
	}

	// If there are no existing items, simply append the new item.
	if len(data.Items) == 0 {
		data.Items = append(data.Items, item)
	}

	return
}

// insertExistBpDataValue inserts a BpItem into the existing sorted BpData.
func (data *BpData) insertExistBpDataValue(item BpItem) {
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
func (data *BpData) split() (node *BpData, err error) {
	// Create a new BpData node to store the items that will be moved.
	node = &BpData{}
	length := len(data.Items)
	// node.Items = data.Items[(length - 2):length]
	node.Items = append(node.Items, data.Items[(length-2):length]...)
	node.Previous = data
	node.Next = data.Next

	// Update the current BpData node to retain the first 'width' items.
	data.Items = data.Items[:(length - 2)]
	data.Next = node

	// No error occurred during the split, so return nil to indicate success.
	return
}
