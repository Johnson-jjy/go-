package main

import (
	"fmt"
	"time"
)

func timeDemo()  {
	now := time.Now() // 获取当前时间
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d-%02d-%02d-%02d", year, month, day, hour, minute, second)
}

func timeStampDemo1(){
	now := time.Now() //获取当前时间
	timeStamp1 := now.Unix() //获取时间戳--时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数
	timeStamp2 := now.UnixNano()  //纳秒时间戳

	fmt.Printf("current time : %v \n", timeStamp1)
	fmt.Printf("current tume : %v \n", timeStamp2)
}

func timeStampDemo2(timeStamp int64){
	timeObj := time.Unix(timeStamp, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d-%02d-%02d-%02d", year, month, day, hour, minute, second)
}

/*const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)*/

/*	time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
	time.Duration表示一段时间间隔，可表示的最长时间段大约290年。
*/

func timeDurationDemo()  {
	fmt.Println("now,go!")
	time.Sleep(5 * time.Second)
	fmt.Println("now,off!")
	n := 7
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("注意！Duration是一个类型！")
}

func timeCalDemo()  {
	now := time.Now()
	then := now.Add(time.Hour)
	fmt.Printf("After an hour : %v \n", then)
	past := then.Sub(now)
	fmt.Printf("10 minutes ago : %v \n", past)
	if now.Before(then){
		fmt.Println("now  before then")
	}
	if then.After(now){
		fmt.Println("then after now")
	}
	if !then.Equal(now) {
		fmt.Println("now !equal then")
	}
}

func timeTickDemo()  {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func timeFormatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	//拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
	}

	//按照指定时区和指定格式解析字符串
	timeStr := "2021/04/13 22:11:00"
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(timeObj)
	fmt.Println(now.Sub(timeObj))
}

func main()  {
	//timeDemo()
	//timeStampDemo1()
	//timeStampDemo2(1618319870)
	//timeDurationDemo()
	//timeCalDemo()
	//timeTickDemo()
	timeFormatDemo()
}