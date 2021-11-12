package input

import piter "github.com/HAo99/orion-idl/common/peekiterator"

type stringSource struct {
	data []rune
	cur  int
}

func NewStringSource(s string) piter.Source {
	return &stringSource{
		data: []rune(s),
		cur:  0,
	}
}

func (s *stringSource) Read() interface{} {
	ret := s.data[s.cur]
	s.cur++
	return ret
}

func (s *stringSource) Empty() bool {
	return len(s.data) <= s.cur
}
