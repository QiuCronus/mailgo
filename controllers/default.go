package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Welcome() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// 登录
func (c *MainController) SignIn() {

}

// 退出
func (c *MainController) LogOut() {

}

// 注册
func (c *MainController) SignUp() {

}
