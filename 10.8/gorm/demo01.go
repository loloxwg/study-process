package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time

}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:12345678@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB,err:=db.DB()
		if err!=nil {
			fmt.Println(err)
		}
		sqlDB.Close()
	}()
	db.AutoMigrate(&User{})
	user := User{Name:"zxw",Age: 20}
	err = db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}


}
