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

package utils

import (
	// Standard Library Imports
	"fmt"

	// Internal Imports
	"github.com/matthewhartstonge/graph"
	"github.com/matthewhartstonge/graph/path"
)

// PrintSolutions given a graph, will solve and print a solution if found.
// It will continue this until no more solutions can be found.
func PrintSolutions(grapher graph.Grapher) {
	solutionCount := 1
	for {
		goalPath := grapher.Search()
		if goalPath == nil {
			// no solutions to be found.
			printSolution(goalPath, solutionCount)
			break
		}

		printSolution(goalPath, solutionCount)
		solutionCount++
	}
}

// printSolution prints out a single solution.
func printSolution(goalPath path.Pather, solutionCount int) {
	if goalPath == nil {
		more := " more"
		if solutionCount == 0 {
			more = ""
		}

		fmt.Printf("No%s solutions found!\n\n", more)
		return

	}

	fmt.Printf("Solution %d:\n", solutionCount)
	for {
		edge := goalPath.Next()
		if edge == nil {
			break
		}

		fmt.Printf("- %s\n", edge)
	}
	fmt.Println()
}
