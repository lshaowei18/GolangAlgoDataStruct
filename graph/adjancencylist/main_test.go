package adjancencylist

import (
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := makeGraph()

	t.Run("Add unique vertex", func(t *testing.T) {
		want := "Tokyo"
		g.addVertex(want)

		if _, ok := g.AdjacencyList[want]; !ok {
			t.Errorf("Key %s not added to AdjancencyList in graph.", want)
		}
	})
	t.Run("add non-unique vertex", func(t *testing.T) {
		want := "Tokyo"
		g.AdjacencyList[want] = append(g.AdjacencyList[want], "Hi")

		g.addVertex(want)

		if len(g.AdjacencyList[want]) == 0 {
			t.Errorf("Repeated key should not replace existing list")
		}
	})
}
