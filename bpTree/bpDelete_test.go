package bpTree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 🧫 loadBasicDeletionExample function is mainly used to generate a fixed B Plus tree with root index []int64{7, 13} and max degree 4.
func loadBasicDeletionExample() (basicDeletionBpTree *BpTree) {
	// Generate continuous data for insertion.
	var basicDeletionNumbers = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

	// Initialize B Plus tree.
	basicDeletionBpTree = NewBpTree(4)

	// Start inserting data.
	for i := 0; i < len(basicDeletionNumbers); i++ {
		// Insert data entries continuously.
		basicDeletionBpTree.InsertValue(BpItem{Key: basicDeletionNumbers[i]})
	}

	// Complete this function.
	return
}

// ⚗️ Test_Check_Basic_BpIndex_Deletion is to load a test B Plus Tree and check the indexes and data.
func Test_Check_Basic_BpIndex_Deletion(t *testing.T) {
	// 🧪 This test is to confirm that the test data is correct.
	t.Run("Load Basic Deletion Example.", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Check the index node of the first level.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index nodes on the second level.
		require.Equal(t, []int64{3, 5}, basicDeletionBpTree.root.IndexNodes[0].Index)
		require.Equal(t, []int64{9, 11}, basicDeletionBpTree.root.IndexNodes[1].Index)
		require.Equal(t, []int64{15, 17, 19}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Now that there are ten data nodes, use Data Head to traverse all the data in the B Plus Tree.

		dataHeader := basicDeletionBpTree.root.IndexNodes[0].DataNodes[0]
		require.Equal(t, int64(1), dataHeader.Items[0].Key)
		require.Equal(t, int64(2), dataHeader.Items[1].Key)

		secondDataNode := dataHeader.Next
		require.Equal(t, int64(3), secondDataNode.Items[0].Key)
		require.Equal(t, int64(4), secondDataNode.Items[1].Key)

		thirdDataNode := secondDataNode.Next
		require.Equal(t, int64(5), thirdDataNode.Items[0].Key)
		require.Equal(t, int64(6), thirdDataNode.Items[1].Key)

		fourthDataNode := thirdDataNode.Next
		require.Equal(t, int64(7), fourthDataNode.Items[0].Key)
		require.Equal(t, int64(8), fourthDataNode.Items[1].Key)

		fifthDataNode := fourthDataNode.Next
		require.Equal(t, int64(9), fifthDataNode.Items[0].Key)
		require.Equal(t, int64(10), fifthDataNode.Items[1].Key)

		sixthDataNode := fifthDataNode.Next
		require.Equal(t, int64(11), sixthDataNode.Items[0].Key)
		require.Equal(t, int64(12), sixthDataNode.Items[1].Key)

		seventhDataNode := sixthDataNode.Next
		require.Equal(t, int64(13), seventhDataNode.Items[0].Key)
		require.Equal(t, int64(14), seventhDataNode.Items[1].Key)

		eighthDataNode := seventhDataNode.Next
		require.Equal(t, int64(15), eighthDataNode.Items[0].Key)
		require.Equal(t, int64(16), eighthDataNode.Items[1].Key)

		ninthDataNode := eighthDataNode.Next
		require.Equal(t, int64(17), ninthDataNode.Items[0].Key)
		require.Equal(t, int64(18), ninthDataNode.Items[1].Key)

		tenthDataNode := ninthDataNode.Next
		require.Equal(t, int64(19), tenthDataNode.Items[0].Key)
		require.Equal(t, int64(20), tenthDataNode.Items[1].Key)
		require.Equal(t, int64(21), tenthDataNode.Items[2].Key)
	})

	// 🧪 This test is to verify that when deleting a Non-Edge Value, the Local-Index is not changed.
	t.Run("When deleting a Non-Edge Value, the Local-Index is not changed.", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Non-Edge-Value 20.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 20})
		require.True(t, deleted)

		// Check the index node of the first level after deleting the Non-Edge-Value 20.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting the Non-Edge-Value 20.
		require.Equal(t, []int64{15, 17, 19}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting the Non-Edge-Value 20.
		require.Equal(t, 2, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items))
		require.Equal(t, int64(19), basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items[0].Key)
		require.Equal(t, int64(21), basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items[1].Key)
	})

	// 🧪 This test is to verify that when an Edge-Value is deleted, the Local-Index needs to be updated.
	t.Run("When an Edge-Value is deleted, the Local-Index needs to be updated.", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Non-Edge-Value 20.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 20})
		require.True(t, deleted)

		// Deleting the Edge-Value 19.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 19})
		require.True(t, deleted)

		// Check the index node of the first level after deleting the Edge-Value 19.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting the Edge-Value 19.
		require.Equal(t, []int64{15, 17, 21}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting the Edge-Value 19.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items))
		require.Equal(t, int64(21), basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items[0].Key)
	})
}

