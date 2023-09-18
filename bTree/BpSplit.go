package bTree

//    .
// ----

//	  .
//	.--
//
// --
func (index *BpIndex) SplitIndex() (item []BpItem) {
	// new index
	// sub := NewBpIndex(index.Data[0].Items[0:BpHalfWidth])
	// index.Index = append(index.Index, sub)
	item = index.Data[0].Items[0:BpHalfWidth]
	// delete half
	index.Data[0].Items = index.Data[0].Items[BpHalfWidth:]
	length := len(index.Data[0].Items)
	index.Intervals[0] = index.Data[0].Items[length-1].Key
	return
}
