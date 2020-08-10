package server

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"web-test/conf"
	"web-test/model"
	"web-test/util"
)

type UserServer struct {
}

func (us *UserServer) Login(ctx iris.Context) *ResultData {
	PriKey, err := util.ParsePKCS1PrivateKey(util.FormatPKCS1PrivateKey(string(conf.PrivateKey)))
	if err != nil {
		PriKey, err = util.ParsePKCS8PrivateKey(util.FormatPKCS8PrivateKey(string(conf.PrivateKey)))
		if err != nil {
			fmt.Errorf("解密私钥文件信息出错 %s", err.Error())
		}
	}
	loginInfo := util.VerifyRsa(ctx.FormValue("sign_data"), PriKey)
	var Users = &model.UserLoginReq{}
	err = json.Unmarshal(loginInfo, &Users)
	if err != nil {
		fmt.Errorf("json 解码错误", err.Error())
	}
	users := &model.Users{}
	users.SelectByAccount("hjs")
	var ret = &ResultData{}
	if users.Account != Users.Pwd {
		ret = Message("2", "密码错误")
		return ret
	}
	ret = Message("0", "登录成功")
	return ret
}
