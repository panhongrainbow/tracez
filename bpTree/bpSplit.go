package bpTree

//    .
// ----

//	  .
//	.--
//
// --
/*func (index *BpIndex2) SplitIndex() (item []BpItem) {
	// new index
	// sub := NewBpIndex(index.DataNodes[0].Items[0:BpHalfWidth])
	// index.IndexNodes = append(index.IndexNodes, sub)
	item = index.DataNodes[0].Items[0:BpHalfWidth]
	// delete half
	index.DataNodes[0].Items = index.DataNodes[0].Items[BpHalfWidth:]
	length := len(index.DataNodes[0].Items)
	index.Intervals[0] = index.DataNodes[0].Items[length-1].Key
	return
}*/
