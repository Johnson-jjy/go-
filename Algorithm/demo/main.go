package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

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
	//创建
	m := make(map[string]int)
	//设置kv
	m["hello"] = 1
	//遍历
	for k, v := range m{
		fmt.Println(k, v)
	}
	//删除k
	delete(m, "hello")
}
/*
注意点
map 键需要可比较，不能为 slice、map、function
map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由 0 变为 1
比较两个 map 需要遍历，其中的 kv 是否相同，因为有默认值关系，所以需要检查 val 和 ok 两个值
*/

//sort
func testSort(){
	//int排序
	sI := []int{3, 2, 1, 6, 5, 4}
	sort.Ints(sI)
	fmt.Printf("Ints: %v\n", sI)

	//字符串排序
	sS := []string{"test", "cxk", "hello"}
	sort.Strings(sS)
	for _, value := range sS{
		fmt.Printf("%s\t", value)
	}

	//自定义排序
	sort.Slice(sI, func(i, j int)bool{return sI[i] < sI[j]})
	fmt.Printf("Ints: %v\n", sI)
}

//math
func testMath() {
	//int32最大最小值
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	//int64最大最小值（int默认是int64）
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)
}

//copy
func testCopy()  {
	//删除a[i](此处i为3)，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	a := []int{1, 2, 3, 4, 5, 6}
	copy(a[3:], a[4:])
	a = a[:len(a) - 1]
	fmt.Println(a)

	//make创建长度，则通过索引赋值
	b := make([]int, 5)
	b[4] = 1
	fmt.Println(b)

	//make长度为0，则通过append()赋值
	c := make([]int, 0)
	c = append(c, 1)
	fmt.Println(c)
}

//skill
func testSkill()  {
	//byte转数字
	s := "123456" //s[0]类型是byte
	num := int(s[0] - '0') //1
	str := string(s[0]) // "1"
	b := byte(num + '0') // '1'
	fmt.Printf("%d\t %s\t %c\n", num, str, b) //111

	//字符串转数字
	num,_ = strconv.Atoi(str)
	str = strconv.Itoa(100)
	fmt.Printf("%v\t%v\n", num, str)
}

func main()  {
	//栈相关
	testStack()
	//队列相关
	testQueue()
	//字典相关
	testDic()
	//sort
	testSort()
	//math
	testMath()
	//copy
	testCopy()
	//skill
	testSkill()
}
