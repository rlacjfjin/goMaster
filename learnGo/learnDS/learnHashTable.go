package learnDS

import "fmt"

/*
什么是哈希表？？

*/

const SIZE = 15


type HashTable struct {
	Table map[int]*Node
	Size int
}
//取余数，相当于求index的值：0，1，2，。。。，size-1
func hashFunction(i, size int) int {
	return (i % size)
}

func insertHash(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverseHash(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookup(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func UsingHashTable(){
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insertHash(hash, i)
	}
	traverseHash(hash)
	for i:=118;i<123; i++{
		if lookup(hash,i){
			fmt.Printf("%d is in the hash table!\n",i)
		}else{
			fmt.Printf("%d is not in the hash table!\n",i)

		}
	}

}
