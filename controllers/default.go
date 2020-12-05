package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
//}
func (c *MainController)Get()  {
	c.Ctx.WriteString("恭喜你,通过云服务器的公网ip访问到了我!")
	c.TplName = "register.html"
}