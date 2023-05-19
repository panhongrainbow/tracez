package tracingLru

import "github.com/panhongrainbow/tracez/model"

type Node struct {
	SpanID   string
	Parent   *Node
	Children []*Node
	Data     model.TracingData
}
