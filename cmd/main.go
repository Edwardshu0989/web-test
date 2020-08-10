package main

import (
	"github.com/kataras/iris"
	"web-test/controller"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "this is a new product")
		ctx.View("product.html")
	})

	app.PartyFunc("user/", func(users iris.Party) {
		users.Post("/login", controller.UserLogin)
	})

	app.Run(iris.Addr(":8080"))
}
