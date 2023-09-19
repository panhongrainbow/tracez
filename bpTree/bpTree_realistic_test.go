package bpTree

import (
	"fmt"
	"testing"
)

func Test_Check_Btree(t *testing.T) {
	root := NewBpTree(3)
	root.InsertValue(BpItem{Key: 5})
	root.InsertValue(BpItem{Key: 6})
	root.InsertValue(BpItem{Key: 5})
	root.InsertValue(BpItem{Key: 7})
	root.InsertValue(BpItem{Key: 4})
	root.InsertValue(BpItem{Key: 3})
	root.InsertValue(BpItem{Key: 1})
	root.InsertValue(BpItem{Key: 4})
	root.InsertValue(BpItem{Key: 1})
	fmt.Println()
}
