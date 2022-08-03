package main

import (
	_ "mailgo/initialize/config" // 加载配置
	_ "mailgo/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
