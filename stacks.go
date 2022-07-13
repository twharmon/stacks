package stacks

import (
	"fmt"
	"sync"
)

type Stack[T any] struct {
	mu    sync.Mutex
	items []T
	len   int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (q *Stack[T]) Push(items ...T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for i := range items {
		if q.len < len(q.items) {
			q.items[q.len] = items[i]
		} else {
			q.items = append(q.items, items[i])
		}
		q.len++
	}
}

func (q *Stack[T]) Pop() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.len == 0 {
		var t T
		return t
	}
	q.len--
	return q.items[q.len]
}

func (q *Stack[T]) Peek() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.len == 0 {
		var t T
		return t
	}
	return q.items[q.len-1]
}

func (q *Stack[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.len
}

func (q *Stack[T]) Slice() []T {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.items[:q.len]
}

func (q *Stack[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = nil
	q.len = 0
}

func (q *Stack[T]) String() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	return fmt.Sprintf("%v", q.items[:q.len])
}
