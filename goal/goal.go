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

package goal

import (
	// Internal Imports
	"github.com/matthewhartstonge/graph"
	"github.com/matthewhartstonge/graph/vertex"
)

// VertexLabelEquals returns a solution based on matching the provided vertex
// label.
func VertexLabelEquals(label string) graph.GoalFunc {
	return func(vertex vertex.Vertexer) bool {
		return vertex.Label() == label
	}
}
