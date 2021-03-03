package stack

// Stack is like an array but with limited functionality.
//
// You can only push to add a new element to the top of the stack,
// pop to remove the element from the top,
// and peek at the top element without popping it off.
type Stack struct {
	array        []interface{}
	currentIndex int
	size         int
}

// New creates a new stack
func New() *Stack {
	s := &Stack{}
	s.size = 10
	s.array = make([]interface{}, s.size)
	s.currentIndex = 0

	return s
}

// Push a new element onto the top of the stack
func (s *Stack) Push(element interface{}) {
	// Check if we need to allocate more space
	s.array[s.currentIndex] = element
	s.currentIndex++
}

func append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}
