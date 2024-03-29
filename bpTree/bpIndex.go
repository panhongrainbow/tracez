package bpTree

import (
	"fmt"
	"sort"
)

// >>>>> >>>>> >>>>> main structure

const (
	statusNormal = iota + 1
	statusProtrudeInode
	statusProtrudeDnode
	statusBorrowFromIndexNode // 执行 borrowFromIndexNode 后
	statusIXMunus
	statusError
	edgeValueInit
	edgeValueNoChanges
	edgeValueChanges
	edgeValueChangesByDelete
	edgeValueUpload
	edgeValueUnDecided
	edgeValuePassBottom
	edgeValueRenew
	statusDeleteItem
	statusDeleteNon
	statusDeProtrude
	statusDeleteProtrude
)

// BpIndex is the index of the B plus tree.
type BpIndex struct {
	Index      []int64    // The maximum values of each group of BpData
	IndexNodes []*BpIndex // Index nodes
	DataNodes  []*BpData  // Data nodes
}

// insertBpDataValue inserts a new index into the BpIndex.
// 经由 BpIndex 直接在新增
func (inode *BpIndex) insertItem(newNode *BpIndex, item BpItem) (popIx int, popKey int64, popNode *BpIndex, status int, err error) {
	var newIndex int64
	var sideDataNode *BpData
	status = statusNormal // status is used to inform the root node that it is not the root node here, so the state becomes Normal !.
	// 状态是用来告知 root 节点，在这里不是 root 节点，所以状态变为 Normal !

	// >>>>> 进入索引结点

	// If there are existing items, insert the new item among them.
	if newNode == nil && len(inode.Index) > 0 {
		// (当索引大于 0，就可以直接开始找位置)

		// Use binary search to find the index(i) where the key should be inserted.
		ix := sort.Search(len(inode.Index), func(i int) bool {
			return inode.Index[i] >= item.Key
		})

		// >>>>> >>>>> >>>>> 进入递归

		if len(inode.IndexNodes) > 0 {
			if len(inode.IndexNodes) != (len(inode.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", inode.Index)
				return
			}

			// If there are index nodes, recursively insert the item into the appropriate node.
			// (这里有递回去找到接近资料切片的地方)
			popIx, popKey, popNode, status, err = inode.IndexNodes[ix].insertItem(nil, item)
			status = statusProtrudeInode
			if popKey != 0 {
				err = inode.mergeUpgradedKeyNode(ix, popKey, popNode)
				popKey = 0
				popNode = nil
			}

			if popNode != nil && popKey == 0 {
				// New index node has been created independently and are going to be upgraded and overwrite inode.
				inode.ackUpgradeIndexNode(ix, popNode) // 在这里同意并覆写 inode
				popNode = nil
			}

			if len(inode.Index) >= BpWidth && len(inode.Index)%2 != 0 { // 进行 pop 和奇数
				popNode, err = inode.protrudeInOddBpWidth()
				return
			} else if len(inode.Index) >= BpWidth && len(inode.Index)%2 == 0 { // 进行 pop 和奇数
				popNode, err = inode.protrudeInEvenBpWidth()
				return
			}

			return // Break here to avoid inserting elsewhere. (立刻中断)
		}

		// If there are data nodes, insert the new item at the determined index.
		if len(inode.DataNodes) > 0 {
			if len(inode.DataNodes) != (len(inode.Index) + 1) {
				err = fmt.Errorf("the number of indexes is incorrect, %v", inode.Index)
				return
			}

			// >>>>> 进入第 1 个资料结点入口

			inode.DataNodes[ix].insert(item) // Insert item at index ix.

			if len(inode.DataNodes[ix].Items) >= BpWidth {
				sideDataNode, err = inode.DataNodes[ix].split()
				if err != nil {
					return
				}

				inode.DataNodes = append(inode.DataNodes, &BpData{})
				copy(inode.DataNodes[(ix+1)+1:], inode.DataNodes[(ix+1):])
				inode.DataNodes[ix+1] = sideDataNode

				inode.insertBpIX(sideDataNode.Items[0].Key)
			}

			if len(inode.Index) >= BpWidth {
				popKey, popNode, err = inode.splitWithDnode()
				status = statusProtrudeDnode
				popIx = ix
				if err != nil {
					return
				}
			}

			return
		}
	}

	// >>>>> 进入第 2 个资料结点入口

	// The length of idx.Index is 0, which only occurs in one scenario where there is only one DataNodesDataNodes.
	// (Idx.Index 的长度为 0，只有在一个状况才会发生，资料分片只有一份)
	if newNode == nil && len(inode.Index) == 0 {
		if len(inode.DataNodes) != 1 {
			// 资料大于1，就会有索引，就不会进入这里
			err = fmt.Errorf("the number of indexes is incorrect initially")
			return
		}
		inode.DataNodes[0].insert(item) // >>>>> (add to DataNodes)

		if inode.DataNodes[0].dataLength() >= BpWidth {
			sideDataNode, err = inode.DataNodes[0].split() // newIndex
			if err != nil {
				return
			}

			inode.DataNodes = append(inode.DataNodes, sideDataNode)
			newIndex = sideDataNode.Items[0].Key
		}
	}

	if sideDataNode != nil {
		inode.insertBpIX(newIndex)

		if len(inode.Index) >= BpWidth && len(inode.Index)%2 != 0 { // 进行 pop 和奇数 (可能没在使用)
			var node *BpIndex
			node, err = inode.protrudeInOddBpWidth()
			*inode = *node
			return
		} else if len(inode.Index) >= BpWidth && len(inode.Index)%2 == 0 { // 进行 pop 和奇数 (可能没在使用)
			var node *BpIndex
			node, err = inode.protrudeInEvenBpWidth()
			*inode = *node
			return
		}

		return // Break here to avoid inserting elsewhere. (立刻中断)
	}

	return
}

// ackUpgradeIndexNode is used by the current layer's index node
// to acknowledge a new independently upgraded index node.
// this function is extracted from insertItem function for testing purposes,
// and it overwrites the original location in the inode.
// (承认新独立的索引结点)
//
//go:inline
func (inode *BpIndex) ackUpgradeIndexNode(ix int, popNode *BpIndex) {
	// 这个函式不能在 root 节点上使用，

	// Insert popNode.Index[0]
	inode.insertBpIX(popNode.Index[0])

	// Create a new BpIndex
	node := &BpIndex{}

	// Split inode.IndexNodes into three parts and then merge them into the new node.IndexNodes
	node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[:ix]...)   // positions are from 0 to ix.
	node.IndexNodes = append(node.IndexNodes, popNode.IndexNodes...)      // position ix is overtired.
	node.IndexNodes = append(node.IndexNodes, inode.IndexNodes[ix+1:]...) // positions are from (ix+1) to the last one.
	// In the above operations, the "ix" position has been overwritten
	//
	// because during the splitting process,
	// the original node is preserved in its entirety, and it is only during the merging that there will be no positional confusion.
	//
	// For example, in this case, the "ix" position is being overwritten, and the original inode has not been modified,
	// so there won't be any confusion.

	// 切割时，会保留原节点的完整，合拼时才不会发生位置错

	// Update inode.IndexNodes
	inode.IndexNodes = node.IndexNodes
}

