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
	"time"

	// Internal Imports
	"github.com/matthewhartstonge/graph"
	"github.com/matthewhartstonge/graph/edge"
	"github.com/matthewhartstonge/graph/goal"
	"github.com/matthewhartstonge/graph/vertex"
)

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
			edge.WithCost(1),
		),
		edge.New(
			v1, v3,
			edge.WithCost(1.2),
		),
		edge.New(
			v2, v3,
			edge.WithCost(1)),
		edge.New(
			v3, v4,
			edge.WithCost(4.8),
			edge.WithUndirected(),
		),
	}

	start := time.Now()
	G := graph.New(
		graph.WithVertices(V),
		graph.WithEdges(E),
		graph.WithGoalFunc(goal.VertexLabelEquals("v4")),
	)

	G.PrintInfo()
	fmt.Printf("took: %s", time.Since(start))
}
