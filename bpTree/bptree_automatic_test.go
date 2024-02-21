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
	randomQuantity = 200

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

			// 显示目前的删除值
			value := numbersForDeleting[i]
			fmt.Println(i, value)

			// Deleting data entries continuously.
			deleted, _, _, err := root.RemoveValue(BpItem{Key: numbersForDeleting[i]})
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
	t.Run("Manually Identify B Plus Tree Operation Errors", func(t *testing.T) {
		// 数量二百的例子
		// Generate random data for insertion.
		var randomNumbers = []int64{808, 1380, 754, 1381, 1720, 1007, 1921, 1911, 725, 305, 765, 535, 1558, 140, 1858, 775, 915, 1564, 1486, 185, 238, 1779, 1729, 1193, 138, 101, 1467, 1906, 223, 1160, 1850, 1125, 677, 1997, 190, 857, 676, 1437, 1088, 1876, 174, 1451, 1437, 885, 1879, 342, 762, 1140, 1672, 1069, 1795, 1858, 1047, 1846, 1490, 1545, 471, 331, 1735, 563, 628, 991, 1859, 1568, 1753, 1133, 956, 324, 867, 1045, 1091, 1992, 336, 615, 593, 1743, 558, 1431, 182, 1194, 1031, 395, 652, 1052, 711, 1404, 242, 450, 765, 817, 898, 64, 1213, 372, 1808, 44, 1270, 1269, 1095, 37, 946, 511, 917, 1486, 363, 173, 1708, 1931, 1696, 1186, 230, 349, 1371, 142, 1623, 397, 363, 1200, 1404, 1523, 1972, 361, 532, 318, 1255, 865, 1488, 1369, 817, 468, 1737, 1070, 293, 311, 1415, 1227, 576, 153, 32, 1930, 84, 1508, 1253, 544, 1991, 459, 1298, 1166, 1487, 349, 1136, 564, 10, 1843, 596, 384, 1732, 984, 1682, 1226, 507, 68, 90, 1746, 1720, 1403, 187, 298, 1110, 1645, 1432, 945, 1245, 1216, 676, 644, 222, 359, 1070, 1622, 831, 727, 1578, 2000, 769, 527, 774, 314, 900, 809, 169, 790, 404, 894, 1681, 425, 1613, 851, 1835, 86}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{10, 397, 984, 775, 676, 1488, 395, 1404, 363, 37, 361, 223, 1931, 527, 596, 64, 1523, 1622, 182, 1166, 900, 1972, 851, 725, 1564, 140, 532, 324, 1737, 1753, 142, 1432, 535, 885, 230, 1876, 1578, 1487, 1843, 363, 349, 677, 90, 1645, 991, 831, 311, 1045, 945, 1415, 1486, 1069, 2000, 1110, 727, 1681, 404, 1735, 1007, 1997, 765, 1729, 898, 563, 1270, 857, 644, 762, 809, 1545, 187, 1906, 336, 1992, 1623, 1808, 1369, 1911, 1490, 342, 558, 1846, 185, 769, 1558, 1732, 222, 946, 1133, 1720, 468, 1930, 1140, 817, 1070, 765, 84, 32, 1486, 1088, 1200, 1255, 1380, 1052, 101, 1269, 1720, 676, 174, 1467, 1682, 1696, 1795, 774, 1298, 1213, 1404, 1859, 817, 1746, 576, 1437, 544, 331, 1672, 1253, 1568, 1879, 1708, 1613, 511, 507, 1226, 242, 425, 298, 1031, 1186, 1227, 459, 1070, 1136, 169, 318, 153, 1245, 790, 1095, 293, 652, 1160, 1125, 1451, 1850, 471, 1779, 917, 372, 1835, 1437, 44, 384, 1858, 867, 628, 86, 1381, 808, 190, 68, 1091, 615, 1193, 138, 1858, 349, 314, 593, 1921, 711, 1508, 1047, 894, 915, 1403, 1431, 956, 1216, 564, 1371, 1991, 1194, 450, 754, 173, 305, 865, 359, 1743, 238}

		// Initialize B plus tree.
		root := NewBpTree(5)

		// Start inserting data.
		for i := 0; i < randomQuantity; i++ {
			// Insert data entries continuously.
			root.InsertValue(BpItem{Key: randomNumbers[i]})
		}

		// Start deleting data.
		for i := 0; i < randomQuantity; i++ {

			// 中断检验
			value := shuffledNumbers[i]
			fmt.Println(i, value)
			if shuffledNumbers[i] == 1437 {
				fmt.Print()
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
