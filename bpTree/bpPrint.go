package bpTree

import "fmt"

func (inode *BpIndex) Print() {
	fmt.Println()
	if len(inode.Index) > 0 {
		fmt.Println("[⭕️IndexNode]:", inode.Index)
	}

	if len(inode.IndexNodes) > 0 {
		for _, indexNode := range inode.IndexNodes {
			indexNode.Print()
		}
	}

	if len(inode.DataNodes) > 0 {
		for _, dataNode := range inode.DataNodes {
			fmt.Printf("[🧺 DataNode]:\n")
			dataNode.Print()
		}
	}
}

func (data *BpData) Print() {
	for _, item := range data.Items {
		fmt.Printf("Key: %d\n", item.Key)
	}
}
