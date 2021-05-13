package main

import (
//config
	"fmt"
	. "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main() {
	db, err := Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	//注意此处若为中文则还需修改mysql配置
	//u1 := UserInfo{1, "cxk", "male", "basketball"}
	//u2 := UserInfo{2, "yy", "female", "football"}
	// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)

	//建议如下写法：更严谨
/*	if err := db.Create(u1).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}*/
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "football")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "soccer")
	// 删除
	db.Delete(&u)
}

