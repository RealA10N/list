package list_test

import (
	"testing"

	"alon.kr/x/list"
)

func TestFromSlice(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	l := list.FromSlice(values)

	if l.Len() != len(values) {
		t.Errorf("expected length %d, got %d", len(values), l.Len())
	}

	i := 0
	for v := range l.Range() {
		if v != values[i] {
			t.Errorf("expected value %d, got %d", values[i], v)
		}
		i++
	}
}

func TestIsEmpty(t *testing.T) {
	l := list.List[int]{}
	if !l.IsEmpty() {
		t.Error("expected list to be empty")
	}

	l.Append(1)
	if l.IsEmpty() {
		t.Error("expected list to not be empty")
	}
}

func TestLen(t *testing.T) {
	l := list.List[int]{}
	if l.Len() != 0 {
		t.Errorf("expected length 0, got %d", l.Len())
	}

	l.Append(1)
	if l.Len() != 1 {
		t.Errorf("expected length 1, got %d", l.Len())
	}
}

func TestAppend(t *testing.T) {
	l := list.List[int]{}
	l.Append(1)
	if l.Head.Value != 1 {
		t.Errorf("expected head value 1, got %d", l.Head.Value)
	}
	if l.Tail.Value != 1 {
		t.Errorf("expected tail value 1, got %d", l.Tail.Value)
	}

	l.Append(2)
	if l.Head.Next.Value != 2 {
		t.Errorf("expected second value 2, got %d", l.Head.Next.Value)
	}
	if l.Tail.Value != 2 {
		t.Errorf("expected tail value 2, got %d", l.Tail.Value)
	}
}

func TestExtend(t *testing.T) {
	l1 := list.FromSlice([]int{1, 2, 3})
	l2 := list.FromSlice([]int{4, 5, 6})

	l1.Extend(&l2)
	if l1.Len() != 6 {
		t.Errorf("expected length 6, got %d", l1.Len())
	}

	expectedValues := []int{1, 2, 3, 4, 5, 6}
	i := 0
	for v := range l1.Range() {
		if v != expectedValues[i] {
			t.Errorf("expected value %d, got %d", expectedValues[i], v)
		}
		i++
	}
}
func TestRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	l := list.FromSlice(slice)
	i := 0
	for v := range l.Range() {
		if v != slice[i] {
			t.Errorf("expected value %d, got %d", slice[i], v)
		}
		i++
	}
}

func TestRange2(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	l := list.FromSlice(slice)
	i := 0
	for idx, v := range l.Range2() {
		if idx != i {
			t.Errorf("expected index %d, got %d", i, idx)
		}
		if v != slice[i] {
			t.Errorf("expected value %d, got %d", slice[i], v)
		}
		i++
	}
}
