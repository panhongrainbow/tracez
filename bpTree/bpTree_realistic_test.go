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
	root.InsertValue(BpItem{Key: 10})
	root.InsertValue(BpItem{Key: 11})
	root.InsertValue(BpItem{Key: 9})
	root.InsertValue(BpItem{Key: 12})
	root.InsertValue(BpItem{Key: 13})
	root.InsertValue(BpItem{Key: 15})
	root.InsertValue(BpItem{Key: 15})
	root.InsertValue(BpItem{Key: 15})
	root.InsertValue(BpItem{Key: 15})
	root.InsertValue(BpItem{Key: 15})
	root.InsertValue(BpItem{Key: 15})
	root.InsertValue(BpItem{Key: 8})
	root.InsertValue(BpItem{Key: 3})
	root.InsertValue(BpItem{Key: 7})
	root.InsertValue(BpItem{Key: 10})
	root.InsertValue(BpItem{Key: 9})
	root.InsertValue(BpItem{Key: 9})
	root.InsertValue(BpItem{Key: 8})
	root.InsertValue(BpItem{Key: 8})
	root.root.Print()
	fmt.Println()
}
