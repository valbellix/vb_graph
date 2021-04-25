package graph

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	nodes int
	edges int
}

func checkProblem(g Graph, p *problem) bool {
	return len(g.Nodes()) == p.nodes && len(g.Edges()) == p.edges
}

func ParseDIMACS(file, edgeMarker string) (Graph, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	g := NewGraph()

	var p *problem = nil

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		if tokens[0] != edgeMarker && tokens[0] != "p" {
			continue
		}

		if len(tokens) < 4 {
			return nil, errors.New("parsing error: dimacs file not well formed")
		}

		if tokens[0] == "p" {
			if tokens[1] == "edge" {
				nodes, err := strconv.Atoi(tokens[2])
				if err != nil {
					return nil, errors.New("parsing error: dimacs files contains a wrong 'problem' line. Field 'nodes' is not an integer")
				}

				edges, err := strconv.Atoi(tokens[3])
				if err != nil {
					return nil, errors.New("parsing error: dimacs files contains a wrong 'problem' line. Field 'edges' is not an integer")
				}

				p = &problem{nodes, edges}
			}
			continue
		}

		// try to parse the weight
		w, err := strconv.Atoi(tokens[3])
		if err != nil {
			return nil, err
		}

		g.AddEdge(tokens[1], tokens[2], w)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	if p != nil {
		if !checkProblem(g, p) {
			return nil, errors.New("error: a problem line exists and it should not match the dimacs file content")
		}
	}

	return g, nil
}
