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
)

func init() {
	// Log actions so you can see what is happening internally.
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	meta := utils.LoadJsonGraph("_examples/mailbot/graph.json")

	start := time.Now()
	V, E := utils.JSONGraphToVE(meta)
	G := graph.New(
		graph.WithVertices(V),
		graph.WithEdges(E),
	)
	log.WithField("took", time.Since(start)).Info()

	G.PrintInfo()
}
