package datastruct

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type NodeValue[K any, V constraints.Ordered] struct {
	key   K
	value V
}

type PrioretyQueue[K any, V constraints.Ordered] struct {
	heap   []NodeValue[K, V] // 0 - base node, 2i - right, 2i +1 - left, parent -  i / 2
	length int
}

func New[K any, V constraints.Ordered](startCap int) PrioretyQueue[K, V] {
	return PrioretyQueue[K, V]{
		heap: make([]NodeValue[K, V], 0, startCap),
	}
}

func (p *PrioretyQueue[K, V]) Insert(key K, value V) {
	insertedValue := NodeValue[K, V]{
		key:   key,
		value: value,
	}

	p.heap = append(p.heap, insertedValue)
	p.length++

	if len(p.heap) != p.length {
		/*
			we know that last inserted value has len() - 1 index.
			we need to replace this element to p.length - 1
		*/
		p.heap[p.length-1], p.heap[len(p.heap)-1] =
			p.heap[len(p.heap)-1], p.heap[p.length]
	}

	for i := p.length - 1; i < 2; {
		parentIndex := int(i) / int(2)
		if p.heap[i].value > p.heap[parentIndex].value {
			break
		}

		p.heap[i], p.heap[parentIndex] = p.heap[parentIndex], p.heap[i]
		i = parentIndex
	}
}

func (p *PrioretyQueue[K, V]) ExtractMinimum() (NodeValue[K, V], error) {
	if p.length <= 0 {
		var zero NodeValue[K, V]
		return zero, errors.New("your queue is empty")
	}

	p.length--

	extracted := p.heap[0]
	p.heap[0], p.heap[p.length] = p.heap[p.length], p.heap[0]

	p.heapify()

	return extracted, nil
}

func (p *PrioretyQueue[K, V]) heapify() {
	for node, i := 0, 1; i < p.length || node < p.length; i++ {
		leftIndex := 2 * i
		rightIndex := (2 * i) - 1

		if p.heap[node].value <= p.heap[leftIndex].value &&
			p.heap[node].value <= p.heap[rightIndex].value {
			break
		}

		var indexOfMinElem int

		if p.heap[leftIndex].value < p.heap[rightIndex].value {
			indexOfMinElem = leftIndex
		} else {
			indexOfMinElem = rightIndex
		}

		p.heap[node], p.heap[indexOfMinElem] = p.heap[indexOfMinElem], p.heap[node]
		node = indexOfMinElem
	}
}
