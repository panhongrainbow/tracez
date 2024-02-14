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
		var randomNumbers = []int64{183, 573, 767, 921, 565, 1949, 594, 1095, 1786, 1700, 809, 1603, 647, 698, 1207, 1794, 1856, 534, 606, 1122, 422, 1474, 243, 264, 1991, 791, 143, 1900, 1145, 936, 188, 523, 798, 200, 1209, 191, 575, 1531, 1338, 102, 1340, 1774, 1602, 1595, 368, 1902, 1418, 48, 245, 655, 226, 1329, 1266, 1388, 219, 337, 440, 547, 813, 1158, 287, 1557, 939, 1262, 1253, 512, 1260, 74, 1773, 1571, 41, 1961, 1727, 1290, 1022, 796, 1351, 521, 992, 867, 1063, 1950, 983, 216, 577, 1890, 1677, 26, 1834, 1283, 1409, 1191, 1392, 414, 74, 1052, 392, 1026, 1383, 1791}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{1052, 1602, 1603, 1351, 523, 1900, 1727, 74, 512, 48, 245, 1290, 243, 1949, 1834, 216, 1773, 791, 74, 1961, 1026, 188, 1791, 577, 226, 1122, 798, 1262, 1794, 767, 1474, 287, 191, 1253, 200, 41, 1266, 1557, 1022, 1191, 647, 992, 1774, 1095, 1383, 1856, 573, 565, 1145, 921, 143, 534, 1991, 698, 809, 1677, 422, 867, 1392, 1531, 796, 368, 264, 1700, 1409, 1329, 1786, 414, 1158, 26, 813, 1388, 655, 1283, 1207, 606, 1063, 1418, 219, 983, 392, 594, 440, 337, 1340, 1595, 939, 183, 102, 1260, 1338, 1890, 575, 1950, 1209, 1902, 521, 936, 547, 1571}

		// 数量二百的例子
		// Generate random data for insertion.
		// var randomNumbers = []int64{240, 662, 205, 1579, 1924, 888, 844, 263, 945, 114, 30, 434, 992, 652, 1774, 1284, 1449, 1398, 698, 1699, 453, 221, 1488, 161, 1423, 1188, 466, 1258, 1829, 671, 807, 1401, 1704, 1618, 944, 892, 1824, 300, 1237, 1078, 448, 1681, 1260, 713, 170, 526, 1859, 500, 1514, 832, 1416, 1095, 1818, 1122, 1991, 1350, 1372, 401, 1237, 797, 476, 1630, 977, 111, 12, 415, 1283, 1866, 984, 1271, 559, 741, 1497, 1956, 1842, 1474, 1272, 726, 516, 347, 1480, 540, 1876, 1832, 779, 673, 1914, 903, 952, 453, 1837, 304, 1460, 44, 172, 972, 1284, 964, 350, 932, 666, 1496, 408, 1226, 763, 968, 1968, 1533, 603, 315, 392, 392, 437, 824, 569, 1431, 1386, 1512, 1073, 1336, 166, 1845, 1114, 491, 1928, 1403, 262, 966, 84, 117, 945, 1883, 80, 1494, 1263, 924, 1392, 1461, 327, 676, 1751, 660, 1568, 1853, 601, 1762, 647, 124, 283, 535, 1992, 1580, 1291, 412, 1769, 37, 1093, 1602, 1218, 487, 1290, 933, 1556, 1176, 1905, 852, 1858, 995, 1734, 1017, 612, 1928, 763, 553, 1342, 1078, 530, 1145, 188, 229, 1490, 100, 1976, 528, 698, 1848, 361, 1636, 1597, 287, 1765, 1359, 1529, 1138, 1016, 432, 1080, 1604, 966, 1767}
		// Generate random data for deletion.
		// var shuffledNumbers = []int64{453, 205, 1138, 287, 1866, 1431, 1145, 1604, 1905, 453, 1597, 797, 166, 952, 844, 1398, 1765, 1284, 1496, 553, 944, 124, 1876, 221, 111, 1580, 528, 559, 1423, 741, 117, 1968, 170, 1291, 1818, 1474, 763, 476, 1829, 992, 1497, 516, 1681, 392, 283, 1734, 603, 1260, 1336, 1602, 1842, 1237, 1488, 1461, 44, 392, 1392, 1751, 30, 995, 1630, 1699, 726, 984, 1386, 1350, 1529, 601, 1078, 1226, 437, 1480, 1272, 824, 698, 240, 892, 1263, 1774, 1416, 1568, 1767, 779, 1533, 535, 304, 263, 491, 415, 903, 966, 1449, 161, 1853, 1991, 763, 361, 1928, 1016, 1832, 1845, 1928, 1403, 1114, 1258, 100, 526, 1490, 1579, 500, 1556, 666, 968, 945, 487, 1017, 327, 37, 412, 1093, 977, 1976, 1290, 408, 932, 1618, 660, 1859, 114, 964, 652, 84, 1284, 972, 1176, 888, 1095, 172, 1824, 832, 1837, 966, 676, 347, 401, 1512, 80, 300, 1956, 1080, 713, 1883, 1848, 1704, 807, 1769, 647, 1858, 1636, 1237, 350, 1460, 1494, 466, 1078, 852, 1073, 1188, 1122, 924, 530, 434, 671, 540, 1924, 1218, 229, 262, 1283, 12, 612, 1359, 945, 1514, 1342, 1992, 569, 1271, 1914, 1762, 673, 1372, 662, 448, 315, 933, 432, 698, 188, 1401}

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
			if shuffledNumbers[i] == 219 { // 在这里要把索引值由 710 改成 715，之后在删除 599 时，会有 这里程式还没写完2 的警告
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
