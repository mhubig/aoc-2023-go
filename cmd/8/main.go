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

	for {
		node = getNextNode(directions[pos%len(directions)], node, nodes)
		if node == nil {
			panic("next node is nil")
		}

		if node.Ident == "ZZZ" {
			break
		}

		pos++
	}

	return pos + 1
}

func walkNodesInParalell(nodes map[string]*Node, directions []string) []int {
	var steps []int

	for _, node := range nodes {
		if strings.HasSuffix(node.Ident, "A") {
			step := 0
			for {
				node = getNextNode(directions[step%len(directions)], node, nodes)
				step++
				if strings.HasSuffix(node.Ident, "Z") {
					break
				}
			}

			steps = append(steps, step)
		}
	}

	return steps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmAll(n []int) int {
	result := n[0]
	for _, b := range n[1:] {
		result = lcm(result, b)
	}

	return result
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

	// Part 1
	result := walkTheNodes(nodes, directions)
	fmt.Printf("Part 1: Total steps %d (took %1.3fs)\n", result, time.Since(tik).Seconds())

	// Part 2
	tik = time.Now()
	result = lcmAll(walkNodesInParalell(nodes, directions))
	fmt.Printf("Part 2: Total steps %d (took %1.3fs)\n", result, time.Since(tik).Seconds())
}
