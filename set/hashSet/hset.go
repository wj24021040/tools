package hashSet

import (
	"fmt"
	"strings"

	. "github.com/wj24021040/tools/set"
)

type Hset map[T]struct{}

func New() Hset {
	return make(Hset)
}

func (h Hset) Add(data T) {
	h[data] = struct{}{}
}
func (h Hset) Remove(data T) {
	if _, ok := h[data]; ok {
		delete(h, data)
	}
}
func (h Hset) Cap() int {
	return len(h)
}
func (h Hset) Equal(dst Set) bool {
	if h.Cap() != dst.Cap() {
		return false
	}

	for k := range h {
		if !dst.Contains(k) {
			return false
		}
	}
	return true
}
func (h Hset) Intersect(dst Set) Set {
	reslut := New()
	if h.Cap() < dst.Cap() {
		for k := range h {
			if dst.Contains(k) {
				reslut.Add(k)
			}
		}
	} else {
		do := func(k T) bool {
			if _, ok := h[k]; ok {
				reslut.Add(k)
			}
			return true
		}
		dst.Each(do)
	}

	return reslut
}
func (h Hset) Difference(dst Set) Set {
	reslut := New()
	for k := range h {
		if !dst.Contains(k) {
			reslut.Add(k)
		}
	}

	do := func(k T) bool {
		if _, ok := h[k]; !ok {
			reslut.Add(k)
		}
		return true
	}
	dst.Each(do)

	return reslut
}
func (h Hset) Clear() {
	for k := range h {
		delete(h, k)
	}
}
func (h Hset) Clone() Set {
	i2 := New()
	for k := range h {
		i2.Add(k)
	}
	return i2
}
func (h Hset) Contains(val ...T) bool {
	for _, v := range val {
		if _, ok := h[v]; !ok {
			return false
		}
	}
	return true
}
func (h Hset) Compare(dst Set) (more, less Set) {
	more = New()
	less = New()

	for k := range h {
		if !dst.Contains(k) {
			more.Add(k)
		}
	}

	do := func(k T) bool {
		if _, ok := h[k]; !ok {
			less.Add(k)
		}
		return true
	}
	dst.Each(do)

	return
}
func (h Hset) IsSubset(dst Set) bool {
	if h.Cap() > dst.Cap() {
		return false
	}

	for k := range h {
		if !dst.Contains(k) {
			return false
		}
	}
	return true
}
func (h Hset) String() string {
	ins := make([]string, 0, len(h))
	for k := range h {
		ins = append(ins, fmt.Sprintf("%v", k))
	}
	return fmt.Sprintf("Set[%s]", strings.Join(ins, ","))
}
func (h Hset) Union(dst Set) Set {
	reslut := New()
	for k := range h {
		reslut.Add(k)
	}
	do := func(k T) bool {
		reslut.Add(k)
		return true
	}
	dst.Each(do)
	return reslut
}
func (h Hset) Each(do func(T) bool) {
	for k := range h {
		if !do(k) {
			break
		}
	}
}

func (h Hset) Pop() T {
	for k := range h {
		delete(h, k)
		return k
	}
	return nil
}
