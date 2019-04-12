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


	g.AdjacencyList[v1] = append(g.AdjacencyList[v1], v2)
	g.AdjacencyList[v2] = append(g.AdjacencyList[v2], v1)
	return nil
}
