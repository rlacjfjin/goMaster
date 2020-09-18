package learnGo

import "fmt"

func getPointer(n *int) {
	*n = *n * *n
}				//不通过返回值来获取变量，直接通过指针的形式改变传入的参数

func returnPointer(n int) *int {
	v := n*n
	return &v
}			 // 返回指针变量

func UsingPoint(){
	i := -10
	j := 25

	pI := &i 		//对非指针变量引用&，获取该变量的地址
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	*pI = 123456
	fmt.Println("i:", i)
	*pI--
	fmt.Println("i:", i)

	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}