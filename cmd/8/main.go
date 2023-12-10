package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type Node struct {
	Ident string
	Left  string
	Right string
}

func getNextNode(direction string, node *Node, nodes map[string]*Node) *Node {
	switch direction {
	case "L":
		return nodes[node.Left]
	case "R":
		return nodes[node.Right]
	default:
		return nil // Should never happen
	}
}

func walkTheNodes(nodes map[string]*Node, directions []string) int {
	node := nodes["AAA"]
	pos := 0

	fmt.Printf("%s = (%s, %s)\n", node.Ident, node.Left, node.Right)

	for {
		node = getNextNode(directions[pos%len(directions)], node, nodes)
		if node == nil {
			panic("next node is nil")
		}

		fmt.Printf("%s = (%s, %s)\n", node.Ident, node.Left, node.Right)

		if node.Ident == "ZZZ" {
			break
		}

		pos++
	}

	return pos + 1
}

func parseDirections(data string) (directions []string) {
	for _, direction := range data {
		directions = append(directions, string(direction))
	}

	return directions
}

func parseNode(data string) (node *Node) {
	return &Node{
		Ident: string(data[0:3]),
		Left:  string(data[7:10]),
		Right: string(data[12:15]),
	}
}

func parseNodes(data string) map[string]*Node {
	rawNodes := strings.Split(data, "\n")
	nodes := make(map[string]*Node)

	for _, rawNode := range rawNodes {
		node := parseNode(rawNode)
		nodes[node.Ident] = node
	}

	return nodes
}

//go:embed data.txt
var input string

func main() {
	tik := time.Now()
	lines := strings.Split(input, "\n\n")

	directions := parseDirections(lines[0])
	nodes := parseNodes(lines[1])
	steps := walkTheNodes(nodes, directions)
	elapsed := time.Since(tik).Seconds()

	fmt.Println("================")
	fmt.Printf("Total steps: %d (took %1.3fs)\n", steps, elapsed)
}
