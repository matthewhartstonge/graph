/*
 * Copyright (C) 2019. Matthew Hartstonge <matt@mykro.co.nz>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package vertex

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

func New(label string) Vertexer {
	return &Vertex{
		label: label,
		// Value: value,
		visited:  false,
		parents:  []Vertexer{},
		children: []Vertexer{},
	}
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
	c, found := addToSet(v.children, vertex)
	if found {
		return v
	}

	v.children = c
	return v
}

func (v Vertex) Parents() []Vertexer {
	return v.parents
}

func (v *Vertex) AddParent(vertex Vertexer) Vertexer {
	p, found := addToSet(v.parents, vertex)
	if found {
		return v
	}

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
