package learnGo

import (
	"fmt"
	DS "github.com/goMaster/learnGo/learnDS"
	"testing"
)

func TestOutputArray(t *testing.T) {
	anArray := []int{1, 2, 4, -4}

	twoD := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	OutputArray(anArray)
	Output2DArray(twoD)

}


func TestSliceGo(t *testing.T) {
	SliceGo()
	SortSlice()
}


func TestMaps(t *testing.T) {
	UsingMaps()
}

func TestUsingPoint(t *testing.T) {
	UsingPoint()
}

func TestUsingStruct(t *testing.T) {
	UsingStruct()
}

func TestUsingTree(t *testing.T){
	DS.UsingTree()
}

func TestUsingHash(t *testing.T)  {
	DS.UsingHashTable()

}

func TestUsingLinkedList(t *testing.T)  {
	DS.UsingLinkedList()

}

func TestUsingQueue(t *testing.T){
	DS.UsingQueue()
}

func TestUsingStack(t *testing.T){
	DS.UsingStack()
}
