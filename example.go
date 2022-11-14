/*
 * @Name: Priority Sorted Queue Go
 * @Author: Max Base
 * @Date: 2022-11-14
 * @Repository: https://github.com/basemax/PrioritySortedQueueGo
 */

package main

import "fmt"

func main() {
	// Create a queue of capacity 5
	queue := NewPriorityQueue(10)

	// Inserting items to the queue
	queue.Enqueue("A", 1)
	queue.Enqueue("B", 20)
	queue.Enqueue("C", 3)
	queue.Enqueue("D", 4)
	queue.Enqueue("E", 5)

	// Print the queue
	fmt.Println(queue.ToString())

	// Dequeue two item
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	// Print the queue
	fmt.Println(queue.ToString())

	// Peek at the front of the queue
	fmt.Println(queue.Peek())

	// Print the queue
	fmt.Println(queue.ToString())

	// Print the size of the queue
	fmt.Println(queue.Size())

	// Clear the queue
	queue.Clear()

	// Print the queue
	fmt.Println(queue.ToString())
}
