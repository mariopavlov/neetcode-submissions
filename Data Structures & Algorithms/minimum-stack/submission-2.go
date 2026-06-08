// MinStack is a stack that supports retrieving the minimum element in O(1) time.
type MinStack struct {
	store    []int
	minStack []int
	size     int
}

// NewMinStack initializes a new MinStack.
func Constructor() MinStack {
	return MinStack{}
}

// Push pushes val onto the stack.
func (s *MinStack) Push(val int) {
	if val < s.GetMin() {
		s.minStack = append(s.minStack, val)
	} else {
		s.minStack = append(s.minStack, s.GetMin())
	}

	s.store = append(s.store, val)
	s.size = s.size + 1
}

// Pop removes the element on top of the stack.
func (s *MinStack) Pop() {
	if s.size == 0 {
		return
	}

	s.store = s.store[:len(s.store)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
	s.size = s.size - 1
}

// Top returns the top element of the stack.
func (s *MinStack) Top() int {
	if s.size == 0 {
		return 0
	}

	elem := s.store[len(s.store)-1]

	return elem
}

// GetMin returns the minimum element in the stack.
func (s *MinStack) GetMin() int {
	if s.size == 0 {
		return math.MaxInt
	}

	return s.minStack[len(s.minStack)-1]
}