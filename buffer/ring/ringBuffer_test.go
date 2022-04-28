package ring

import (
	"testing"
)

func TestRingBuffer1(t *testing.T) {
	b, _ := New(1)
	b.Push(2)
	v, _ := b.Pop()
	if v != 2 {
		t.Error("expected get 2,but ", v)
	}

	b.Push(2)
	b.Push(3)
	b.Push(4)
	b.Push(5)
	b.Push(6)
	b.Push(7)
	v, _ = b.Pop()
	if v != 2 {
		t.Error("expected get 2,but ", v)
	}
	v, _ = b.Pop()
	if v != 3 {
		t.Error("expected get 3,but ", v)
	}
	v = b.Len()
	if v != 4 {
		t.Error("expected get 4,but ", 0)
	}
	b.Push(8)
	v, _ = b.Pop()
	if v != 4 {
		t.Error("expected get 5,but ", v)
	}
	v, _ = b.Pop()
	if v != 5 {
		t.Error("expected get 6,but ", v)
	}
	v, _ = b.Pop()
	if v != 6 {
		t.Error("expected get 7,but ", v)
	}
	v, _ = b.Pop()
	if v != 7 {
		t.Error("expected get 7,but ", v)
	}
	v, _ = b.Pop()
	if v != 8 {
		t.Error("expected get 7,but ", v)
	}
	v = b.Len()
	if v != 0 {
		t.Error("expected get 4,but ", v)
	}

	_, err := b.Pop()
	if err == nil {
		t.Error("expected get the empty err, but not ")
	}
}

type Tt struct {
	A int
}

func TestRingBuffer2(t *testing.T) {

	t1 := Tt{A: 1}
	t2 := Tt{A: 2}
	t3 := Tt{A: 3}
	t4 := Tt{A: 4}
	t5 := Tt{A: 5}
	t6 := Tt{A: 6}
	t7 := Tt{A: 7}
	t8 := Tt{A: 8}

	b, _ := New(2)

	b.Push(t1)
	v, _ := b.Pop()
	//vl := v.(Tt)
	if v.(Tt).A != 1 {
		t.Error("expected get 1,but ", v.(Tt).A)
	}

	v, err := b.Pop()
	if err != EmptyErr {
		t.Error("expected get EmptyErr,but ", err.Error())
	}

	b.Push(t2)
	b.Push(t3)
	b.Push(t4)
	b.Push(t5)
	b.Push(t6)
	b.Push(t7)

	l := b.Capacity()
	if l != 8 {
		t.Error("expected get 8,but ", l)
	}
	v, _ = b.Pop()
	if v.(Tt).A != 2 {
		t.Error("expected get 2,but ", v.(Tt).A)
	}
	v, _ = b.Pop()
	if v.(Tt).A != 3 {
		t.Error("expected get 3,but ", v.(Tt).A)
	}
	l = b.Len()
	if l != 4 {
		t.Error("expected get 4,but ", l)
	}
	b.Push(t8)
	v, _ = b.Pop()
	if v.(Tt).A != 4 {
		t.Error("expected get 5,but ", v.(Tt).A)
	}
	v, _ = b.Pop()
	if v.(Tt).A != 5 {
		t.Error("expected get 6,but ", v.(Tt).A)
	}
	v, _ = b.Pop()
	if v.(Tt).A != 6 {
		t.Error("expected get 7,but ", v.(Tt).A)
	}
	v, _ = b.Pop()
	if v.(Tt).A != 7 {
		t.Error("expected get 7,but ", v.(Tt).A)
	}
	v, _ = b.Pop()
	if v.(Tt).A != 8 {
		t.Error("expected get 7,but ", v.(Tt).A)
	}
	l = b.Len()
	if l != 0 {
		t.Error("expected get 4,but ", l)
	}
}
