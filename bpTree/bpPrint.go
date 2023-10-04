package bpTree

import "fmt"

func (inode *BpIndex) Print() {
	fmt.Println()
	fmt.Println("[⭕️IndexNode]:", inode.Index)

	if len(inode.IndexNodes) > 0 {
		for _, indexNode := range inode.IndexNodes {
			indexNode.Print()
		}
	}

	for _, dataNode := range inode.DataNodes {
		fmt.Printf("[🟣 DataNode]:\n")
		dataNode._print()
	}
}

func (data *BpData) _print() {
	for _, item := range data.Items {
		fmt.Printf("Key: %d\n", item.Key)
	}
}

func (inode *BpIndex) BpDataHead() (head *BpData) {
	current := inode
	for {
		if len(current.DataNodes) == 0 {
			current = current.IndexNodes[0]
		} else {
			return current.DataNodes[0]
		}
	}
}

func (data *BpData) Print() {
	current := data
	i := 0

	for current != nil {
		fmt.Printf("[🟣 DataNode]: NO %d \n", i)
		for _, item := range current.Items {
			fmt.Printf("Key: %d\n", item.Key)
		}

		i++
		current = current.Next
	}
}

func (data *BpData) PrintNodeKeys(number int) (keys []int64) {
	current := data
	i := 0

	for current != nil {
		if i == number {
			for _, item := range current.Items {
				keys = append(keys, item.Key)
			}
			return
		}

		i++
		current = current.Next
	}

	return
}
