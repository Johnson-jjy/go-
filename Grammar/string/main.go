package main

import (
	"fmt"
	"strings"
)

// 字符串比较
func testCompare() {
	a := "compareTest"
	b := "CompareTEST"
	c := "compare_test"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(b, b))
	fmt.Println(strings.Compare(a, c))
	fmt.Println(strings.EqualFold(a, b)) // 不区分大小写
	fmt.Println(strings.EqualFold(a, a))
	fmt.Println(strings.EqualFold(b, c))
}

// 是否存在某个字符或字串
func testContain() {
	a := "Go contain test"
	fmt.Println(strings.Contains(a, "test"))
	fmt.Println(strings.Contains(a, "cxk"))
	fmt.Println(strings.Contains(a, "")) // true!
	fmt.Println(strings.ContainsAny(a, "test"))
	fmt.Println(strings.ContainsAny(a, "cxk"))
	fmt.Println(strings.ContainsAny(a, "")) // false!
	fmt.Println(strings.ContainsRune(a, 't'))
	fmt.Println(strings.ContainsAny("", "")) //false!
}

// 字符串出现次数（字符串匹配）
func testCount() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("谷歌中国", "")) // 都是5！wht？！待补充！
	fmt.Println(strings.Count("test", ""))
	fmt.Println(strings.Count("ababa", "aba")) // 不重叠
}

// 字符串分割
func testSplit()  {

}

// 字符串是否有某个前缀或后缀
func testFix()  {
	a := "TestPreFix"
	b := "TestSufFix"
	fmt.Println(strings.HasPrefix(a, "Test"))
	fmt.Println(strings.HasPrefix(a, "test"))
	fmt.Println(strings.HasPrefix(a, "")) // true!
	fmt.Println(strings.HasSuffix(b, "Fix"))
	fmt.Println(strings.HasSuffix(b, "fix"))
	fmt.Println(strings.HasSuffix(b, "")) // true!
}

func main() {
	//string类型
	var address string = "hello"
	fmt.Println(address)
	fmt.Println(address[0], len(address)) //打印h的utf-8编码值
	var s = "我想看看这个输出"
	for i := 0; i < len(address); i++ {
		fmt.Printf("1:%v\n",address[i])
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("2:%v\n", s[i])
	}

	//str3 := `
	//	func main() {
	//		fmt.Println(address)
	//	}

	//fmt.Println(str3)
	//4. 字符串拼接方式 或 fmt.Sprintf
	//str4 := "xiao"
	//str4 += "xiao"
	//fmt.Println(str4)
	//newstr4 := fmt.Sprintf("%s%s", str4, "new")
	//fmt.Println(newstr4)

	//5. 当一行字符串太长时，需要使用到多行字符串，可以如下处理
	//str5 := "adadas" + "adsadasdd" +
	//	"sdfsfsdf" + "dfdfsdf" +
	//	"dsfsfsdf"
	//
	//fmt.Println(str5)

	//6. 转义 \
	//eg：打印一个Windows平台下的文件路径
	//fmt.Println("str := \"c:\\Code\\test\\go.exe\"")

	//7. 分割 strings.Split

	// 字符串比较
	fmt.Println("Test Compare")
	//testCompare()

	// 字符串匹配
	fmt.Println("Test Count")
	//testCount()

	// 判断是否包含 strings.contains
	fmt.Println("Test Contains")
	//testContain()

	// 前缀/后缀判断 strings.HasPreFix, strings.HasSuffix
	fmt.Println("Test Fix")
	//testFix()

	// 字串出现的位置 strings.Index(),strings.LastIndex()

	// join操作 strings.Join(a[]string, sep string)

	//
}
