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
	randomQuantity = 100

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
			deleted, _, _, err := root.root.delFromRoot(BpItem{Key: numbersForDeleting[i]})
			if deleted == false {
				fmt.Println("Breakpoint: Data deletion not successful. 💢 The number is ", numbersForDeleting[i], i)
			}
			if err != nil {
				fmt.Println("Breakpoint: Deletion encountered an error. 💢 The number is ", numbersForDeleting[i], i)
			}
		}
	})
	// Automated random testing for B+ tree.
	t.Run("Manually Identify B+ Tree Operation Errors", func(t *testing.T) {
		// Generate random data for insertion.
		var randomNumbers = []int64{366, 100, 303, 47, 24, 43, 186, 243, 500, 44, 226, 486, 176, 50, 35, 17, 330, 256, 128, 205, 19, 209, 438, 417, 416, 101, 53, 208, 42, 377, 316, 360, 15, 381, 232, 458, 17, 80, 466, 61, 165, 123, 73, 419, 58, 223, 341, 63, 484, 330, 399, 268, 182, 244, 86, 62, 423, 23, 14, 126, 446, 344, 484, 104, 332, 336, 231, 84, 483, 319, 450, 33, 487, 167, 360, 403, 43, 362, 136, 434, 143, 82, 426, 313, 369, 95, 274, 465, 45, 347, 295, 495, 113, 231, 32, 100, 228, 328, 454, 85}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{328, 50, 466, 19, 86, 319, 330, 15, 14, 123, 426, 62, 100, 313, 101, 256, 483, 465, 486, 417, 369, 45, 43, 416, 231, 61, 113, 136, 58, 128, 243, 495, 17, 446, 228, 458, 205, 126, 53, 208, 438, 450, 33, 226, 344, 186, 366, 403, 244, 176, 500, 209, 84, 487, 23, 336, 104, 434, 484, 423, 165, 47, 377, 35, 63, 143, 332, 362, 295, 360, 43, 32, 316, 44, 454, 17, 381, 82, 330, 85, 303, 399, 274, 347, 167, 100, 80, 42, 231, 268, 182, 24, 360, 341, 419, 232, 223, 484, 73, 95}
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
			if i == 92 { // 51
				fmt.Println()
			}
			deleted, _, _, err := root.RemoveValue(BpItem{Key: shuffledNumbers[i]})
			if deleted == false {
				fmt.Println("Breakpoint: Data deletion not successful. 💢 The number is ", shuffledNumbers[i], i)
			}
			if err != nil {
				fmt.Println("Breakpoint: Deletion encountered an error. 💢 The number is ", shuffledNumbers[i], i)
			}
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
