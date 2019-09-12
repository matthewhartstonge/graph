package csp

func NewDirectedEdge(a, b *Vertex) *Edge {
	return &Edge{
		Vertex1:  a,
		Vertex2:  b,
		Directed: true,
	}
}

func NewUndirectedEdge(a, b *Vertex) *Edge {
	return &Edge{
		Vertex1:  a,
		Vertex2:  b,
		Directed: false,
	}
}

type Edger interface {
}

// Edge provides the data structure for a curve, or arc, which details a path
// between two given vertices within a graph.
//
// When solving a state-space problem, an edge represents an action.
type Edge struct {
	// Cost provides a weight for deducing edge satisfiability for heuristic,
	// or cost/weight constrained solutions.
	Cost float64
	// Directed specifies whether the edge is either bi-directional, or
	// directed, that is Vertex1 -> Vertex2.
	Directed bool
	// Label provides context to an Edge, for example, it may provide the
	// action that will take the agent from one vertex to another.
	Label string
	// Vertex1 provides the first vertex on an edge.
	// On a directed edge, Vertex1 is the starting vertex.
	Vertex1 *Vertex
	// Vertex2 provides the second vertex on an edge.
	// On a directed edge, Vertex2 is the endpoint.
	Vertex2 *Vertex
}
