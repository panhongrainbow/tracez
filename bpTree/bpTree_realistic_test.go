package bpTree

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_Check_Btree_In_Real_World(t *testing.T) {
	t.Run("Tests inserting B-tree with a width of 3.", func(t *testing.T) {
		// Initialize B-tree.
		root := NewBpTree(3)
		// Insert 50 data entries continuously.
		root.InsertValue(BpItem{Key: 40})
		root.InsertValue(BpItem{Key: 38})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 81})
		root.InsertValue(BpItem{Key: 98})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 30})
		root.InsertValue(BpItem{Key: 67})
		root.InsertValue(BpItem{Key: 35})
		root.InsertValue(BpItem{Key: 89})
		root.InsertValue(BpItem{Key: 96})
		root.InsertValue(BpItem{Key: 78})
		root.InsertValue(BpItem{Key: 95})
		root.InsertValue(BpItem{Key: 86})
		root.InsertValue(BpItem{Key: 19})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 99})
		root.InsertValue(BpItem{Key: 59})
		root.InsertValue(BpItem{Key: 49})
		root.InsertValue(BpItem{Key: 65})
		root.InsertValue(BpItem{Key: 37})
		root.InsertValue(BpItem{Key: 73})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 29})
		root.InsertValue(BpItem{Key: 97})
		root.InsertValue(BpItem{Key: 77})
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 18})
		root.InsertValue(BpItem{Key: 69})
		root.InsertValue(BpItem{Key: 46})
		root.InsertValue(BpItem{Key: 72})
		root.InsertValue(BpItem{Key: 6})
		root.InsertValue(BpItem{Key: 36})
		root.InsertValue(BpItem{Key: 22})
		root.InsertValue(BpItem{Key: 56})
		root.InsertValue(BpItem{Key: 62})
		root.InsertValue(BpItem{Key: 23})
		root.InsertValue(BpItem{Key: 94})
		root.InsertValue(BpItem{Key: 11})
		root.InsertValue(BpItem{Key: 71})
		root.InsertValue(BpItem{Key: 34})
		root.InsertValue(BpItem{Key: 13})
		root.InsertValue(BpItem{Key: 100})
		root.InsertValue(BpItem{Key: 60})
		root.InsertValue(BpItem{Key: 24})
		root.InsertValue(BpItem{Key: 91})
		root.InsertValue(BpItem{Key: 25})
		root.InsertValue(BpItem{Key: 66})
		root.InsertValue(BpItem{Key: 50})
		root.InsertValue(BpItem{Key: 80})

		// Check the distribution of the entire B-tree.
		assert.Equal(t, []int64{22, 40}, root.root.Index, "Top-level index is incorrect")

		assert.Equal(t, []int64{10}, root.root.IndexNodes[0].Index, "Index error for the 1st index node on the 1st level down")
		assert.Equal(t, []int64{30}, root.root.IndexNodes[1].Index, "Index error for the 2nd index node on the 1st level down")
		assert.Equal(t, []int64{67, 77}, root.root.IndexNodes[2].Index, "Index error for the 3rd index node on the 1st level down")

		assert.Equal(t, []int64{5}, root.root.IndexNodes[0].IndexNodes[0].Index, "Index error for the 1st index node on the 2nd level down")
		assert.Equal(t, []int64{13}, root.root.IndexNodes[0].IndexNodes[1].Index, "Index error for the 2nd index node on the 2nd level down")
		assert.Equal(t, []int64{24}, root.root.IndexNodes[1].IndexNodes[0].Index, "Index error for the 3rd index node on the 2nd level down")
		assert.Equal(t, []int64{36}, root.root.IndexNodes[1].IndexNodes[1].Index, "Index error for the 4th index node on the 2nd level down")
		assert.Equal(t, []int64{59}, root.root.IndexNodes[2].IndexNodes[0].Index, "Index error for the 5th index node on the 2nd level down")
		assert.Equal(t, []int64{71}, root.root.IndexNodes[2].IndexNodes[1].Index, "Index error for the 6th index node on the 2nd level down")
		assert.Equal(t, []int64{89, 96}, root.root.IndexNodes[2].IndexNodes[2].Index, "Index error for the 7th index node on the 2nd level down")

		assert.Equal(t, []int64{4}, root.root.IndexNodes[0].IndexNodes[0].IndexNodes[0].Index, "Index error for the 1st index node on the 3rd level down")
		assert.Equal(t, []int64{6}, root.root.IndexNodes[0].IndexNodes[0].IndexNodes[1].Index, "Index error for the 2nd index node on the 3rd level down")
		assert.Equal(t, []int64{11}, root.root.IndexNodes[0].IndexNodes[1].IndexNodes[0].Index, "Index error for the 3rd index node on the 3rd level down")
		assert.Equal(t, []int64{19}, root.root.IndexNodes[0].IndexNodes[1].IndexNodes[1].Index, "Index error for the 4th index node on the 3rd level down")
		assert.Equal(t, []int64{23}, root.root.IndexNodes[1].IndexNodes[0].IndexNodes[0].Index, "Index error for the 5th index node on the 3rd level down")
		assert.Equal(t, []int64{25}, root.root.IndexNodes[1].IndexNodes[0].IndexNodes[1].Index, "Index error for the 6th index node on the 3rd level down")
		assert.Equal(t, []int64{35}, root.root.IndexNodes[1].IndexNodes[1].IndexNodes[0].Index, "Index error for the 7th index node on the 3rd level down")
		assert.Equal(t, []int64{38}, root.root.IndexNodes[1].IndexNodes[1].IndexNodes[1].Index, "Index error for the 8th index node on the 3rd level down")
		assert.Equal(t, []int64{49, 50}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[0].Index, "Index error for the 9th index node on the 3rd level down")
		assert.Equal(t, []int64{62, 65}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[1].Index, "Index error for the 10th index node on the 3rd level down")
		assert.Equal(t, []int64{69}, root.root.IndexNodes[2].IndexNodes[1].IndexNodes[0].Index, "Index error for the 11th index node on the 3rd level down")
		assert.Equal(t, []int64{73}, root.root.IndexNodes[2].IndexNodes[1].IndexNodes[1].Index, "Index error for the 12th index node on the 3rd level down")
		assert.Equal(t, []int64{78, 81}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[0].Index, "Index error for the 13th index node on the 3rd level down")
		assert.Equal(t, []int64{94}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[1].Index, "Index error for the 14th index node on the 3rd level down")
		assert.Equal(t, []int64{98, 99}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[2].Index, "Index error for the 15th index node on the 3rd level down")

		assert.Equal(t, []BpItem{{Key: 1}}, root.root.IndexNodes[0].IndexNodes[0].IndexNodes[0].DataNodes[0].Items, "Data error for the 1st data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 4}}, root.root.IndexNodes[0].IndexNodes[0].IndexNodes[0].DataNodes[1].Items, "Data error for the 2nd data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 5}}, root.root.IndexNodes[0].IndexNodes[0].IndexNodes[1].DataNodes[0].Items, "Data error for the 1st data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 6}, {Key: 9}}, root.root.IndexNodes[0].IndexNodes[0].IndexNodes[1].DataNodes[1].Items, "Data error for the 2nd data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 10}}, root.root.IndexNodes[0].IndexNodes[1].IndexNodes[0].DataNodes[0].Items, "Data error for the 3rd data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 11}}, root.root.IndexNodes[0].IndexNodes[1].IndexNodes[0].DataNodes[1].Items, "Data error for the 4th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 13}, {Key: 18}}, root.root.IndexNodes[0].IndexNodes[1].IndexNodes[1].DataNodes[0].Items, "Data error for the 5th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 19}}, root.root.IndexNodes[0].IndexNodes[1].IndexNodes[1].DataNodes[1].Items, "Data error for the 6th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 22}}, root.root.IndexNodes[1].IndexNodes[0].IndexNodes[0].DataNodes[0].Items, "Data error for the 7th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 23}}, root.root.IndexNodes[1].IndexNodes[0].IndexNodes[0].DataNodes[1].Items, "Data error for the 8th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 24}}, root.root.IndexNodes[1].IndexNodes[0].IndexNodes[1].DataNodes[0].Items, "Data error for the 9th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 25}, {Key: 29}}, root.root.IndexNodes[1].IndexNodes[0].IndexNodes[1].DataNodes[1].Items, "Data error for the 10th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 30}, {Key: 34}}, root.root.IndexNodes[1].IndexNodes[1].IndexNodes[0].DataNodes[0].Items, "Data error for the 11th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 35}}, root.root.IndexNodes[1].IndexNodes[1].IndexNodes[0].DataNodes[1].Items, "Data error for the 12th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 36}, {Key: 37}}, root.root.IndexNodes[1].IndexNodes[1].IndexNodes[1].DataNodes[0].Items, "Data error for the 13th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 38}}, root.root.IndexNodes[1].IndexNodes[1].IndexNodes[1].DataNodes[1].Items, "Data error for the 14th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 40}, {Key: 46}}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[0].DataNodes[0].Items, "Data error for the 15th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 49}}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[0].DataNodes[1].Items, "Data error for the 16th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 50}, {Key: 56}}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[0].DataNodes[2].Items, "Data error for the 17th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 59}, {Key: 60}}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[1].DataNodes[0].Items, "Data error for the 18th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 62}}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[1].DataNodes[1].Items, "Data error for the 19th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 65}, {Key: 66}}, root.root.IndexNodes[2].IndexNodes[0].IndexNodes[1].DataNodes[2].Items, "Data error for the 20th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 67}}, root.root.IndexNodes[2].IndexNodes[1].IndexNodes[0].DataNodes[0].Items, "Data error for the 21th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 69}}, root.root.IndexNodes[2].IndexNodes[1].IndexNodes[0].DataNodes[1].Items, "Data error for the 22th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 71}, {Key: 72}}, root.root.IndexNodes[2].IndexNodes[1].IndexNodes[1].DataNodes[0].Items, "Data error for the 23th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 73}}, root.root.IndexNodes[2].IndexNodes[1].IndexNodes[1].DataNodes[1].Items, "Data error for the 24th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 77}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[0].DataNodes[0].Items, "Data error for the 25th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 78}, {Key: 80}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[0].DataNodes[1].Items, "Data error for the 26th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 81}, {Key: 86}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[0].DataNodes[2].Items, "Data error for the 26th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 89}, {Key: 91}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[1].DataNodes[0].Items, "Data error for the 27th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 94}, {Key: 95}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[1].DataNodes[1].Items, "Data error for the 28th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 96}, {Key: 97}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[2].DataNodes[0].Items, "Data error for the 29th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 98}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[2].DataNodes[1].Items, "Data error for the 30th data node on the 4th level down")
		assert.Equal(t, []BpItem{{Key: 99}, {Key: 100}}, root.root.IndexNodes[2].IndexNodes[2].IndexNodes[2].DataNodes[2].Items, "Data error for the 31th data node on the 4th level down")

		// Retrieve the head node of the bottom-level Link List.
		head := root.root.BpDataHead()

		// Check the continuity of the bottom-level Link List.
		tests := []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{1}},
			{1, []int64{4}},
			{2, []int64{5}},
			{3, []int64{6, 9}},
			{4, []int64{10}},
			{5, []int64{11}},
			{6, []int64{13, 18}},
			{7, []int64{19}},
			{8, []int64{22}},
			{9, []int64{23}},
			{10, []int64{24}},
			{11, []int64{25, 29}},
			{12, []int64{30, 34}},
			{13, []int64{35}},
			{14, []int64{36, 37}},
			{15, []int64{38}},
			{16, []int64{40, 46}},
			{17, []int64{49}},
			{18, []int64{50, 56}},
			{19, []int64{59, 60}},
			{20, []int64{62}},
			{21, []int64{65, 66}},
			{22, []int64{67}},
			{23, []int64{69}},
			{24, []int64{71, 72}},
			{25, []int64{73}},
			{26, []int64{77}},
			{27, []int64{78, 80}},
			{28, []int64{81, 86}},
			{29, []int64{89, 91}},
			{30, []int64{94, 95}},
			{31, []int64{96, 97}},
			{32, []int64{98}},
			{33, []int64{99, 100}},
		}

		// Start checking the continuity of the Link List.
		for _, test := range tests {
			actualKeys := head.PrintNodeAscent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// Testing the continuity of the data in the opposite direction.
		tail := root.root.BpDataTail()

		tests = []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{100, 99}},
			{1, []int64{98}},
			{2, []int64{97, 96}},
			{3, []int64{95, 94}},
			{4, []int64{91, 89}},
			{5, []int64{86, 81}},
			{6, []int64{80, 78}},
			{7, []int64{77}},
			{8, []int64{73}},
			{9, []int64{72, 71}},
			{10, []int64{69}},
			{11, []int64{67}},
			{12, []int64{66, 65}},
			{13, []int64{62}},
			{14, []int64{60, 59}},
			{15, []int64{56, 50}},
			{16, []int64{49}},
			{17, []int64{46, 40}},
			{18, []int64{38}},
			{19, []int64{37, 36}},
			{20, []int64{35}},
			{21, []int64{34, 30}},
			{22, []int64{29, 25}},
			{23, []int64{24}},
			{24, []int64{23}},
			{25, []int64{22}},
			{26, []int64{19}},
			{27, []int64{18, 13}},
			{28, []int64{11}},
			{29, []int64{10}},
			{30, []int64{9, 6}},
			{31, []int64{5}},
			{32, []int64{4}},
			{33, []int64{1}},
		}

		for _, test := range tests {
			actualKeys := tail.PrintNodeDescent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// You can use the following functions to print the entire B-tree distribution.
		// root.root.Print() // Print the entire B Plus tree.
		// head.PrintAscent() // Print continuous data in the forward direction.
		// tail.PrintDescent() // Print continuous data in the reverse direction.
	})
	// For B-tree with a width of 4, perform tests inserting.
	t.Run("Tests inserting B-tree with a width of 4.", func(t *testing.T) {
		// Initialize B-tree.
		root := NewBpTree(4)
		// Insert 50 data entries continuously.
		root.InsertValue(BpItem{Key: 40})
		root.InsertValue(BpItem{Key: 38})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 81})
		root.InsertValue(BpItem{Key: 98})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 30})
		root.InsertValue(BpItem{Key: 67})
		root.InsertValue(BpItem{Key: 35})
		root.InsertValue(BpItem{Key: 89})
		root.InsertValue(BpItem{Key: 96})
		root.InsertValue(BpItem{Key: 78})
		root.InsertValue(BpItem{Key: 95})
		root.InsertValue(BpItem{Key: 86})
		root.InsertValue(BpItem{Key: 19})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 99})
		root.InsertValue(BpItem{Key: 59})
		root.InsertValue(BpItem{Key: 49})
		root.InsertValue(BpItem{Key: 65})
		root.InsertValue(BpItem{Key: 37})
		root.InsertValue(BpItem{Key: 73})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 29})
		root.InsertValue(BpItem{Key: 97})
		root.InsertValue(BpItem{Key: 77})
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 18})
		root.InsertValue(BpItem{Key: 69})
		root.InsertValue(BpItem{Key: 46})
		root.InsertValue(BpItem{Key: 72})
		root.InsertValue(BpItem{Key: 6})
		root.InsertValue(BpItem{Key: 36})
		root.InsertValue(BpItem{Key: 22})
		root.InsertValue(BpItem{Key: 56})
		root.InsertValue(BpItem{Key: 62})
		root.InsertValue(BpItem{Key: 23})
		root.InsertValue(BpItem{Key: 94})
		root.InsertValue(BpItem{Key: 11})
		root.InsertValue(BpItem{Key: 71})
		root.InsertValue(BpItem{Key: 34})
		root.InsertValue(BpItem{Key: 13})
		root.InsertValue(BpItem{Key: 100})
		root.InsertValue(BpItem{Key: 60})
		root.InsertValue(BpItem{Key: 24})
		root.InsertValue(BpItem{Key: 91})
		root.InsertValue(BpItem{Key: 25})
		root.InsertValue(BpItem{Key: 66})
		root.InsertValue(BpItem{Key: 50})
		root.InsertValue(BpItem{Key: 80})

		// Check the distribution of the entire B-tree.
		assert.Equal(t, []int64{40, 72}, root.root.Index, "Top-level index is incorrect")
		assert.Equal(t, []int64{19, 30}, root.root.IndexNodes[0].Index, "Index error for the 1st index node on the 1st level down")
		assert.Equal(t, []int64{62}, root.root.IndexNodes[1].Index, "Index error for the 2nd index node on the 1st level down")
		assert.Equal(t, []int64{81, 96}, root.root.IndexNodes[2].Index, "Index error for the 3rd index node on the 1st level down")

		assert.Equal(t, []int64{5, 10, 13}, root.root.IndexNodes[0].IndexNodes[0].Index, "Index error for the 1st data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 1}, {Key: 4}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[0].Items, "Data error for the 1st data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 5}, {Key: 6}, {Key: 9}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[1].Items, "Data error for the 2nd data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 10}, {Key: 11}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[2].Items, "Data error for the 3rd data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 13}, {Key: 18}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[3].Items, "Data error for the 4th data node on the 3rd level down")

		assert.Equal(t, []int64{23, 25}, root.root.IndexNodes[0].IndexNodes[1].Index, "Index error for the 2nd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 19}, {Key: 22}}, root.root.IndexNodes[0].IndexNodes[1].DataNodes[0].Items, "Data error for the 5th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 23}, {Key: 24}}, root.root.IndexNodes[0].IndexNodes[1].DataNodes[1].Items, "Data error for the 6th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 25}, {Key: 29}}, root.root.IndexNodes[0].IndexNodes[1].DataNodes[2].Items, "Data error for the 7th data node on the 3rd level down")

		assert.Equal(t, []int64{35, 37}, root.root.IndexNodes[0].IndexNodes[2].Index, "Index error for the 3rd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 30}, {Key: 34}}, root.root.IndexNodes[0].IndexNodes[2].DataNodes[0].Items, "Data error for the 8th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 35}, {Key: 36}}, root.root.IndexNodes[0].IndexNodes[2].DataNodes[1].Items, "Data error for the 9th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 37}, {Key: 38}}, root.root.IndexNodes[0].IndexNodes[2].DataNodes[2].Items, "Data error for the 10th data node on the 3rd level down")

		assert.Equal(t, []int64{49, 59}, root.root.IndexNodes[1].IndexNodes[0].Index, "Index error for the 4th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 40}, {Key: 46}}, root.root.IndexNodes[1].IndexNodes[0].DataNodes[0].Items, "Data error for the 11th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 49}, {Key: 50}, {Key: 56}}, root.root.IndexNodes[1].IndexNodes[0].DataNodes[1].Items, "Data error for the 12th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 59}, {Key: 60}}, root.root.IndexNodes[1].IndexNodes[0].DataNodes[2].Items, "Data error for the 13th data node on the 3rd level down")

		assert.Equal(t, []int64{67}, root.root.IndexNodes[1].IndexNodes[1].Index, "Index error for the 5th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 62}, {Key: 65}, {Key: 66}}, root.root.IndexNodes[1].IndexNodes[1].DataNodes[0].Items, "Data error for the 14th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 67}, {Key: 69}, {Key: 71}}, root.root.IndexNodes[1].IndexNodes[1].DataNodes[1].Items, "Data error for the 15th data node on the 3rd level down")

		assert.Equal(t, []int64{77}, root.root.IndexNodes[2].IndexNodes[0].Index, "Index error for the 6th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 72}, {Key: 73}}, root.root.IndexNodes[2].IndexNodes[0].DataNodes[0].Items, "Data error for the 16th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 77}, {Key: 78}, {Key: 80}}, root.root.IndexNodes[2].IndexNodes[0].DataNodes[1].Items, "Data error for the 17th data node on the 3rd level down")

		assert.Equal(t, []int64{89, 94}, root.root.IndexNodes[2].IndexNodes[1].Index, "Index error for the 7th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 81}, {Key: 86}}, root.root.IndexNodes[2].IndexNodes[1].DataNodes[0].Items, "Data error for the 18th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 89}, {Key: 91}}, root.root.IndexNodes[2].IndexNodes[1].DataNodes[1].Items, "Data error for the 19th data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 94}, {Key: 95}}, root.root.IndexNodes[2].IndexNodes[1].DataNodes[2].Items, "Data error for the 20th data node on the 3rd level down")

		assert.Equal(t, []int64{98}, root.root.IndexNodes[2].IndexNodes[2].Index, "Index error for the 8th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 96}, {Key: 97}}, root.root.IndexNodes[2].IndexNodes[2].DataNodes[0].Items, "Data error for the 21st data node on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 98}, {Key: 99}, {Key: 100}}, root.root.IndexNodes[2].IndexNodes[2].DataNodes[1].Items, "Data error for the 22nd data node on the 3rd level down")

		// Retrieve the head node of the bottom-level Link List.
		head := root.root.BpDataHead()

		// Check the continuity of the bottom-level Link List.
		tests := []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{1, 4}},
			{1, []int64{5, 6, 9}},
			{2, []int64{10, 11}},
			{3, []int64{13, 18}},
			{4, []int64{19, 22}},
			{5, []int64{23, 24}},
			{6, []int64{25, 29}},
			{7, []int64{30, 34}},
			{8, []int64{35, 36}},
			{9, []int64{37, 38}},
			{10, []int64{40, 46}},
			{11, []int64{49, 50, 56}},
			{12, []int64{59, 60}},
			{13, []int64{62, 65, 66}},
			{14, []int64{67, 69, 71}},
			{15, []int64{72, 73}},
			{16, []int64{77, 78, 80}},
			{17, []int64{81, 86}},
			{18, []int64{89, 91}},
			{19, []int64{94, 95}},
			{20, []int64{96, 97}},
			{21, []int64{98, 99, 100}},
		}

		// Start checking the continuity of the Link List.
		for _, test := range tests {
			actualKeys := head.PrintNodeAscent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// Testing the continuity of the data in the opposite direction.
		tail := root.root.BpDataTail()

		tests = []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{100, 99, 98}},
			{1, []int64{97, 96}},
			{2, []int64{95, 94}},
			{3, []int64{91, 89}},
			{4, []int64{86, 81}},
			{5, []int64{80, 78, 77}},
			{6, []int64{73, 72}},
			{7, []int64{71, 69, 67}},
			{8, []int64{66, 65, 62}},
			{9, []int64{60, 59}},
			{10, []int64{56, 50, 49}},
			{11, []int64{46, 40}},
			{12, []int64{38, 37}},
			{13, []int64{36, 35}},
			{14, []int64{34, 30}},
			{15, []int64{29, 25}},
			{16, []int64{24, 23}},
			{17, []int64{22, 19}},
			{18, []int64{18, 13}},
			{19, []int64{11, 10}},
			{20, []int64{9, 6, 5}},
			{21, []int64{4, 1}},
		}

		for _, test := range tests {
			actualKeys := tail.PrintNodeDescent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// You can use the following functions to print the entire B-tree distribution.
		// root.root.Print() // Print the entire B Plus tree.
		// head.PrintAscent() // Print continuous data in the forward direction.
		// tail.PrintDescent() // Print continuous data in the reverse direction.
	})
	t.Run("Tests inserting B-tree with a width of 5.", func(t *testing.T) {
		// Initialize B-tree.
		root := NewBpTree(5)
		// Insert 50 data entries continuously.
		root.InsertValue(BpItem{Key: 40})
		root.InsertValue(BpItem{Key: 38})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 81})
		root.InsertValue(BpItem{Key: 98})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 30})
		root.InsertValue(BpItem{Key: 67})
		root.InsertValue(BpItem{Key: 35})
		root.InsertValue(BpItem{Key: 89})
		root.InsertValue(BpItem{Key: 96})
		root.InsertValue(BpItem{Key: 78})
		root.InsertValue(BpItem{Key: 95})
		root.InsertValue(BpItem{Key: 86})
		root.InsertValue(BpItem{Key: 19})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 99})
		root.InsertValue(BpItem{Key: 59})
		root.InsertValue(BpItem{Key: 49})
		root.InsertValue(BpItem{Key: 65})
		root.InsertValue(BpItem{Key: 37})
		root.InsertValue(BpItem{Key: 73})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 29})
		root.InsertValue(BpItem{Key: 97})
		root.InsertValue(BpItem{Key: 77})
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 18})
		root.InsertValue(BpItem{Key: 69})
		root.InsertValue(BpItem{Key: 46})
		root.InsertValue(BpItem{Key: 72})
		root.InsertValue(BpItem{Key: 6})
		root.InsertValue(BpItem{Key: 36})
		root.InsertValue(BpItem{Key: 22})
		root.InsertValue(BpItem{Key: 56})
		root.InsertValue(BpItem{Key: 62})
		root.InsertValue(BpItem{Key: 23})
		root.InsertValue(BpItem{Key: 94})
		root.InsertValue(BpItem{Key: 11})
		root.InsertValue(BpItem{Key: 71})
		root.InsertValue(BpItem{Key: 34})
		root.InsertValue(BpItem{Key: 13})
		root.InsertValue(BpItem{Key: 100})
		root.InsertValue(BpItem{Key: 60})
		root.InsertValue(BpItem{Key: 24})
		root.InsertValue(BpItem{Key: 91})
		root.InsertValue(BpItem{Key: 25})
		root.InsertValue(BpItem{Key: 66})
		root.InsertValue(BpItem{Key: 50})
		root.InsertValue(BpItem{Key: 80})

		// Check the distribution of the entire B-tree.
		assert.Equal(t, []int64{30, 59, 81}, root.root.Index, "Top-level index is incorrect")

		assert.Equal(t, []int64{9, 18, 22, 24}, root.root.IndexNodes[0].Index, "Index error for the 1st data node on the 1st level down")
		assert.Equal(t, []int64{36, 40, 49}, root.root.IndexNodes[1].Index, "Index error for the 2nd data node on the 1st level down")
		assert.Equal(t, []int64{62, 67, 73}, root.root.IndexNodes[2].Index, "Index error for the 3rd data node on the 1st level down")
		assert.Equal(t, []int64{89, 95, 97}, root.root.IndexNodes[3].Index, "Index error for the 4th data node on the 1st level down")

		assert.Equal(t, []BpItem{{Key: 1}, {Key: 4}, {Key: 5}, {Key: 6}}, root.root.IndexNodes[0].DataNodes[0].Items, "Data error for the 1st data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 9}, {Key: 10}, {Key: 11}, {Key: 13}}, root.root.IndexNodes[0].DataNodes[1].Items, "Data error for the 2nd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 18}, {Key: 19}}, root.root.IndexNodes[0].DataNodes[2].Items, "Data error for the 3rd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 22}, {Key: 23}}, root.root.IndexNodes[0].DataNodes[3].Items, "Data error for the 4th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 24}, {Key: 25}, {Key: 29}}, root.root.IndexNodes[0].DataNodes[4].Items, "Data error for the 5th data node on the 2nd level down")

		assert.Equal(t, []BpItem{{Key: 30}, {Key: 34}, {Key: 35}}, root.root.IndexNodes[1].DataNodes[0].Items, "Data error for the 6th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 36}, {Key: 37}, {Key: 38}}, root.root.IndexNodes[1].DataNodes[1].Items, "Data error for the 7th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 40}, {Key: 46}}, root.root.IndexNodes[1].DataNodes[2].Items, "Data error for the 8th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 49}, {Key: 50}, {Key: 56}}, root.root.IndexNodes[1].DataNodes[3].Items, "Data error for the 9th data node on the 2nd level down")

		assert.Equal(t, []BpItem{{Key: 59}, {Key: 60}}, root.root.IndexNodes[2].DataNodes[0].Items, "Data error for the 10th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 62}, {Key: 65}, {Key: 66}}, root.root.IndexNodes[2].DataNodes[1].Items, "Data error for the 11th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 67}, {Key: 69}, {Key: 71}, {Key: 72}}, root.root.IndexNodes[2].DataNodes[2].Items, "Data error for the 12th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 73}, {Key: 77}, {Key: 78}, {Key: 80}}, root.root.IndexNodes[2].DataNodes[3].Items, "Data error for the 13th data node on the 2nd level down")

		assert.Equal(t, []BpItem{{Key: 81}, {Key: 86}}, root.root.IndexNodes[3].DataNodes[0].Items, "Data error for the 14th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 89}, {Key: 91}, {Key: 94}}, root.root.IndexNodes[3].DataNodes[1].Items, "Data error for the 15th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 95}, {Key: 96}}, root.root.IndexNodes[3].DataNodes[2].Items, "Data error for the 16th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 97}, {Key: 98}, {Key: 99}, {Key: 100}}, root.root.IndexNodes[3].DataNodes[3].Items, "Data error for the 17th data node on the 2nd level down")

		// Retrieve the head node of the bottom-level Link List.
		head := root.root.BpDataHead()

		// Check the continuity of the bottom-level Link List.
		tests := []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{1, 4, 5, 6}},
			{1, []int64{9, 10, 11, 13}},
			{2, []int64{18, 19}},
			{3, []int64{22, 23}},
			{4, []int64{24, 25, 29}},
			{5, []int64{30, 34, 35}},
			{6, []int64{36, 37, 38}},
			{7, []int64{40, 46}},
			{8, []int64{49, 50, 56}},
			{9, []int64{59, 60}},
			{10, []int64{62, 65, 66}},
			{11, []int64{67, 69, 71, 72}},
			{12, []int64{73, 77, 78, 80}},
			{13, []int64{81, 86}},
			{14, []int64{89, 91, 94}},
			{15, []int64{95, 96}},
			{16, []int64{97, 98, 99, 100}},
		}

		// Start checking the continuity of the Link List.
		for _, test := range tests {
			actualKeys := head.PrintNodeAscent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// Testing the continuity of the data in the opposite direction.
		tail := root.root.BpDataTail()

		tests = []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{100, 99, 98, 97}},
			{1, []int64{96, 95}},
			{2, []int64{94, 91, 89}},
			{3, []int64{86, 81}},
			{4, []int64{80, 78, 77, 73}},
			{5, []int64{72, 71, 69, 67}},
			{6, []int64{66, 65, 62}},
			{7, []int64{60, 59}},
			{8, []int64{56, 50, 49}},
			{9, []int64{46, 40}},
			{10, []int64{38, 37, 36}},
			{11, []int64{35, 34, 30}},
			{12, []int64{29, 25, 24}},
			{13, []int64{23, 22}},
			{14, []int64{19, 18}},
			{15, []int64{13, 11, 10, 9}},
			{16, []int64{6, 5, 4, 1}},
		}

		for _, test := range tests {
			actualKeys := tail.PrintNodeDescent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// You can use the following functions to print the entire B-tree distribution.
		// root.root.Print() // Print the entire B Plus tree.
		// head.PrintAscent() // Print continuous data in the forward direction.
		// tail.PrintDescent() // Print continuous data in the reverse direction.
	})
	t.Run("Tests inserting B-tree with a width of 6.", func(t *testing.T) {
		// Initialize B-tree.
		root := NewBpTree(6)
		// Insert 50 data entries continuously.
		root.InsertValue(BpItem{Key: 40})
		root.InsertValue(BpItem{Key: 38})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 81})
		root.InsertValue(BpItem{Key: 98})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 30})
		root.InsertValue(BpItem{Key: 67})
		root.InsertValue(BpItem{Key: 35})
		root.InsertValue(BpItem{Key: 89})
		root.InsertValue(BpItem{Key: 96})
		root.InsertValue(BpItem{Key: 78})
		root.InsertValue(BpItem{Key: 95})
		root.InsertValue(BpItem{Key: 86})
		root.InsertValue(BpItem{Key: 19})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 99})
		root.InsertValue(BpItem{Key: 59})
		root.InsertValue(BpItem{Key: 49})
		root.InsertValue(BpItem{Key: 65})
		root.InsertValue(BpItem{Key: 37})
		root.InsertValue(BpItem{Key: 73})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 29})
		root.InsertValue(BpItem{Key: 97})
		root.InsertValue(BpItem{Key: 77})
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 18})
		root.InsertValue(BpItem{Key: 69})
		root.InsertValue(BpItem{Key: 46})
		root.InsertValue(BpItem{Key: 72})
		root.InsertValue(BpItem{Key: 6})
		root.InsertValue(BpItem{Key: 36})
		root.InsertValue(BpItem{Key: 22})
		root.InsertValue(BpItem{Key: 56})
		root.InsertValue(BpItem{Key: 62})
		root.InsertValue(BpItem{Key: 23})
		root.InsertValue(BpItem{Key: 94})
		root.InsertValue(BpItem{Key: 11})
		root.InsertValue(BpItem{Key: 71})
		root.InsertValue(BpItem{Key: 34})
		root.InsertValue(BpItem{Key: 13})
		root.InsertValue(BpItem{Key: 100})
		root.InsertValue(BpItem{Key: 60})
		root.InsertValue(BpItem{Key: 24})
		root.InsertValue(BpItem{Key: 91})
		root.InsertValue(BpItem{Key: 25})
		root.InsertValue(BpItem{Key: 66})
		root.InsertValue(BpItem{Key: 50})
		root.InsertValue(BpItem{Key: 80})

		// Check the distribution of the entire B-tree.
		assert.Equal(t, []int64{36, 65}, root.root.Index, "Top-level index is incorrect")

		assert.Equal(t, []int64{10, 22, 30}, root.root.IndexNodes[0].Index, "Index error for the 1st data node on the 1st level down")
		assert.Equal(t, []int64{40, 56}, root.root.IndexNodes[1].Index, "Index error for the 2nd data node on the 1st level down")
		assert.Equal(t, []int64{72, 78, 89, 97}, root.root.IndexNodes[2].Index, "Index error for the 3rd data node on the 1st level down")

		assert.Equal(t, []BpItem{{Key: 1}, {Key: 4}, {Key: 5}, {Key: 6}, {Key: 9}}, root.root.IndexNodes[0].DataNodes[0].Items, "Data error for the 1st data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 10}, {Key: 11}, {Key: 13}, {Key: 18}, {Key: 19}}, root.root.IndexNodes[0].DataNodes[1].Items, "Data error for the 2nd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 22}, {Key: 23}, {Key: 24}, {Key: 25}, {Key: 29}}, root.root.IndexNodes[0].DataNodes[2].Items, "Data error for the 3rd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 30}, {Key: 34}, {Key: 35}}, root.root.IndexNodes[0].DataNodes[3].Items, "Data error for the 4th data node on the 2nd level down")

		assert.Equal(t, []BpItem{{Key: 36}, {Key: 37}, {Key: 38}}, root.root.IndexNodes[1].DataNodes[0].Items, "Data error for the 6th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 40}, {Key: 46}, {Key: 49}, {Key: 50}}, root.root.IndexNodes[1].DataNodes[1].Items, "Data error for the 7th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 56}, {Key: 59}, {Key: 60}, {Key: 62}}, root.root.IndexNodes[1].DataNodes[2].Items, "Data error for the 8th data node on the 2nd level down")

		assert.Equal(t, []BpItem{{Key: 65}, {Key: 66}, {Key: 67}, {Key: 69}, {Key: 71}}, root.root.IndexNodes[2].DataNodes[0].Items, "Data error for the 10th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 72}, {Key: 73}, {Key: 77}}, root.root.IndexNodes[2].DataNodes[1].Items, "Data error for the 11th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 78}, {Key: 80}, {Key: 81}, {Key: 86}}, root.root.IndexNodes[2].DataNodes[2].Items, "Data error for the 12th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 89}, {Key: 91}, {Key: 94}, {Key: 95}, {Key: 96}}, root.root.IndexNodes[2].DataNodes[3].Items, "Data error for the 13th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 97}, {Key: 98}, {Key: 99}, {Key: 100}}, root.root.IndexNodes[2].DataNodes[4].Items, "Data error for the 13th data node on the 2nd level down")

		// Retrieve the head node of the bottom-level Link List.
		head := root.root.BpDataHead()

		// Check the continuity of the bottom-level Link List.
		tests := []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{1, 4, 5, 6, 9}},
			{1, []int64{10, 11, 13, 18, 19}},
			{2, []int64{22, 23, 24, 25, 29}},
			{3, []int64{30, 34, 35}},
			{4, []int64{36, 37, 38}},
			{5, []int64{40, 46, 49, 50}},
			{6, []int64{56, 59, 60, 62}},
			{7, []int64{65, 66, 67, 69, 71}},
			{8, []int64{72, 73, 77}},
			{9, []int64{78, 80, 81, 86}},
			{10, []int64{89, 91, 94, 95, 96}},
			{11, []int64{97, 98, 99, 100}},
		}

		// Start checking the continuity of the Link List.
		for _, test := range tests {
			actualKeys := head.PrintNodeAscent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// Testing the continuity of the data in the opposite direction.
		tail := root.root.BpDataTail()

		tests = []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{100, 99, 98, 97}},
			{1, []int64{96, 95, 94, 91, 89}},
			{2, []int64{86, 81, 80, 78}},
			{3, []int64{77, 73, 72}},
			{4, []int64{71, 69, 67, 66, 65}},
			{5, []int64{62, 60, 59, 56}},
			{6, []int64{50, 49, 46, 40}},
			{7, []int64{38, 37, 36}},
			{8, []int64{35, 34, 30}},
			{9, []int64{29, 25, 24, 23, 22}},
			{10, []int64{19, 18, 13, 11, 10}},
			{11, []int64{9, 6, 5, 4, 1}},
		}

		for _, test := range tests {
			actualKeys := tail.PrintNodeDescent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// You can use the following functions to print the entire B-tree distribution.
		// root.root.Print() // Print the entire B Plus tree.
		// head.PrintAscent() // Print continuous data in the forward direction.
		// tail.PrintDescent() // Print continuous data in the reverse direction.
	})
	t.Run("Tests inserting B-tree with a width of 7.", func(t *testing.T) {
		// Initialize B-tree.
		root := NewBpTree(7)
		// Insert 50 data entries continuously.
		root.InsertValue(BpItem{Key: 40})
		root.InsertValue(BpItem{Key: 38})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 81})
		root.InsertValue(BpItem{Key: 98})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 30})
		root.InsertValue(BpItem{Key: 67})
		root.InsertValue(BpItem{Key: 35})
		root.InsertValue(BpItem{Key: 89})
		root.InsertValue(BpItem{Key: 96})
		root.InsertValue(BpItem{Key: 78})
		root.InsertValue(BpItem{Key: 95})
		root.InsertValue(BpItem{Key: 86})
		root.InsertValue(BpItem{Key: 19})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 99})
		root.InsertValue(BpItem{Key: 59})
		root.InsertValue(BpItem{Key: 49})
		root.InsertValue(BpItem{Key: 65})
		root.InsertValue(BpItem{Key: 37})
		root.InsertValue(BpItem{Key: 73})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 29})
		root.InsertValue(BpItem{Key: 97})
		root.InsertValue(BpItem{Key: 77})
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 18})
		root.InsertValue(BpItem{Key: 69})
		root.InsertValue(BpItem{Key: 46})
		root.InsertValue(BpItem{Key: 72})
		root.InsertValue(BpItem{Key: 6})
		root.InsertValue(BpItem{Key: 36})
		root.InsertValue(BpItem{Key: 22})
		root.InsertValue(BpItem{Key: 56})
		root.InsertValue(BpItem{Key: 62})
		root.InsertValue(BpItem{Key: 23})
		root.InsertValue(BpItem{Key: 94})
		root.InsertValue(BpItem{Key: 11})
		root.InsertValue(BpItem{Key: 71})
		root.InsertValue(BpItem{Key: 34})
		root.InsertValue(BpItem{Key: 13})
		root.InsertValue(BpItem{Key: 100})
		root.InsertValue(BpItem{Key: 60})
		root.InsertValue(BpItem{Key: 24})
		root.InsertValue(BpItem{Key: 91})
		root.InsertValue(BpItem{Key: 25})
		root.InsertValue(BpItem{Key: 66})
		root.InsertValue(BpItem{Key: 50})
		root.InsertValue(BpItem{Key: 80})

		// Check the distribution of the entire B-tree.
		assert.Equal(t, []int64{59}, root.root.Index, "Top-level index is incorrect")

		assert.Equal(t, []int64{6, 19, 30, 38}, root.root.IndexNodes[0].Index, "Index error for the 1st data node on the 1st level down")
		assert.Equal(t, []int64{69, 73, 81, 95}, root.root.IndexNodes[1].Index, "Index error for the 2nd data node on the 1st level down")

		assert.Equal(t, []BpItem{{Key: 1}, {Key: 4}, {Key: 5}}, root.root.IndexNodes[0].DataNodes[0].Items, "Data error for the 1st data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 6}, {Key: 9}, {Key: 10}, {Key: 11}, {Key: 13}, {Key: 18}}, root.root.IndexNodes[0].DataNodes[1].Items, "Data error for the 2nd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 19}, {Key: 22}, {Key: 23}, {Key: 24}, {Key: 25}, {Key: 29}}, root.root.IndexNodes[0].DataNodes[2].Items, "Data error for the 3rd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 30}, {Key: 34}, {Key: 35}, {Key: 36}, {Key: 37}}, root.root.IndexNodes[0].DataNodes[3].Items, "Data error for the 4th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 38}, {Key: 40}, {Key: 46}, {Key: 49}, {Key: 50}, {Key: 56}}, root.root.IndexNodes[0].DataNodes[4].Items, "Data error for the 5th data node on the 2nd level down")

		assert.Equal(t, []BpItem{{Key: 59}, {Key: 60}, {Key: 62}, {Key: 65}, {Key: 66}, {Key: 67}}, root.root.IndexNodes[1].DataNodes[0].Items, "Data error for the 6th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 69}, {Key: 71}, {Key: 72}}, root.root.IndexNodes[1].DataNodes[1].Items, "Data error for the 7th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 73}, {Key: 77}, {Key: 78}, {Key: 80}}, root.root.IndexNodes[1].DataNodes[2].Items, "Data error for the 8th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 81}, {Key: 86}, {Key: 89}, {Key: 91}, {Key: 94}}, root.root.IndexNodes[1].DataNodes[3].Items, "Data error for the 9th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 95}, {Key: 96}, {Key: 97}, {Key: 98}, {Key: 99}, {Key: 100}}, root.root.IndexNodes[1].DataNodes[4].Items, "Data error for the 10th data node on the 2nd level down")

		// Retrieve the head node of the bottom-level Link List.
		head := root.root.BpDataHead()

		// Check the continuity of the bottom-level Link List.
		tests := []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{1, 4, 5}},
			{1, []int64{6, 9, 10, 11, 13, 18}},
			{2, []int64{19, 22, 23, 24, 25, 29}},
			{3, []int64{30, 34, 35, 36, 37}},
			{4, []int64{38, 40, 46, 49, 50, 56}},
			{5, []int64{59, 60, 62, 65, 66, 67}},
			{6, []int64{69, 71, 72}},
			{7, []int64{73, 77, 78, 80}},
			{8, []int64{81, 86, 89, 91, 94}},
			{9, []int64{95, 96, 97, 98, 99, 100}},
		}

		// Start checking the continuity of the Link List.
		for _, test := range tests {
			actualKeys := head.PrintNodeAscent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// Testing the continuity of the data in the opposite direction.
		tail := root.root.BpDataTail()

		tests = []struct {
			position     int
			expectedKeys []int64
		}{
			{0, []int64{100, 99, 98, 97, 96, 95}},
			{1, []int64{94, 91, 89, 86, 81}},
			{2, []int64{80, 78, 77, 73}},
			{3, []int64{72, 71, 69}},
			{4, []int64{67, 66, 65, 62, 60, 59}},
			{5, []int64{56, 50, 49, 46, 40, 38}},
			{6, []int64{37, 36, 35, 34, 30}},
			{7, []int64{29, 25, 24, 23, 22, 19}},
			{8, []int64{18, 13, 11, 10, 9, 6}},
			{9, []int64{5, 4, 1}},
		}

		for _, test := range tests {
			actualKeys := tail.PrintNodeDescent(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// You can use the following functions to print the entire B-tree distribution.
		// root.root.Print() // Print the entire B Plus tree.
		// head.PrintAscent() // Print continuous data in the forward direction.
		// tail.PrintDescent() // Print continuous data in the reverse direction.
	})
}
