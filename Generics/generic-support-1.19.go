package Generics

import "sort"

type GenericType[T sort.Interface] struct {
	t []T
}

func (g *GenericType[T]) printSort() {
	g := &GenericType[int]{
		t: []int{1, 2, 3},
	}
}
