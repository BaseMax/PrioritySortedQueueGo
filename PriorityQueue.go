/*
 * @Name: Priority Sorted Queue Go
 * @Author: Max Base
 * @Date: 2022-11-14
 * @Repository: https://github.com/basemax/PrioritySortedQueueGo
 */

package main

import "fmt"

type QueueItem struct {
	Value    interface{}
	Priority int
}

type PriorityQueue struct {
	front int
	rear  int
	size  int
	// sorted list
	items []QueueItem
}

// Create, O(1)
func NewPriorityQueue(size int) *PriorityQueue {
	return &PriorityQueue{
		front: -1,
		rear:  -1,
		size:  size,
		items: make([]QueueItem, size),
	}
}

// IsFull, O(1)
func (q *PriorityQueue) IsFull() bool {
	return q.rear == q.size-1
}

// IsEmpty, O(1)
func (q *PriorityQueue) IsEmpty() bool {
	return q.front == -1 || q.front > q.rear
}

// Enqueue, O(n)
func (q *PriorityQueue) Enqueue(value interface{}, priority int) {
	if q.IsFull() {
		return
	}
	q.rear++
	q.items[q.rear] = QueueItem{
		Value:    value,
		Priority: priority,
	}
	if q.front == -1 {
		q.front = 0
	}
	// sort
	for i := q.rear; i > 0; i-- {
		if q.items[i].Priority > q.items[i-1].Priority {
			q.items[i], q.items[i-1] = q.items[i-1], q.items[i]
		}
	}
}

// Dequeue, O(1)
func (q *PriorityQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := q.items[q.front]
	q.front++
	return item.Value
}

// Peek, O(1)
func (q *PriorityQueue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[q.front].Value
}

// Size, O(1)
func (q *PriorityQueue) Size() int {
	return q.rear - q.front + 1
}

// Clear, O(1)
func (q *PriorityQueue) Clear() {
	q.front = -1
	q.rear = -1
}

// ToString, O(n)
func (q *PriorityQueue) ToString() string {
	if q.IsEmpty() {
		return ""
	}
	str := ""
	for i := q.front; i <= q.rear; i++ {
		str += fmt.Sprintf("%v ", q.items[i].Value)
	}
	return str
}
