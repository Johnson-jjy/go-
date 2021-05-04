package main

import (
	"fmt"
	"os"
)
//1.向外输出

//1.1 Print
//Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。

//func Print(a ...interface{}) (n int, err error)
//func Printf(format string, a ...interface{}) (n int, err error)
//func Println(a ...interface{}) (n int, err error)
func testPrint(){
	fmt.Print("在终端打印信息：")
	name1 := "cxk"
	fmt.Printf("我是：%s\n",name1)
	fmt.Println("在终端打印单独一行显示")
}

//1.2 Fprint
//将内容输出到一个io.Writer接口类型的变量w中，通常用来往文件中写入内容。
//注意，只要满足io.Writer接口的类型都支持写入。

//func Fprint(w io.Writer, a ...interface{}) (n int, err error)
//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

func testFprint()  {
	//向标注输出写入内容
	fmt.Fprintln(os.Stdout,"向标注输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		fmt.Println("打开文件出错，err：", err)
		return
	}
	name2 := "沙河蔡徐坤"
	//向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写入信息：%s", name2)
}

//1.3 Sprint
//把传入的数据生成并返回一个字符串。
//func Sprint(a ...interface{}) string
//func Sprintf(format string, a ...interface{}) string
//func Sprintln(a ...interface{}) string
func testSprint() {
	s1 := fmt.Sprint("沙河蔡徐坤")
	name3 := "沙河张艺兴"
	age := 18
	s2 := fmt.Sprintf("name:%s, age:%d", name3, age)
	s3 := fmt.Sprintln("沙河马思维")
	fmt.Println(s1, s2, s3)
}

//1.4 Errorf
//根据format参数生成格式化字符串并返回一个包含该字符串的错误。
//通常使用这种方式来自定义错误类型
//func Errorf(format string, a ...interface{}) error
//err := fmt.Errorf("这是一个错误")
//e := errors.New("原始错误e")
//w := fmt.Errorf("Wrap了一个错误%w", e)

func testPrintf()  {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct {
		name string
	}{"小王子"}
	fmt.Printf("%v\n", o)//值的默认格式表示
	fmt.Printf("%#v\n", o)//值的Go语法表示
	fmt.Printf("%T\n", o)//打印值的类型
	fmt.Printf("100%%\n")//百分号

	n := 65
	fmt.Printf("%b\n", n)//表示为二进制
	fmt.Printf("%c\n", n)//该值对应的unicode码值
	fmt.Printf("%d\n", n)//表示为十进制
	fmt.Printf("%o\n", n)//表示为八进制
	fmt.Printf("%x\n", n)//表示为十六进制，使用a-f
	fmt.Printf("%X\n", n)//表示为十六进制，使用A-F

	f := 12.34
	fmt.Printf("%b\n", f)//无小数部分、二进制指数的科学计数法，如-123456p-78
	fmt.Printf("%e\n", f)//科学计数法，如-1234.456e+78
	fmt.Printf("%E\n", f)//科学计数法，如-1234.456E+78
	fmt.Printf("%f\n", f)//有小数部分但无指数部分，如123.456
	fmt.Printf("%g\n", f)//根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	fmt.Printf("%G\n", f)//根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	s := "小王子"
	fmt.Printf("%s\n", s)//直接输出字符串或者[]byte
	fmt.Printf("%q\n", s)//该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	fmt.Printf("%x\n", s)//每个字节用两字符十六进制数表示（使用a-f
	fmt.Printf("%X\n", s)//每个字节用两字符十六进制数表示（使用A-F）

	a := 10
	fmt.Printf("%p\n", &a)//表示为十六进制，并加上前导的0x
	fmt.Printf("%#p\n", &a)

	n2 := 12.34
	fmt.Printf("%f\n", n2)//默认宽度，默认精度
	fmt.Printf("%9f\n", n2)//宽度9，默认精度
	fmt.Printf("%.2f\n", n2)//默认宽度，精度2
	fmt.Printf("%9.2f\n", n2)//宽度9，精度2
	fmt.Printf("%9.f\n", n2)//宽度9，精度0

	s2 := "小王子"
	fmt.Printf("%s\n", s2)
	fmt.Printf("%5s\n", s2)
	fmt.Printf("%-5s\n", s2)//总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
	fmt.Printf("%5.7s\n", s2)//对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
	fmt.Printf("%-5.7s\n", s2)//在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
	fmt.Printf("%5.2s\n", s2)//八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
	fmt.Printf("%05s\n", s2)//使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
}

func main()  {
	//print系列
	testPrint()
	//Fprint系列，网格中输入中写东西
	testFprint()
	//Sprint系列，打印并返回字符串
	testSprint()
	//printf系列，进一步探索各不同类型
	testPrintf()
	//scan输入系列（待未来补全）

}
