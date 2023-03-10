package main

import (
	"ginchat/model"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@(43.139.195.17:8686)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 迁移 schema
	db.AutoMigrate(&model.User{})

	// Create

	//// Read
	//var User model.User
	////db.First(&User, 1) // 根据整型主键查找
	//db.First(&User, "Username", "D42") // 查找 code 字段值为 D42 的记录
	////
	////Update - 将 model.User 的 price 更新为 200
	//db.Model(&User).Update("phone", 1234567890)
	////// Update - 更新多个字段
	//db.Model(&User).Updates(model.User{UserName: "hehe"}) // 仅更新非零值字段
	//db.Model(&User).Updates(map[string]interface{}{"UserName": "niuma", "Phone": "181664848799"})
	//
	//// Delete - 删除 model.User
	//db.Delete(&User, 1)

}
