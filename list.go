package list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

// Create a new list from a slice of values.
func FromSlice[T any](values []T) (l List[T]) {
	for _, v := range values {
		l.Append(v)
	}
	return
}

// Create a new list from a single value.
func FromSingle[T any](value T) List[T] {
	node := Node[T]{Value: value}
	return List[T]{Head: &node, Tail: &node}
}

func (l *List[T]) ToSlice() (values []T) {
	for node := l.Head; node != nil; node = node.Next {
		values = append(values, node.Value)
	}
	return
}

// Returns true if the list consists of exactly zero nodes.
func (l *List[T]) IsEmpty() bool {
	return l.Head == nil
}

// Iterates over the list and returns the node count.
func (l *List[T]) Len() int {
	var n int
	for node := l.Head; node != nil; node = node.Next {
		n++
	}
	return n
}

// Erase the contents of the list, making it empty.
func (l *List[T]) Erase() {
	l.Head = nil
	l.Tail = nil
}

// Iterate over the list, yielding each value.
func (l *List[T]) Range() func(func(T) bool) {
	return func(yield func(T) bool) {
		for node := l.Head; node != nil; node = node.Next {
			if !yield(node.Value) {
				return
			}
		}
	}
}

// Iterate over the list, yielding each value and its index.
func (l *List[T]) Range2() func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		var i int
		for node := l.Head; node != nil; node = node.Next {
			if !yield(i, node.Value) {
				return
			}
			i++
		}
	}
}

// Append a value to the end of the list.
func (l *List[T]) Append(value T) {
	node := &Node[T]{Value: value}
	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

// Extend the list with the contents of another list.
// Erases the second list, making it empty.
func (l *List[T]) Extend(l2 *List[T]) {
	if l2 == nil || l2.IsEmpty() {
		return
	}

	if l.IsEmpty() {
		l.Head = l2.Head
		l.Tail = l2.Tail
	} else {
		l.Tail.Next = l2.Head
		l.Tail = l2.Tail
	}

	l2.Erase()
}
