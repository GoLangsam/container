// package lsb provides a lazy string buffer: efficient, and NOT concurrency safe!
package lsb // LazyStringBuffer

type Friendly interface {
	Index(i int) byte // read a previously appended byte
	Append(c byte)    // extend by adding c at the end
	Unpend()          // undo last Append - Backtracking - track back
	String() string   // the final string - result
	Pos() int         // bytes written
}

// found in path/path.go - chapeaux to Rob Pike

// I would not mind to see this below "strings" :-)

// added:
//  methods Pos & Unpend
//  interface Friendly
// edited:
//  balanced if with } else { - for easier reading.
//  renamed: s => ori - for easier reading.
//  inverted Index: b.buf == nil first - for symmetry

// TODO: Do we need some panics re m? E.g. underflow upon Unpend.
// Or shall we leave things to subsequent panics upon illegal access to ori resp. buf?

// A LazyStringBuffer is a lazily constructed path buffer.
// It supports append, reading previously appended bytes,
// and retrieving the final string. It does not allocate a buffer
// to hold the output until that output diverges from s.
type LazyStringBuffer struct {
	ori string // original string
	buf []byte // buffer, if need
	w   int    // bytes written
}

func New(s string) *LazyStringBuffer {
	return &LazyStringBuffer{ori: s}
}

func (b *LazyStringBuffer) Index(i int) byte {
	if b.buf == nil {
		return b.ori[i]
	} else {
		return b.buf[i]
	}
}

func (b *LazyStringBuffer) Append(c byte) {
	if b.buf == nil {
		if b.w < len(b.ori) && b.ori[b.w] == c {
			b.w++
			return
		}
		b.buf = make([]byte, len(b.ori))
		copy(b.buf, b.ori[:b.w])
	}
	b.buf[b.w] = c
	b.w++
}

func (b *LazyStringBuffer) Unpend() {
	b.w--
}

func (b *LazyStringBuffer) String() string {
	if b.buf == nil {
		return b.ori[:b.w]
	} else {
		return string(b.buf[:b.w])
	}
}

func (b *LazyStringBuffer) Pos() int {
	return b.w
}
