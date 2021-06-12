package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

//GET请求示例
func testGet()  {
	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))
}

//

func main()  {
	//GET请求实例
	testGet()
}