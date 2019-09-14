package csp

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
