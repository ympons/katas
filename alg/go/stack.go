package main

type Stack interface{
	Push(int)
	Pop() int
	IsEmpty() bool
}

type snode struct {
	value int
	next *snode
}

type stack_ struct {
	top *snode
	n int
}

func NewStack() *stack_ {
	return &stack_{}
}

func (s *stack_) Push(v int) {
	s.top = &snode{v, s.top}
	s.n++
}

func (s *stack_) Pop() (value int, ok bool) {
	if ok = s.n > 0; ok {
		value, s.top = s.top.value, s.top.next
		s.n--
	}
	return
}

func (s stack_) IsEmpty() bool {
	return s.n == 0
}
