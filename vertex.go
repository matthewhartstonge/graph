package csp

import (
	"github.com/sirupsen/logrus"
)

func NewVertex(label string) Vertexer {
	return &Vertex{
		label: label,
		// Value: value,
		visited:  false,
		parents:  []Vertexer{},
		children: []Vertexer{},
	}
}

type Vertexer interface {
	Label() string
	SetLabel(label string)
	Children() []Vertexer
	AddChild(vertexer Vertexer) Vertexer
	Parents() []Vertexer
	AddParent(vertexer Vertexer) Vertexer
	Visited() bool
	SetVisited(visited bool)
}

// Vertex provides the data structure for a node, or point, within a given
// graph.
//
// When solving a state-space problem, a vertex represents a state.
type Vertex struct {
	// Label provides the name of the vertex.
	label string
	// children contains the vertices this vertex is a parent of.
	children []Vertexer
	// parents contains the vertices that this vertex links to.
	parents []Vertexer
	// visited specifies if the node has been visited.
	visited bool

	// Value provides the vertex's value.
	Value int
}

func (v Vertex) Label() string {
	return v.label
}

func (v *Vertex) SetLabel(label string) {
	v.label = label
}

func (v Vertex) Children() []Vertexer {
	return v.children
}

func (v *Vertex) AddChild(vertex Vertexer) Vertexer {
	logger := logrus.WithFields(logrus.Fields{
		"struct":  "Vertex",
		"method":  "AddChild",
		"vertex1": v.Label(),
		"vertex2": vertex.Label(),
	})

	c, found := addToSet(v.children, vertex)
	if found {
		logger.Debug("child exists")
		return v
	}

	logger.Debug("added child")
	v.children = c
	return v
}

func (v Vertex) Parents() []Vertexer {
	return v.parents
}

func (v *Vertex) AddParent(vertex Vertexer) Vertexer {
	logger := logrus.WithFields(logrus.Fields{
		"struct":  "Vertex",
		"method":  "AddParent",
		"vertex1": vertex.Label(),
		"vertex2": v.Label(),
	})

	p, found := addToSet(v.parents, vertex)
	if found {
		logger.Debug("parent exists")
		return v
	}

	logger.Debug("added parent")
	v.parents = p
	return v
}

func (v Vertex) Visited() bool {
	return v.visited
}

func (v *Vertex) SetVisited(visited bool) {
	v.visited = visited
}

// addToSet ensures no duplicates are added to the array based on key.
func addToSet(src []Vertexer, vertex Vertexer) (dst []Vertexer, found bool) {
	for _, v := range src {
		if v.Label() == vertex.Label() {
			// node already added to parents.
			return src, true
		}
	}

	return append(src, vertex), false
}
