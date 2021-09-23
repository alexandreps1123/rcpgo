package test

import (
	"fmt"

	myStructures "rcpgo/datastructures"
)

func TestStack() {

	s := myStructures.NewStack()

	s.Push(&myStructures.Node{1})
	s.Push(&myStructures.Node{2})
	s.Push(&myStructures.Node{3})
	s.Push(&myStructures.Node{4})
	s.Push(&myStructures.Node{5})

	fmt.Println("Saída esperada: 5 4 3 2 1")
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop(), s.Pop())
}
