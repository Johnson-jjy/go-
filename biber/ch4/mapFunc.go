package ch4

import "fmt"

// 通过循环实现map间的比较
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

// 使用map来记录提交相同的字符串列表的次数
var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list) //%q记录每个字符串元素的信息
}

func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}

// 待完成：比较字符串时忽略大小写