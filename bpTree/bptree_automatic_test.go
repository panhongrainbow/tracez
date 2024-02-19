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
	randomQuantity = 150

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
		// 数量一百的例子
		// Generate random data for insertion.
		// var randomNumbers = []int64{183, 573, 767, 921, 565, 1949, 594, 1095, 1786, 1700, 809, 1603, 647, 698, 1207, 1794, 1856, 534, 606, 1122, 422, 1474, 243, 264, 1991, 791, 143, 1900, 1145, 936, 188, 523, 798, 200, 1209, 191, 575, 1531, 1338, 102, 1340, 1774, 1602, 1595, 368, 1902, 1418, 48, 245, 655, 226, 1329, 1266, 1388, 219, 337, 440, 547, 813, 1158, 287, 1557, 939, 1262, 1253, 512, 1260, 74, 1773, 1571, 41, 1961, 1727, 1290, 1022, 796, 1351, 521, 992, 867, 1063, 1950, 983, 216, 577, 1890, 1677, 26, 1834, 1283, 1409, 1191, 1392, 414, 74, 1052, 392, 1026, 1383, 1791}
		// Generate random data for deletion.
		// var shuffledNumbers = []int64{1052, 1602, 1603, 1351, 523, 1900, 1727, 74, 512, 48, 245, 1290, 243, 1949, 1834, 216, 1773, 791, 74, 1961, 1026, 188, 1791, 577, 226, 1122, 798, 1262, 1794, 767, 1474, 287, 191, 1253, 200, 41, 1266, 1557, 1022, 1191, 647, 992, 1774, 1095, 1383, 1856, 573, 565, 1145, 921, 143, 534, 1991, 698, 809, 1677, 422, 867, 1392, 1531, 796, 368, 264, 1700, 1409, 1329, 1786, 414, 1158, 26, 813, 1388, 655, 1283, 1207, 606, 1063, 1418, 219, 983, 392, 594, 440, 337, 1340, 1595, 939, 183, 102, 1260, 1338, 1890, 575, 1950, 1209, 1902, 521, 936, 547, 1571}

		// 数量二百的例子
		// Generate random data for insertion.
		var randomNumbers = []int64{892, 87, 1192, 1268, 872, 596, 1115, 481, 609, 815, 392, 368, 1799, 1177, 1027, 1488, 262, 422, 1107, 1312, 1527, 1866, 1114, 1227, 1267, 1027, 305, 531, 1601, 1055, 385, 545, 1584, 884, 1348, 1614, 1851, 1683, 497, 40, 1181, 1108, 820, 229, 91, 1132, 1307, 881, 31, 869, 819, 412, 1945, 1274, 230, 1460, 673, 1879, 1124, 60, 1759, 115, 1798, 1910, 886, 889, 1775, 1067, 812, 1531, 1178, 158, 907, 482, 537, 1765, 1865, 146, 1285, 803, 129, 1164, 578, 614, 1455, 553, 1181, 948, 1366, 1581, 500, 301, 1373, 172, 633, 840, 1731, 428, 711, 165, 812, 1116, 1210, 410, 304, 432, 1383, 1872, 1625, 1222, 402, 1788, 313, 1459, 667, 1389, 1671, 1120, 1465, 743, 905, 1932, 333, 1409, 1168, 591, 210, 1633, 1116, 1698, 1953, 1651, 78, 1848, 1868, 1801, 163, 247, 478, 1672, 1936, 84, 571, 1065, 1863, 1977, 1170, 164, 141, 1900}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{609, 1116, 1164, 812, 1312, 392, 819, 1055, 1115, 1651, 892, 1348, 545, 164, 478, 1953, 1977, 1459, 1868, 1192, 578, 1765, 1409, 1601, 163, 869, 812, 304, 905, 1788, 305, 1465, 1177, 1801, 1067, 84, 1114, 91, 1170, 230, 1107, 1460, 872, 1731, 743, 301, 881, 1181, 1268, 1210, 333, 1120, 1698, 428, 667, 1798, 571, 1581, 889, 313, 165, 146, 210, 1307, 129, 78, 1614, 1366, 1274, 1866, 1872, 141, 481, 711, 840, 1759, 482, 884, 1027, 1932, 1900, 115, 1863, 907, 1527, 247, 422, 803, 596, 673, 1775, 1178, 1285, 500, 1389, 1672, 1181, 1116, 60, 158, 1671, 614, 1222, 412, 1531, 633, 1065, 172, 368, 262, 1799, 432, 1910, 1455, 1936, 1373, 886, 1865, 1945, 402, 1168, 591, 1488, 1124, 1848, 1132, 948, 1851, 815, 1879, 497, 531, 1625, 1267, 1108, 1227, 87, 537, 1383, 31, 1027, 553, 229, 410, 1584, 40, 1633, 385, 820, 1683}

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
			if shuffledNumbers[i] == 1625 {
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
