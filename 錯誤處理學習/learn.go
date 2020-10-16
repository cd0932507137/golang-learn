package main

import "fmt"

func main() {
	// 處理錯誤就必須要加上 defer 這樣才能正確的抓到
	defer func() {
		fmt.Println("first")
		// 這裡要注意的是 Go 的 panic 的錯誤真的是錯誤，程式一定會發生錯誤然後關閉，除非你有使用 recover 會回覆他的數值。
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end")
	}()
	f()
}

func f() {
	fmt.Println("test")
	panic(1)
	fmt.Println("test2")
}
