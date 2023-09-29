// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package tree

// Abstract interface
type INoder interface {
	CalSortValue() int
}

// Decorator
type node struct {
	Noder INoder
}

// Implement sorting algorithm
type nodes []*node

// Sort by sort value
func (ns nodes) Len() int { return len(ns) }
func (ns nodes) Less(i, j int) bool {
	return ns[i].Noder.CalSortValue() > ns[j].Noder.CalSortValue()
}
func (ns nodes) Swap(i, j int) { ns[i], ns[j] = ns[j], ns[i] }
func (ns *nodes) Push(x interface{}) {
	*ns = append(*ns, x.(*node))
}
func (ns *nodes) Pop() interface{} {
	old := *ns
	n := len(old)
	x := old[n-1]
	*ns = old[0 : n-1]
	return x
}

// Create node
func NewNode(n INoder) *node {
	return &node{Noder: n}
}
