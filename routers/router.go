package routers

import (
	"myproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// var filterFunc = func(ctx *context.Context) {
	// 	uid := ctx.Input.Session("uid")
	// 	if uid == nil {
	// 		ctx.Redirect(302, "/login.tpl")
	// 	}
	// }
	// beego.InsertFilter("/getUser/*", beego.BeforeRouter, filterFunc)

	beego.Router("/", &controllers.MainController{})
	beego.Router("user", &controllers.UserController{})
	beego.Router("register", &controllers.AuthController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("login", &controllers.AuthController{}, "get:ShowLogin;post:HandleLogin")

	// //初始化 namespace
	// ns :=
	// beego.NewNamespace("/api",
	// 	// beego.NSCond(func(ctx *context.Context) bool {
	// 	// 	if ctx.Input.Domain() == "api.beego.me" {
	// 	// 		return true
	// 	// 	}
	// 	// 	return false
	// 	// }),
	// 	beego.NSBefore(auth),
	// 	beego.NSNamespace("/user",
	// 		beego.NSBefore(sentry),
	// 		beego.NSGet("/:id", &controllers.UserController{}, "get:Info")
	// 	),
	// )
	// //注册 namespace
	// beego.AddNamespace(ns)
}
