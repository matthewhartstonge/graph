package edge

import (
	// Internal Imports
	"github.com/matthewhartstonge/graph/vertex"
)

type Edger interface {
	Cost() float64
	SetCost(cost float64)
	Directed() bool
	SetDirected(directed bool)
	Label() string
	SetLabel(label string)
	Vertex1() vertex.Vertexer
	SetVertex1(vertex1 vertex.Vertexer)
	Vertex2() vertex.Vertexer
	SetVertex2(vertex2 vertex.Vertexer)
}

func New(vertex1 vertex.Vertexer, vertex2 vertex.Vertexer, opts ...Option) Edger {
	edge := &Edge{
		cost:     0,
		directed: false,
		label:    "",
		vertex1:  vertex1,
		vertex2:  vertex2,
	}

	for _, opt := range opts {
		opt(edge)
	}

	return newEdge(edge, vertex1, vertex2)
}

// newEdge performs setup required for linking vertices together, so the links
// are setup in such a way that any path can be traversed the an edge.
func newEdge(edge *Edge, v1, v2 vertex.Vertexer) Edger {
	edge.SetVertex1(v1)
	edge.SetVertex2(v2)

	return edge
}

type Option func(edge Edger)

func WithDirected(directed bool) Option {
	return func(edge Edger) {
		edge.SetDirected(directed)
	}
}

func WithCost(cost float64) Option {
	return func(edge Edger) {
		edge.SetCost(cost)
	}
}

// Edge provides the data structure for a curve, or arc, which details a path
// between two given vertices within a graph.
//
// When solving a state-space problem, an edge represents an action.
type Edge struct {
	// Cost provides a weight for deducing edge satisfiability for heuristic,
	// or cost/weight constrained solutions.
	cost float64
	// Directed specifies whether the edge is either bi-directional, or
	// directed, that is Vertex1 -> Vertex2.
	directed bool
	// Label provides context to an Edge, for example, it may provide the
	// action that will take the agent from one vertex to another.
	label string
	// Vertex1 provides the first vertex on an edge.
	// On a directed edge, Vertex1 is the starting vertex.
	vertex1 vertex.Vertexer
	// Vertex2 provides the second vertex on an edge.
	// On a directed edge, Vertex2 is the endpoint.
	vertex2 vertex.Vertexer
}

func (e Edge) Cost() float64 {
	return e.cost
}

func (e *Edge) SetCost(cost float64) {
	e.cost = cost
}

func (e Edge) Directed() bool {
	return e.directed
}

func (e *Edge) SetDirected(directed bool) {
	e.directed = directed
}

func (e Edge) Label() string {
	return e.label
}

func (e *Edge) SetLabel(label string) {
	e.label = label
}

func (e Edge) Vertex1() vertex.Vertexer {
	return e.vertex1
}

func (e *Edge) SetVertex1(vertex1 vertex.Vertexer) {
	if e.vertex1 == nil {
		return
	}

	e.vertex1 = vertex1
	setFamily(e)
}

func (e Edge) Vertex2() vertex.Vertexer {
	return e.vertex2
}

func (e *Edge) SetVertex2(vertex2 vertex.Vertexer) {
	if e.vertex2 == nil {
		return
	}

	e.vertex2 = vertex2
	setFamily(e)
}

func setFamily(edger Edger) {
	if edger.Vertex1() != nil && edger.Vertex2() != nil {
		switch edger.Directed() {
		case false:
			edger.Vertex2().AddChild(edger.Vertex1())
			edger.Vertex1().AddParent(edger.Vertex2())
			fallthrough

		case true:
			edger.Vertex1().AddChild(edger.Vertex2())
			edger.Vertex2().AddParent(edger.Vertex1())
		}
	}
}
