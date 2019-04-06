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

func TestAddEdge(t *testing.T) {
	g := makeGraph()
	vertex1 := "Tokyo"
	vertex2 := "China"
	g.addVertex(vertex1)
	g.addVertex(vertex2)
	g.addEdge(vertex1, vertex2)

	actual := g.AdjacencyList[vertex1][0]
	want := vertex2
	if actual != vertex2 {
		t.Errorf("Vertex %s should have %s as first edge, but got %s instead",
			vertex1, want, actual)
	}

	actual = g.AdjacencyList[vertex2][0]
	if actual != vertex1 {
		t.Errorf("Vertex %s should have %s as first edge, but got %s instead",
			vertex2, want, actual)
	}
}
