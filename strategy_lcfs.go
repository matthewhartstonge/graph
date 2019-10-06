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
	// Standard Library Imports
	"container/heap"

	// Internal Imports
	"github.com/matthewhartstonge/graph/path"
)

// NewLowestCostFirstSearch returns a lowest-cost-first search strategy.
func NewLowestCostFirstSearch() *LCFS {
	pQueue := &priorityQueue{}
	heap.Init(pQueue)

	return &LCFS{
		queue: pQueue,
	}
}

// LCFS provides a lowest-cost-first search strategy. It always selects a path
// from the frontier with teh lowest cost. It is a priority queue ordered by
// path cost. When path costs are equal, it performs a breadth-first search.
type LCFS struct {
	queue *priorityQueue
}

// Len returns the current number of paths stored in the queue.
func (l LCFS) Len() int {
	return len(*l.queue)
}

// Add enqueues a path to the end of the queue.
func (l *LCFS) Add(newPath path.Pather) {
	heap.Push(l.queue, newPath)
}

// Next dequeues a path from the front of the queue.
func (l *LCFS) Next() (nextPath path.Pather) {
	if l.Len() > 0 {
		return heap.Pop(l.queue).(path.Pather)
	}

	return
}

// A priorityQueue implements heap.Interface and provides a lowest-cost-first
// priority queue.
type priorityQueue []path.Pather

// Len implements sort.Interface.
func (pq priorityQueue) Len() int { return len(pq) }

// Less implements sort.Interface.
func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest priority so we use
	// less than here.
	return pq[i].Cost() < pq[j].Cost()
}

// Swap implements sort.Interface.
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push implements heap.Interface.
func (pq *priorityQueue) Push(x interface{}) {
	item := x.(path.Pather)
	*pq = append(*pq, item)
}

// Pop implements heap.Interface.
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
