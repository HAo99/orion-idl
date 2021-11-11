package peekiterator

import (
	"errors"

	"github.com/HAo99/orion-idl/common/linkedlist"
)

// PeekIterator defines a type of iterator that can be put back and peek.
type PeekIterator interface {
	HasNext() bool
	Next() interface{}
	Peek() interface{}
	PutBack()
}

// Source is the source for PeekIterator to fetch data.
type Source interface {
	Read() interface{}
	Empty() bool
}

type peekIterator struct {
	src          Source
	putBackStack linkedlist.LinkedList
	cacheQueue   linkedlist.LinkedList
}

const (
	CACHE_SIZE = 10
)

var (
	ErrSourceEmpty = errors.New("PeekIterator: source is empty")
)

// New contructs a PerrkIterator object.
func New(src Source) PeekIterator {
	return &peekIterator{
		src:          src,
		putBackStack: linkedlist.New(),
		cacheQueue:   linkedlist.New(),
	}
}

// HasNext returns true if the source of iterator not empty.
func (x *peekIterator) HasNext() bool {
	return !x.putBackStack.Empty() || !x.src.Empty()
}

// Next returns the next element and move forward.
func (x *peekIterator) Next() (ret interface{}) {
	if !x.putBackStack.Empty() {
		ret = x.putBackStack.PopBack()
	} else {
		if x.src.Empty() {
			panic(ErrSourceEmpty)
		}
		ret = x.src.Read()
	}
	for x.cacheQueue.Len() > CACHE_SIZE {
		x.cacheQueue.PopFront()
	}
	x.cacheQueue.PushBack(ret)
	return
}

// Peek returns the next element but will not move forward.
func (x *peekIterator) Peek() interface{} {
	if !x.putBackStack.Empty() {
		return x.putBackStack.Back()
	}
	ret := x.Next()
	x.PutBack()
	return ret
}

// PutBack moves back a element so that Next can get it again.
func (x *peekIterator) PutBack() {
	if !x.cacheQueue.Empty() {
		x.putBackStack.PushBack(x.cacheQueue.PopBack())
	}
}
