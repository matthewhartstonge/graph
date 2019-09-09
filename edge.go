package csp

func NewDirectedEdge(a, b *Vertex) *Edge {
	return &Edge{
		A:          a,
		B:          b,
		Undirected: false,
	}
}

func NewUndirectedEdge(a, b *Vertex) *Edge {
	return &Edge{
		A:          a,
		B:          b,
		Undirected: true,
	}
}

type Edger interface {
}

// Edge provides the data structure for a graph's Edge.
type Edge struct {
	A          *Vertex
	B          *Vertex
	Undirected bool
}
