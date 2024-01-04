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
		var randomNumbers = []int64{276, 309, 152, 462, 477, 322, 250, 491, 208, 109, 211, 256, 413, 455, 478, 363, 423, 395, 185, 331, 404, 172, 396, 471, 305, 35, 213, 153, 20, 235, 341, 386, 323, 392, 296, 300, 82, 195, 309, 103, 211, 438, 187, 106, 474, 135, 220, 249, 349, 189}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{423, 309, 395, 211, 349, 474, 20, 195, 220, 455, 396, 296, 309, 103, 185, 109, 413, 276, 211, 386, 404, 106, 341, 363, 35, 322, 208, 300, 331, 235, 256, 172, 491, 392, 323, 249, 471, 438, 213, 135, 187, 477, 82, 189, 152, 462, 478, 153, 305, 250}
		// var shuffledNumbers = []int64{423, 309, 395, 211, 349, 474, 20, 195, 220, 455, 396, 296, 309, 103, 185, 109, 413, 276, 211, 386, 404, 106, 341, 363, 35, 322, 208, 300, 331, 235, 256, 172, 491, 392, 323, 249, 471, 438, 213, 462, 187, 477, 82, 189, 152, 135, 478, 153, 305, 250}

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
			if shuffledNumbers[i] == 82 { // 135
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
