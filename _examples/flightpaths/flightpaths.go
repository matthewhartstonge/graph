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

package main

import (
	// Standard Library Imports
	"fmt"

	// Internal Imports
	"github.com/matthewhartstonge/graph"
	"github.com/matthewhartstonge/graph/_examples/utils"
	"github.com/matthewhartstonge/graph/edge"
	"github.com/matthewhartstonge/graph/goal"
	"github.com/matthewhartstonge/graph/vertex"
)

func solveFlightPath(searchStrategy graph.Strategizer) graph.Grapher {
	christchurch := vertex.New("Christchurch")
	auckland := vertex.New("Auckland")
	wellington := vertex.New("Wellington")
	goldCoast := vertex.New("Gold Coast")

	vertices := []vertex.Vertexer{
		christchurch,
		auckland,
		wellington,
		goldCoast,
	}

	edges := []edge.Edger{
		edge.New(christchurch, goldCoast),
		edge.New(christchurch, auckland),
		edge.New(christchurch, wellington),
		edge.New(wellington, goldCoast),
		edge.New(wellington, auckland),
		edge.New(auckland, goldCoast),
	}

	return graph.New(
		// graph.WithTraceLogging(),
		graph.WithVertices(vertices),
		graph.WithEdges(edges),
		graph.WithStartingVertices(christchurch),
		graph.WithSearchStrategy(searchStrategy),
		graph.WithGoalFunc(goal.VertexLabelEquals(goldCoast.Label())),
	)
}

func main() {
	fmt.Println("Depth-First Solution:")
	dfs := graph.NewDepthFirstSearch()
	utils.PrintSolutions(solveFlightPath(dfs))

	fmt.Println("Breadth-First Solution:")
	bfs := graph.NewBreadthFirstSearch()
	utils.PrintSolutions(solveFlightPath(bfs))
}
