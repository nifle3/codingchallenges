package dataStruct

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	Freq  int
	Elem  rune
	Left  *Node
	Right *Node
}

func (n *Node) PrintTree() {
	var defaultValue rune
	if n.Elem != defaultValue {
		fmt.Printf("%s ", strconv.QuoteRune(n.Elem))
		return
	}

	fmt.Printf("%d ", n.Freq)

	if n.Left != nil {
		n.Left.PrintTree()
	}

	if n.Right != nil {
		n.Right.PrintTree()
	}
}

type PrioretyQueue struct {
	heap   []Node // 0 - base node, 2i - right, 2i +1 - left, parent -  i / 2
	length int
}

func New(startCap int) PrioretyQueue {
	if startCap < 0 {
		startCap = 0
	}

	return PrioretyQueue{
		heap: make([]Node, 0, startCap),
	}
}

func (p *PrioretyQueue) Insert(node Node) {
	p.heap = append(p.heap, node)
	p.length++

	if len(p.heap) != p.length {
		/*
			we know that last inserted value has len() - 1 index.
			we need to replace this element to p.length - 1
		*/
		p.heap[p.length-1], p.heap[len(p.heap)-1] =
			p.heap[len(p.heap)-1], p.heap[p.length]
	}

	for i := p.length - 1; i > 2; {
		parentIndex := int(i) / int(2)
		if p.heap[i].Freq > p.heap[parentIndex].Freq {
			break
		}

		p.heap[i], p.heap[parentIndex] = p.heap[parentIndex], p.heap[i]
		i = parentIndex
	}
}

func (p *PrioretyQueue) ExtractMinimum() (Node, error) {
	if p.length <= 0 {
		return Node{}, errors.New("your queue is empty")
	}

	p.length--

	extracted := p.heap[0]
	p.heap[0], p.heap[p.length] = p.heap[p.length], p.heap[0]

	p.heapify()

	return extracted, nil
}

func (p *PrioretyQueue) heapify() {
	for i := 0; i < p.length; i++ {
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		smallest := i

		if leftIndex < p.length && p.heap[leftIndex].Freq < p.heap[smallest].Freq {
			smallest = leftIndex
		}
		if rightIndex < p.length && p.heap[rightIndex].Freq < p.heap[smallest].Freq {
			smallest = rightIndex
		}

		if smallest == i {
			break
		}

		p.heap[i], p.heap[smallest] = p.heap[smallest], p.heap[i]
		i = smallest
	}
}

func (p *PrioretyQueue) SeeMinimum() (Node, error) {
	if p.length <= 0 {
		return Node{}, errors.New("your queue is empty")
	}

	return p.heap[0], nil
}

func (p *PrioretyQueue) Length() int {
	return p.length
}
