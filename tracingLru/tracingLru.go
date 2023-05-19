package tracingLru

type Node struct {
	SpanID   string
	Parent   *Node
	Children []*Node
}
