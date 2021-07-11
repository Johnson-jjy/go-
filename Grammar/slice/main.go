package main

import (
	"fmt"
	"sync"
)

// 对扩容进行一些测试
func testAppend() {
	a := []byte{1, 0}
	fmt.Printf("len of a is: %d,cap of a is:%d\n",len(a), cap(a))
	a = append(a, 1, 1, 1)
	fmt.Printf("len of a is: %d,cap of a is:%d\n",len(a), cap(a))

	b := []int{23, 51}
	b = append(b, 4, 5, 6)
	fmt.Printf("len of b is: %d,cap of b is:%d\n",len(b), cap(b))

	c := []int32{1, 23}
	c = append(c, 2, 5, 6)
	fmt.Printf("len of c is: %d,cap of c is:%d\n",len(c), cap(c))

	type D struct{
		age byte
		name string
	}
	d := []D{
		{1,"123"},
		{2,"234"},
	}

	d = append(d,D{4,"456"},D{5,"567"},D{6,"678"})
	fmt.Printf("len of d is: %d,cap of d is:%d\n",len(d), cap(d))
}

func main()  {
	//切片的内部结构：地址+长度+容量

	//切片的初始化--声明的基本语法：var name []T
	//0.切片的直接声明与初始化
	var t0 []string
	var t1 = []int{} //并完成了初始化
	var t2 = []bool{true, false}
	fmt.Println(t0)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t0 == nil) //只能与nil比较--	切片比较--nil--无底层构造
	fmt.Println(t1 == nil) //虽空，但完成了初始化 --故判空应该对比len与0，而nil是没有底层构造
	fmt.Println(t2 ==nil)

	//1.切数组
	var origin = [...]int{1, 2, 3, 4, 5}
	var slice1 = origin[1:3]
	slice2 := origin[1:]
	slice3 := origin[:2]
	slice4 := origin[:]
	fmt.Printf("slice1:%v;len(slice1):%v;cap(slice1):%v \n",slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2:%v;len(slice2):%v;cap(slice2):%v \n",slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3:%v;len(slice3):%v;cap(slice3):%v \n",slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4:%v;len(slice4):%v;cap(slice4):%v \n",slice4, len(slice4), cap(slice4))
	//完整的切片表达式
	t1 = origin[1:3:4] //注意t1与slice1的区别
	fmt.Printf("t1:%v;len(t1):%v;cap(t1):%v \n",t1, len(t1), cap(t1))


	//2.切切片
	//high的上限边界是切片的容量cap(a)，而不是长度。常量索引必须是非负的
	s1 := slice2[1:3]
	fmt.Printf("s1:%v;len(s1):%v;cap(s1):%v \n",s1, len(s1), cap(s1))

	//3.make构造--make([]T, size, cap)
	var sm1 = make([]int, 2, 7)
	fmt.Printf("sm1:%v  len(sm1):%v   cap(sm1):%v\n",sm1, len(sm1), cap(sm1))
	sm1[0] = 1
	sm1[1] = 2
	//sm1[2] = 3 --不能如此赋值，此时len始终是2
	sm1 = append(sm1, 3) //只能用append动态添加元素--超cap时还会涉及扩容
	fmt.Printf("new!---sm1:%v  len(sm1):%v   cap(sm1):%v\n",sm1, len(sm1), cap(sm1))

	//切片赋值构造
	sm2 := sm1 //sm2会与sm1公用一个底层数组
	sm2[2] = 100
	fmt.Println("sm1:",sm1)
	fmt.Println("sm2:",sm2)

	//copy切片 --copy(destSlice, srcSlice []T)
	sm3 := make([]int, 4, 5)
	copy(sm3, sm1) // sm3的len必须大于sm1
	fmt.Println("sm3:",sm3)
	sm3[2] = 200
	fmt.Println("new! sm1:",sm1)
	fmt.Println("now! sm3:",sm3)
	//sm4 := [5]int{1,2,3,4,5} -- 若使用数组则报错

	//遍历切片--与遍历数组相同
	for i := 0; i < len(sm1); i++{
		fmt.Println(i,sm1[i])
	}
	for index, value := range sm1{
		fmt.Println(index, value)
	}

	//append方法相关
	//1.追加与多次追加
	var sa1 []int
	//var sa2 = make([]int)--会报错说没声明len

	sa1 = append(sa1, 1) //没有必要初始化！
	//sa2 = append(sa2, 2, 3, 4)
	fmt.Println("sa1:", sa1)

	//2.类似拼接--...
	sa3 := []int{5, 6, 7}
	sa1 = append(sa1, sa3...)
	fmt.Println("new! sa1:", sa1)

	//3.删除元素
	sa1 = append(sa1[:1], sa1[3:]...)//相当于删除了索引为1、2的元素
	fmt.Println("now! sa1:", sa1)

	//扩容策略
	fmt.Println("扩容策略")
	var numSlice []int
	for i := 0; i < 10; i++{
		numSlice = append(numSlice, i)
		fmt.Printf("%v	len:%d	cap:%d	ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	//小test__解释代码的输出结果？？？
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)

	//sort_test_对数组var a = [...]int{3, 7, 8, 9, 1}进行排序


	// 对扩容进行一些测试
	testAppend()

	// 对函数参数传入结构体做的一些测试
	testSlicePointer()

	// 线程安全相关的测试
	//x := []string{"Start"} //初始化时，slice的容量为1
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	y := append(x, "Hello", "World")
	//	fmt.Printf("y slice len:%d, cap:%d, value:%+v\n", len(y), cap(y), y)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	z := append(x, "Java", "Golang", "React")
	//	fmt.Printf("z slice len:%d, cap:%d, value:%+v\n", len(z), cap(z), z)
	//}()
	//wg.Wait()

	x := make([]string, 0, 6)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := make([]string, 0, len(x) + 2) //分配新的内存
		y = append(y, x...)
		y = append(y, "Hello", "World", "!")
		fmt.Printf("y slice len:%d, cap:%d, value:%+v\n", len(y), cap(y), y)
	}()

	go func() {
		defer wg.Done()
		z := make([]string, 0, len(x) + 2) //分配新的内存
		z = append(z, x...)
		z = append(z, "PHP", "Go", "Java")
		fmt.Printf("z slice len:%d, cap:%d, value:%+v\n", len(z), cap(z), z)
	}()
	wg.Wait()
}

