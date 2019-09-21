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

// TODO: Document module
package path

import (
	// Internal Imports
	"github.com/matthewhartstonge/graph/edge"
)

const emptyIndex = -1

type Pather interface {
	Append(edge edge.Edger)
	Copy() Pather
	Cost() float64
	Current() edge.Edger
	Prev() edge.Edger
	Next() edge.Edger
	Last() edge.Edger
	Reset()
}

func New(opts ...Option) *Path {
	path := &Path{
		current: emptyIndex,
		path:    []edge.Edger{},
	}

	for _, opt := range opts {
		opt(path)
	}

	return path
}

type Option func(pather Pather)

func WithEdge(edger edge.Edger) Option {
	return func(pather Pather) {
		pather.Append(edger)
	}
}

type Path struct {
	cost    float64
	current int
	path    []edge.Edger
}

func (p *Path) Append(pathExtension edge.Edger) {
	p.cost = pathExtension.Cost()
	p.path = append(p.path, pathExtension)
}

func (p Path) Cost() float64 {
	return p.cost
}

// Copy deep copies the path.
func (p Path) Copy() Pather {
	return &Path{
		cost:    p.cost,
		current: p.current,
		path:    append([]edge.Edger{}, p.path...),
	}
}

func (p Path) Current() edge.Edger {
	if p.current == 0 || len(p.path) >= p.current {
		return nil
	}

	return p.path[p.current]
}

func (p *Path) Next() edge.Edger {
	p.current++
	lenPath := len(p.path)
	if p.current >= lenPath {
		p.current = lenPath
		return nil
	}

	return p.path[p.current]
}

func (p Path) Last() edge.Edger {
	pathLen := len(p.path)
	if pathLen == 0 {
		return nil
	}

	return p.path[pathLen-1]
}

func (p *Path) Prev() edge.Edger {
	p.current--
	if p.current <= emptyIndex {
		p.current = emptyIndex
		return nil
	}

	return p.path[p.current]
}

func (p *Path) Reset() {
	p.current = emptyIndex
}
