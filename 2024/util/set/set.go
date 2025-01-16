package set

import "maps"

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, i := range items {
		s.Add(i)
	}
	return s
}

func (s Set[T]) Exists(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Add(val ...T) {
	for _, v := range val {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}

func (s Set[T]) Intersects(another Set[T]) bool {
	for i := range s {
		if _, found := another[i]; found {
			return true
		}
	}

	return false
}

func (s Set[T]) Combine(another Set[T]) Set[T] {
	newSet := maps.Clone(s)
	maps.Copy(newSet, another)
	return newSet
}

func (s Set[T]) Extend(another Set[T]) {
	maps.Copy(s, another)
}

func (s Set[T]) ForAll(fn func(item T) bool) bool {
	for i := range s {
		if !fn(i) {
			return false
		}
	}
	return true
}
