package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// LinkedList defines a couple of linked list operations.
// When updating or popping a node, if the node doesn't exist.
// Such as index is out of range, or popping a node when the list is empty.
// These methods will panic rather than return an error.
// So please check the Empty and Len before you doing this stuff.
type LinkedList interface {
	Front() interface{}
	Back() interface{}

	PushBack(e interface{})
	PushFront(e interface{})

	PopBack() interface{}
	PopFront() interface{}

	Empty() bool
	Len() int
}

const _INT_MAX = int(^uint(0) >> 1)

var (
	ErrListEmpty = errors.New("LinkedList: cannot get front element when list is empty")
	ErrListFull  = errors.New("LinkedList: the list is full, cannot emplace any element")
)

type linkedList struct {
	len        int
	head, tail *node
}

type node struct {
	val        interface{}
	next, prev *node
}

// New constructs an empty LinkedList object.
func New() LinkedList {
	vNode := &node{}
	return &linkedList{
		len:  0,
		head: vNode,
		tail: vNode,
	}
}

// Front returns the first element of the LinkedList.
func (x *linkedList) Front() interface{} {
	if x.Empty() {
		panic(ErrListEmpty)
	}
	return x.head.next.val
}

// Back returns the last element of the LinkedList.
func (x *linkedList) Back() interface{} {
	if x.Empty() {
		panic(ErrListEmpty)
	}
	return x.tail.val
}

// checkEmpty panic if list is empty.
func (x *linkedList) checkEmpty() {
	if x.Empty() {
		panic(ErrListEmpty)
	}
}

// checkFull panic if list if full.
func (x *linkedList) checkFull() {
	if x.Len() == _INT_MAX {
		panic(ErrListFull)
	}
}

// PushBack links a element to the end of the list. It's an O(1) operation.
func (x *linkedList) PushBack(e interface{}) {
	x.checkFull()
	x.tail.next = &node{val: e, prev: x.tail}
	x.tail = x.tail.next
	x.len++
}

// PushFront links a element to the front of the list. It's an O(1) operation.
func (x *linkedList) PushFront(e interface{}) {
	x.checkFull()
	ret := x.head.next
	x.head.next = &node{val: e, next: ret, prev: x.head}
	ret.prev = x.head.next
	x.len++
}

// PopFront removes the first element of the list and returns its value. It's an O(1) operation.
func (x *linkedList) PopFront() (e interface{}) {
	x.checkEmpty()
	e = x.head.next.val
	x.head.next = x.head.next.next
	x.len--
	if !x.Empty() {
		x.head.next.prev = x.head
	}
	return
}

// PopBack removes the last element of the list and returns its value. It's an O(1) operation.
func (x *linkedList) PopBack() (e interface{}) {
	x.checkEmpty()
	e = x.tail.val
	x.tail = x.tail.prev
	x.tail.next = nil
	x.len--
	return
}

// Len returns the length of the LinkedList. It's an O(1) operation.
func (x *linkedList) Len() int {
	return x.len
}

// Empty returns true if list does not contain any elements. It's an O(1) operation.
func (x *linkedList) Empty() bool {
	return x.Len() == 0
}

func (x *linkedList) String() string {
	var (
		first = true
		cur   = x.head
		sb    = strings.Builder{}
	)
	sb.WriteByte('[')
	for cur.next != nil {
		cur = cur.next
		if first {
			sb.WriteString(fmt.Sprintf("%v", cur.val))
			first = false
			continue
		}
		sb.WriteString(fmt.Sprintf(" -> %v", cur.val))
	}
	sb.WriteByte(']')
	return sb.String()
}
