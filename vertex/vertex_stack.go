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

package vertex

// NewStack returns a new vertex stack.
func NewStack() Stack {
	return Stack{
		len:   0,
		stack: []Vertexer{},
	}
}

// Stack provides a stack data structure for vertices.
type Stack struct {
	len   int
	stack []Vertexer
}

func (v Stack) Len() int {
	return v.len
}

func (v *Stack) Push(vertexer Vertexer) {
	v.stack = append(v.stack, vertexer)
	v.len++
}

func (v *Stack) Pop() (vertexer Vertexer) {
	if v.len > 0 {
		vertexer = v.stack[v.len-1]
		v.stack = v.stack[:v.len-1]
		v.len--
	}

	return
}
