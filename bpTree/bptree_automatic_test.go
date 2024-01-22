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
		var randomNumbers = []int64{1615, 1138, 1423, 1569, 1076, 76, 1799, 341, 1647, 1599, 825, 1195, 1712, 1554, 272, 1610, 1965, 490, 1626, 1145, 1002, 991, 975, 12, 891, 458, 1651, 1254, 1321, 1347, 1222, 1787, 1518, 720, 535, 1641, 1534, 742, 1453, 230, 544, 551, 56, 1505, 1848, 1780, 1746, 1645, 1478, 720, 1082, 1042, 604, 1806, 1641, 1541, 612, 1920, 1502, 1092, 1396, 1231, 384, 830, 923, 201, 456, 659, 161, 592, 1668, 893, 197, 564, 551, 1245, 316, 1444, 982, 817, 1952, 1931, 589, 279, 616, 869, 204, 406, 1010, 1066, 99, 1055, 1574, 132, 1839, 356, 915, 1194, 1103, 1502}
		// Generate random data for deletion.
		var shuffledNumbers = []int64{161, 1245, 197, 490, 1965, 1839, 1626, 1092, 1347, 99, 1423, 1712, 991, 1138, 551, 1541, 341, 1610, 1534, 230, 564, 1799, 1076, 279, 1254, 1195, 201, 1641, 923, 893, 1002, 825, 384, 1444, 612, 1103, 592, 1066, 659, 1641, 76, 1651, 817, 1396, 1453, 830, 456, 1806, 272, 1920, 1746, 1042, 1194, 56, 720, 720, 316, 1668, 1478, 616, 604, 1145, 1599, 1574, 1615, 1554, 1931, 1082, 975, 1502, 1569, 551, 1647, 1645, 1518, 1848, 891, 589, 1505, 1231, 1952, 132, 1055, 356, 1010, 1787, 1780, 458, 204, 1502, 742, 535, 982, 1321, 1222, 406, 12, 869, 544, 915}

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
