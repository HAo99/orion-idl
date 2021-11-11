package peekiterator

import "testing"

type arraySource struct {
	cur    int
	buffer []int
}

func (s *arraySource) Read() interface{} {
	ret := s.buffer[s.cur]
	s.cur++
	return ret
}

func (s *arraySource) Empty() bool { return s.cur == len(s.buffer) }

func TestPeekIterator(t *testing.T) {
	iter := New(&arraySource{cur: 0, buffer: []int{1, 2, 3, 4, 5}})
	if iter.HasNext() != true {
		t.Error("iter.HasNext() expect true; but got false")
	}
	if e := iter.Peek(); e != 1 {
		t.Errorf("iter.Peek() expect 1; but got %v", e)
	}
	if e := iter.Next(); e != 1 {
		t.Errorf("iter.Next() expect 1; but got %v", e)
	}
	if e := iter.Next(); e != 2 {
		t.Errorf("iter.Next() expect 2; but got %v", e)
	}
	iter.PutBack()
	if e := iter.Next(); e != 2 {
		t.Errorf("iter.Next() expect 2; but got %v", e)
	}
}
