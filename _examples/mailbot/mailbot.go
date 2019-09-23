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
	"os"
	"time"

	// External Imports
	log "github.com/sirupsen/logrus"

	// Internal Imports
	"github.com/matthewhartstonge/graph"
	"github.com/matthewhartstonge/graph/_examples/utils"
	"github.com/matthewhartstonge/graph/goal"
)

func init() {
	// Log actions so you can see what is happening internally.
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	meta := utils.LoadJsonGraph("_examples/mailbot/graph.json")

	start := time.Now()

	// Convert the JSON representation of vertices and edges to our internal
	// types, ready to build a graph.
	V, E := utils.JSONGraphToVE(meta)

	// set starting vertex to room 'o103'
	startVertex := V[2]

	// Build a graph with a default DFS search strategy.
	G := graph.New(
		// graph.WithTraceLogging(),
		graph.WithVertices(V),
		graph.WithEdges(E),
		graph.WithStartingVertices(startVertex),
		graph.WithGoalFunc(goal.VertexLabelEquals("r123")),
	)

	utils.PrintSolutions(G)
	log.WithField("took", time.Since(start)).Info()
}
