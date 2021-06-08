package main

import "fmt"

func main()  {
	//指针的常见使用（两符号）
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
	//make与new的使用
	/*这一段为引发panic，因为没分配空间
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)*/
	var d *int
	d = new(int)
	*d = 10
	fmt.Println(*d)

	var e map[string]int
	e = make(map[string]int, 10)
	e["沙河娜扎"] = 100
	fmt.Println(e)

	//自定义类型基础使用（大写加注释）

	//类型别名

}