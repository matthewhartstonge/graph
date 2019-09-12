package csp

import (
	"fmt"
)

type Grapher interface {
	GetFrontier()
	GetPaths()
	AddEdge()
	AddNode()
}

// NewGraph creates a new graph.
func NewGraph(options ...GraphOption) *Graph {
	g := &Graph{
		V:        []Vertexer{},
		E:        []Edger{},
		digraph:  false,
		Strategy: nil,
		Goal:     nil,
	}

	for _, option := range options {
		option(g)
	}

	return g
}

// GraphOption provides variadic options when creating a new Graph.
type GraphOption func(g *Graph)

// GoalFunc provides the algorithm to check if this is the goal vertex.
type GoalFunc func(v Vertexer) bool

// TODO: use Grapher interface instead.
// TODO: document SearchStrategy.
type SearchStrategy func(V []Vertexer, E []Edger) *Vertex

// WithVertices provides a way to supply your own vertices when creating a new
// graph.
func WithVertices(vertices []Vertexer) GraphOption {
	return func(g *Graph) {
		g.V = vertices
	}
}

// WithEdges provides a way to supply your own edges when creating a new graph.
func WithEdges(edges []Edger) GraphOption {
	return func(g *Graph) {
		g.E = edges
	}
}

// WithSearchStrategy provides a way to inject a graph search algorithm.
func WithSearchStrategy(f SearchStrategy) GraphOption {
	return func(g *Graph) {
		g.Strategy = f
	}
}

// WithGoalFunc provides a way to supply your own algorithm in order to satisfy
// the graph search.
func WithGoalFunc(f GoalFunc) GraphOption {
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
	V []Vertexer
	// E contains a set of edges, also called links.
	E []Edger

	// Goal contains the algorithm to check if a given vertex satisfies the
	// goal state.
	Goal GoalFunc
	// Strategy provides the search strategy algorithm.
	Strategy SearchStrategy
}

func (g *Graph) preprocess() {
	// for _, e := range g.E {
	// 	if e.Undirected {
	// 		g.digraph = false
	// 		break
	// 	}
	// }
}

func (g Graph) PrintInfo() {
	// graphType := "undirected"
	// if g.digraph {
	// 	graphType = "directed"
	// }
	// fmt.Printf("Graph:\n- is a %s graph\n", graphType)

	// Print Links
	fmt.Println("Lineage:")
	for _, v := range g.V {
		fmt.Println(printDescendants(v))
	}

	fmt.Println("\nAncestors:")
	for _, v := range g.V {
		fmt.Println(printHeritage(v))
	}
}

func printDescendants(vertex Vertexer) string {
	line := ""
	for _, child := range vertex.Children() {
		line = fmt.Sprintf(" = parent of => %s", printDescendants(child))
	}
	return fmt.Sprintf("(%s)%s", vertex.Label, line)
}

func printHeritage(vertex Vertexer) string {
	line := ""
	for _, parent := range vertex.Parents() {
		line = fmt.Sprintf(" = child of => %s", printHeritage(parent))
	}
	return fmt.Sprintf("(%s)%s", vertex.Label, line)
}
