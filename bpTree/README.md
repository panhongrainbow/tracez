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

###  mergeWithDnode

To merge these three components

- the old index node (named **inode**)
- the new key (named **key**)
- and the new index node (named **side**)

into a new index node using **function mergeWithDnode**.

<img src="../assets/image-20231020101419151.png" alt="image-20231020101419151" style="zoom:95%;" />

