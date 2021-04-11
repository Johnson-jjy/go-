package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile(){
	//1. 打开和关闭文件
	//os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件

	//只读方式打开当前目录下的main.go文件
	file, err := os.Open("./test.txt")
	if err != nil{
		fmt.Println("open file failed!", err)
		return
	}
	//关闭文件
	defer file.Close() //为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。

	//2. 读取文件file.Read()
	//func (f *File) Read(b []byte) (n int, err error)
	var tmp1 = make([] byte, 128)
	n, err := file.Read(tmp1)
	if err == io.EOF{
		fmt.Println("读完啦！")
		return
	}
	if err != nil{
		fmt.Println("read file failed, err")
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp1[:n]))
}

func readAllFiles(){
	//1. 打开和关闭文件
	//os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件

	//只读方式打开当前目录下的main.go文件
	file, err := os.Open("./test.txt")
	if err != nil{
		fmt.Println("open file failed!", err)
		return
	}
	//关闭文件
	defer file.Close() //为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。

	//3. 循环读取
	var content []byte
	for{
		var tmp2 = make([]byte, 128)
		n, err := file.Read(tmp2)
		if err == io.EOF {
			fmt.Println("文件读完了！")
			fmt.Println(string(tmp2[:n]))
			break
		}
		if err != nil{
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp2[:n]...)
		fmt.Printf("读取了%d字节数据\n", n)
		fmt.Println(string(tmp2[:n]))
	}
	fmt.Println(string(content))
}

//bufio是在file的基础上封装了一层API，支持更多的功能。
func readByBufio(){
	file, err := os.Open("./test.txt")
	if err != nil{
		fmt.Println("open file failed, err", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF{
			if len(line) != 0{
				fmt.Println(line)
			}
			fmt.Println("文件读完了！")
			break
		}
		if err != nil{
			fmt.Println("read file failed, err", err)
			return
		}
		fmt.Print(line)
	}
}

func readByIoutil()  {
	//io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil{
		fmt.Println("read file failed! err:", err)
		return
	}
	fmt.Println(string(content))
}

//os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
//func OpenFile(name string, flag int, perm FileMode) (*File, error) {
//	...
//}
func write()  {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil{
		fmt.Println("read file failed, err:", err)
		return
	}
	defer file.Close()
	str := "test write!"
	file.Write([]byte (str)) //写入字节切片数据
	file.WriteString(str)  //直接写入字符串数据
}

func writeByBufio()  {
	file, err := os.OpenFile("./newtest.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil{
		fmt.Println("read file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("new test i") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func writeByIoutil()  {
	str := "hello new test"
	err := ioutil.WriteFile("./newtest.txt", []byte(str), 0666)
	if err != nil{
		fmt.Println("write file failed! err:", err)
	}
}

func main()  {
	//readFile()
	//readAllFiles()
	//readByBufio
	//readByIoutil
	//write()
	//writeByBufio()
	writeByIoutil()
}
