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

package edge

import (
	// Standard Library Imports
	"fmt"

	// Internal Imports
	"github.com/matthewhartstonge/graph/vertex"
)

// Edger provides an implementation for an edge within a graph.
//
// Edges lead from one vertex to another, and as such, provide a way to traverse
// a graph. This traversal throughout a graph provides the hope that one of the
// next head vertices might be able to satisfy the goal state.
type Edger interface {
	// Cost returns the cost of traversing the edge.
	Cost() float64
	// SetCost enables the edge cost to be set.
	SetCost(cost float64)

	// Directed returns true if the edge is a directed edge. This informs us
	// that, if the edge is directed, the edge can only be traversed from from
	// the tail vertex, to the head vertex. An undirected edge means that the
	// edge can be traversed from either way.
	Directed() bool
	// SetDirected enables setting the edge as either being directed or
	// undirected.
	SetDirected(directed bool)

	// Label provides a name for the edge.
	Label() string
	// SetLabel sets a name for the edge.
	SetLabel(label string)

	// Tail returns the starting point of the edge.
	// If the tail is nil, this edge is the starting of a path.
	Tail() vertex.Vertexer
	// SetTail enables setting the start point of an edge.
	SetTail(tail vertex.Vertexer)

	// Head returns the end point of the edge.
	// If the head vertex holds no neighbours, this edge actually contains an
	// isolated vertex with no edges.
	Head() vertex.Vertexer
	// SetHead enables setting the end point of an edge.
	SetHead(head vertex.Vertexer)
}

// New returns a new directed edge. The edge can be mutated by providing
// variadic options.
func New(tail vertex.Vertexer, head vertex.Vertexer, opts ...Option) Edger {
	edge := &Edge{
		cost:     0,
		directed: true,
		label:    "",
		tail:     tail,
		head:     head,
	}

	for _, opt := range opts {
		opt(edge)
	}

	return newEdge(edge, tail, head)
}

// newEdge performs setup required for linking vertices together, so the links
// are setup in such a way that any path can be traversed the an edge.
func newEdge(edge *Edge, v1, v2 vertex.Vertexer) Edger {
	edge.SetTail(v1)
	edge.SetHead(v2)

	return edge
}

// Option provides options to mutate a given edge on initialization.
type Option func(edge Edger)

// WithCost sets the cost of the edge.
// By default, an edge costs nothing to traverse.
func WithCost(cost float64) Option {
	return func(edge Edger) {
		edge.SetCost(cost)
	}
}

// WithLabel sets the label for a given edge.
// By default, an edge has no label.
func WithLabel(label string) Option {
	return func(edge Edger) {
		edge.SetLabel(label)
	}
}

// WithUndirected sets the edge as being undirected, allowing traversal from
// both sides.
// By default, an edge is marked as being directed.
func WithUndirected() Option {
	return func(edge Edger) {
		edge.SetDirected(false)
	}
}

// Edge provides the concrete data structure for an Edger, that is a curve, or
// arc, which details a path between two given vertices within a graph.
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

	// tail provides the first vertex on an edge.
	// On a directed edge, tail is the starting vertex.
	tail vertex.Vertexer

	// head provides the second vertex on an edge.
	// On a directed edge, head is the endpoint.
	head vertex.Vertexer
}

// Cost returns the cost of traversing the edge.
func (e Edge) Cost() float64 {
	return e.cost
}

// SetCost sets the cost of traversing the edge.
func (e *Edge) SetCost(cost float64) {
	e.cost = cost
}

// Directed returns true if the edge is a directed edge.
func (e Edge) Directed() bool {
	return e.directed
}

// SetDirected enables setting the edge as either being directed or
// undirected.
func (e *Edge) SetDirected(directed bool) {
	e.directed = directed
}

// Label returns the name set for the edge.
func (e Edge) Label() string {
	return e.label
}

// SetLabel sets a name for the edge.
func (e *Edge) SetLabel(label string) {
	e.label = label
}

// Tail returns the starting point of the edge.
// If the tail is nil, this edge is the starting of a path.
func (e Edge) Tail() vertex.Vertexer {
	return e.tail
}

// SetTail enables setting the start point of an edge.
func (e *Edge) SetTail(tail vertex.Vertexer) {
	if e.tail == nil {
		return
	}

	e.tail = tail
	setFamily(e)
}

// Head returns the end point of the edge.
// If the head vertex holds no neighbours, this edge actually contains an
// isolated vertex with no edges.
func (e Edge) Head() vertex.Vertexer {
	return e.head
}

// SetHead enables setting the end point of an edge.
func (e *Edge) SetHead(head vertex.Vertexer) {
	if e.head == nil {
		return
	}

	e.head = head
	setFamily(e)
}

// setFamily, based on knowledge provided by the edge, binds in neighbours to
// build an adjacency list on the vertices.
func setFamily(edger Edger) {
	if edger.Tail() != nil && edger.Head() != nil {
		switch edger.Directed() {
		case false:
			edger.Head().AddChild(edger.Tail())
			edger.Tail().AddParent(edger.Head())
			fallthrough

		case true:
			edger.Tail().AddChild(edger.Head())
			edger.Head().AddParent(edger.Tail())
		}
	}
}

// String implements Stringer.
// Edge returns the edge's label, if set, or meta information about the edge to
// help provide context, or a greater understanding of part of a path.
func (e Edge) String() string {
	if e.Label() != "" {
		return e.Label()
	}

	return e.edgeMeta()
}

// edgeMeta generates meta information about the edge to help understand the
// part this edge plays within a path.
func (e Edge) edgeMeta() string {
	cost := ""
	if e.Cost() > 0 {
		cost = fmt.Sprintf("(%.1f)", e.Cost())
	}

	tailDirection := "<"
	headDirection := ">"
	if e.Directed() {
		tailDirection = ""
	}

	return fmt.Sprintf(
		"(%s) %s-%s-%s (%s)",
		e.Tail().Label(),
		tailDirection, cost, headDirection,
		e.Head().Label(),
	)
}
