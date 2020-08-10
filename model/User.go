package model

import "fmt"

type Users struct {
	Account string `gorm:"type:char(191);unique";json:"account"`
	Pwd     string `json:"pwd"`
}

type UserLoginReq struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
}

func (u *Users) SelectByAccount(account string) {
	fmt.Println(Db)
	Db.Debug().Select("pwd").Where("account =?", account).First(&u)
}
