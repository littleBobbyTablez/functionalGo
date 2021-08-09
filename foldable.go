package foldable

type T interface{}

type Foldable interface {
	Fold(init T, f func(result, next T) T) T
	Init() Foldable
	Append(item T) Foldable
}

type List []T
