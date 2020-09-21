package learnGo

import "fmt"

/*
1. 创建struct
2. 初始化函数（返回指针）
3. New()函数-------new returns a pointer
   Note: make 与 New区别： make是初始化，返回的不是指针，用在maps，channels，slices
*/


type XYZ struct {
	X int
	Y int
	Z int
}




func UsingStruct(){
	var s XYZ
	fmt.Println(s.X)
	fmt.Println(s)

	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}

	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlice)

	s1 := createStruct("Mihalis", "Tsoukalos", 123)
	s2 := retStructure("Mihalis", "Tsoukalos", 123)
	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

}


type myStructure struct{
	Name string
	Surname string
	Height int32
}
//返回指针类型的函数
func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}

}
//返回类型为不是指针
func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}

}

