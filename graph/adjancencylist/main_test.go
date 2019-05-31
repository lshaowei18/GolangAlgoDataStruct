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
	t.Run("Both vertexes are added", func(t *testing.T) {
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
	})

	t.Run("One vertex is not added ", func(t *testing.T) {
		g := makeGraph()
		g.addVertex("China")
		err := g.addEdge("China", "Korea")
		if err == nil {
			t.Errorf("Should return an error since Korea is not added as a vertex.")
		}
	})
}

func TestRemoveEdge(t *testing.T) {
	t.Run("Both vertices and edges are added", func(t *testing.T) {
		g := makeGraph()
		v1 := "China"
		v2 := "Korea"
		g.addVertex(v1)
		g.addVertex(v2)
		g.addEdge(v1, v2)

		g.removeEdge(v1, v2)

		if len(g.AdjacencyList[v1]) != 0 {
			t.Errorf("Length of %v should be 0 after removeEdge.", v1)
		}
		if len(g.AdjacencyList[v2]) != 0 {
			t.Errorf("Length of %v should be 0 after removeEdge.", v2)
		}

	})
}

func TestRemoveVertex(t *testing.T) {
	t.Run("Remove vertice", func(t *testing.T) {
		//Make a graph
		g := makeGraph()
		v1 := "China"
		v2 := "Korea"
		g.addVertex(v1)
		g.addVertex(v2)
		g.addEdge(v1, v2)

		//Remove vertex
		g.removeVertex(v1)

		//Check if v1 is removed
		if _, ok := g.AdjacencyList[v1]; ok {
			t.Errorf("%v should not be in the list.", v1)
		}

		//Check if v1 is inside v2's list of edges
		if len(g.AdjacencyList[v2]) != 0 {
			t.Errorf("%v's list of edges should not contain %v.", v2, v1)
		}
	})
}
