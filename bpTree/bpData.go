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
func (data *BpData) split(width int) (err error) {
	// Check if the number of items in the BpData is less than or equal to the specified width.
	if len(data.Items) <= width {
		// If it's not greater than the width, return an error.
		return fmt.Errorf("cannot split BpData node with less than or equal to %d items", width)
	}

	// Create a new BpData node to store the items that will be moved.
	node := &BpData{}
	node.Items = data.Items[width:]
	node.Previous = data
	node.Next = data.Next

	// Update the current BpData node to retain the first 'width' items.
	data.Items = data.Items[0:width]
	data.Next = node

	// No error occurred during the split, so return nil to indicate success.
	return nil
}
