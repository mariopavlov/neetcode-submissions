
// MinStack is a stack that supports retrieving the minimum element in O(1) time.
type MinStack struct {
	store []int
	min   int
}

// NewMinStack initializes a new MinStack.
func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt,
	}
}

// Push pushes val onto the stack.
func (s *MinStack) Push(val int) {
	if val < s.min {
		s.min = val
	}

	s.store = append(s.store, val)
}

// Pop removes the element on top of the stack.
func (s *MinStack) Pop() {
	top := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]

	if top == s.min {
		s.min = math.MaxInt
		for _, cur := range s.store {
			if cur < s.min {
				s.min = cur
			}
		}
	}
}

// Top returns the top element of the stack.
func (s *MinStack) Top() int {
	elem := s.store[len(s.store)-1]

	return elem
}

// GetMin returns the minimum element in the stack.
func (s *MinStack) GetMin() int {
	return s.min
}
