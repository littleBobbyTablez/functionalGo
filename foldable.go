package f

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

func (list List) ForEach(f func(elem T)) {
	list.forEachRec(0, f)
}

func (list List) forEachRec(i int, f func(elem T)) {
	if len(list) == i {
		return
	}
	f(list[i])
	list.forEachRec(i+1, f)
}

func (list List) Head() T {
	return list[0]
}

func (list List) Tail() T {
	return list[1:]
}

func (list List) Last() T {
	return list[len(list)-1]
}

func (list List) Init() T {
	if len(list) == 0 {
		return List{}
	}
	return list[:len(list)-1]
}
