# B Plus Tree

## Operations on Data Nodes

## Operations on Index Nodes

### Split and merge with BpData

> When splitting and merging at the **bottom-level** index node here, **BpData will be split along with it**.

#### splitWithDnode

Split the bottom-level Index Node

```go
func (inode *BpIndex) splitWithDnode() (key int64, side *BpIndex, err error)
func Test_Check_inode_splitWithDnode(t *testing.T)
```

Before the function is executed:

This is a 3-node B+ tree. A new independent node is allocated with 2 data node slices, starting from position pos2. 
A total of 2 data nodes are cut out, [40] and [81, 98].

Similarly, the index is also cut starting from pos2, resulting in a single index, 81. Therefore, the entire program is correct.

(重点就是同时用 Pos 这个位置去切割 index 切片和 BpData 切片都不会有错误)

<img src="../assets/image-20231014214243496.png" alt="Before Execution" style="zoom:80%;" />

After the function is executed:

**function splitWithDnode** will divide the index node into three parts:

- the old index node (named **inode**)
- the new key (named **key**)
- the new index node (named **side**)

and then reassemble them afterward.

<img src="../assets/image-20231014221906550.png" alt="After Execution" style="zoom:115%;" />

####  mergeWithDnode

Combines split index nodes into a new node, overwriting the original inode's address.

```go
func (inode *BpIndex) mergeWithDnode(podKey int64, side *BpIndex) error
func Test_Check_inode_mergeWithDnode(t *testing.T)
```

To merge these three components

- the old index node (named **inode**)
- the new key (named **key**)
- and the new index node (named **side**)

into a new index node using **function mergeWithDnode**.

<img src="../assets/image-20231020101419151.png" alt="image-20231020101419151" style="zoom:95%;" />

### Protrude index node

#### protrudeInOddBpWidth

protrudeInOddBpWidth performs index upgrade; when the middle value of the index slice pops out, it gets upgraded to the upper-level index.

Applicable to an **odd** number of top-level indexes.

```go
func (inode *BpIndex) protrudeInOddBpWidth() (middle *BpIndex, err error)
Test_Check_inode_protrudeInOddBpWidth(t *testing.T)
```

The initial index node was too large, and now some sub-index nodes need to be upgraded.

<img src="../assets/image-20231027114351856.png" alt="image-20231027114351856" style="zoom:150%;" />

The upgraded result is as follows: index 40 has been upgraded to the upper-level node.

<img src="../assets/image-20231027114009154.png" alt="image-20231027114009154" style="zoom:150%;" />

The subsequent encounters will be handled in two ways:

1. If this node itself is the root node, the entire tree will become a new root.

2. If this node is not the root node, it will merge the upgraded node. (refer to **ackUpgradeIndexNode**)

(是根就是新根，不是就升级合拼)

#### protrudeInEvenBpWidth

protrudeInOddBpWidth performs index upgrade; when the middle value of the index slice pops out, it gets upgraded to the upper-level index.

Applicable to an **even** number of top-level indexes.

```go
func (inode *BpIndex) protrudeInEvenBpWidth() (popMiddleNode *BpIndex, err error)
func Test_Check_inode_protrudeInEvenBpWidth(t *testing.T)
```

The initial index node was too large, and now some sub-index nodes need to be upgraded.

<img src="../assets/image-20231030023003143.png" alt="image-20231030023003143" style="zoom:100%;" />

The upgraded result is as follows: index 72 has been upgraded to the upper-level node.

<img src="../assets/image-20231030025022988.png" alt="image-20231030025022988" style="zoom:100%;" />

Condensed the previous two charts as follows: (上两张图精简如下)

<img src="../assets/image-20231030031635639.png" alt="image-20231030031635639" style="zoom:150%;" />

The subsequent encounters will be handled in two ways:

1. If this node itself is the root node, the entire tree will become a new root.

2. If this node is not the root node, it will merge the upgraded node. (refer to **ackUpgradeIndexNode**)

(是根就是新根，不是就升级合拼)

#### ackUpgradeIndexNode

**ackUpgradeIndexNode** is used by the current layer's index node to acknowledge a new independently upgraded index node.
This function is extracted from insertItem function for testing purposes, and it overwrites the original location in the inode.

(承认新独立的索引结点)

```go
func (inode *BpIndex) ackUpgradeIndexNode(ix int, popNode *BpIndex)
Test_Check_inode_ackUpgradeIndexNode(t *testing.T)
```

This is with a width of 3.

Before the execution of the **ackUpgradeIndexNode** function, as shown in the diagram, when the index slice []int64{67, 77, 89} is too large, it needs to be split.

After splitting, it is then overwrites back to the original position.

For example, if the parameter ix = 1 is passed, it means it should be overwrited at position 1, which is also the second index node []int64{67, 77, 89} under the root node. 

To avoid position confusion, the original node (named inode) is not deleted; instead, it is overwrited at the ix position.

(为避免位置混淆，不会删除，而是在 ix 上进行覆盖)

<img src="../assets/image-20231103135403218.png" alt="image-20231103135403218" style="zoom:100%;" />

The results after overwriting are as follows:

<img src="../assets/image-20231103140154185.png" alt="image-20231103140154185" style="zoom:100%;" />

The above process can be simplified as shown in the following diagram.

Before the execution of the ackUpgradeIndexNode function

<img src="../assets/image-20231103135710211.png" alt="image-20231103135710211" style="zoom:100%;" />

After the execution of the ackUpgradeIndexNode function

![image-20231103141032488](../assets/image-20231103141032488.png)

### Split and merge with upgraded key and node

#### mergeUpgradedKeyNode

Merges the to-be-upgraded Key and the to-be-upgraded Inode into the parent index node.

```go
func (inode *BpIndex) mergeUpgradedKeyNode(insertAfterPosition int, key int64, side *BpIndex) (err error)
// insertAfterPosition indicates where it should be inserted after.

func Test_Check_inode_mergeUpgradedKeyNode(t *testing.T)
```

Currently, three parts are going to be merged:

- index node 1 (the original index node)
- key 1 (the part of the index to be upgraded)
- and index node 2 (the index node to be upgraded)

<img src="../assets/image-20231022224138744.png" alt="image-20231022224138744" style="zoom:80%;" />

The merged result is as follows.

<img src="../assets/image-20231023012429846.png" alt="image-20231023012429846" style="zoom:80%;" />

#### insertAfterPosition parameter

insertAfterPosition indicates where it should be inserted after. (插在什么位置之后)

For example, if insertAfterPosition is at pos0(0), insert the upgraded node after pos0(0), which is at pos1(1). (传入0，在位置0之后，就是1)

<img src="../assets/image-20231023021615643.png" alt="image-20231023021615643" style="zoom:80%;" />



