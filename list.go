package list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l *List[T]) Len() int {
	var n int
	for node := l.Head; node != nil; node = node.Next {
		n++
	}
	return n
}

func (l *List[T]) Range() func(func(T) bool) {
	return func(yield func(T) bool) {
	for node := l.Head; node != nil; node = node.Next {
			if !yield(node.Value) {
				return
			}
		}
	}
}

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

func (l *List[T]) Extend(l2 *List[T]) {
	l.Tail.Next = l2.Head
	l.Tail = l2.Tail
}
