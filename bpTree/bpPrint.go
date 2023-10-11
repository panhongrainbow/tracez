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

func (inode *BpIndex) BpDataTail() (head *BpData) {
	current := inode
	for {
		if len(current.DataNodes) == 0 {
			length := len(current.IndexNodes)
			current = current.IndexNodes[length-1]
		} else {
			length := len(current.DataNodes)
			return current.DataNodes[length-1]
		}
	}
}

func (data *BpData) PrintAscent() {
	current := data
	nodeNumber := 0

	for current != nil {
		fmt.Printf("[🟣 DataNode]: NO %d \n", nodeNumber)
		length := len(current.Items)
		for i := 0; i < length; i++ {
			fmt.Printf("Key: %d\n", current.Items[i].Key)
		}

		nodeNumber++
		current = current.Next
	}
}

func (data *BpData) PrintDescent() {
	current := data
	nodeNumber := 0

	for current != nil {
		fmt.Printf("[🟣 DataNode]: NO %d \n", nodeNumber)
		length := len(current.Items)
		for i := length - 1; i >= 0; i-- {
			fmt.Printf("Key: %d\n", current.Items[i].Key)
		}

		nodeNumber++
		current = current.Previous
	}
}

func (data *BpData) PrintNodeAscent(number int) (keys []int64) {
	current := data
	nodeNumber := 0

	for current != nil {
		if nodeNumber == number {
			length := len(current.Items)
			for i := 0; i < length; i++ {
				keys = append(keys, current.Items[i].Key)
			}
			return
		}

		nodeNumber++
		current = current.Next
	}

	return
}

func (data *BpData) PrintNodeDescent(number int) (keys []int64) {
	current := data
	nodeNumber := 0

	for current != nil {
		if nodeNumber == number {
			length := len(current.Items)
			for i := length - 1; i >= 0; i-- {
				keys = append(keys, current.Items[i].Key)
			}
			return
		}

		nodeNumber++
		current = current.Previous
	}

	return
}
