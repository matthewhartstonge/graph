package csp

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
