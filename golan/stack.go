package main

import "fmt"

type StackEmptyError struct{}

func (e *StackEmptyError) Error() string {
	return "Exception: empty stack"
}

type Node struct {
	Value int
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Push(v int) *Stack {
	s.size += 1
	if s.head == nil {
		s.head = &Node{Value: v}
	} else {
		node := Node{v, s.head}
		s.head = &node
	}
	return s
}

func (s *Stack) Peak() (int, error) {
	if s.head == nil {
		return -1, &StackEmptyError{}
	}
	return s.head.Value, nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Empty() bool {
	return s.head == nil
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return -1, &StackEmptyError{}
	}
	s.size += 1
	v := s.head.Value
	s.head = s.head.next

	return v, nil
}

func (n Node) String() string {
	return fmt.Sprint(n.Value)
}

func build_string(s string, v int) string {
	return fmt.Sprint(v) + " " + s
}

func (s Stack) String() string {
	repr := traverse(&s, build_string, "")

	return "{ " + repr + "}"
}

func traverse_node[T any](n *Node, f func(accum T, v int) T, init T) T {
	if n.next == nil {
		return f(init, n.Value)
	}
	return f(traverse_node(n.next, f, init), n.Value)
}

func traverse[T any](s *Stack, f func(accum T, v int) T, init T) T {
	if (s.Empty()) {
        return init
    }
     return traverse_node (s.head, f, init)
}
