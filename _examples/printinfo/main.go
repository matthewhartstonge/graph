package main

import (
	// Standard Library Imports
	"fmt"
	"time"

	// Internal Imports
	"github.com/matthewhartstonge/graph"
	"github.com/matthewhartstonge/graph/edge"
	"github.com/matthewhartstonge/graph/vertex"
)

func solvent(v vertex.Vertexer) bool {
	if v.Label() == "v4" {
		return true
	}

	return false
}

func main() {
	v1 := vertex.New("v1")
	v2 := vertex.New("v2")
	v3 := vertex.New("v3")
	v4 := vertex.New("v4")

	V := []vertex.Vertexer{
		v1,
		v2,
		v3,
		v4,
	}

	E := []edge.Edger{
		edge.New(
			v1, v2,
			edge.WithDirected(true),
			edge.WithCost(1),
		),
		edge.New(
			v1, v3,
			edge.WithDirected(true),
			edge.WithCost(1.2),
		),
		edge.New(
			v2, v3,
			edge.WithDirected(true),
			edge.WithCost(1)),
		edge.New(
			v3, v4,
			edge.WithCost(4.8),
		),
	}

	start := time.Now()
	G := graph.New(
		graph.WithVertices(V),
		graph.WithEdges(E),
		graph.WithGoalFunc(solvent),
	)

	G.PrintInfo()
	fmt.Printf("took: %s", time.Since(start))
}
