package generator

import (
	"fmt"
	"github.com/guilhermocc/test-case-generator/internal/parser"
	"strings"
)

var statesCounter = map[string]int{}

func GenerateTestCases(inputFilePath string) {
	initialState, events, transitionsMap, err := parser.ParseStateMachineCsv(inputFilePath)
	if err != nil {
		panic(err)
	}

	transitionsTree := generateTransitionsTree(initialState, events, transitionsMap)

	fmt.Println("** Transitions Tree **")
	printTree(transitionsTree)
	fmt.Println("** Transitions Tree **")

	fmt.Println()

	fmt.Println("** Basic Paths **")
	printAllPaths(transitionsTree, []string{})
	fmt.Println("** Basic Paths **")
}

func printTree(tree *Node) {
	if tree.Children == nil {
		fmt.Println(tree.StateCount)
		return
	}

	nodesQueue := []*Node{tree}

	for len(nodesQueue) > 0 {
		node := nodesQueue[0]
		nodesQueue = nodesQueue[1:]

		for _, child := range node.Children {
			fmt.Printf("%s, %s \n", node.StateCount, child.StateCount)
			nodesQueue = append(nodesQueue, child)
		}
	}
}

func printAllPaths(tree *Node, path []string) {
	if tree.Children == nil {
		path = append(path, tree.StateCount)
		fmt.Println(strings.Join(path, " -> "))
		return
	}

	path = append(path, tree.StateCount)

	for _, child := range tree.Children {
		printAllPaths(child, path)
	}
}

func generateTransitionsTree(initialState string, events []string, transitionsMap map[string][]string) *Node {
	initialNode := &Node{
		State:      initialState,
		StateCount: fmt.Sprintf("%s_%d", initialState, statesCounter[initialState]),
	}
	statesCounter[initialState]++

	nodesQueue := []*Node{initialNode}

	for len(nodesQueue) > 0 {
		node := nodesQueue[0]
		nodesQueue = nodesQueue[1:]

		for _, transition := range transitionsMap[node.State] {
			if transition != "" {
				if statesCounter[transition] != 0 {
					child := &Node{State: transition, StateCount: fmt.Sprintf("%s_%d", transition, statesCounter[transition])}
					node.Children = append(node.Children, child)
					statesCounter[transition]++
				} else {
					child := &Node{State: transition, StateCount: fmt.Sprintf("%s_%d", transition, statesCounter[transition])}
					node.Children = append(node.Children, child)
					statesCounter[transition]++
					nodesQueue = append(nodesQueue, child)
				}

			}
		}

	}

	return initialNode
}

type Node struct {
	State      string
	StateCount string
	Children   []*Node
}
