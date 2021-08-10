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
	return list.filterRec(0, List{}, f)
}

func (list List) filterRec(i int, result List, f func(elem T) bool) T {
	if len(list) == i {
		return result
	}
	if f(list[i]) {
		return list.filterRec(i+1, append(result, list[i]), f)
	}
	return list.filterRec(i+1, result, f)
}

func (list List) Map(f func(elem T) T) T {
	return list.mapRec(0, List{}, f)
}

func (list List) mapRec(i int, result List, f func(elem T) T) T {
	if len(list) == len(result) {
		return result
	}
	updated := append(result, f(list[i]))
	return list.mapRec(i+1, updated, f)
}

func (list List) Reduce(f func(i T, j T) T) T {
	return list[1:].reduceRec(list[0], list[1], f)
}

func (list List) reduceRec(x T, y T, f func(i T, j T) T) T {
	if len(list) <= 1 {
		return f(x, y)
	}
	t := f(x, y)
	newList := list[1:]
	return newList.reduceRec(t, newList[0], f)
}
