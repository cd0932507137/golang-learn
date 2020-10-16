package main

import "fmt"

func add(x, y int) int {
	return x + y
}

// 多數值的返回
func swap(x, y string) (string, string) {
	return x, y
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 一個沒有參數的 return 語句，會將當前的值作為返回值返回。
	return
}

func main() {
	fmt.Println(add(5, 5))
	// 多數值的返回
	a, b := swap("demo", "string")
	fmt.Println(a, b)
	// 命名返回值
	fmt.Println(split(17))
}
