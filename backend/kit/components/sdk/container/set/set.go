package set

import (
	"sync"
)

type Set[T ISetItem] struct {
	c  map[string]T
	lc sync.Locker
}

type ISetItem interface {
	UniqueKey() string
}

func NewSet[T ISetItem](items ...T) Set[T] {

	s := Set[T]{
		c:  make(map[string]T, len(items)),
		lc: &sync.Mutex{},
	}

	for _, x := range items {
		s.Add(x)
	}

	return s
}

func (s Set[T]) Add(item T) bool {

	s.lc.Lock()

	var put bool
	key := item.UniqueKey()
	if _, found := s.c[key]; found {
		put = false
	} else {
		s.c[key] = item
		put = true
	}
	s.lc.Unlock()
	return put
}

func (s Set[T]) ToSlice() []T {

	s.lc.Lock()
	rsp := make([]T, 0, len(s.c))
	for _, v := range s.c {
		rsp = append(rsp, v)
	}

	s.lc.Unlock()
	return rsp
}

func (s Set[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for _, elem := range s.c {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}
