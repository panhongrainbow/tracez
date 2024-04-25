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
		var randomNumbers = []int64{189, 1916, 1293, 1856, 226, 748, 180, 1714, 1863, 1906, 120, 878, 91, 895, 1314, 1583, 1585, 1834, 31, 958, 1171, 861, 1024, 1448, 800, 1762, 980, 177, 1318, 1542, 1849, 38, 202, 1477, 901, 1474, 1581, 294, 1924, 1569, 278, 922, 1559, 454, 692, 874, 1157, 1684, 1166, 957, 713, 707, 1401, 1509, 1374, 1810, 1300, 1303, 1412, 885, 729, 1034, 1304, 1201, 1237, 1439, 1594, 1530, 1112, 446, 365, 1215, 1855, 1622, 515, 1022, 1747, 1343, 1879, 1480, 1611, 1380, 1379, 1849, 95, 905, 1055, 1065, 925, 1533, 733, 393, 1451, 189, 784, 116, 1934, 725, 229, 888, 1648, 1854, 1382, 1289, 1731, 273, 308, 1123, 777, 950, 1890, 1790, 366, 1957, 1887, 200, 626, 635, 269, 1296, 473, 1246, 1017, 773, 1071, 1377, 880, 229, 235, 706, 423, 1008, 1367, 532, 129, 1341, 292, 349, 1404, 960, 1634, 1528, 779, 1216, 1408, 1109, 81, 1122, 797, 1544, 378, 867, 503, 1395, 1276, 1661, 227, 11, 997, 780, 1673, 351, 1823, 1546, 707, 979, 731, 1922, 1455, 1799, 1038, 117, 1573, 244, 1443, 518, 896, 1255, 107, 1175, 381, 1904, 1869, 42, 1574, 1454, 592, 877, 900, 533, 629, 1672, 1386, 833, 10, 984, 46, 1323, 1320, 1717}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{1065, 227, 1581, 1318, 177, 729, 235, 473, 1343, 381, 874, 1166, 515, 1834, 349, 1304, 1451, 880, 1583, 1922, 635, 454, 38, 91, 1569, 901, 1201, 706, 1382, 1293, 1533, 1684, 867, 878, 877, 957, 777, 922, 1237, 1957, 1673, 780, 997, 592, 800, 1320, 1480, 692, 1112, 1528, 1717, 189, 365, 861, 129, 1890, 269, 229, 773, 11, 46, 626, 707, 1157, 229, 1408, 1672, 797, 1916, 707, 1401, 1790, 1341, 1594, 1404, 1869, 1255, 200, 1648, 1443, 1904, 244, 116, 1008, 308, 273, 779, 784, 950, 1879, 226, 1246, 1799, 1574, 1661, 351, 1017, 748, 1823, 1022, 1477, 1530, 984, 1849, 1542, 1034, 1474, 895, 518, 446, 107, 731, 81, 1323, 1448, 1573, 1024, 1055, 1855, 1810, 120, 366, 202, 423, 1849, 1854, 278, 393, 1455, 1747, 1123, 713, 1386, 1374, 958, 980, 10, 1544, 896, 888, 1296, 925, 1071, 117, 1887, 1856, 1175, 905, 833, 725, 979, 1559, 180, 1377, 960, 1509, 1300, 31, 189, 1122, 1863, 1109, 1634, 1439, 1276, 292, 1289, 532, 1934, 885, 1546, 1303, 1216, 1622, 378, 1380, 1395, 1924, 1038, 1714, 1171, 1314, 1412, 629, 1906, 1379, 42, 1454, 533, 900, 1762, 1367, 733, 1611, 95, 1215, 1731, 294, 503, 1585}

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
			if shuffledNumbers[i] == 960 {
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

