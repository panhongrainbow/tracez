package tracingLru

import "github.com/panhongrainbow/tracez/model"

type TracingInfo struct {
	root     *Node
	err      *Node
	traceIDs map[string]*Node
	spanIDs  map[string]*Node
}

type Node struct {
	SpanID   string
	Parent   *Node
	Children []*Node
	Data     model.TracingData
}

func (root *Node) IsRoot() (isRoot bool) {
	if root.SpanID == "root" {
		isRoot = true
	}
	return
}

func (root *Node) ListTraceID() {
	// n.Children = append(n.Children, child)
}
