package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err := gorm.Open("mysql", "root:ai132818@tcp(localhost:3306)/web-test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Errorf("数据库连接错误.err %s", err.Error()))
	}
	Db.DB().SetMaxIdleConns(20)
	Db.DB().SetMaxOpenConns(50)
	Db.AutoMigrate(&Users{})
	fmt.Println(Db)
}
