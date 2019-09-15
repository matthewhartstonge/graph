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

package path

import (
	// Internal Imports
	"github.com/matthewhartstonge/graph/edge"
)

const emptyIndex = -1

func New() *Path {
	return &Path{
		current: emptyIndex,
		edges:   []edge.Edger{},
	}
}

type Pather interface {
	Append(edge edge.Edger)
	Cost() float64
	Next() edge.Edger
	Prev() edge.Edger
	Reset()
}

type Path struct {
	cost    float64
	current int
	edges   []edge.Edger
}

func (p *Path) Append(edge edge.Edger) {
	p.cost += edge.Cost()
	p.edges = append(p.edges, edge)
}

func (p Path) Cost() float64 {
	return p.cost
}

func (p *Path) Next() edge.Edger {
	p.current++
	lenEdges := len(p.edges)
	if p.current >= lenEdges {
		p.current = lenEdges
		return nil
	}

	return p.edges[p.current]
}

func (p *Path) Prev() edge.Edger {
	p.current--
	if p.current <= emptyIndex {
		p.current = emptyIndex
		return nil
	}

	return p.edges[p.current]
}

func (p *Path) Reset() {
	p.current = emptyIndex
}
