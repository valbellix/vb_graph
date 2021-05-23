package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
	"vb_graph/graph"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	var file string
	var edgeMarker string
	var cpuprofile string

	flag.StringVar(&file, "file", "", "path to a valid DIMACS file")
	flag.StringVar(&edgeMarker, "marker", "e", "edge marker")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")

	flag.Parse()

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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

	g, err := graph.ParseDIMACS(file, edgeMarker)
	end := makeTimestamp()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Building the graph took %v ms\n", (end - beginning))
	fmt.Printf("The graph has %v nodes and %v edges.\n", len(g.Nodes()), len(g.Edges()))

	beginning = makeTimestamp()
	_, _, err = graph.ShortestPath(g, g.GetRoot())
	if err != nil {
		fmt.Println(err)
		return
	}
	end = makeTimestamp()
	fmt.Printf("Calculating the shortest path from the root to all nodes of the graph took %v ms\n", (end - beginning))
}
