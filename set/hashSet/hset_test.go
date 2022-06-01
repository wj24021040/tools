package hashSet

import (
	"testing"
)

var (
	srcStrSli = []string{"1", "2", "3", "1", "4", "5", "6", "7", "8", "8", "9", "10", "11", "12"}
	dstStrSli = []string{"2", "3", "6", "7", "8", "8", "12", "15"}
	srcl      = len(srcStrSli) - 2
	dstl      = len(dstStrSli) - 1
	more      = Hset(map[interface{}]struct{}{
		"1":  struct{}{},
		"4":  struct{}{},
		"5":  struct{}{},
		"9":  struct{}{},
		"10": struct{}{},
		"11": struct{}{},
	})
	less = Hset(map[interface{}]struct{}{
		"15": struct{}{},
	})
	u = Hset(map[interface{}]struct{}{
		"1":  struct{}{},
		"2":  struct{}{},
		"3":  struct{}{},
		"4":  struct{}{},
		"5":  struct{}{},
		"6":  struct{}{},
		"7":  struct{}{},
		"8":  struct{}{},
		"9":  struct{}{},
		"10": struct{}{},
		"11": struct{}{},
		"12": struct{}{},
		"15": struct{}{},
	})
	i = Hset(map[interface{}]struct{}{
		"2":  struct{}{},
		"3":  struct{}{},
		"6":  struct{}{},
		"7":  struct{}{},
		"8":  struct{}{},
		"12": struct{}{},
	})
	d = Hset(map[interface{}]struct{}{
		"1":  struct{}{},
		"4":  struct{}{},
		"5":  struct{}{},
		"9":  struct{}{},
		"10": struct{}{},
		"11": struct{}{},
		"15": struct{}{}})
)

func TestHashSetStr(t *testing.T) {
	srcSet := New()
	for _, v := range srcStrSli {
		srcSet.Add(v)
	}

	if srcSet.Cap() != srcl {
		t.Error("the hset cap is wrong, expect: ", srcl, "but active: ", srcSet.Cap())
	}
	dstSet := New()
	for _, v := range dstStrSli {
		dstSet.Add(v)
	}
	if dstSet.Cap() != dstl {
		t.Error("the hset cap is wrong, expect: ", dstl, "but active: ", dstSet.Cap())
	}

	srcCp := srcSet.Clone()
	if !srcCp.Equal(srcSet) {
		t.Error("the copy make two set which not equal, cp: ", srcCp.String())
	}

	srcCp.Clear()
	if srcCp.Cap() != 0 {
		t.Error("the hset clear,but has items: ", srcCp.String())
	}
	srcSet.Add("9999")
	if !srcSet.Contains("9999") {
		t.Error("after add 9999, but the hset don't Contains  9999!")
	}

	srcSet.Remove("9999")
	if srcSet.Contains("9999") {
		t.Error("after delete 9999， but the hset Contains  9999!")
	}

	m, l := srcSet.Compare(dstSet)
	if !m.Equal(more) || !l.Equal(less) {
		t.Error("hset compare 1 failed, ", m.String(), "   ", l.String())
	}

	m, l = dstSet.Compare(srcSet)
	if !m.Equal(less) || !l.Equal(more) {
		t.Error("hset compare 2 failed")
	}

	iner := srcSet.Intersect(dstSet)
	if !iner.Equal(i) {
		t.Error("hset Intersect failed")
	}

	iner = dstSet.Intersect(srcSet)
	if !iner.Equal(i) {
		t.Error("hset Intersect failed")
	}

	if !iner.IsSubset(srcSet) || !iner.IsSubset(dstSet) {
		t.Error("the Intersect set isn't the Subset!")
	}

	uni := srcSet.Union(dstSet)
	if !uni.Equal(u) {
		t.Error("hset Union failed")
	}

	uni = dstSet.Union(srcSet)
	if !uni.Equal(u) {
		t.Error("hset Union failed")
	}

	di := srcSet.Difference(dstSet)
	if !di.Equal(d) {
		t.Error("hset Difference failed: ", di.String())
	}

	di = dstSet.Difference(srcSet)
	if !di.Equal(d) {
		t.Error("hset Difference failed")
	}

	v := srcSet.Pop()
	if v == nil || srcSet.Cap() != (srcl-1) {
		t.Error("hset Pop failed")
	}

	srcSet.String()
}
