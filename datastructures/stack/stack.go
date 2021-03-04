package stack

// Stack is like an array but with limited functionality.
//
// You can only push to add a new element to the top of the stack,
// pop to remove the element from the top,
// and peek at the top element without popping it off.
type Stack []interface{}

// IsEmpty will check if the stack has no elements
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new element onto the top of the stack
func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

// Pop will return the top element, and remove it from the stack
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true
}

// Top simply peeks at the top element, but does not pop it
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	return element
}
