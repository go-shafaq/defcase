package smap

type Map[V any] interface {
	Set(string, *V)
	SetVal(string, V)
	GetOrSet(k string, v *V) *V
	Get(string) (*V, bool)
}
