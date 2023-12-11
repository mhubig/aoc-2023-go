package main

import (
	"testing"
)

func TestGetNextNode_GivenThreeNodes_GetLeftOne(t *testing.T) {
	nodesGiven := map[string]*Node{
		"AAA": {Ident: "AAA", Left: "BBB", Right: "CCC"},
		"BBB": {Ident: "BBB", Left: "DDD", Right: "EEE"},
		"CCC": {Ident: "CCC", Left: "ZZZ", Right: "GGG"},
	}

	nodeGiven := nodesGiven["AAA"]
	nodeExpected := nodesGiven["BBB"]
	nodeActual := getNextNode("L", nodeGiven, nodesGiven)

	if nodeActual != nodeExpected {
		t.Errorf("getNextNode() = %v, want %v", nodeActual, nodeExpected)
	}
}

func TestGetNextNode_GivenThreeNodes_GetRightOne(t *testing.T) {
	nodesGiven := map[string]*Node{
		"AAA": {Ident: "AAA", Left: "BBB", Right: "CCC"},
		"BBB": {Ident: "BBB", Left: "DDD", Right: "EEE"},
		"CCC": {Ident: "CCC", Left: "ZZZ", Right: "GGG"},
	}

	nodeGiven := nodesGiven["AAA"]
	nodeExpected := nodesGiven["CCC"]
	nodeActual := getNextNode("R", nodeGiven, nodesGiven)

	if nodeActual != nodeExpected {
		t.Errorf("getNextNode() = %v, want %v", nodeActual, nodeExpected)
	}
}

func TestGetNextNode_GivenThreeNodes_GetNil(t *testing.T) {
	nodesGiven := map[string]*Node{
		"AAA": {Ident: "AAA", Left: "BBB", Right: "CCC"},
		"BBB": {Ident: "BBB", Left: "DDD", Right: "EEE"},
		"CCC": {Ident: "CCC", Left: "ZZZ", Right: "GGG"},
	}

	nodeGiven := nodesGiven["AAA"]
	nodeActual := getNextNode("X", nodeGiven, nodesGiven)

	if nodeActual != nil {
		t.Errorf("getNextNode() = %v, want %v", nodeActual, nil)
	}
}

// RL
// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)
// ZZZ = (ZZZ, ZZZ)
func TestWalkTheNodes_GivenNodesFromTestInput1_WalkThem(t *testing.T) {
	nodesGiven := map[string]*Node{
		"AAA": {Ident: "AAA", Left: "BBB", Right: "CCC"},
		"BBB": {Ident: "BBB", Left: "DDD", Right: "EEE"},
		"CCC": {Ident: "CCC", Left: "ZZZ", Right: "GGG"},
		"DDD": {Ident: "DDD", Left: "DDD", Right: "DDD"},
		"EEE": {Ident: "EEE", Left: "EEE", Right: "EEE"},
		"GGG": {Ident: "GGG", Left: "GGG", Right: "GGG"},
		"ZZZ": {Ident: "ZZZ", Left: "ZZZ", Right: "ZZZ"},
	}

	directionsGiven := []string{"R", "L"}

	steps := walkTheNodes(nodesGiven, directionsGiven)

	if steps != 2 {
		t.Errorf("walkTheNodes() = %v, want %v", steps, 2)
	}
}

// LLR
// AAA = (BBB, BBB)
// BBB = (AAA, ZZZ)
// ZZZ = (ZZZ, ZZZ)
func TestWalkTheNodes_GivenNodesFromTestInput2_WalkThem(t *testing.T) {
	nodesGiven := map[string]*Node{
		"AAA": {Ident: "AAA", Left: "BBB", Right: "BBB"},
		"BBB": {Ident: "BBB", Left: "AAA", Right: "ZZZ"},
		"ZZZ": {Ident: "ZZZ", Left: "ZZZ", Right: "ZZZ"},
	}

	directionsGiven := []string{"L", "L", "R"}

	steps := walkTheNodes(nodesGiven, directionsGiven)

	if steps != 6 {
		t.Errorf("walkTheNodes() = %v, want %v", steps, 6)
	}
}

func TestParseNode_GivenString_ReturnNode(t *testing.T) {
	nodeGiven := "AAA = (BBB, CCC)"
	nodeExpected := &Node{Ident: "AAA", Left: "BBB", Right: "CCC"}
	nodeActual := parseNode(nodeGiven)

	if nodeActual.Ident != nodeExpected.Ident {
		t.Errorf("parseNode() = %v, want %v", nodeActual.Ident, nodeExpected.Ident)
	}

	if nodeActual.Left != nodeExpected.Left {
		t.Errorf("parseNode() = %v, want %v", nodeActual.Left, nodeExpected.Left)
	}

	if nodeActual.Right != nodeExpected.Right {
		t.Errorf("parseNode() = %v, want %v", nodeActual.Right, nodeExpected.Right)
	}
}

func TestParseNodes_GivenStringWithTreeNodes_ReturnThreeNodes(t *testing.T) {
	nodesGiven := "AAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)"
	nodesExpected := map[string]*Node{
		"AAA": {Ident: "AAA", Left: "BBB", Right: "CCC"},
		"BBB": {Ident: "BBB", Left: "DDD", Right: "EEE"},
		"CCC": {Ident: "CCC", Left: "ZZZ", Right: "GGG"},
	}

	nodesActual := parseNodes(nodesGiven)

	if len(nodesActual) != len(nodesExpected) {
		t.Errorf("parseNodes() = %v, want %v", len(nodesActual), len(nodesExpected))
	}
}

func TestParseDirections_GivenString_ReturnSliceOfStrings(t *testing.T) {
	directionsGiven := "RL"
	directionsExpected := []string{"R", "L"}

	directionsActual := parseDirections(directionsGiven)

	if len(directionsActual) != len(directionsExpected) {
		t.Errorf("parseDirections() = %v, want %v", len(directionsActual), len(directionsExpected))
	}

	if directionsActual[0] != directionsExpected[0] {
		t.Errorf("parseDirections() = %v, want %v", directionsActual[0], directionsExpected[0])
	}

	if directionsActual[1] != directionsExpected[1] {
		t.Errorf("parseDirections() = %v, want %v", directionsActual[1], directionsExpected[1])
	}
}
