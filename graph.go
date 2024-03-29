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

package graph

import (
	// Standard Library Imports
	"fmt"

	// External Imports
	log "github.com/sirupsen/logrus"

	// Internal Imports
	"github.com/matthewhartstonge/graph/edge"
	"github.com/matthewhartstonge/graph/path"
	"github.com/matthewhartstonge/graph/vertex"
)

// Grapher is the interface that wraps the graph search algorithms.
//
// Search traverses a graph in order to find a solution. If no solution is
// found, nil will be returned. Search should be implemented in such a way that
// multiple calls to search can continue to find other solutions from where it
// left off.
type Grapher interface {
	Search() (goalPath path.Pather)
}

// New creates a new graph that to solve a graph search.
func New(options ...Option) *Graph {
	g := &Graph{
		digraph: false,
		V:       []vertex.Vertexer{},
		E:       []edge.Edger{},

		Frontier:         NewDepthFirstSearch(),
		StartingVertices: []vertex.Vertexer{},
		Goal:             nil,
	}

	for _, option := range options {
		option(g)
	}

	return g.preprocess()
}

// Option provides variadic options when creating a new Graph.
type Option func(g *Graph)

// WithVertices provides a way to supply your own vertices when creating a new
// graph.
func WithVertices(vertices []vertex.Vertexer) Option {
	return func(g *Graph) {
		g.V = vertices
	}
}

// WithEdges provides a way to supply your own edges when creating a new graph.
func WithEdges(edges []edge.Edger) Option {
	return func(g *Graph) {
		g.E = edges
	}
}

// WithStartingVertices provides a way to inject vertices to start with.
func WithStartingVertices(vertices ...vertex.Vertexer) Option {
	return func(g *Graph) {
		g.StartingVertices = vertices
	}
}

// WithSearchStrategy provides a way to inject a search strategy, that is, the
// way in which a frontier is expanded in order to find the goal state.
func WithSearchStrategy(strategy Strategizer) Option {
	return func(g *Graph) {
		g.Frontier = strategy
	}
}

// WithGoalFunc provides a way to supply your own algorithm in order to satisfy
// the graph search.
func WithGoalFunc(f GoalFunc) Option {
	return func(g *Graph) {
		g.Goal = f
	}
}

// WithTraceLogging provides a way to enable algorithm trace logging.
func WithTraceLogging() Option {
	return func(g *Graph) {
		log.SetLevel(log.TraceLevel)
		g.traceLog = true
	}
}

// GoalFunc provides the algorithm to check if the provided vertex satisfies
// the goal.
type GoalFunc func(vertex vertex.Vertexer) bool

// Graph provides the data structure for a Graph.
// G = (V, E)
type Graph struct {
	// digraph provides a check to see if the graph is a directed or
	// undirected.
	digraph bool
	// traceLog provides a way to turn on trace logging.
	// Useful for testing a new algorithm or
	// understanding the process
	traceLog bool

	// V contains a set of vertices, also called nodes.
	V []vertex.Vertexer
	// E contains a set of edges, also called links.
	E []edge.Edger

	// Frontier provides the paths that have been, or may yet to be expanded.
	// The way in which the frontier returns paths is known as the search
	// strategy.
	Frontier Strategizer
	// StartingVertices contain the vertices to start solving the graph search
	// from.
	StartingVertices []vertex.Vertexer
	// Goal contains the algorithm to check if a given vertex satisfies the
	// goal state.
	Goal GoalFunc
}

// preprocess performs any upfront initialization required, for example,
// ensuring the frontier has paths to enable solving.
func (g *Graph) preprocess() *Graph {
	g.digraph = true
	for _, e := range g.E {
		// Detect if a directed graph.
		if !e.Directed() {
			g.digraph = false
			break
		}
	}

	// First off, we need to add the starting vertices to the frontier so we
	// have some starting points to attempt to solve the graph search.
	for _, startingVertex := range g.StartingVertices {
		// We add the vertices as a path, but with the tail being null, to
		// enable detection of the start of a path.
		g.Frontier.Add(path.New(path.WithEdge(
			edge.New(
				nil, startingVertex,
				edge.WithLabel("start"),
			),
		)))
	}

	return g
}

