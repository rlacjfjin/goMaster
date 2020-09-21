package learnDS

import "fmt"

/*
两个操作：
	1. Pop
	2. Push
*/

var queue_size = 0
var queue = new(Node)

func Push(t *Node, v int) bool {
	if queue == nil {
		queue = &Node{v, nil}
		queue_size++
		return true
	}
	t = &Node{v, nil}
	t.Next = queue
	queue = t
	queue_size++
	return true
}

func Pop(t *Node) (int, bool) {
	if queue_size == 0 {
		return 0, false
	}
	if queue_size == 1 {
		queue = nil
		queue_size--
		return t.Value, true
	}
	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}
	v := (temp.Next).Value
	temp.Next = nil
	queue_size--
	return v, true
}

func traverseQueue(t *Node) {
	if queue_size == 0 {
		fmt.Println("Empty Queue!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func UsingQueue(){
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", queue_size)
	traverseQueue(queue)
	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queue_size)

	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	traverseQueue(queue)
	fmt.Println("Size:", queue_size)
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queue_size)
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queue_size)
	traverseQueue(queue)
}
