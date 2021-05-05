package main

import "fmt"

//栈相关
func testStack()  {
	//创建栈
	stack := make([]int, 0)
	//push压入
	stack = append(stack, 10)
	//pop弹出
	v := stack[len(stack)-1]
	stack = stack[:len(stack) - 1]
	fmt.Printf("stack: %v\n", v)
	//检查栈空
	if len(stack) == 0{
		fmt.Println("Empty!")
	}
}

//队列
func testQueue(){
	//创建队列
	queue := make([]int, 0)
	//enqueue入队
	queue = append(queue, 10)
	//dequeue出队
	v := queue[0]
	queue = queue[1:]
	fmt.Printf("queue: %v\n", v)
	//长度0为空
	if len(queue)==0 {
		fmt.Println("Empty!")
	}
}
/*
注意点:
参数传递，只能修改，不能新增或者删除原始数据
默认 s=s[0:len(s)]，取下限不取上限，数学表示为：[)
*/

//字典
func testDic()  {
	//
	m := make(map[string]int)
	//
	m["hello"] = 1
	//
	delete(m, "hello")
	//
	for k, v := range m{
		println(k, v)
	}
}
/*
注意点
map 键需要可比较，不能为 slice、map、function
map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由 0 变为 1
比较两个 map 需要遍历，其中的 kv 是否相同，因为有默认值关系，所以需要检查 val 和 ok 两个值
*/

//sort


func main()  {

}
