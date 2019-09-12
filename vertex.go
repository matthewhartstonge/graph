package csp

func NewVertex(label string) *Vertex {
	return &Vertex{
		Label: label,
		// Value: value,
		visited:  false,
		parents:  []*Vertex{},
		children: []*Vertex{},
	}
}

type Vertexer interface {
	Visited() bool
	SetVisited()
	AddChild(v *Vertex) *Vertex
	AddParent(v *Vertex) *Vertex
}

// Vertex provides the data structure for a node, or point, within a given
// graph.
//
// When solving a state-space problem, a vertex represents a state.
type Vertex struct {
	// Label provides the name of the vertex.
	Label string
	// Value provides the vertex's value.
	Value int
	// visited specifies if the node has been visited.
	visited bool
	// parents contains the vertices that this vertex links to.
	parents []*Vertex
	// children contains the vertices this vertex is a parent of.
	children []*Vertex
}

func (v Vertex) Visited() bool {
	return v.visited
}

func (v *Vertex) SetVisited() {
	v.visited = true
}

func (v *Vertex) AddParent(vertex *Vertex) *Vertex {
	p, found := addToSet(v.parents, vertex)
	if found {
		return v
	}

	v.parents = p
	return vertex.AddChild(v)
}

func (v *Vertex) AddChild(vertex *Vertex) *Vertex {
	c, found := addToSet(v.children, vertex)
	if found {
		return v
	}

	v.children = c
	return vertex.AddParent(v)
}

// addToSet ensures no duplicates are added to the array based on key.
func addToSet(src []*Vertex, vertex *Vertex) (dst []*Vertex, found bool) {
	for _, v := range src {
		if v.Label == vertex.Label {
			// node already added to parents.
			return src, true
		}
	}

	return append(src, vertex), false
}
