package graphs

type Graph interface {
	AddVertex(v string)
	AddEdge(v1, v2 string)
	GetNeighbors(v string) []string
	GetVertices() []string
}

type BasicGraph struct {
	vertices map[string]struct{}
	edges    map[string]map[string]struct{}
}

func NewSimpleGraph() *BasicGraph {
	return &BasicGraph{
		vertices: make(map[string]struct{}),
		edges:    make(map[string]map[string]struct{}),
	}
}

func (g *BasicGraph) AddVertex(v string) {
	g.vertices[v] = struct{}{}
	if _, exists := g.edges[v]; !exists {
		g.edges[v] = make(map[string]struct{})
	}
}

func (g *BasicGraph) AddEdge(v1, v2 string) {
	if _, exists := g.vertices[v1]; !exists {
		g.AddVertex(v1)
	}
	if _, exists := g.vertices[v2]; !exists {
		g.AddVertex(v2)
	}
	g.edges[v1][v2] = struct{}{}
	g.edges[v2][v1] = struct{}{}
}

func (g *BasicGraph) GetNeighbors(v string) []string {
	neighbors := []string{}
	if neighborsSet, exists := g.edges[v]; exists {
		for neighbor := range neighborsSet {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (g *BasicGraph) GetVertices() []string {
	vertices := []string{}
	for vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}
