package foldable

type T interface{}
type Int int

type Foldable interface {
	Fold(init T, f func(result, next T) T) T
	Init() Foldable
	Append(item T) Foldable
}

type Mappable interface {
	Map(f func(t T) T) Mappable
}

type List []T

func (list List) Fold(acc T, f func(acc T, i T) T) T {
	if len(list) == 0 {
		return acc
	}
	t := f(acc, list[0])
	return list[1:].Fold(t, f)
}

func (list List) Filter(f func(elem T) bool) T {
	return list.FilterRec(0, f)
}

func (list List) FilterRec(i int, f func(elem T) bool) T {
	if len(list) == 0 {
		return list
	}
	if f(list[i]) {
		return list.FilterRec(i+1, f)
	}
	return append(list[:i], list[i:]).FilterRec(i, f)
}

func (list List) Map(f func(elem T) T) T {
	return list.MapRec(0, f)
}

func (list List) MapRec(i int, f func(elem T) T) T {
	if len(list) == 0 {
		return list
	}
	updated := append(list[:i], f(list[i]), list[i:])
	return updated.MapRec(i+1, f)
}
