package main

import "fmt"

func main() {
	//数组的声明
	var a1 [3]int
	fmt.Println(a1)

	//数组的初始化
	//1.标准初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)

	//2.编译器自动识别的初始化
	var schoolArray = [...]string{"UESTC", "MIT", "CMU"}
	fmt.Println(schoolArray)

	//3.通过索引下表进行初始化
	var testArray = [...]int{1: 1, 3: 3, 5: 5}
	fmt.Println(testArray)

	//数组的遍历
	//1.标准for
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	//2.for_range
	for index, value := range schoolArray {
		fmt.Println(index, value)
	}

	for index := range schoolArray {
		fmt.Println(index)
	}

	for _, value := range schoolArray {
		fmt.Println(value)
	}

	//二维（多维）数组
	var company = [3][2]string{
		{"tx", "ali"},
		{"by", "baidu"},
		{"huawei", "xiaomi"},
	}//注：仅外层可用...

	for _, v1 := range company {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	fmt.Println(company[2][0])//支持索引访问

	//数组是值类型
	numTest:=[3]int{2,6,10}
	fmt.Println(numTest[1])
	numTest[1]=8
	fmt.Println(numTest[1])//同一个函数域内真实修改

	test(numTest)
	fmt.Println(numTest[1])//在func_test中修改的只是t的副本
}

func test(t [3]int)  {
	t[1]=100
}
