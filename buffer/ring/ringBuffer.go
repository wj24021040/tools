package ring

import (
	"errors"
)

var (
	ParamErr = errors.New("the param is wrong")
	FullErr  = errors.New("the buf is full and enlarge failed")
	EmptyErr = errors.New("the buffer is empty")
)

type T = interface{}

type Ring struct {
	buffer []T
	cap    int
	len    int
	header int
	tail   int
}

func New(c int) (*Ring, error) {
	if c < 0 {
		return nil, ParamErr
	}

	return &Ring{
		buffer: make([]T, c),
		cap:    c,
	}, nil
}

func (r *Ring) Push(t T) error {
	if r.len == r.cap {
		r.autoscalar()
		if r.len == r.cap {
			return FullErr
		}
	}

	r.buffer[r.tail] = t
	r.len++
	r.tail++
	if r.tail >= r.cap {
		r.tail = 0
	}
	return nil
}

func (r *Ring) Pop() (t T, err error) {
	t = nil
	err = nil
	if r.len == 0 {
		err = EmptyErr
		return
	}

	t = r.buffer[r.header]
	r.len--
	r.header++
	if r.header >= r.cap {
		r.header = 0
	}
	return
}
func (r *Ring) autoscalar() {
	b := make([]T, 2*r.cap)
	copy(b[0:], r.buffer[r.header:])
	copy(b[(r.cap-r.header):], r.buffer[0:r.header])

	r.buffer = b
	r.cap = 2 * r.cap
	r.header = 0
	r.tail = r.len
}

func (r *Ring) Len() int {
	return r.len
}

func (r *Ring) Capacity() int {
	return r.cap
}
