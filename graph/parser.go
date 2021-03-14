package graph

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func ParseDIMACS(file, edgeMarker string) (Graph, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	g := NewGraph()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		if tokens[0] != edgeMarker {
			continue
		}

		if len(tokens) < 4 {
			return nil, errors.New("parsing error: dimacs file not well formed")
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

	return g, nil
}
