package test

import (
	"fmt"
	"math/rand"
	"rcpgo/sort"
)

func TestBubbleSort() {
	var a []int

	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(100))
	}

	sort.BubbleSort(a)

	fmt.Println(a)

}

func TestQuickSort() {
	var a []int

	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(100))
	}

	//(array, indice_primeiro_elemento, indice_ultimo_elemento)
	sort.QuickSort(a)

	fmt.Println(a)
}
