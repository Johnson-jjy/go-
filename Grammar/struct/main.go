package main

import "fmt"

func test() {
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	//fmt.Printf("test-stus: %p\n", &stus[0])
	//fmt.Printf("test-stus: %p\n", &stus[1])

	for _, stu := range stus {
		m[stu.name] = &stu
		fmt.Printf("test-stu: %p\n", &stu) //range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Printf("test: %v\n", m)

	for i := 0; i < len(stus); i++ {
		m[stus[i].name] = &stus[i]
		fmt.Printf("test-stu: %p\n", &stus[i]) //获取真实地址直接取
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Printf("test: %v\n", m)
}

func main()  {
	//自定义类型基础使用（大写加注释）

	//类型别名

	//结构的定义

	//结构体的初始化与结构体指针

	//1.结构体的基本初始化（与指针语法糖）

	//2.键值对初始化

	//3.值的列表初始化

	//匿名结构体

	//补充知识点：结构体的内存对齐

	//结构体构造函数




	//方法的定义（接收者：值与指针）

	//修改的辨析

	//为任意类型添加方法--基于基本类型添加自己的方法



	//匿名结构体

	//嵌套结构体

	//匿名嵌套

	//匿名嵌套体数据的直接访问与间接访问

	//访问冲突的情况

	//结构体的继承




	//结构体字段的可见性

	//JSON序列化

	//JSON反序列化

	//加tag

	test()
}
