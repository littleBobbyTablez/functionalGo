package foldable

import (
	"fmt"
	"reflect"
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
			"Can map int to string",
			List{1, 2, 3},
			convertInts,
			List{"01: 1", "02: 2", "03: 3"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func convertInts(elem T) T {
	i := elem.(int)
	msg := fmt.Sprintf("%2d: %d", i, i)
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
