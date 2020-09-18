package learnGo

import (
	"fmt"
	"sort"
)

/*
1. capacity  cap()
2. length  len()
3. append
4. copy
5. sort.Slice()
*/

func SliceGo(){
	aSliceLiteral := []int{1, 2, 3, 4, 5}   //方式1
	integer := make([]int, 20)     //方式2

	//append
	fmt.Printf("Cap: %d, Length: %d\n", cap(integer), len(integer))
	integer = append(integer,444)
	fmt.Println("after append:",integer)
	fmt.Printf("Cap: %d, Length: %d\n", cap(integer), len(integer))

	//reslice
	reSlice := integer[1:3]
	fmt.Println("part of slices:",reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println("update of part slices:",reSlice)
	fmt.Println("orginal slices:",integer)

	//empty
	fmt.Println(aSliceLiteral)
	aSliceLiteral = nil   //
	fmt.Println("after empty:",aSliceLiteral)


	//copy
	fmt.Println("---------copy------------")
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)

	//twoD slice
	twoD := make([][]int, 3)
	fmt.Println(twoD)

	//sorting
}



type aStructure struct {
	person string
	height int
	weight int
}

func SortSlice(){
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}