// ⚗️ Test_Check_borrowFromDataNode_Function is primarily used to test the borrowFromDataNode function.
// More details can be found in Chapter 2.3.1 `Borrow from Neighbor` in the documentation.
func Test_Check_borrowFromDataNode_Function(t *testing.T) {
	// 🧪 This test is mainly used to test the scenario of Status 1.
	t.Run("Status 1 in Chapter 2.3.1", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Non-Edge-Value 14.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 14})
		require.True(t, deleted)

		// Deleting the Non-Edge-Value 13.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 13})
		require.True(t, deleted)

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{7, 15}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{16, 17, 19}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting data.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items))
		require.Equal(t, int64(15), basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[1].Items))
		require.Equal(t, int64(16), basicDeletionBpTree.root.IndexNodes[2].DataNodes[1].Items[0].Key)
	})

	// 🧪 This test is mainly used to test the scenario of Status 2-1.
	t.Run("Status 2-1 in Chapter 2.3.1", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Non-Edge-Value 20.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 20})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 19.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 19})
		require.True(t, deleted)

		// Deleting the inner-Edge-Value 21.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 21})
		require.True(t, deleted)

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{15, 17, 18}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting data.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[2].Items))
		require.Equal(t, int64(17), basicDeletionBpTree.root.IndexNodes[2].DataNodes[2].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items))
		require.Equal(t, int64(18), basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items[0].Key)
	})

	// 🧪 This test is mainly used to test the scenario of Status 2-2.
	t.Run("Status 2-2 in Chapter 2.3.1", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Non-Edge-Value 20.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 20})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 19.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 19})
		require.True(t, deleted)

		// Deleting the inner-Edge-Value 18.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 18})
		require.True(t, deleted)

		// Deleting the inner-Edge-Value 17.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 17})
		require.True(t, deleted)

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{15, 16, 21}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting data.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[1].Items))
		require.Equal(t, int64(15), basicDeletionBpTree.root.IndexNodes[2].DataNodes[1].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[2].Items))
		require.Equal(t, int64(16), basicDeletionBpTree.root.IndexNodes[2].DataNodes[2].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items))
		require.Equal(t, int64(21), basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items[0].Key)
	})

	// 🧪  This test is mainly used to test the scenario of Status 3.
	t.Run("Status 3 in Chapter 2.3.1", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Non-Edge-Value 18.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 18})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 17.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 17})
		require.True(t, deleted)

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{15, 19, 20}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting data.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[2].Items))
		require.Equal(t, int64(19), basicDeletionBpTree.root.IndexNodes[2].DataNodes[2].Items[0].Key)

		require.Equal(t, 2, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items))
		require.Equal(t, int64(20), basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items[0].Key)
		require.Equal(t, int64(21), basicDeletionBpTree.root.IndexNodes[2].DataNodes[3].Items[1].Key)
	})
}

