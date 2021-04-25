package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"vb_graph/graph"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	var file string
	flag.StringVar(&file, "file", "", "path to a valid DIMACS file")

	if file == "" {
		fmt.Println("File is not specified")
		return
	}

	if _, err := os.Stat(file); err != nil {
		fmt.Printf("File %v does not exists", file)
		return
	}

	fmt.Println("Parsing the graph from a DIMACS file")
	beginning := makeTimestamp()

	g, err := graph.ParseDIMACS(file, "e")
	end := makeTimestamp()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Building the graph took %v ms\n", (end - beginning))

	beginning = makeTimestamp()
	_, _, err = graph.ShortestPath(g, g.GetRoot())
	if err != nil {
		fmt.Println(err)
		return
	}
	end = makeTimestamp()
	fmt.Printf("Calculating the shortest path from the root to all nodes of the graph took %v ms\n", (end - beginning))
}
