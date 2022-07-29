package routers

import (
	"mailgo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Welcome")

	// 登录
	beego.Router("/login", &controllers.MainController{}, "post:SignIn")
	// 注册
	beego.Router("/singup", &controllers.MainController{}, "post:SignUp")
	// 推出
	beego.Router("/logout", &controllers.MainController{}, "post:LogOut")

	// 用户管理
	beego.Router("/user/new", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/user/delete", &controllers.UserController{}, "delete:DeleteUser")
	beego.Router("/user/update", &controllers.UserController{}, "put,post:UpdateUser")
	beego.Router("/user", &controllers.UserController{}, "get:ListUsers")

	// sender增删改查

	// contact增删改查

	// proxy增删改查

	// sender_group
	// contact_group
	// proxy_group

	// mail

	// task

	// job

	// setting
}
