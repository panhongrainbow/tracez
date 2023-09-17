package bTree

import (
	"sync"
)

var BpWidth int
var BpHalfWidth int
var BpMiddle int

// BpTree 為B加樹的根
type BpTree struct {
	mutex sync.Mutex
	root  *BpIndex
	width int // 表示 B 加樹的階
	halfw int // B樹的最小存放量
}

func NewBpTree(width int) *BpTree {
	BpWidth = width
	BpHalfWidth = int((float32(BpWidth) + 0.1) / 2)
	BpMiddle = int((float32(BpWidth) - 0.1) / 2)
	tree := &BpTree{
		width: width,
		halfw: width / 2,
		root: &BpIndex{
			Data:   make([]*BpData, width),
			Isleaf: true,
		},
	}
	for i := 0; i < width; i++ {
		tree.root.Data[i] = &BpData{
			Items: make([]BpItem, 0, width),
		}
	}
	return tree
}

func (tree *BpTree) InsertValue(item BpItem) {
	tree.root.insertIndexValue(item)
	return
}
