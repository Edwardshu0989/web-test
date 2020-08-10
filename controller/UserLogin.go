package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
	"web-test/server"
)

type UserController struct {
	UserServer *server.UserServer
}

func UserLogin(ctx iris.Context) {
	uc := &UserController{}
	ret := uc.UserServer.Login(ctx)
	retJson := context.JSON{
		StreamingJSON: false,
		UnescapeHTML:  false,
		Indent:        ret.Code,
		Prefix:        ret.Message,
	}
	ctx.JSON(http.StatusOK, retJson)
}
