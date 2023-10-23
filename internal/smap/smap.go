package smap

import (
	"sync"
)

func NewMap[V any]() Map[V] {
	return sm[V]{new(sync.Map)}
}

type sm[V any] struct {
	m *sync.Map
}

func (s sm[V]) Set(k string, v *V) {
	s.m.Store(k, v)
}

func (s sm[V]) SetVal(k string, v V) {
	s.Set(k, &v)
}

func (s sm[V]) GetOrSet(k string, v *V) *V {
	val, _ := s.m.LoadOrStore(k, v)
	if val == nil {
		return nil
	}

	return val.(*V)
}

func (s sm[V]) Get(k string) (*V, bool) {
	v, ok := s.m.Load(k)
	if v == nil {
		return nil, ok
	}

	return v.(*V), ok
}
