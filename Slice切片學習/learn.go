package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	// Slice 有兩種數值，一個是容量一個是長度
	// 如果要自訂容量大小
	// slice2 := make([]int, 2, 10)
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
