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

// NewQueue returns a new vertex queue.
func NewQueue() Queue {
	return Queue{
		len:   0,
		queue: []Vertexer{},
	}
}

// Queue provides a queue data structure for vertices.
type Queue struct {
	len   int
	queue []Vertexer
}

func (v Queue) Len() int {
	return v.len
}

func (v *Queue) Enqueue(vertexer Vertexer) {
	v.queue = append(v.queue, vertexer)
	v.len++
}

func (v *Queue) Dequeue() (vertexer Vertexer) {
	if v.len > 0 {
		vertexer = v.queue[0]
		v.queue = v.queue[1:]
		v.len--
	}

	return
}
