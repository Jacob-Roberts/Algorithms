package stack

import "testing"

func TestEmpty(t *testing.T) {
	var stack Stack

	if !stack.IsEmpty() {
		t.Fatalf("Stack should have been empty, but it wasn't")
	}
	if len(stack) != 0 {
		t.Fatalf("Stack should have been empty, but it wasn't")
	}
	if stack.Top() != nil {
		t.Fatalf("Stack should have been empty, but it wasn't")
	}
	if val, succeeded := stack.Pop(); val != nil || succeeded {
		t.Fatalf("Stack should have been empty, but it wasn't")
	}
}

func TestOneElement(t *testing.T) {
	var stack Stack

	stack.Push(123)

	if stack.IsEmpty() {
		t.Fatalf("Stack should not have been empty, but it was")
	}
	if len(stack) != 1 {
		t.Fatalf("Stack should not have been empty, but it was")
	}
	if stack.Top() != 123 {
		t.Fatalf("Stack should not have been empty, but it was")
	}

	result, ok := stack.Pop()
	if !ok {
		t.Fatalf("Stack failed while popping a value")
	}
	if result != 123 {
		t.Fatalf("Stack popped the wrong value")
	}
	if !stack.IsEmpty() {
		t.Fatalf("Stack didn't decrement after popping")
	}
	if len(stack) != 0 {
		t.Fatalf("Stack didn't decrement after popping")
	}
	if stack.Top() != nil {
		t.Fatalf("Stack wasn't empty")
	}

	result, ok = stack.Pop()
	if ok {
		t.Fatalf("Stack should fail popping when empty")
	}
	if result != nil {
		t.Fatalf("Stack didn't return nil when it was empty")
	}
}

func TestTwoElement(t *testing.T) {
	var stack Stack

	stack.Push(123)

	if stack.IsEmpty() {
		t.Fatalf("Stack should not have been empty, but it was")
	}
	if len(stack) != 1 {
		t.Fatalf("Stack should not have been empty, but it was")
	}
	if stack.Top() != 123 {
		t.Fatalf("Stack should not have been empty, but it was")
	}

	result, ok := stack.Pop()
	if !ok {
		t.Fatalf("Stack failed while popping a value")
	}
	if result != 123 {
		t.Fatalf("Stack popped the wrong value")
	}
	if !stack.IsEmpty() {
		t.Fatalf("Stack didn't decrement after popping")
	}
	if len(stack) != 0 {
		t.Fatalf("Stack didn't decrement after popping")
	}
	if stack.Top() != nil {
		t.Fatalf("Stack wasn't empty")
	}

	result, ok = stack.Pop()
	if ok {
		t.Fatalf("Stack should fail popping when empty")
	}
	if result != nil {
		t.Fatalf("Stack didn't return nil when it was empty")
	}
}
