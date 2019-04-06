package adjancencylist

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
