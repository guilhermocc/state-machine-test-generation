package generator

import (
	"encoding/json"
	"fmt"
	"github.com/guilhermocc/test-case-generator/internal/parser"
	"os"
)

var statesCounter = map[string]int{}
var basicPaths = [][]byte{}

func (p Path) String() string {
	path := p.StateCount
	if p.Event != "" {
		path += fmt.Sprintf(" - |%s|", p.Event)
	}
	if p.NextPath != nil {
		path += " - " + p.NextPath.String()

	}
	return path
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GenerateTestCases(eventsMappingTablePath, transitionsTablePath string) {
	initialState, events, transitionsMap, err := parser.ParseStateMachineCsv(transitionsTablePath)
	check(err)

	eventsActions, err := parser.ParseEventsActionsCsv(eventsMappingTablePath)
	check(err)

	file, err := os.Create("result.md")
	check(err)
	defer file.Close()

	file.Write([]byte("# Test generation result \n"))

	fmt.Println("Generating transitions tree...")
	transitionsTree := generateTransitionsTree(initialState, events, transitionsMap)
	printTree(file, transitionsTree)

	fmt.Println("Traversing tree to generate basic paths...")
	traverseBasicPaths(eventsActions, transitionsTree, nil)
	printBasicPaths(file)

	fmt.Println("Generating test scripts...")
	printAllTestCasesPaths(file, eventsActions)
}

func printBasicPaths(file *os.File) {
	_, err := file.Write([]byte("## Basic Paths\n```\n"))
	check(err)

	for _, path := range basicPaths {
		var finalPath Path
		json.Unmarshal(path, &finalPath)
		file.Write([]byte(fmt.Sprintf("%s\n", finalPath.String())))
	}
	_, err = file.Write([]byte("```\n"))
	check(err)
}

func traverseBasicPaths(eventsActions map[string]string, tree *Node, currentPath *Path) {
	if tree.Children == nil {
		if currentPath == nil {
			finalPath, _ := json.Marshal(Path{State: tree.State, StateCount: tree.StateCount, Event: ""})
			basicPaths = append(basicPaths, finalPath)
		} else {
			path := currentPath
			for path.NextPath != nil {
				path = path.NextPath
			}
			path.NextPath = &Path{State: tree.State, StateCount: tree.StateCount, Event: ""}
			finalPath, _ := json.Marshal(currentPath)
			basicPaths = append(basicPaths, finalPath)
			path.NextPath = nil
		}
	}

	for _, child := range tree.Children {

		if currentPath == nil {
			currentPath = &Path{State: tree.State, StateCount: tree.StateCount, Event: child.Event}
			traverseBasicPaths(eventsActions, child.Node, currentPath)
		} else {
			path := currentPath
			for path.NextPath != nil {
				path = path.NextPath
			}
			path.NextPath = &Path{State: tree.State, StateCount: tree.StateCount, Event: child.Event}
			traverseBasicPaths(eventsActions, child.Node, currentPath)
			path.NextPath = nil
		}

	}

}

func printTree(file *os.File, tree *Node) {
	_, err := file.Write([]byte("## Transitions Tree\n```\n"))
	check(err)

	if tree.Children == nil {
		fmt.Println(tree.StateCount)
		return
	}

	nodesQueue := []*Node{tree}

	for len(nodesQueue) > 0 {
		node := nodesQueue[0]
		nodesQueue = nodesQueue[1:]

		for _, child := range node.Children {
			_, err := file.Write([]byte(fmt.Sprintf("%s - |%s| -  %s \n", node.StateCount, child.Event, child.Node.StateCount)))
			check(err)
			nodesQueue = append(nodesQueue, child.Node)
		}
	}
	_, err = file.Write([]byte("```\n"))
	check(err)
}

func printAllTestCasesPaths(file *os.File, eventsActions map[string]string) {
	_, err := file.Write([]byte("## Generated test script\n``` go\n"))
	check(err)

	for i, path := range basicPaths {
		var finalPath Path
		json.Unmarshal(path, &finalPath)
		file.Write([]byte(fmt.Sprintf("t.Run(\"TestPath%d\", func(t *testing.T) {", i+1)))
		file.Write([]byte("\n\tdevice := &Device{State: \"OFF\"}"))

		nextPath := &finalPath

		for nextPath != nil {
			file.Write([]byte(fmt.Sprintf("\n\tassert.Equal(t, \"%s\", device.GetState())", nextPath.State)))
			if nextPath.Event != "" {
				file.Write([]byte(fmt.Sprintf("\n\t%s", eventsActions[nextPath.Event])))
			}
			nextPath = nextPath.NextPath
		}

		file.Write([]byte("\n})\n"))

	}
	_, err = file.Write([]byte("```\n"))
	check(err)

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

		for i, transition := range transitionsMap[node.State] {
			if transition != "" {
				if statesCounter[transition] != 0 {
					child := &NodeChild{
						Event: events[i],
						Node:  &Node{State: transition, StateCount: fmt.Sprintf("%s_%d", transition, statesCounter[transition])},
					}
					node.Children = append(node.Children, child)
					statesCounter[transition]++
				} else {
					child := &NodeChild{
						Event: events[i],
						Node:  &Node{State: transition, StateCount: fmt.Sprintf("%s_%d", transition, statesCounter[transition])},
					}
					node.Children = append(node.Children, child)
					statesCounter[transition]++
					nodesQueue = append(nodesQueue, child.Node)
				}

			}
		}

	}

	return initialNode
}

type Path struct {
	State      string `json:"State"`
	StateCount string `json:"StateCount"`
	Event      string `json:"Event"`
	NextPath   *Path  `json:"NextPath"`
}

type Node struct {
	State      string
	StateCount string
	Children   []*NodeChild
}

type NodeChild struct {
	Event string
	Node  *Node
}
