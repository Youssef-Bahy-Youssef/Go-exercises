package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x <= 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
	// select {}
}

// func main() {
// 	naturals := make(chan int)
// 	squares := make(chan int)
// 	// Counter
// 	go func() {
// 		for x := 0; x < 100; x++ {
// 			naturals <- x
// 		}
// 		close(naturals)
// 	}()
// 	// Squarer
// 	go func() {
// 		for x := range naturals {
// 			squares <- x * x
// 		}
// 		close(squares)
// 	}()
// 	// Printer (in main goroutine)
// 	for x := range squares {
// 		fmt.Println(x)
// 	}

// }
