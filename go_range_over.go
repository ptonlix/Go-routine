package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	//close(queue)
	//test := <-queue
	select {
	case test := <- queue:
		fmt.Println(test)
	default:
		fmt.Println("default")
	}
	//fmt.Println(test)
	//for elem := range queue {
	//	fmt.Println(elem)
	//}
}