func testSlicePointer()  {
	oldS := make([]int, 3, 4)
	oldS[0] = 1
	oldS[1] = 1
	oldS[2] = 1

	fmt.Printf("oldS1--ptr1:%p\n", &oldS)
	fmt.Printf("oldS2--data-ptr2:%p\n", oldS)
	newS := valuePass(oldS)
	fmt.Printf("newS1--ptr1:%p\n", &newS)
	fmt.Printf("newS2--data-ptr2:%p\n", newS)

	fmt.Println(oldS)
	fmt.Println(newS)

	ptrPass(&oldS)
	fmt.Println(oldS)
}

func valuePass(oldS []int) []int {
	// 这里 oldS 虽然改变了，但并不会影响外层函数的 oldS
	fmt.Printf("valuePass--ptr1:%p\n", &oldS)
	fmt.Printf("valuePass--data-ptr2:%p\n", oldS)
	oldS = append(oldS, 100)
	oldS[1] = 2
	return oldS
}

func ptrPass(old *[]int) {
	// 会改变外层 old 本身
	fmt.Printf("ptrPass--ptr0:%p\n", &old)
	fmt.Printf("ptrPass--ptr1:%p\n", old)
	fmt.Printf("ptrPass--data-ptr2:%p\n", *old)
	*old = append(*old, 100)
	(*old)[2] = 3 //注意优先级
	return
}
