//	queue: LIFO - last in, first out
//	pilha: ultimo a entrar, primeiro a sair
package structures

import "fmt"

type Node struct {
	Value int
}

type Stack struct {
	nodes []*Node
	count int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
