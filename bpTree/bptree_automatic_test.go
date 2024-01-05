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
	randomQuantity = 30

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
			test, _, _, _ := root.root.delRoot(BpItem{Key: numbersForDeleting[i]})
			fmt.Println(test)
		}

		fmt.Println()
	})
	// Automated random testing for B+ tree.
	t.Run("Manually Identify B+ Tree Operation Errors", func(t *testing.T) {
		// Generate random data for insertion.
		var randomNumbers = []int64{228, 279, 164, 187, 225, 147, 150, 248, 163, 277, 230, 291, 173, 19, 141, 491, 355, 83, 178, 169, 132, 382, 444, 458, 215, 308, 138, 221, 10, 45, 54, 438, 485, 453, 106, 423, 448, 172, 45, 58, 357, 167, 37, 352, 421, 165, 492, 48, 225, 340, 311, 52, 246, 500, 75, 16, 185, 135, 71, 342, 378, 296, 47, 132, 82, 333, 483, 98, 147, 116, 101, 51, 223, 35, 372, 284, 247, 80, 427, 30, 455, 360, 382, 464, 353, 344, 405, 333, 156, 272, 312, 120, 366, 490, 144, 417, 347, 312, 460, 499}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{333, 353, 144, 225, 352, 344, 455, 116, 225, 448, 357, 277, 378, 291, 47, 485, 45, 19, 185, 138, 83, 460, 491, 246, 272, 360, 438, 423, 223, 75, 279, 427, 308, 120, 421, 340, 187, 382, 490, 228, 98, 311, 248, 499, 453, 80, 82, 141, 284, 54, 165, 230, 52, 417, 312, 215, 178, 444, 347, 163, 172, 48, 464, 37, 372, 51, 132, 500, 135, 247, 101, 30, 458, 483, 150, 333, 164, 71, 45, 58, 35, 342, 106, 312, 405, 221, 296, 147, 492, 169, 147, 16, 156, 382, 173, 132, 355, 167, 366, 10}
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
			/*if shuffledNumbers[i] == 10 {
				fmt.Println()
			}*/
			root.root.delRoot(BpItem{Key: shuffledNumbers[i]})
		}

		fmt.Println()
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
