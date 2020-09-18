package learnGo

import (
	"fmt"
)

/*
缺点：
1. cannot change its size
2. .....
3. .....

总之，用array速度慢！

Because of their disadvantages, arrays are rarely used in Go!

*/

func OutputArray(array []int){
	fmt.Println("size of array:",len(array))
	for i,value := range array{
		fmt.Println("index:", i, "value: ", value)
	}
}

func Output2DArray(array [][]int){
	fmt.Println("size of array:",len(array))
	for i,v := range array{
		for j,value := range v{
		fmt.Println("index:(", i,j, ")  value: ", value)
		}
	}
}