package set

type T = interface{}

type Set interface {
	Add(data T)
	Remove(data T)
	Cap() int
	Equal(dst Set) bool
	Intersect(dst Set) Set
	Difference(dst Set) Set
	Clear()
	Clone() Set
	Contains(val ...T) bool
	Compare(dst Set) (more, less Set)
	IsSubset(dst Set) bool
	String() string
	Union(dst Set) Set
	Each(do func(T) bool)
	Pop() T
}
