# Priority Sorted Queue Go (PrioritySortedQueueGo)

This is a priority sorted queue written in Go (Golang).

## Features

-  Sorted list
-  Support for custom priority
-  Support for custom data type

## Installation

```bash
$ go get github.com/basemax/PrioritySortedQueueGo
```

## Usage

```go
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
```

© Copyright 2022, Max Base
