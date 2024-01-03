package bpTree

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var (
	// randomQuantity represents the number of elements to be generated for random testing.
	randomQuantity = 50

	// randomMax represents the maximum value for generating random numbers.
	randomMax = 500

	// randomMin represents the minimum value for generating random numbers.
	randomMin = 10
)

// Test_Check_BpTree_Automatic is used for automated testing, generating test data with random numbers for B+ tree insertion and deletion.
func Test_Check_BpTree_Automatic(t *testing.T) {
	// Automated random testing for B+ tree.
	t.Run("Automated Testing Section", func(t *testing.T) {
		// Set up randomization.
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)

		// Generate random data for insertion.
		numbersForAdding := make([]int64, randomQuantity)
		for i := 0; i < randomQuantity; i++ {
			num := int64(random.Intn(randomMax-randomMin+1) + randomMin)
			numbersForAdding[i] = num
		}
		fmt.Println("Random data for insertion:", numbersForAdding)

		// Generate random data for deletion.
		numbersForDeleting := make([]int64, randomQuantity)
		copy(numbersForDeleting, numbersForAdding)
		shuffleSlice(numbersForDeleting, random)
		fmt.Println("Random data for deletion:", numbersForDeleting)

		// Generate sorted data.
		sortedNumbers := make([]int64, randomQuantity)
		copy(sortedNumbers, numbersForAdding)
		sort.Slice(sortedNumbers, func(i, j int) bool {
			return sortedNumbers[i] < sortedNumbers[j]
		})
		fmt.Println("Sorted data:", sortedNumbers)

		// Initialize B-tree.
		root := NewBpTree(5)

		// Start inserting data.
		for i := 0; i < randomQuantity; i++ {
			// Insert data entries continuously.
			root.InsertValue(BpItem{Key: numbersForAdding[i]})
		}

		// Start deleting data.
		for i := 0; i < randomQuantity; i++ {
			// Deleting data entries continuously.
			root.root.delRoot(BpItem{Key: numbersForDeleting[i]})
		}
	})
	// Automated random testing for B+ tree.
	t.Run("Manually Identify B+ Tree Operation Errors", func(t *testing.T) {
		// Generate random data for insertion.
		var randomNumbers = []int64{416, 33, 150, 426, 42, 301, 462, 214, 300, 331, 163, 389, 170, 482, 117, 170, 127, 339, 163, 32, 132, 323, 308, 149, 486, 341, 338, 119, 92, 472, 27, 39, 108, 281, 375, 471, 206, 109, 299, 265, 112, 247, 264, 51, 118, 176, 321, 292, 158, 227}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{482, 112, 32, 170, 108, 163, 341, 265, 117, 150, 227, 299, 158, 486, 92, 206, 339, 301, 214, 462, 375, 247, 300, 389, 281, 471, 42, 127, 149, 323, 321, 472, 119, 176, 118, 51, 426, 33, 27, 132, 163, 331, 170, 308, 292, 338, 264, 109, 416, 39}

		// Initialize B-tree.
		root := NewBpTree(5)

		// Start inserting data.
		for i := 0; i < randomQuantity; i++ {
			// Insert data entries continuously.
			root.InsertValue(BpItem{Key: randomNumbers[i]})
		}

		// Start deleting data.
		for i := 0; i < randomQuantity; i++ {
			fmt.Println(i, shuffledNumbers[i])
			// Deleting data entries continuously.
			if shuffledNumbers[i] == 109 {
				fmt.Println()
			}
			root.root.delRoot(BpItem{Key: shuffledNumbers[i]})
		}
	})
}

// shuffleSlice randomly shuffles the elements in the slice.
func shuffleSlice(slice []int64, rng *rand.Rand) {
	// Iterate through the slice in reverse order, starting from the last element.
	for i := len(slice) - 1; i > 0; i-- {
		// Generate a random index 'j' between 0 and i (inclusive).
		j := rng.Intn(i + 1)

		// Swap the elements at indices i and j.
		slice[i], slice[j] = slice[j], slice[i]
	}
}
