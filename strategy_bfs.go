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

// NewBreadthFirstSearch returns a breadth-first search strategy.
func NewBreadthFirstSearch() *BFS {
	return &BFS{
		queue: []path.Pather{},
	}
}

// BFS provides a breadth-first search strategy. It always selects one of the
// earliest paths added to the frontier.
type BFS struct {
	queue []path.Pather
}

// Len returns the current number of paths stored in the queue.
func (b BFS) Len() int {
	return len(b.queue)
}

// Add enqueues a path to the end of the queue.
func (b *BFS) Add(path path.Pather) {
	b.queue = append(b.queue, path)
}

// Next dequeues a path from the front of the queue.
func (b *BFS) Next() (path path.Pather) {
	if b.Len() > 0 {
		path = b.queue[0]
		b.queue = b.queue[1:]
	}

	return path
}
