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
	randomMax = 2000

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

		fmt.Println()
	})
	// Automated random testing for B+ tree.
	t.Run("Manually Identify B+ Tree Operation Errors", func(t *testing.T) {
		// Generate random data for insertion.
		var randomNumbers = []int64{1573, 900, 1826, 1082, 1670, 289, 1109, 658, 789, 857, 1069, 1263, 46, 624, 607, 846, 602, 928, 755, 1359, 1981, 2000, 411, 1720, 176, 110, 1125, 113, 995, 1363, 240, 967, 1288, 1906, 913, 1492, 1980, 1974, 914, 1112, 1567, 601, 700, 772, 1246, 848, 1407, 227, 971, 86, 406, 482, 1244, 1192, 1783, 1710, 1901, 1628, 1977, 542, 1624, 874, 271, 649, 587, 548, 679, 88, 1137, 1898, 574, 1605, 1813, 721, 1472, 1410, 445, 1405, 1072, 1261, 725, 384, 506, 965, 440, 505, 390, 166, 1023, 439, 1440, 937, 1358, 1479, 1234, 1168, 1324, 38, 1322, 1088}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{725, 1624, 1670, 965, 1405, 1137, 1898, 406, 1479, 574, 1720, 445, 1234, 86, 1288, 439, 928, 271, 506, 1906, 874, 440, 1322, 772, 2000, 1168, 1901, 913, 624, 1023, 390, 1088, 1072, 1826, 971, 1440, 113, 846, 289, 679, 1977, 1974, 1628, 548, 110, 1980, 848, 1981, 1363, 700, 38, 937, 1112, 176, 1082, 607, 240, 755, 166, 1407, 1710, 658, 967, 411, 1359, 1261, 505, 1492, 88, 1358, 482, 995, 789, 384, 1069, 587, 601, 1192, 602, 227, 46, 914, 1783, 721, 1244, 1109, 1324, 857, 1573, 1410, 1813, 1567, 649, 1605, 1246, 542, 1263, 1125, 900, 1472}

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
			if shuffledNumbers[i] == 900 { // 检验
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