// >>>>> >>>>> >>>>> insert method

// insertBpIX inserts a new index at the correct position using binary search.
// Just inserting an index slice won't result in any errors, so it doesn't return an error.
func (inode *BpIndex) insertBpIX(newIx int64) {
	// Use binary search to find the position where the index should be inserted.
	ix := sort.Search(len(inode.Index), func(i int) bool {
		return inode.Index[i] >= newIx
	})

	// Expand the slice to accommodate the new item.
	inode.Index = append(inode.Index, 0)

	// Shift the elements to the right to make space for the new item.
	copy(inode.Index[ix+1:], inode.Index[ix:])

	// Insert the new item at the correct position.
	inode.Index[ix] = newIx

	return
}

// protrudeInOddBpWidth performs index upgrade; when the middle value of the index slice pops out, it gets upgraded to the upper-level index.
// This is used when the width of BpWidth is odd.
// (进行索引升级，当索引切片的中间值会弹出升级成上层的索引)
func (inode *BpIndex) protrudeInOddBpWidth() (middle *BpIndex, err error) {
	// At the beginning, a check is performed.
	// This function is designed to handle cases where the BpWidth is an odd number,
	// meaning the length of the Index slice is odd,
	// and the length of the IndexNodes slice is even,
	// with a difference of 1 in the lengths.(Index 切片 和 IndexNodes 切片长度 差 1)
	if len(inode.Index)%2 != 1 || len(inode.IndexNodes)%2 != 0 {
		err = fmt.Errorf("in the case of an odd width, protruding oversized index nodes results in an error")
		return
	}

	// Calculate the current index lengths for splitting.
	indexLen := (len(inode.Index) - 1) / 2
	indexNodeLen := len(inode.IndexNodes) / 2

	// Create a new left node.
	leftNode := &BpIndex{
		Index:      append([]int64{}, inode.Index[:indexLen]...),
		IndexNodes: append([]*BpIndex{}, inode.IndexNodes[:indexNodeLen]...),
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Create a new right node.
	rightNode := &BpIndex{
		Index:      append([]int64{}, inode.Index[indexLen+1:]...),
		IndexNodes: append([]*BpIndex{}, inode.IndexNodes[indexNodeLen:]...),
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Create a new middle node.
	middle = &BpIndex{
		Index:      inode.Index[indexLen : indexLen+1],
		IndexNodes: []*BpIndex{leftNode, rightNode},
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Return the error, regardless of whether there is an error or not.
	return
}

// protrudeInOddBpWidth performs index upgrade; when the middle value of the index slice pops out, it gets upgraded to the upper-level index.
// This is used when the width of BpWidth is even.
// (进行索引升级，当索引切片的中间值会弹出升级成上层的索引)
func (inode *BpIndex) protrudeInEvenBpWidth() (popMiddleNode *BpIndex, err error) {
	// At the beginning, a check is performed.
	// This function is designed to handle cases where the BpWidth is an odd number,
	// meaning the length of the Index slice is even,
	// and the length of the IndexNodes slice is odd,
	// with a difference of 1 in the lengths.(Index 切片 和 IndexNodes 切片长度 差 1)
	if len(inode.Index)%2 != 0 || len(inode.IndexNodes)%2 != 1 {
		err = fmt.Errorf("in the case of an odd width, protruding oversized index nodes results in an error")
		return
	}

	// Calculate the current index lengths for splitting.
	indexLen := (len(inode.Index)) / 2
	indexNodeLen := (len(inode.IndexNodes) - 1) / 2

	// Create a new left node.
	leftNode := &BpIndex{
		Index:      append([]int64{}, inode.Index[:indexLen]...),
		IndexNodes: append([]*BpIndex{}, inode.IndexNodes[:indexNodeLen+1]...),
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Create a new right node.
	rightNode := &BpIndex{
		Index:      append([]int64{}, inode.Index[indexLen+1:]...),
		IndexNodes: append([]*BpIndex{}, inode.IndexNodes[indexNodeLen+1:]...),
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Create a new middle node.
	popMiddleNode = &BpIndex{
		Index:      inode.Index[indexLen : indexLen+1],
		IndexNodes: []*BpIndex{leftNode, rightNode},
		// DataNode slice is set to nil directly. It should not be used later.
	}

	// Return the error, regardless of whether there is an error or not.
	return
}

// >>>>> >>>>> >>>>> split and merge the bottom-level index node.

// splitWithDnode splits the bottom-level index node effectively and returns a new independent key and index node.
func (inode *BpIndex) splitWithDnode() (key int64, side *BpIndex, err error) {
	// Check if both IndexNodes and DataNodes have data,
	// which is incorrect as we don't know the type of node.
	if len(inode.IndexNodes) != 0 && len(inode.DataNodes) != 0 {
		err = fmt.Errorf("both IndexNodes and DataNodes have data, we cannot determine the type of node")
		return
	}

	// Handle splitting based on DataNodes.
	if len(inode.DataNodes) != 0 {
		// Create a new node named side.
		side = &BpIndex{}
		length := len(inode.DataNodes)

		// Append a portion of the Index and DataNodes to the 'side' structure.
		side.Index = append(side.Index, inode.Index[(length-BpHalfWidth):]...)
		// This is equivalent to side.Index = append(side.Index, inode.Index[(length-BpHalfWidth):len(inode.Index)])
		// 这里等于 side.Index = append(side.Index, inode.Index[(length-BpHalfWidth):len(inode.Index)])

		side.DataNodes = append(side.DataNodes, inode.DataNodes[(length-BpHalfWidth):]...)
		// This is equivalent to side.DataNodes = append(side.DataNodes, inode.DataNodes[(length-BpHalfWidth):len(inode.DataNodes)]),
		// where len(inode.DataNodes) will be one more than len(inode.Index)
		// Hence, side.DataNodes will be one more than side.Index, so the slicing operation is correct.

		// 这里等于 side.DataNodes = append(side.DataNodes, inode.DataNodes[(length-BpHalfWidth):len(inode.DataNodes)])，len(inode.DataNodes) 会比 len(inode.Index) 多 1 个
		// 最后 side.DataNodes 会比 side.Index 多 1 个，所以切割操作正确

		// The logic here is a bit complex, where the length is the length of the DataNode slice,
		// and the expression [(length-BpHalfWidth):] determines how much data the new node should take.
		// When [(length-BpHalfWidth):] is applied to the index code, side.Index = append(side.Index, inode.Index[(length-BpHalfWidth):]...),
		// the length will be one less than side.DataNodes. This ensures that DataNodes has one more element than Index,
		// so the overall logic is correct.

		// 这里的程式码有点复杂，其中长度 length 为 DataNode 切片的长度，那式子 [(length-BpHalfWidth):] 中的 BpHalfWidth 意思就为新节点要取多少笔资料，
		// 再把 [(length-BpHalfWidth):] 套上 index 的代码中，side.Index = append(side.Index, inode.Index[(length-BpHalfWidth):]...)，长度会比 side.DataNodes 少 1 个
		// 这样就符合 DataNodes 的切片长度比 Index 多 1，整个逻辑是正确的

		// Update the 'key' assignment with a value from the original Index.
		key = inode.Index[length-BpHalfWidth-1]

		// Update the original Index and DataNodes by removing the appended portion.
		inode.Index = inode.Index[0 : length-BpHalfWidth-1]
		inode.DataNodes = inode.DataNodes[0 : length-BpHalfWidth]
	}

	// Just return and don't worry about anything.
	return
}

// mergeWithDnode combines the newly split index nodes created by splitWithDnode into a new node,
// overwriting the original inode's address.
func (inode *BpIndex) mergeWithDnode(podKey int64, side *BpIndex) error {
	// Create a new BpIndex structure.
	originAndSide := &BpIndex{
		Index: []int64{podKey},
	}

	// Copy the current inode's Index, IndexNodes, and DataNodes to the new structure.
	copyInode := &BpIndex{
		Index:      append([]int64{}, inode.Index...),
		IndexNodes: append([]*BpIndex{}, inode.IndexNodes...),
		DataNodes:  append([]*BpData{}, inode.DataNodes...),
	}

	// Add copyInode to originAndSide.IndexNodes.
	originAndSide.IndexNodes = append(originAndSide.IndexNodes, copyInode)

	// Add side to originAndSide.IndexNodes.
	originAndSide.IndexNodes = append(originAndSide.IndexNodes, side)

	// Assign the value of originAndSide to inode.
	*inode = *originAndSide

	// Return nil to indicate no error.
	return nil
}

// >>>>> >>>>> >>>>> merge the upgraded key and upgraded index node.

// mergeUpgradedKeyNode merges the to-be-upgraded Key and the to-be-upgraded Inode into the parent higher-level index node.
func (inode *BpIndex) mergeUpgradedKeyNode(insertAfterPosition int, key int64, side *BpIndex) (err error) {
	// The B Plus tree builds an index, and when some indexes become independent, they turn into keys.
	// Merging these keys into other index nodes is not difficult.
	// It's just a matter of sorting.
	insertAfterPosition = insertAfterPosition + 1
	inode.insertBpIX(key)

	// Store the upgraded index node named side at the appropriate position in the IndexNodes slice.
	inode.IndexNodes = append(inode.IndexNodes, &BpIndex{})
	copy(inode.IndexNodes[insertAfterPosition+1:], inode.IndexNodes[insertAfterPosition:])
	inode.IndexNodes[insertAfterPosition] = side

	// Return the error, regardless of whether there is an error or not.
	return err
}
