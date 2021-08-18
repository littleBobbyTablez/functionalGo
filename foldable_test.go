package f

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestList_Fold(t *testing.T) {
	type args struct {
		acc T
		f   func(acc T, i T) T
	}
	tests := []struct {
		name string
		list List
		args args
		want T
	}{
		{"can find Max in List",
			List{15, 4, 20},
			args{0, max},
			20},
		{"can concatenate Strings",
			List{"a", "b", "c"},
			args{"", concatenate},
			"abc"},
		{"can count non empty strings",
			List{"a", "b", ""},
			args{0, countNonEmpty},
			2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Fold(tt.args.acc, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Map(t *testing.T) {
	tests := []struct {
		name string
		list List
		arg  func(elem T) T
		want List
	}{
		{
			"can map int to string",
			List{1, 2, 3},
			convertInts,
			List{"01: 1", "02: 2", "03: 3"},
		},
		{
			"can double entries",
			List{1, 2, 3},
			double,
			List{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Map(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Filter(t *testing.T) {
	tests := []struct {
		name string
		list List
		arg  func(elem T) bool
		want List
	}{
		{
			"can filter uneaven numbers",
			List{1, 2, 3},
			isUneaven,
			List{1, 3},
		},
		{
			"can filter for strings",
			List{"John", "Sandy", "Peter", "Sarah"},
			startsWithS,
			List{"Sandy", "Sarah"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Filter(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Reduce(t *testing.T) {
	type args struct {
		acc T
		f   func(acc T, i T) T
	}
	tests := []struct {
		name string
		list List
		args args
		want T
	}{
		{"can find Max in List",
			List{15, 4, 20},
			args{0, max},
			20},
		{"can concatenate Strings",
			List{"a", "b", "c"},
			args{"", concatenate},
			"abc"},
		{"can multiply",
			List{1, 5, 6},
			args{0, multiply},
			30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Reduce(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_ForEach(t *testing.T) {
	ch := make(chan int, 3)
	list := List{1, 2, 3}
	want := 6
	wg := sync.WaitGroup{}
	wg.Add(1)
	var got int
	go forEachAsync(list, ch, &wg)
	wg.Wait()
	close(ch)

	for i := range ch {
		got += i
	}

	if got != want {
		t.Errorf("ForEach() = %v, want %v", got, want)
	}

}

func TestList_Head(t *testing.T) {
	list := List{1, 2, 3}
	want := 1
	if got := list.Head(); got != want {
		t.Errorf("ForEach() = %v, want %v", got, want)
	}
}

func TestList_Last(t *testing.T) {
	list := List{1, 2, 3}
	want := 3
	if got := list.Last(); got != want {
		t.Errorf("ForEach() = %v, want %v", got, want)
	}
}

func TestList_Tail(t *testing.T) {
	list := List{1, 2, 3}
	want := List{2, 3}
	if got := list.Tail(); !reflect.DeepEqual(got, want) {
		t.Errorf("ForEach() = %v, want %v", got, want)
	}
}

func TestList_Init(t *testing.T) {
	list := List{1, 2, 3}
	want := List{1, 2}
	if got := list.Init(); !reflect.DeepEqual(got, want) {
		t.Errorf("ForEach() = %v, want %v", got, want)
	}
}

func forEachAsync(list List, ch chan int, wg *sync.WaitGroup) {
	list.ForEach(func(i T) { ch <- i.(int) })
	wg.Done()
}

func multiply(i T, j T) T {
	return i.(int) * j.(int)
}

func startsWithS(elem T) bool {
	return elem.(string)[0] == 'S'
}

func isUneaven(elem T) bool {
	return elem.(int)%2 == 1
}

func double(elem T) T {
	return elem.(int) * 2
}

func convertInts(elem T) T {
	i := elem.(int)
	msg := fmt.Sprintf("%02d: %d", i, i)
	return msg
}

func max(i T, j T) T {
	if i.(int) <= j.(int) {
		return j
	}
	return i
}

func concatenate(i T, j T) T {
	return i.(string) + j.(string)
}

func countNonEmpty(i T, j T) T {
	if j.(string) != "" {
		return i.(int) + 1
	}
	return i.(int)
}
