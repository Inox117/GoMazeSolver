package model

import "testing"

func TestBuildFromEdgesAndNodes(t *testing.T) {
	node1 := Node{Id: 1, isEntry: true, isExit: false}
	node2 := Node{Id: 2, isEntry: false, isExit: false}
	node3 := Node{Id: 3, isEntry: false, isExit: false}
	node4 := Node{Id: 4, isEntry: false, isExit: true}
	nodes := []Node{node1, node2, node3, node4}
	edges := []Edge{
		{from: 1, to: 2},
		{from: 1, to: 3},
		{from: 2, to: 3},
		{from: 3, to: 4},
	}
	graph := BuildFromEdgesAndNodes(nodes, edges)
	if graph.getNbEdges() != 4 || graph.getNbNodes() != 4 {
		t.Fail()
	}
}

func TestBuildFromMatrix(t *testing.T) {
	matrix := [][]int{{0, 0, 0}, {2, 1, 3}, {0, 0, 0}}
	graph := BuildFromMatrix(matrix)
	if graph.getNbEdges() != 4 || graph.getNbNodes() != 3 {
		t.Errorf("Got %v edges instead of 6 and got %v nodes instead of 3", graph.getNbEdges(), graph.getNbNodes())
	}
}