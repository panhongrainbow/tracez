package bpTree

import (
	"fmt"
	"testing"
)

// Test_Check_Simple_BpTree is used for the simplest of confirmations, confirming some basic behaviors.
func Test_Check_Simple_BpTree(t *testing.T) {
	// Automated random testing for B Plus tree.
	t.Run("Manually Identify B Plus Tree Operation Errors", func(t *testing.T) {
		// Simplest B Plus Tree with Max Degree of 4.
		// Information is in consecutive numbers 1 to 21

		// Preparation of continuous information for addition.
		var insertedNumbers = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

		// Initialize B Plus Tree. Max Degree is 4.
		root := NewBpTree(4)

		// Start inserting data.
		for i := 0; i < len(insertedNumbers); i++ {
			// Insert data entries continuously.
			root.InsertValue(BpItem{Key: insertedNumbers[i]})
		}

		// Try deleting 1 edge value first.
		deleted, _, _, err := root.RemoveValue(BpItem{Key: 13})
		fmt.Println(deleted, err)
	})
}
