package learnDS

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Graph: G (V, E)
		V-----set of vertices
		E-----set of edges
Types:
	1. cyclic graphs（循环图）
	2. acyclic graphs（非循环图）


Advantages of binary trees:
	1. represent hierarchical data
	2. 有序的构造，因此不需要重新排序（相应的，删除元素很麻烦）
	3. If a binary tree is balanced: search  insert  delete  ----- log(n) steps
	4. height of a balanced binary tree is approximately log2(n)

Disadvantage of binary trees:
	1. shape of the tree depends on the order in which its elements were inserted
	2. if a tree is not balanced, then the performance of the tree will be unpredictable.
*/

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

//递归遍历所有节点
//遍历顺序：左孩子-根结点-右孩子

func traverse(t *Tree){
	if t == nil{
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

//构造函数：创建一棵树
func create(n int) *Tree{
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}
//插入数据：有大小顺序，左<根<右
func insert(t *Tree, v int) *Tree{
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func UsingTree(){
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is",tree.Value)
}