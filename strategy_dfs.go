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
	// Internal Imports
	"github.com/matthewhartstonge/graph/path"
)

// NewDepthFirstSearch returns a depth-first search strategy.
func NewDepthFirstSearch() *DFS {
	return &DFS{
		stack: []path.Pather{},
	}
}

// DFS provides a depth-first search strategy. It always selects the last
// path added to the frontier.
type DFS struct {
	stack []path.Pather
}

// Len returns the current number of paths stored in the stack.
func (d DFS) Len() int {
	return len(d.stack)
}

// Add pushes a path on to the top of the stack.
func (d *DFS) Add(path path.Pather) {
	d.stack = append(d.stack, path)
}

// Next pops the path that is sitting on top of the stack.
func (d *DFS) Next() (path path.Pather) {
	lenPaths := d.Len()
	if lenPaths > 0 {
		path = d.stack[lenPaths-1]
		d.stack = d.stack[:lenPaths-1]
	}

	return path
}

var _ Strategizer = &DFS{}
