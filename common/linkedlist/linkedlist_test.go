package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := New()
	t.Log(ll)
	if ll.Empty() != true {
		t.Error("ll.empty expect: true; but got false")
	}
	ll.PushBack(4)
	ll.PushBack(5)
	ll.PushFront(2)
	ll.PushFront(1)
	t.Log(ll)
	if e := ll.Front(); e != 1 {
		t.Errorf("ll.PopBack() expect: 1; but got %d", e)
	}
	if e := ll.Back(); e != 5 {
		t.Errorf("ll.PopBack() expect: 1; but got %d", e)
	}
	if e := ll.PopBack(); e != 5 {
		t.Errorf("ll.PopBack() expect: 5; but got %d", e)
	}
	t.Log(ll)
	if e := ll.PopFront(); e != 1 {
		t.Errorf("ll.PopFront() expect: 1; but got %d", e)
	}
	t.Log(ll)
	if e := ll.PopBack(); e != 4 {
		t.Errorf("ll.PopBack() expect: 4; but got %d", e)
	}
	t.Log(ll)
	if e := ll.PopFront(); e != 2 {
		t.Errorf("ll.PopFront() expect: 2; but got %d", e)
	}
	t.Log(ll)
}
