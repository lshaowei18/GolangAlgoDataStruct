package adjancencylist

import (
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := Graph{}
	want := "Tokyo"
	g.addVertex(want)

	if _, ok := g.AdjacencyList[want]; !ok {
		t.Errorf("Key %s not added to AdjancencyList in graph.", want)
	}
}
