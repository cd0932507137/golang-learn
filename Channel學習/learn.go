package main

import "fmt"

func main() {
	s := make(chan string) //宣告一個 channel 變數
	go func() {
		fmt.Println(111)
		s <- "hello" //寫入 channel (sender)
		fmt.Println(s)
		fmt.Println(222)
	}()

	val := <-s //讀取 channel (receiver)
	fmt.Println(val)
}
