package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/EnriqueL8/adventofcode/utils"
)

type Graph map[string][]string

// ReadLines reads a file, splits it into lines by the given splitter
// and returns an array of strings
func ReadLines(path, splitOn string) ([]string, error) {
	content, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return nil, readFileErr
	}
	lines := strings.Split(string(content), splitOn)

	return lines, nil
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

	b, err := json.MarshalIndent(graph[searchDep], "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
