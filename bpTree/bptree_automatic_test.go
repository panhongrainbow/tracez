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

			// 显示目前的删除值
			value := numbersForDeleting[i]
			fmt.Println(i, value)

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
		// 数量一百的例子
		// Generate random data for insertion.
		var randomNumbers = []int64{1876, 1612, 189, 914, 948, 1500, 682, 1646, 1472, 1320, 810, 1163, 221, 435, 344, 1658, 93, 1891, 1281, 1660, 589, 1269, 1397, 1760, 1860, 80, 1450, 1548, 894, 1939, 1639, 1290, 805, 402, 337, 380, 590, 1024, 1151, 1610, 853, 680, 368, 632, 1101, 660, 1422, 1752, 938, 1552, 196, 1243, 1079, 395, 587, 622, 429, 1211, 1759, 1940, 1984, 1126, 592, 1308, 618, 1194, 186, 776, 170, 122, 975, 1966, 1060, 1419, 1592, 694, 1345, 763, 1632, 247, 621, 1413, 1083, 455, 1975, 824, 1786, 1272, 1734, 608, 168, 167, 895, 1646, 975, 276, 514, 1261, 1415, 1812}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{514, 1194, 1163, 763, 622, 337, 1413, 824, 1734, 1632, 402, 1450, 589, 948, 186, 455, 1320, 1860, 1269, 1472, 80, 1610, 1024, 1261, 587, 435, 1308, 853, 894, 1422, 167, 590, 1759, 122, 1126, 1548, 1281, 1060, 247, 1552, 1975, 170, 1658, 1079, 592, 1272, 1500, 1345, 276, 1760, 1612, 1101, 1786, 380, 618, 1812, 344, 660, 975, 805, 429, 1419, 938, 1891, 1639, 1984, 395, 1397, 93, 189, 1243, 196, 975, 1415, 1660, 1876, 682, 608, 1290, 1151, 895, 1592, 168, 914, 1083, 680, 694, 1646, 1646, 221, 1940, 1966, 1752, 632, 1211, 1939, 368, 621, 810, 776}

		// 数量二百的例子
		// Generate random data for insertion.
		// var randomNumbers = []int64{240, 662, 205, 1579, 1924, 888, 844, 263, 945, 114, 30, 434, 992, 652, 1774, 1284, 1449, 1398, 698, 1699, 453, 221, 1488, 161, 1423, 1188, 466, 1258, 1829, 671, 807, 1401, 1704, 1618, 944, 892, 1824, 300, 1237, 1078, 448, 1681, 1260, 713, 170, 526, 1859, 500, 1514, 832, 1416, 1095, 1818, 1122, 1991, 1350, 1372, 401, 1237, 797, 476, 1630, 977, 111, 12, 415, 1283, 1866, 984, 1271, 559, 741, 1497, 1956, 1842, 1474, 1272, 726, 516, 347, 1480, 540, 1876, 1832, 779, 673, 1914, 903, 952, 453, 1837, 304, 1460, 44, 172, 972, 1284, 964, 350, 932, 666, 1496, 408, 1226, 763, 968, 1968, 1533, 603, 315, 392, 392, 437, 824, 569, 1431, 1386, 1512, 1073, 1336, 166, 1845, 1114, 491, 1928, 1403, 262, 966, 84, 117, 945, 1883, 80, 1494, 1263, 924, 1392, 1461, 327, 676, 1751, 660, 1568, 1853, 601, 1762, 647, 124, 283, 535, 1992, 1580, 1291, 412, 1769, 37, 1093, 1602, 1218, 487, 1290, 933, 1556, 1176, 1905, 852, 1858, 995, 1734, 1017, 612, 1928, 763, 553, 1342, 1078, 530, 1145, 188, 229, 1490, 100, 1976, 528, 698, 1848, 361, 1636, 1597, 287, 1765, 1359, 1529, 1138, 1016, 432, 1080, 1604, 966, 1767}
		// Generate random data for deletion.
		// var shuffledNumbers = []int64{453, 205, 1138, 287, 1866, 1431, 1145, 1604, 1905, 453, 1597, 797, 166, 952, 844, 1398, 1765, 1284, 1496, 553, 944, 124, 1876, 221, 111, 1580, 528, 559, 1423, 741, 117, 1968, 170, 1291, 1818, 1474, 763, 476, 1829, 992, 1497, 516, 1681, 392, 283, 1734, 603, 1260, 1336, 1602, 1842, 1237, 1488, 1461, 44, 392, 1392, 1751, 30, 995, 1630, 1699, 726, 984, 1386, 1350, 1529, 601, 1078, 1226, 437, 1480, 1272, 824, 698, 240, 892, 1263, 1774, 1416, 1568, 1767, 779, 1533, 535, 304, 263, 491, 415, 903, 966, 1449, 161, 1853, 1991, 763, 361, 1928, 1016, 1832, 1845, 1928, 1403, 1114, 1258, 100, 526, 1490, 1579, 500, 1556, 666, 968, 945, 487, 1017, 327, 37, 412, 1093, 977, 1976, 1290, 408, 932, 1618, 660, 1859, 114, 964, 652, 84, 1284, 972, 1176, 888, 1095, 172, 1824, 832, 1837, 966, 676, 347, 401, 1512, 80, 300, 1956, 1080, 713, 1883, 1848, 1704, 807, 1769, 647, 1858, 1636, 1237, 350, 1460, 1494, 466, 1078, 852, 1073, 1188, 1122, 924, 530, 434, 671, 540, 1924, 1218, 229, 262, 1283, 12, 612, 1359, 945, 1514, 1342, 1992, 569, 1271, 1914, 1762, 673, 1372, 662, 448, 315, 933, 432, 698, 188, 1401}

		// Initialize B-tree.
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
			if shuffledNumbers[i] == 123 {
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
