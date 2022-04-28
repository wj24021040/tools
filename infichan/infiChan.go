package infiChan

import (
	"sync/atomic"

	"github.com/wj24021040/tools/buffer/ring"
)

type T = interface{}

type InfiChan struct {
	I   chan<- T
	O   <-chan T
	buf *ring.Ring
	len int32
}

func New(initCap int) *InfiChan {
	b, err := ring.New(initCap)
	if err != nil {
		panic("create infichan failed: " + err.Error())
		return nil
	}
	in := make(chan T)
	out := make(chan T)
	ic := &InfiChan{
		I:   in,
		O:   out,
		buf: b,
		len: 0,
	}
	go ic.synchronous(in, out)

	return ic
}

func (ic *InfiChan) synchronous(in, out chan T) {
	defer func() {
		close(out)
	}()
	var ov T = nil
	for {
	LOOP:
		if ov != nil {
			for {
				var err error
				select {
				case v, ok := <-in:
					if !ok { //close
						goto clean
					}
					atomic.AddInt32(&(ic.len), 1)
					err = ic.buf.Push(v)
					if err != nil {
						close(in)
					}
				case out <- ov:
					atomic.AddInt32(&(ic.len), -1)
					ov, err = ic.buf.Pop()
					if err != nil {
						goto LOOP
					}
				}
			}
		} else {
			v, ok := <-in
			if !ok { //close
				goto clean
			}
			atomic.AddInt32(&(ic.len), 1)
			ov = v
		}
	}
clean:
	if ov != nil {
		out <- ov
		atomic.AddInt32(&(ic.len), -1)
	}
	for ic.buf.Len() > 0 {
		ov, _ := ic.buf.Pop()
		out <- ov
		atomic.AddInt32(&(ic.len), -1)
	}
	return
}

func (ic *InfiChan) Close() {
	close(ic.I)
}

func (ic *InfiChan) Len() int {
	return int(atomic.LoadInt32(&(ic.len)))
}
