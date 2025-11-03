package queues

import "fmt"

func NewDeque() *Deque {
	return &Deque{elements: make([]interface{}, 0)}
}

func (deque *Deque) Print() {
	if deque.Size() == 0 {
		fmt.Println("The queue is empty")
		return
	}

	for _, element := range deque.elements {
		fmt.Printf("[%d] ", element)
	}

}

func (deque *Deque) Append(data interface{}) {
	deque.elements = append(deque.elements, data)
}

func (deque *Deque) Size() int {
	return len(deque.elements)
}

func (deque *Deque) IsEmpty() bool {
	return deque.Size() == 0
}

func (deque *Deque) Pop() interface{} {
	if deque.IsEmpty() {
		return nil
	}

	lastElementIndex := len(deque.elements) - 1
	lastElement := deque.elements[lastElementIndex]

	if deque.Size() == 1 {
		deque.elements = []interface{}{}
		return lastElement
	}

	slicedQueue := deque.elements[:lastElementIndex]
	deque.elements = slicedQueue
	return lastElement
}

func (deque *Deque) PopLeft() interface{} {
	if len(deque.elements) == 0 {
		return nil
	}

	elementToRemove := deque.elements[0]

	if len(deque.elements) == 1 {
		deque.elements = []interface{}{}
		return elementToRemove
	}

	slicedQueue := deque.elements[1:]
	deque.elements = slicedQueue
	return elementToRemove
}

func (deque *Deque) Insert(data interface{}) {
	if len(deque.elements) == 0 {
		deque.elements = append(deque.elements, data)
		return
	}

	newQueue := append([]interface{}{data}, deque.elements...)
	deque.elements = newQueue
}
