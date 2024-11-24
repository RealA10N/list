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

func (l *List[T]) Len() uint {
	var n uint
	for node := l.Head; node != nil; node = node.Next {
		n++
	}
	return n
}

func (l *List[T]) Range(f func(T)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node.Value)
	}
}

func (l *List[T]) Range2(f func(uint, T)) {
	var i uint
	for node := l.Head; node != nil; node = node.Next {
		f(i, node.Value)
		i++
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