// ⚗️ Test_Check_borrowFromBottomIndexNode_Function will verify the following process:
// it will borrow data from lower-level index nodes.
// However, the process is complex, with at least six scenarios that need to be analyzed one by one.
func Test_Check_borrowFromBottomIndexNode_Function(t *testing.T) {
	// 🧪  This test is mainly used to test the scenario in Chapter 2.3.2.
	t.Run("Borrow data from the Right first and not damage Neighbor", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Inner-Edge-Value 7.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 7})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 8.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 8})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 9.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 9})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 10.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 10})
		require.True(t, deleted)

		// 🏗️ Here, the test environment is formed.

		// Deleting the new Inner-Edge-Value 11.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 11}) // ⚔️ Delete crucial data and test.
		require.True(t, deleted)

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{12, 14}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{13}, basicDeletionBpTree.root.IndexNodes[1].Index)
		require.Equal(t, []int64{15, 17, 19}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting data.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items))
		require.Equal(t, int64(12), basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items))
		require.Equal(t, int64(13), basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items))
		require.Equal(t, int64(14), basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items[0].Key)
	})
	// 🧪  This test is mainly used to test the scenario in Chapter 2.3.3.
	t.Run("Scenario in Chapter 2.3.3", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Inner-Edge-Value 14.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 14})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 7.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 7})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 8.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 8})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 9.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 9})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 10.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 10})
		require.True(t, deleted)

		// 🏗️ Here, the test environment is formed.

		// Check the test environment.

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{11, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{12}, basicDeletionBpTree.root.IndexNodes[1].Index)
		require.Equal(t, []int64{15, 17, 19}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting data.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items))
		require.Equal(t, int64(11), basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items))
		require.Equal(t, int64(12), basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items))
		require.Equal(t, int64(13), basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items[0].Key)

		// Deleting the new Inner-Edge-Value 11.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 11}) // ⚔️ Delete crucial data and test.
		require.True(t, deleted)

		// 🩻 Check Test Results.

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{12, 15}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{13}, basicDeletionBpTree.root.IndexNodes[1].Index)
		require.Equal(t, []int64{17, 19}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the data nodes of the third level after deleting data.
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items))
		require.Equal(t, int64(12), basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items[0].Key)

		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items))
		require.Equal(t, int64(13), basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items[0].Key)

		require.Equal(t, 2, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items))
		require.Equal(t, int64(15), basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items[0].Key)
		require.Equal(t, int64(16), basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items[1].Key)
	})
	// 🧪  This test is mainly used to test the scenario in Chapter 2.3.4.
	t.Run("Scenario in Chapter 2.3.4", func(t *testing.T) {
		// Load a simple B Plus Tree where max degree is 4.
		basicDeletionBpTree := loadBasicDeletionExample()

		// Deleting the Inner-Edge-Value 8.
		deleted, _, _, _ := basicDeletionBpTree.RemoveValue(BpItem{Key: 8})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 10.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 10})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 12.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 12})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 14.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 14})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 18.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 18})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 20.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 20})
		require.True(t, deleted)

		// Deleting the new Inner-Edge-Value 21.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 21})
		require.True(t, deleted)

		// Neither the Outer-Edge-Values nor the Inner-Edge-Values are deleted, so the index will not change.

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{9, 11}, basicDeletionBpTree.root.IndexNodes[1].Index)
		require.Equal(t, []int64{15, 17, 19}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Deleting the boundary values will change the index significantly.

		// Deleting the Inner-Edge-Value 11.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 11})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 17.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 17})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 19.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 19})
		require.True(t, deleted)

		// Deleting the Inner-Edge-Value 16.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 16})
		require.True(t, deleted)

		// Deleting key value 16 and adding it again will remove index 16.
		basicDeletionBpTree.InsertValue(BpItem{Key: 16})

		// 🏗️ Here, the test environment is formed.

		// Check the test environment.

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{7, 13}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{9}, basicDeletionBpTree.root.IndexNodes[1].Index)
		require.Equal(t, []int64{15}, basicDeletionBpTree.root.IndexNodes[2].Index)

		// Check the number of data in the 2nd index node.
		require.Equal(t, 2, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes))
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items))
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items))

		// Check the number of data in the 3rd index node.
		require.Equal(t, 2, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes))
		require.Equal(t, 1, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items))
		require.Equal(t, 2, len(basicDeletionBpTree.root.IndexNodes[2].DataNodes[1].Items))

		// Check the contents of the data in the 2nd index node.
		require.Equal(t, int64(7), basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items[0].Key)
		require.Equal(t, int64(9), basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items[0].Key)

		// Check the contents of the data in the 3rd index node.
		require.Equal(t, int64(13), basicDeletionBpTree.root.IndexNodes[2].DataNodes[0].Items[0].Key)
		require.Equal(t, int64(15), basicDeletionBpTree.root.IndexNodes[2].DataNodes[1].Items[0].Key)
		require.Equal(t, int64(16), basicDeletionBpTree.root.IndexNodes[2].DataNodes[1].Items[1].Key)

		// Delete critical data, key 9, triggering node merging.
		deleted, _, _, _ = basicDeletionBpTree.RemoveValue(BpItem{Key: 9}) // ⚔️ Delete crucial data and test.

		// 🩻 Check Test Results.

		// Check the index node of the first level after deleting data.
		require.Equal(t, []int64{7}, basicDeletionBpTree.root.Index)

		// Check the index node of the second level after deleting data.
		require.Equal(t, []int64{3, 5}, basicDeletionBpTree.root.IndexNodes[0].Index)
		require.Equal(t, []int64{13, 15}, basicDeletionBpTree.root.IndexNodes[1].Index)

		// Check the contents of the data in the 2nd index node.
		require.Equal(t, int64(7), basicDeletionBpTree.root.IndexNodes[1].DataNodes[0].Items[0].Key)
		require.Equal(t, int64(13), basicDeletionBpTree.root.IndexNodes[1].DataNodes[1].Items[0].Key)
		require.Equal(t, int64(15), basicDeletionBpTree.root.IndexNodes[1].DataNodes[2].Items[0].Key)
		require.Equal(t, int64(16), basicDeletionBpTree.root.IndexNodes[1].DataNodes[2].Items[1].Key)
	})
}
