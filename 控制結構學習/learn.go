package main

import (
	"fmt"
	"math"
)

// if 和 else
func pow(x, y, z float64) float64 {
	if v := math.Pow(x, y); v < z {
		return v
	} else {
		fmt.Printf("%g>= %g\n", v, z)
	}
	// can't use v here, though
	return z
}

// switch
func choose(i int) {
	switch i {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4:
		// 如果使用「 fallthrough 」關鍵字，會執行下一個 Case
		fallthrough
	case 5:
		fmt.Println("5")
	}
}

func anotherChoose(i int) {
	switch {
	case i > 5:
		fmt.Println(i, "> 5")
	case i < 5:
		fmt.Println(i, "< 5")
	case i == 5:
		fmt.Println(i, "= 5")
	}
}

func main() {
	// 第一種迴圈寫法
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

	// 第二種迴圈寫法
	two := 1
	for two < 1000 {
		two += two
	}
	fmt.Println("two =", two)

	// if 和 else
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Switch
	choose(1)
	choose(3)
	choose(4)

	anotherChoose(6)
}
