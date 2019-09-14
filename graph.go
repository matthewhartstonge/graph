package graph

import (
	// Standard Library Imports
	"fmt"

	// Internal Imports
	"github.com/matthewhartstonge/graph/edge"
	"github.com/matthewhartstonge/graph/vertex"
)

// TODO: what does the graph say?
type Grapher interface {
	GetFrontier()
	GetPaths()
	AddEdge()
	AddNode()
	Solve()
}

// New creates a new graph.
func New(options ...Option) *Graph {
	g := &Graph{
		V:        []vertex.Vertexer{},
		E:        []edge.Edger{},
		digraph:  false,
		Strategy: nil,
		Goal:     nil,
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

// WithSearchStrategy provides a way to inject a graph search algorithm.
func WithSearchStrategy(f SearchStrategy) Option {
	return func(g *Graph) {
		g.Strategy = f
	}
}

// WithGoalFunc provides a way to supply your own algorithm in order to satisfy
// the graph search.
func WithGoalFunc(f GoalFunc) Option {
	return func(g *Graph) {
		g.Goal = f
	}
}

// Graph provides the data structure for a Graph.
// G = (V, E)
type Graph struct {
	// digraph provides a check to see if the graph is a directed or
	// undirected.
	digraph bool

	// V contains a set of vertices, also called nodes.
	V []vertex.Vertexer
	// E contains a set of edges, also called links.
	E []edge.Edger

	// Goal contains the algorithm to check if a given vertex satisfies the
	// goal state.
	Goal GoalFunc
	// Strategy provides the search strategy algorithm.
	Strategy SearchStrategy
}

// GoalFunc provides the algorithm to check if this is the goal vertex.
type GoalFunc func(v vertex.Vertexer) bool

// TODO: use Grapher interface instead.
// TODO: document SearchStrategy.
type SearchStrategy func(V []vertex.Vertexer, E []edge.Edger) vertex.Vertexer

func (g *Graph) preprocess() *Graph {
	g.digraph = true
	for _, e := range g.E {
		// Detect if a directed graph.
		if !e.Directed() {
			g.digraph = false
			break
		}
	}

	return g
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
