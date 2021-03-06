package adjancencylist

import "fmt"

type Graph struct {
	AdjacencyList map[string][]string
}

func makeGraph() Graph {
	g := Graph{}
	g.AdjacencyList = make(map[string][]string)
	return g
}

func (g *Graph) addVertex(name string) {
	if _, ok := g.AdjacencyList[name]; ok {
		return
	}
	g.AdjacencyList[name] = []string{}
}

func (g *Graph) addEdge(v1 string, v2 string) error {
	if _, ok := g.AdjacencyList[v1]; !ok {
		return fmt.Errorf("%s is not a vertex.", v1)
	}
	if _, ok := g.AdjacencyList[v2]; !ok {
		return fmt.Errorf("%s is not a vertex.", v2)
	}

	g.AdjacencyList[v1] = append(g.AdjacencyList[v1], v2)
	g.AdjacencyList[v2] = append(g.AdjacencyList[v2], v1)
	return nil
}

func (g *Graph) removeEdge(v1 string, v2 string) {
	// Loop through the vertex's list of edges
	for i, vertex := range g.AdjacencyList[v1] {
		if vertex == v2 {
			//remove vertex from the slice
			g.AdjacencyList[v1] = append(g.AdjacencyList[v1][:i], g.AdjacencyList[v1][i+1:]...)
		}
	}

	//Loop through the vertex's list of edges
	for i, vertex := range g.AdjacencyList[v2] {
		if vertex == v1 {
			//remove vertex from the slice
			g.AdjacencyList[v2] = append(g.AdjacencyList[v2][:i], g.AdjacencyList[v2][i+1:]...)
		}
	}
}

func (g *Graph) removeVertex(v string) {
	//Delete the vertex
	delete(g.AdjacencyList, v)
	//Loop through all the vertexes to remove the vertex from the edges
	for v2, _ := range g.AdjacencyList {
		g.removeEdge(v, v2)
	}
}
