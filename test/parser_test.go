package test

import (
	"testing"
	"vb_graph/graph"

	"github.com/stretchr/testify/assert"
)

func TestDIMACSParser(t *testing.T) {
	// no problem line
	t.Log("No 'problem' line")
	_, err := graph.ParseDIMACS("data/dummy.dimacs", "e")
	assert.Nil(t, err, "error should be nil")

	t.Log("Correct 'problem' line")
	_, err = graph.ParseDIMACS("data/test_path.dimacs", "e")
	assert.Nil(t, err, "error should be nil")

	t.Log("Wrong 'problem' line")
	_, err = graph.ParseDIMACS("data/ugly.dimacs", "e")
	assert.NotNil(t, err, "error should be set to a not-nil value")
}
