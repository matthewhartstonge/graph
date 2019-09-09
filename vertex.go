package csp

func NewVertex(k string) *Vertex {
	return &Vertex{
		K: k,
		// V: v,
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

type Vertex struct {
	// K provides the name of the node.
	K string
	// V provides the nodes value.
	V int
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
		if v.K == vertex.K {
			// node already added to parents.
			return src, true
		}
	}

	return append(src, vertex), false
}
