package bpTree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_Check_Btree(t *testing.T) {
	t.Run("3 Width", func(t *testing.T) {
		root := NewBpTree(3)
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 6})
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 7})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 3})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 11})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 12})
		root.InsertValue(BpItem{Key: 13})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 8})
		root.InsertValue(BpItem{Key: 3})
		root.InsertValue(BpItem{Key: 7})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 8})
		root.InsertValue(BpItem{Key: 8})

		root.root.deleteItem(nil, BpItem{Key: 6})
		root.root.deleteItem(nil, BpItem{Key: 7})

		root.root.deleteItem(nil, BpItem{Key: 1})

		root.root.Print()
		fmt.Println()
	})
	// For B-tree with a width of 4, perform tests inserting.
	t.Run("tests inserting B-tree with a width of 4.", func(t *testing.T) {
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
		assert.Equal(t, []BpItem{{Key: 1}, {Key: 4}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[0].Items, "Data error for the 1st data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 5}, {Key: 6}, {Key: 9}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[1].Items, "Data error for the 2nd data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 10}, {Key: 11}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[2].Items, "Data error for the 3rd data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 13}, {Key: 18}}, root.root.IndexNodes[0].IndexNodes[0].DataNodes[3].Items, "Data error for the 4th data slice on the 3rd level down")

		assert.Equal(t, []int64{23, 25}, root.root.IndexNodes[0].IndexNodes[1].Index, "Index error for the 2nd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 19}, {Key: 22}}, root.root.IndexNodes[0].IndexNodes[1].DataNodes[0].Items, "Data error for the 5th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 23}, {Key: 24}}, root.root.IndexNodes[0].IndexNodes[1].DataNodes[1].Items, "Data error for the 6th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 25}, {Key: 29}}, root.root.IndexNodes[0].IndexNodes[1].DataNodes[2].Items, "Data error for the 7th data slice on the 3rd level down")

		assert.Equal(t, []int64{35, 37}, root.root.IndexNodes[0].IndexNodes[2].Index, "Index error for the 3rd data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 30}, {Key: 34}}, root.root.IndexNodes[0].IndexNodes[2].DataNodes[0].Items, "Data error for the 8th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 35}, {Key: 36}}, root.root.IndexNodes[0].IndexNodes[2].DataNodes[1].Items, "Data error for the 9th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 37}, {Key: 38}}, root.root.IndexNodes[0].IndexNodes[2].DataNodes[2].Items, "Data error for the 10th data slice on the 3rd level down")

		assert.Equal(t, []int64{49, 59}, root.root.IndexNodes[1].IndexNodes[0].Index, "Index error for the 4th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 40}, {Key: 46}}, root.root.IndexNodes[1].IndexNodes[0].DataNodes[0].Items, "Data error for the 11th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 49}, {Key: 50}, {Key: 56}}, root.root.IndexNodes[1].IndexNodes[0].DataNodes[1].Items, "Data error for the 12th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 59}, {Key: 60}}, root.root.IndexNodes[1].IndexNodes[0].DataNodes[2].Items, "Data error for the 13th data slice on the 3rd level down")

		assert.Equal(t, []int64{67}, root.root.IndexNodes[1].IndexNodes[1].Index, "Index error for the 5th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 62}, {Key: 65}, {Key: 66}}, root.root.IndexNodes[1].IndexNodes[1].DataNodes[0].Items, "Data error for the 14th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 67}, {Key: 69}, {Key: 71}}, root.root.IndexNodes[1].IndexNodes[1].DataNodes[1].Items, "Data error for the 15th data slice on the 3rd level down")

		assert.Equal(t, []int64{77}, root.root.IndexNodes[2].IndexNodes[0].Index, "Index error for the 6th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 72}, {Key: 73}}, root.root.IndexNodes[2].IndexNodes[0].DataNodes[0].Items, "Data error for the 16th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 77}, {Key: 78}, {Key: 80}}, root.root.IndexNodes[2].IndexNodes[0].DataNodes[1].Items, "Data error for the 17th data slice on the 3rd level down")

		assert.Equal(t, []int64{89, 94}, root.root.IndexNodes[2].IndexNodes[1].Index, "Index error for the 7th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 81}, {Key: 86}}, root.root.IndexNodes[2].IndexNodes[1].DataNodes[0].Items, "Data error for the 18th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 89}, {Key: 91}}, root.root.IndexNodes[2].IndexNodes[1].DataNodes[1].Items, "Data error for the 19th data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 94}, {Key: 95}}, root.root.IndexNodes[2].IndexNodes[1].DataNodes[2].Items, "Data error for the 20th data slice on the 3rd level down")

		assert.Equal(t, []int64{98}, root.root.IndexNodes[2].IndexNodes[2].Index, "Index error for the 8th data node on the 2nd level down")
		assert.Equal(t, []BpItem{{Key: 96}, {Key: 97}}, root.root.IndexNodes[2].IndexNodes[2].DataNodes[0].Items, "Data error for the 21st data slice on the 3rd level down")
		assert.Equal(t, []BpItem{{Key: 98}, {Key: 99}, {Key: 100}}, root.root.IndexNodes[2].IndexNodes[2].DataNodes[1].Items, "Data error for the 22nd data slice on the 3rd level down")

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
			actualKeys := head.PrintNodeKeys(test.position)
			assert.Equal(t, test.expectedKeys, actualKeys, "Bottom-level Link list at position "+strconv.Itoa(test.position)+" is not continuous")
		}

		// You can use the following functions to print the entire B-tree distribution.
		// head.Print()
		// root.root.Print()
	})
	t.Run("5 Width", func(t *testing.T) {
		root := NewBpTree(5)
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 6})
		root.InsertValue(BpItem{Key: 5})
		root.InsertValue(BpItem{Key: 7})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 3})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 4})
		root.InsertValue(BpItem{Key: 1})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 11})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 12})
		root.InsertValue(BpItem{Key: 13})
		/*root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})*/
		/*root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 15})
		root.InsertValue(BpItem{Key: 8})
		root.InsertValue(BpItem{Key: 3})
		root.InsertValue(BpItem{Key: 7})
		root.InsertValue(BpItem{Key: 10})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 9})
		root.InsertValue(BpItem{Key: 8})
		root.InsertValue(BpItem{Key: 8})*/
		root.root.Print()
		fmt.Println()
	})
}

func Test_Check_Btree2(t *testing.T) {

}
