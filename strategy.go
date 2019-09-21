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

// Strategizer provides a search strategy.
// A search strategy defines the way in which the underlying frontier is
// expanded and traversed.
type Strategizer interface {
	// Len returns the current number of paths stored in the frontier.
	Len() int
	// Next returns the next path in the frontier to process, or nil if there
	// are no more paths to expand.
	Next() path.Pather
	// Add stores an expanded path into the frontier.
	Add(path path.Pather)
}
