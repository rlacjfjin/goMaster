package learnDS

import "fmt"

var stack_size = 0
var stack = new(Node)

func stackPush(v int) bool {
	if stack == nil {
		stack = &Node{v, nil}
		stack_size = 1
		return true
	}
	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	stack_size++
	return true
}

func stackPop(t *Node) (int, bool) {
	if stack_size == 0 {
		return 0, false
	}
	if stack_size == 1 {
		stack_size = 0
		stack = nil
		return t.Value, true
	}
	stack = stack.Next
	stack_size--
	return t.Value, true
}

func traverseStack(t *Node) {
	if stack_size == 0 {
		fmt.Println("Empty Stack!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func UsingStack(){
	stack = nil
	v, b := Pop(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() failed!")
	}
	stackPush(100)
	traverseStack(stack)
	stackPush(200)
	traverseStack(stack)
	for i := 0; i < 10; i++ {
		stackPush(i)
	}
	for i := 0; i < 15; i++ {
		v, b := stackPop(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	traverseStack(stack)
}