// Search implements a generic search algorithm: given a graph, starting
// vertices and a goal, incrementally explore edges from the start vertices.
func (g *Graph) Search() (goalPath path.Pather) {
	// Begin looping over all paths within the frontier so we can
	// attempt to expand a vertex's neighbours in order to build a path to find
	// a solution.
	for {
		// From the frontier pull out the next potential goal path to try and
		// derive a solution.
		goalPath = g.Frontier.Next()
		if goalPath == nil {
			// If we have no more paths left in the frontier, we have found no
			// solution, and as such, need to return home with our tails
			// between our legs.
			return
		}
		g.logTrace(traceRemovePath, goalPath)

		// Given a potential goal path, we need to get the last vertex along
		// the path to check to see if it satisfies the goal.
		headVertex := goalPath.Last().Head()
		if g.Goal(headVertex) {
			// If we manage to find a solution, we will be a good Dobby and
			// tell our master that we did the good.
			return goalPath
		}

		// Otherwise, search the known edges for neighbours to the current
		// vertex.
		for _, knownEdge := range g.E {
			// we've found an edge which has the current vertex as the starting
			// point..
			if knownEdge.Tail() == headVertex {
				// if the edge is not labelled, give it something to show the
				// direction.
				// knownEdge.SetLabel(generateEdgeLabel(knownEdge))

				// Get a deep copy of the path, so that we don't mung data.
				potentialGoalPath := goalPath.Copy()

				// Expand the potential goal path with the vertex's new found
				// neighbour and add the path to the frontier for later
				// processing.
				potentialGoalPath.Append(knownEdge)
				g.Frontier.Add(potentialGoalPath)
				g.logTrace(traceAddPath, potentialGoalPath)
			}
		}
	}
}

// traceAction provides a specific type for tracing actions performed while
// solving a graph search.
type traceAction string

const (
	traceAddPath    traceAction = "+"
	traceRemovePath traceAction = "-"
)

// traceAction enables tracing the operations happening throughout the
// lifecycle of the graph solve.
func (g Graph) logTrace(action traceAction, path path.Pather) {
	if g.traceLog {
		fields := log.Fields{}
		if path.Cost() != 0 {
			fields["cost"] = path.Cost()
		}

		log.WithFields(fields).Trace(fmt.Sprintf("%s %s", action, path))
	}
}

// PrintInfo prints information about the graphs directionality, parents and
// children.
func (g Graph) PrintInfo() {
	graphType := "undirected graph"
	if g.digraph {
		graphType = "digraph"
	}
	fmt.Printf("Graph:\n- is a %s.\n\n", graphType)

	// Print Links
	fmt.Println("Lineage:")
	for _, v := range g.V {
		for _, v := range g.V {
			// Reset visited status
			v.SetVisited(false)
		}

		fmt.Println(printDescendants(v))
	}

	fmt.Println("\nAncestors:")
	for _, v := range g.V {
		for _, v := range g.V {
			// Reset visited status
			v.SetVisited(false)
		}

		fmt.Println(printHeritage(v))
	}
}

func printDescendants(vertex vertex.Vertexer) string {
	line := fmt.Sprintf("(%s)\n", vertex.Label())
	vertex.SetVisited(true)

	for _, child := range vertex.Children() {
		if child.Visited() {
			continue
		}

		child.SetVisited(true)
		line = fmt.Sprintf("%s|- ancestor of -> %s", line, printDescendants(child))
	}

	return fmt.Sprintf("%s", line)
}

func printHeritage(vertex vertex.Vertexer) string {
	line := fmt.Sprintf("(%s)\n", vertex.Label())
	vertex.SetVisited(true)

	for _, parent := range vertex.Parents() {
		if parent.Visited() {
			continue
		}

		line = fmt.Sprintf("%s|- descendant of -> %s", line, printHeritage(parent))
	}

	return fmt.Sprintf("%s", line)
}
