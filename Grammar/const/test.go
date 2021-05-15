package main

import "fmt"

const pi= 3.14
const e = 2.7

const(
	test1 = 1
	test2 = 2
)

const(
	n1 = 10
	n2 //10
	n3 //10
)

const(
	show0 = iota //0
	show1 //1
	_
	show3 //3
)

const(
	num0 = iota //0
	_
	num2 = 100 //100
	num3 = iota + 2  //5
)

const num4 = iota //0

const(
	a, b = iota + 1, iota + 2 //1, 2
	c, d //2, 3
	f, h //3, 4
)

func main(){
	fmt.Println(pi, e)
	fmt.Println(test1, test2)
	fmt.Println(n1, n2, n3)
	fmt.Println(show0, show1, show3)
	fmt.Println(num0, num2, num3, num4)
	fmt.Println(a, b, c, d, f, h)
}


