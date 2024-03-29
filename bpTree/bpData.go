package bpTree

import (
	"fmt"
	"sort"
)

// ➡️ basic struct

// BpData represents the data structure for a B+ tree node.
type BpData struct {
	Previous         *BpData  // Pointer to the previous BpData node.
	Next             *BpData  // Pointer to the next BpData node.
	Items            []BpItem // Slice to store BpItem elements.
	ShouldRenewIndex bool     // Flag indicating whether index renewal is needed.
}

// BpItem is used to record key-value pairs.
type BpItem struct {
	Key  int64       // The key used for indexing.
	Val  interface{} // The associated value.
	Mask bool        // Deleted, but unable to update the index on time.
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

// ➡️ insert operation

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
	// Create a new BpData node to store the items that will be moved.移动资料了
	side = &BpData{} // It is the new node.
	length := len(data.Items)
	side.Items = append(side.Items, data.Items[(length-BpHalfWidth):length]...) // Add the last BpHalfWidth items from data.Items to the new node.后半部的旧资料移动到新节点

	// Adjust pointers for the first old node and the new node.
	side.Previous = data  // The previous node of the new node is the first old node.新节点 的上一个节点为 第1旧节点
	side.Next = data.Next // The next node of the new node is the next node of the current node.新节点 的下一个节点为 第2旧节点

	// Reduce the data in the original node (first old node)
	data.Items = data.Items[:(length - BpHalfWidth)] // Remove the last BpHalfWidth items from data.Items to the end in the first old node.上面一行切到 length-BpHalfWidth 為基準
	data.Next = side                                 // Set the next node of the first old node to the new node.第1旧节点 的下一个节点为 新节点

	// Correct the connections between nodes!
	// There is an error here, so it needs to be corrected.
	// The situation is that when a new node is created between two old nodes, the connection between the first old node and the new node is updated.
	// However, the connection between the new node and the second-old node still needs to be updated.新节点 和 第二旧节点 之间还是要更新
	// [First old node (named data)] <--> [New node (named side)]
	//       \_______________________________________________________ [Second old node (got the value by data.Next.Next)]

	if data.Next != nil && data.Next.Next != nil {
		data.Next.Next.Previous = side // Update the previous node of the second-old node to the new node.第2旧节点 的上一个节点为 新节点
	}

	// No error
	return
}

// ➡️ delete operation

const (
	deleteNoThing = iota + 1
	deleteMiddleOne
	deleteLeftOne
	deleteRightOne
	maskLeftOne
	maskRightOne
)

// delete is a method of the BpData type that attempts to delete a BpItem from the BpData.
// It first checks the current node and then navigates to the appropriate neighbor node if needed.
/*func (data *BpData) delete(item BpItem, considerMark bool) (deleted bool, direction int) {
	// Initialize variables to track deletion status and index.
	var ix int
	deleted, ix = data._delete(item)
	direction = deleteNoThing // Set initial value.

	// Here, testing is being conducted (测试用).
	fmt.Println("in Data", ix)

	// Simultaneously search for and delete.
	if deleted {
		// If the item is successfully deleted in the current node, return.
		direction = deleteMiddleOne
	} else if !deleted {
		// When the quantity is 1, search for data in the left and right nodes.
		// 当数量为1，就左右邻居都找资料 ‼️
		if ix > 0 || len(data.Items) == 1 {
			// If the item is not found in the current node and has a non-zero index,
			// attempt deletion in the next (right) neighbor node.
			if data.Next != nil {
				if considerMark {
					if deleted, _ = data.Next._mask(item); deleted {
						direction = maskRightOne
						return
					}
				} else {
					if deleted, _ = data.Next._delete(item); deleted {
						direction = deleteRightOne
						return
					}
				}
			}
		}
		if ix == 0 || len(data.Items) == 1 {
			// If the item is not found in the current node and has an index of zero,
			// attempt deletion in the previous (left) neighbor node.
			if data.Previous != nil {
				if considerMark {
					if deleted, _ = data.Previous._mask(item); deleted {
						direction = maskLeftOne
						return
					}
				} else {
					if deleted, _ = data.Previous._delete(item); deleted {
						direction = deleteLeftOne
						return
					}
				}
			}
		}
	}

	// After going through the above process, proceed with the return.
	return
}*/

// _delete is a helper method of the BpData type that performs the actual deletion of a BpItem.
// It uses binary search to find the index where the item should be deleted.
// (真正执行删除的地方 ‼️)
func (data *BpData) _delete(item BpItem) (deleted bool, ix int, edgeValue int64, status int) {
	// 初始化回传值，data.Items 的长度不可能会为 0，因为在删除资料前，早就会进行资料合拼
	edgeValue = data.Items[0].Key
	status = edgeValueNoChanges

	// Use binary search to find the index where the item should be deleted.
	ix = sort.Search(len(data.Items), func(i int) bool {
		return data.Items[i].Key >= item.Key
	})

	// If the item is found in the current node, perform deletion and update the slice.
	if ix <= len(data.Items)-1 && ix < len(data.Items) && data.Items[ix].Key == item.Key {
		copy(data.Items[ix:], data.Items[ix+1:])
		data.Items = data.Items[:len(data.Items)-1]
		deleted = true

		// When ix is 0, it is a data edge node (这为资料边界节点)
		if ix == 0 {
			// When the edge node is empty, the edge value cannot be determined and the status becomes edgeValueUnDecided.
			// (边界节点为空)
			status = edgeValueUnDecided
			if len(data.Items) > 0 && edgeValue != data.Items[0].Key {
				edgeValue = data.Items[0].Key
				// When the edge value changes, the status changes to edgeValueChanges.
				status = edgeValueChangesByDelete
			}
		}
	}

	// If the item is not found, return without performing deletion.
	return
}
