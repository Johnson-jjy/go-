package main

import (
	"fmt"
	"runtime"
	"time"
)

//var wg sync.WaitGroup
//
//func hello(i int)  {
//	defer wg.Done() // -1
//	fmt.Println("Hello Goroutine!", i)
//}
//
//func main()  {
//	for i := 0; i < 10; i++{
//		wg.Add(1) // +1
//		go hello(i)
//	}
//	wg.Wait()
//}


func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}