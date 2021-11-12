package input

import (
	"errors"
	"io"

	piter "github.com/HAo99/orion-idl/common/peekiterator"
)

type readerSource struct {
	reader io.Reader
	buffer []rune
	cur    int
	end    bool
}

const (
	_BUFFER_SIZE = 512
)

func NewReaderSource(in io.Reader) piter.Source {
	return &readerSource{
		buffer: make([]rune, 0),
		cur:    0,
		end:    false,
		reader: in,
	}
}

func (x *readerSource) Read() interface{} {
	if x.cur >= len(x.buffer) {
		x.readNew()
	}
	ret := x.buffer[x.cur]
	x.cur++
	return ret
}

func (x *readerSource) readNew() {
	bs := make([]byte, _BUFFER_SIZE)
	n, err := x.reader.Read(bs)
	if err != nil {
		if errors.Is(io.EOF, err) {
			x.end = true
			return
		}
		panic(err)
	}
	if n == 0 {
		x.end = true
		return
	}
	x.buffer = []rune(string(bs[:n]))
	x.cur = 0
}

func (x *readerSource) Empty() bool {
	if x.cur >= len(x.buffer) {
		if x.end {
			return true
		}
		x.readNew()
		return x.Empty()
	}
	return false
}
