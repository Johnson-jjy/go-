package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main()  {
	//map初始化--与nil

	//map声明
	//1.直接初始化
	//2.make赋值
	scoreMap := map[string]int {
		"Jay" : 90,
		"John" : 100,
	}
	scoreMap["Alex"] = 66
	fmt.Printf("type of a %T\n", scoreMap)

	//判断某个键是否存在
	value, ok := scoreMap["Alex"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}

	//map的遍历 -- 注：遍历map时的元素顺序与添加键值对的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//按照指定的顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var testSeq = make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		testSeq[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 10)
	for testKey := range testSeq {
		keys = append(keys, testKey)
	}
	//对切片排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, testSeq[key])
	}

	//删除键值对
	delete(scoreMap, "Jay")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}


	//自补充：增加键值对--扩容问题？


	//元素为map类型的切片--注意初始化问题（下同）
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("Atfer Init")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["cxk"] = "basketball"
	mapSlice[0]["address"] = "beijing"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}


	//值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("After Init")
	country := "China"
	city, ok := sliceMap[country]
	if !ok {
		city = make([]string, 0, 2)
	}
	city = append(city, "CDC", "CSC")
	sliceMap[country] = city
	fmt.Println(sliceMap)


	//练习题：统计一个字符串中每个单词出现的次数
}