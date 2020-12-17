package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/EnriqueL8/adventofcode/utils"
)

type Graph map[string][]string

func ReadLines(path, splitOn string) ([]string, error) {
	content, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return nil, readFileErr
	}
	lines := strings.Split(string(content), splitOn)

	return lines, nil
}

func tvc(graph Graph, find string) (results [][]string) {
	visited := []string{}
	for m, modules := range graph {
		if len(modules) == 0 {
			continue
		}

		for _, subTree := range tvcModules(graph, visited, modules, find) {
			subTree = append(subTree, m)
			results = append(results, subTree)
		}

	}

	return results
}

func visitedAleady(visited []string, module string) bool {
	for _, m := range visited {
		if m == module {
			return true
		}
	}
	return false
}

func tvcModules(graph Graph, visited []string, modules []string, find string) (results [][]string) {
	for _, module := range modules {
		if visitedAleady(visited, module) {
			continue
		}
		visited = append(visited, module)
		if strings.Contains(module, find) {
			return [][]string{[]string{module}}
		}

		if subModules, ok := graph[module]; ok {
			for _, subTree := range tvcModules(graph, visited, subModules, find) {
				subTree = append(subTree, module)
				results = append(results, subTree)
			}
		}
	}

	return results
}

func getLongestPath(paths [][]string) []string {
	longestPath := []string{}
	for _, path := range paths {
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}
	return longestPath
}

func prettyPrint(path []string) {
	for i, m := range path {
		for v := 0; v < i; v++ {
			fmt.Printf("  ")
		}
		fmt.Printf("|__ %s\n", m)
	}
}

func revertPath(path []string) (rPath []string) {
	rPath = []string{}
	for i := len(path) - 1; i >= 0; i-- {
		rPath = append(rPath, path[i])
	}

	return
}

func main() {
	filename := os.Args[1]
	if filename == "" {
		filename = "graph.txt"
	}
	searchDep := os.Args[2]
	lines, _ := utils.ReadLines(filename, "\n")

	fmt.Printf("Searching dependencies for %s: \n", searchDep)

	graph := Graph{}
	for _, line := range lines {
		values := strings.Split(line, " ")
		if len(values) != 2 {
			continue
		}
		key := values[0]
		value := values[1]
		if _, ok := graph[key]; !ok {
			graph[key] = []string{value}
		}

		graph[key] = append(graph[key], value)
	}

	paths := tvc(graph, searchDep)
	longestPath := getLongestPath(paths)
	path := revertPath(longestPath)
	prettyPrint(path)

}
