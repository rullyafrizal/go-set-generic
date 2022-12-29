package main

import "golang.org/x/exp/constraints"

type SetType interface {
	constraints.Ordered
}

type Set[T SetType] interface {
	add(val T)
	contains(key T) bool
	size() int
	remove(v T)
	indexOf(v T) int
}

type SetImpl[T SetType] struct {
	Cache   map[T]int
	Content []T
}

func NewSetString[T SetType]() Set[T] {
	return &SetImpl[T]{
		Cache:   map[T]int{},
		Content: make([]T, 0),
	}
}

func (s *SetImpl[T]) add(val T) {
	if s.contains(val) {
		return
	}

	(*s).Content = append((*s).Content, val)
	(*s).Cache[val] = s.size() - 1
	return
}

func (s *SetImpl[T]) contains(key T) bool {
	if _, ok := s.Cache[key]; ok {
		return true
	}
	return false
}

func (s *SetImpl[T]) size() int {
	return len((*s).Content)
}

func (s *SetImpl[T]) remove(v T) {
	if s.contains(v) {
		var removedIdx int = s.indexOf(v)
		(*s).Content = append((*s).Content[:removedIdx], (*s).Content[removedIdx+1:]...)
		delete((*s).Cache, v)

		return
	}
	return
}

func (s *SetImpl[T]) indexOf(v T) int {
	return (*s).Cache[v]
}
