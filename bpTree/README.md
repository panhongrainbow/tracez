# B Plus Tree

## Operations on Data Nodes

## Operations on Index Nodes

### BpIndex.splitWithDnode

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

![Before Execution](../assets/image-20231014214243496.png)

After the function is executed:

![After Execution](../assets/image-20231014221906550.png)